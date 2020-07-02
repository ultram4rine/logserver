import Vue from "vue";

import router from "./routes/router";
import store from "./store/store";
import vuetify from "./plugins/vuetify";

import App from "./App.vue";

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  vuetify,

  render: (h) => h(App),
}).$mount("#app");
