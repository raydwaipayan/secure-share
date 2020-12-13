package handlers

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/raydwaipayan/secure-share/crypto"
)

const (
	// KeySize is the AES key size
	KeySize int = 32
)

// Filedata wraps over the file object
type Filedata struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func genFileID(filename string) string {
	t := time.Now().UnixNano()
	fileid := fmt.Sprintf("%s_%v", filename, t)
	return fileid
}

func getFileName(fileid string) (string, error) {
	arr := strings.Split(fileid, "_")
	if len(arr) < 2 {
		return "", errors.New("Invalid fileid")
	}
	return strings.Join(arr[:len(arr)-1], ""), nil
}

func genKey() ([]byte, error) {
	key := make([]byte, KeySize)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return key, err
	}

	return key, nil
}

func writeFile(data []byte, filepath string) error {
	w, err := os.Create(filepath)
	defer w.Close()
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func readFile(filepath string) ([]byte, error) {
	r, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, r)
	return buf.Bytes(), nil
}

// Submit handler takes care of encrypting and storing the file
func Submit(c *gin.Context) {
	var filedata Filedata
	if err := c.ShouldBind(&filedata); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	file, err := filedata.File.Open()
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid file format")
		return
	}

	data := make([]byte, filedata.File.Size)
	_, err = file.Read(data)
	if err != nil {
		c.String(http.StatusBadRequest, "Corrupted file")
		return
	}

	key, err := genKey()
	if err != nil {
		c.String(http.StatusInternalServerError, "Generator error")
		log.Print(err)
		return
	}

	encdata, err := crypto.Encrypt(data, key)
	if err != nil {
		c.String(http.StatusInternalServerError, "Crypto Error")
		log.Print(err)
		return
	}

	fileid := genFileID(filedata.File.Filename)
	wd, _ := os.Getwd()
	err = os.MkdirAll(filepath.Join(wd, "data"), 0777)
	if err != nil {
		c.String(http.StatusInternalServerError, "Filesystem access denied")
		log.Print(err)
		return
	}

	err = writeFile(encdata, filepath.Join(wd, "data", fileid))
	if err != nil {
		c.String(http.StatusInternalServerError, "Filesystem write Error")
		log.Print(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"key":    hex.EncodeToString(key),
		"fileid": fileid,
	})
}

// Retrieve handler takes care of decrypting and returning the file
func Retrieve(c *gin.Context) {
	fileid := c.Query("file")
	filename, err := getFileName(fileid)
	if err != nil {
		c.String(http.StatusInternalServerError, "Invalid fileid")
		log.Print(err)
		return
	}

	hexkey := c.Query("key")
	wd, _ := os.Getwd()

	data, err := readFile(filepath.Join(wd, "data", fileid))
	if err != nil {
		c.String(http.StatusInternalServerError, "File doesn't exist")
		log.Print(err)
		return
	}

	key, err := hex.DecodeString(hexkey)
	if err != nil {
		c.String(http.StatusInternalServerError, "Invalid key format")
		log.Print(err)
		return
	}

	decdata, err := crypto.Decrypt(data, key)
	if err != nil {
		c.String(http.StatusInternalServerError, "Wrong key")
		log.Print(err)
		return
	}

	// Delete the file after first access
	err = os.Remove(filepath.Join(wd, "data", fileid))
	if err != nil {
		log.Print(err)
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", fmt.Sprintf("%d", len(decdata)))
	c.Writer.Write(decdata)
}
