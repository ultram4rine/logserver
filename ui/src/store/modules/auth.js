import axios from "axios";
import Cookies from "js-cookie";

import { AUTH_LOGIN, AUTH_LOGOUT } from "@/store/actions";

const state = {
  token: Cookies.get("info"),
};

const getters = {
  isAuthenticated: (state) => {
    return !!state.token;
  },
};

const actions = {
  [AUTH_LOGIN]: (context, user) => {
    return new Promise((resolve, reject) => {
      context.commit(AUTH_LOGIN);
      axios
        .post("/auth", user)
        .then((resp) => {
          state.token = Cookies.get("info");
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
};

const mutations = {
  [AUTH_LOGOUT]: (state) => {
    state.token = "";
  },
};

const auth = {
  state,
  getters,
  actions,
  mutations,
};

export default auth;
