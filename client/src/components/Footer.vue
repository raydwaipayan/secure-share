<template>
  <v-footer v-bind="localAttrs" :padless="padless" class="ma-0 pa-0">
    <v-card flat tile width="100%" class="grey darken-4 text-center">
      <v-card-text class="pb-1">
        <v-btn
          v-for="icon in icons"
          :key="icon"
          :href="icon.link"
          class="mx-4"
          icon
        >
          <v-icon size="24px">
            {{ icon.code }}
          </v-icon>
        </v-btn>
      </v-card-text>
      <v-card-text class="white--text pa-0 pb-1">
        {{ new Date().getFullYear() }} — <strong>Dwai</strong>
      </v-card-text>
    </v-card>
  </v-footer>
</template>

<script>
export default {
  data: () => ({
    icons: [
      { code: "mdi-home", link: process.env.VUE_APP_CLIENT_URL },
      { code: "mdi-github", link: "https://github.com/raydwaipayan" },
    ],
    items: ["default", "absolute", "fixed"],
    padless: false,
    variant: "default",
  }),
  computed: {
    localAttrs() {
      const attrs = {};

      if (this.variant === "default") {
        attrs.absolute = false;
        attrs.fixed = false;
      } else {
        attrs[this.variant] = true;
      }
      return attrs;
    },
  },
};
</script>
