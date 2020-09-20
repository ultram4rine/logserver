import Vue from "vue";
import VueRouter from "vue-router";

import store from "@/store/store";

import Home from "@/views/Home.vue";
import Login from "@/views/Login.vue";

Vue.use(VueRouter);

const router = new VueRouter({
  mode: "history",
  routes: [
    {
      path: "/login",
      component: Login,
      meta: { skipIfAuth: true },
    },
    {
      path: "/",
      component: Home,
      meta: { requiresAuth: true },
    },
  ],
});

router.beforeEach((to, _from, next) => {
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    if (!store.getters["isAuthenticated"]) {
      next({
        path: "/login",
        query: { redirect: to.fullPath },
      });
    } else {
      next();
    }
  } else if (to.matched.some((record) => record.meta.skipIfAuth)) {
    if (store.getters.isAuthenticated) {
      next({ path: "/" });
    } else {
      next();
    }
  } else {
    next();
  }
});

export default router;
