import Vue from "vue";
import VueRouter from "vue-router";

import Home from "./components/Home";
import Login from "./components/Login";

Vue.use(VueRouter);

export default new VueRouter({
  mode: "history",
  routes: [
    { path: "/", component: Home, meta: { requiresAuth: true } },
    { path: "/login", component: Login },
  ],
});
