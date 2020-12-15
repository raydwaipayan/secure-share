<div align="center">

<img src="https://socialify.git.ci/raydwaipayan/secure-share/image?description=1&descriptionEditable=A%20go%20web%20app%20to%20securely%20share%20files%20over%20network&font=Rokkitt&forks=1&issues=1&language=1&owner=1&pattern=Brick%20Wall&pulls=1&stargazers=1&theme=Light" />

<b>secure-share</b>
<br/>
<a href="https://goreportcard.com/report/github.com/raydwaipayan/secure-share">
		<img src="https://goreportcard.com/badge/github.com/raydwaipayan/secure-share" alt="Go Report Card" />
</a>
<p>
Share files securely over network.
</br>
Written with Golang + Vue.
</div>

## Inspiration
I needed a file sharing service which had the following:

- Minimal but simple interface
- Data is encrypted using strong cryptographic algorithm on server side
- Only a user with the crypto key could unlock the file
- File would be deleted after first download
- Material Dark design is a plus

Secure-share is built with all of these in mind.

## Cryptographic Details
The cryptographic algorithm used is AES-256.
GCM mode is used. Since GCM is a stream cipher mode,
a nonce is required. Both the nonce and keys are randomly generated using `crypto/rand`.

What the server stores:
- Unencrypted metadata - Only filename and size
- Encrypted file data

The keys are never stored physically. If the user loses the
key, it is impossible to recover the file.

## Development
The application can be easily started using the makefile
The following construct starts both the client and server:

```lang=bash
make run-all
```

Note that the .env file must exist under both root and client/
directories. The make constructs automatically does that for you.
For manual runs:

```lang=bash
cp sample.env .env
cp .env ./client/.env
```

## Deployment
The application can be deployed using docker-compose.

```lang=bash
docker-compose up -d
```

