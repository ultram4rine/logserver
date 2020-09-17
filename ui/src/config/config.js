export let config;

if (!process.env.NODE_ENV || process.env.NODE_ENV === "production") {
  config = {
    apiURL: "/api",
  };
} else {
  config = {
    apiURL: "/api",
  };
}
