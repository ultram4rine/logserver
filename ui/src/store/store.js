import { createStore } from "vuex";

import axios from "axios";
import Cookies from "js-cookie";

import { AUTH_LOGIN, AUTH_LOGOUT } from "@/store/actions";

export default createStore({
  state: {
    token: Cookies.get("info"),
  },
  actions: {
    [AUTH_LOGIN]: (context, user) => {
      return new Promise((resolve, reject) => {
        axios
          .post("/auth", user)
          .then((resp) => {
            context.commit(AUTH_LOGIN, Cookies.get("info"));
            resolve(resp);
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    [AUTH_LOGOUT]: (context) => {
      return new Promise((resolve, reject) => {
        context.commit(AUTH_LOGOUT);
        axios.post("/logout").catch((err) => {
          reject(err);
        });
        resolve();
      });
    },
  },
  mutations: {
    [AUTH_LOGIN]: (state, token) => {
      state.token = token;
    },
    [AUTH_LOGOUT]: (state) => {
      state.token = "";
    },
  },
});
