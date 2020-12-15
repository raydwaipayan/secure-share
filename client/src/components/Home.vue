<template>
  <v-main class="black">
    <v-container>
      <v-row>
        <v-col>
          <v-sheet
            min-height="70vh"
            rounded="lg"
            class="grey darken-4 py-16 px-6"
          >
            <v-file-input
              v-model="files"
              color="deep-purple accent-4 dark"
              class="text-h6"
              counter
              label="File input"
              placeholder="Select your file"
              prepend-icon="mdi-paperclip"
              outlined
              :show-size="1000"
            >
              <template v-slot:selection="{ index, text }">
                <v-chip
                  v-if="index < 2"
                  color="deep-purple accent-4"
                  dark
                  label
                >
                  {{ text }}
                </v-chip>
              </template>
            </v-file-input>
            <v-btn
              large
              :loading="loading"
              :disabled="loading || !files"
              color="purple darken-2"
              class="white--text ml-8"
              @click="upload()"
            >
              Upload
              <v-icon right dark>
                mdi-cloud-upload
              </v-icon>
            </v-btn>
            <div class="mt-10"></div>
            <v-sheet
              class="ml-7 mr-7 pa-2 text-h6"
              min-height="20vh"
              rounded="lg"
            >
              <div v-if="responseData.key">
                <p>
                  Link: <a :href="getLink()">{{ getLink() }}</a>
                </p>
                <v-btn
                  class="ml-0 my-4"
                  color="success"
                  dark
                  @click.prevent="copy()"
                >
                  Copy
                  <v-icon dark right>
                    mdi-checkbox-marked-circle
                  </v-icon>
                </v-btn>
                <v-snackbar v-model="copydialog" :timeout="timeout">
                  Copied to Clipboard!

                  <template v-slot:action="{ attrs }">
                    <v-btn
                      color="blue"
                      text
                      v-bind="attrs"
                      @click="snackbar = false"
                    >
                      Close
                    </v-btn>
                  </template>
                </v-snackbar>
              </div>
              <div v-if="responseError">
                <p class="red--text">
                  {{ responseError }}
                </p>
              </div>
            </v-sheet>
          </v-sheet>
        </v-col>
      </v-row>
    </v-container>
  </v-main>
</template>

<script>
import Vue from "vue";
export default {
  data: () => ({
    files: null,
    loading: false,
    responseData: {
      key: null,
      fileid: null,
    },
    responseError: "",
    timeout: 3000,
    copydialog: false,
  }),
  methods: {
    async upload() {
      if (!this.files) return;
      this.loading = true;
      this.responseData.key = null;
      this.responseData.fileid = null;
      this.responseError = null;
      let formdata = new FormData();
      formdata.append("file", this.files);

      Vue.axios({
        method: "post",
        url: "/file/submit",
        data: formdata,
        headers: { "Content-Type": "multipart/form-data" },
      })
        .then((response) => {
          const { key, fileid } = response.data;
          this.responseData.key = key;
          this.responseData.fileid = fileid;
        })
        .catch((error) => {
          this.responseError = error.message;
        })
        .finally(() => {
          this.loading = false;
        });
    },
    getLink() {
      return (
        this.$clientUri +
        "/get?fileid=" +
        this.responseData.fileid +
        "&key=" +
        this.responseData.key
      );
    },
    copy() {
      const link = this.getLink();
      this.$copyText(link).then(() => {
        this.copydialog = true;
      });
    },
  },
};
</script>

<style scoped>
p {
  word-break: break-all;
}
</style>
