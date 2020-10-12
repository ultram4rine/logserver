<template>
  <v-app>
    <router-view></router-view>
  </v-app>
</template>

<script>
import axios from "axios";

import { AUTH_LOGOUT } from "./store/actions";

export default {
  name: "App",

  setup() {},

  created() {
    axios.interceptors.response.use(
      (response) => {
        return response;
      },
      (err) => {
        return new Promise(() => {
          if (err.response.status === 401) {
            this.$store.dispatch(AUTH_LOGOUT);
            this.$router.push("/login");
          }
          throw err;
        });
      }
    );
  },
};
</script>
