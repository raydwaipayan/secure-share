<template>
  <v-app>
    <v-app-bar app color="dark" dark> </v-app-bar>
    <v-spacing></v-spacing>
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
                :disabled="loading"
                color="purple darken-2"
                class="white--text ml-8"
                @click="upload()"
              >
                Upload
                <v-icon right dark>
                  mdi-cloud-upload
                </v-icon>
              </v-btn>
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import Vue from "vue";
export default {
  name: "App",
  components: {},

  data: () => ({
    files: null,
    loading: false,
  }),
  methods: {
    upload() {
      if (!this.files) return;

      let formdata = new FormData();
      formdata.append("file", this.files);

      Vue.axios({
        method: "post",
        url: "/file/put",
        data: formdata,
        headers: { "Content-Type": "multipart/form-data" },
      })
        .then(function(response) {
          //handle success
          console.log(response);
        })
        .catch(function(response) {
          //handle error
          console.log(response);
        });
    },
  },
};
</script>
