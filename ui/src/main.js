import Vue from "vue";
import VueCompositionAPI from "@vue/composition-api";

import router from "@/routes/router";
import vuetify from "@/plugins/vuetify";

import App from "@/App.vue";

Vue.use(VueCompositionAPI);
Vue.config.productionTip = false;

new Vue({
  router,
  vuetify,

  render: (h) => h(App),
}).$mount("#app");
