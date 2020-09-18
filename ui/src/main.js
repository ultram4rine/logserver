import Vue from "vue";
import VueCompositionAPI from "@vue/composition-api";

import axios from "axios";

import router from "@/routes/router";
import store from "@/store/store";
import vuetify from "@/plugins/vuetify";

import App from "@/App.vue";

Vue.use(VueCompositionAPI);
Vue.config.productionTip = false;

const token = localStorage.getItem("user-token");
if (token) {
  axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;
}

new Vue({
  router,
  store,
  vuetify,

  render: (h) => h(App),
}).$mount("#app");
