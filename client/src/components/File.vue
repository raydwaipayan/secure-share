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
            <div v-if="error">
              <p class="text-h5 red--text text-center">
                Error: The requested file could not be found.
              </p>
              <div class="d-flex justify-center my-10">
                <v-btn
                  color="error"
                  dark
                  @click.prevent="$router.push({ path: '/' })"
                >
                  Go Back
                  <v-icon dark right>
                    mdi-checkbox-marked-circle
                  </v-icon>
                </v-btn>
              </div>
            </div>
            <div v-else>
              <v-row>
                <v-spacer></v-spacer>
                <v-col cols="12" sm="6">
                  <v-card elevation="2" tile min-height="50vh">
                    <v-card-title class="text-h5">
                      File Metadata:
                    </v-card-title>
                    <v-card-text class="text-h6 pa-0 pl-6">
                      Filename:
                      <p class="purple--text accent-4">
                        {{ data.name }}
                      </p>
                    </v-card-text>
                    <v-card-text class="text-h6 pa-0 pl-6">
                      Size in bytes:
                      <p class="purple--text accent-4">
                        {{ data.size }}
                      </p>
                    </v-card-text>
                    <v-btn
                      large
                      :loading="loading"
                      :disabled="loading"
                      color="purple darken-2"
                      class="white--text ml-6"
                      @click="download()"
                    >
                      Download
                      <v-icon right dark>
                        mdi-cloud-download
                      </v-icon>
                    </v-btn>
                    <div class="ma-6">
                      <v-progress-linear
                        v-model="downloadprogress"
                        :active="loading"
                        rounded
                        height="5"
                        color="purple --accent-4"
                        value="15"
                      ></v-progress-linear>
                    </div>
                    <v-card-text v-if="downloaderror" class="red--text text-h5">
                      {{ downloaderror }}
                      <div>
                        <a
                          class="red--text text-h6"
                          @click="$router.push({ path: '/' })"
                        >
                          Go Back?</a
                        >
                      </div>
                    </v-card-text>
                  </v-card>
                </v-col>
                <v-spacer></v-spacer>
              </v-row>
            </div>
          </v-sheet>
        </v-col>
      </v-row>
    </v-container>
  </v-main>
</template>

<script>
import Vue from "vue";
export default {
  data() {
    return {
      error: null,
      data: null,
      query: null,
      downloaderror: null,
      loading: false,
      downloadprogress: 0,
    };
  },
  beforeRouteEnter(to, from, next) {
    const params = {
      fileid: to.query.fileid,
    };

    Vue.axios("/file/info", {
      params: params,
    })
      .then((response) => {
        next((vm) => vm.setData(null, response.data, to.query));
      })
      .catch((error) => {
        next((vm) => vm.setData(error, null, to.query));
      });
  },

  beforeRouteUpdate(to, from, next) {
    this.data = null;
    const params = {
      fileid: to.query.fileid,
    };

    Vue.axios("/file/info", {
      params: params,
    })
      .then((response) => {
        this.setData(null, response.data, to.query);
        next();
      })
      .catch((error) => {
        this.setData(error, null, to.query);
      });
  },
  methods: {
    setData(err, data, query) {
      if (err) {
        this.error = err.toString();
      } else {
        this.data = data;
        this.query = query;
      }
    },
    download() {
      this.loading = true;
      let formdata = new FormData();
      formdata.append("fileid", this.query.fileid);
      formdata.append("key", this.query.key);
      Vue.axios({
        url: "/file/retrieve",
        method: "post",
        data: formdata,
        responseType: "blob",
        onDownloadProgress: (progressEvent) => {
            var percentCompleted = Math.round((progressEvent.loaded * 100) / progressEvent.total);
            this.downloadprogress=percentCompleted
        },
      })
        .then((response) => {
          var fileURL = window.URL.createObjectURL(new Blob([response.data]));
          var fileLink = document.createElement("a");

          fileLink.href = fileURL;
          fileLink.setAttribute("download", this.data.name);
          document.body.appendChild(fileLink);

          fileLink.click();
        })
        .catch(async (error) => {
          const text = await new Response(error.response.data).text();
          this.downloaderror = text;
        })
        .finally(() => {
          this.loading = false;
        });
    },
  },
};
</script>
