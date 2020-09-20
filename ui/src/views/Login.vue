<template>
  <v-main>
    <v-container class="fill-height" fluid>
      <v-row align="center" justify="center">
        <v-col cols="12" sm="8" md="4">
          <v-card class="elevation-12">
            <v-toolbar dark flat>
              <v-toolbar-title>Login</v-toolbar-title>
            </v-toolbar>
            <v-card-text>
              <v-form>
                <v-text-field
                  v-model="username"
                  label="Name"
                  type="text"
                  required
                  :prepend-icon="this.mdiAccount"
                ></v-text-field>
                <v-text-field
                  v-model="password"
                  label="Password"
                  :type="show ? 'text' : 'password'"
                  required
                  :prepend-icon="this.mdiKey"
                  :append-icon="show ? this.mdiEye : this.mdiEyeOff"
                  @click:append="show = !show"
                ></v-text-field>
              </v-form>
            </v-card-text>
            <v-card-actions>
              <v-spacer />
              <v-btn color="primary" class="mr-4" @click="login">
                Sign in
                <v-icon right>{{ mdiLogin }}</v-icon>
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </v-main>
</template>

<script>
import { mdiEye, mdiEyeOff, mdiAccount, mdiKey, mdiLogin } from "@mdi/js";

export default {
  name: "Login",

  data() {
    return {
      mdiEye: mdiEye,
      mdiEyeOff: mdiEyeOff,
      mdiAccount: mdiAccount,
      mdiKey: mdiKey,
      mdiLogin: mdiLogin,
      show: false,
      username: "",
      password: "",
    };
  },

  methods: {
    login: function () {
      const { username, password } = this;
      this.$store
        .dispatch("AUTH_LOGIN", { username, password })
        .then(() => {
          this.$router.push("/");
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
};
</script>
