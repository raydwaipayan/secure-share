import Vue from "vue";
import VueRouter from "vue-router";
import App from "./App.vue";
import vuetify from "./plugins/vuetify";
import axios from "axios";
import VueAxios from "vue-axios";
import VueClipboard from "vue-clipboard2";

//import views to serve through router view
import Home from "./components/Home";
import File from "./components/File";

Vue.use(VueRouter);
Vue.use(VueAxios, axios);
Vue.use(VueClipboard);

axios.defaults.baseURL = process.env.baseURL || "http://localhost:8080";
Vue.prototype.$clientUri = process.env.clientURL || "http://localhost:3000";
Vue.config.productionTip = false;

const routes = [
  {
    path: "/",
    component: Home,
  },
  {
    path: "/get",
    component: File,
  },
];

const router = new VueRouter({
  mode: "history",
  routes,
});

new Vue({
  vuetify,
  router,
  render: (h) => h(App),
}).$mount("#app");
