import Vue from "vue";
import VueRouter from "vue-router";

import Home from "../components/Home.vue";
import Login from "../components/Login.vue";

Vue.use(VueRouter);

export default new VueRouter({
  mode: "history",
  routes: [
    { path: "/", component: Home, meta: { requiresAuth: true } },
    { path: "/login", component: Login },
  ],
});
