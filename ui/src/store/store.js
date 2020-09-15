import Vue from "vue";
import Vuex from "vuex";

import auth from "@/store/modules/auth";

Vue.use(Vuex);

const debug = process.env.NODE_ENV !== "production";

const store = new Vuex.Store({
  modules: {
    auth: auth,
  },
  strict: debug,
});

export default store;
