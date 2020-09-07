const path = require("path");
const { CleanWebpackPlugin } = require("clean-webpack-plugin");

module.exports = {
  outputDir: path.resolve(__dirname, "../public"),
  assetsDir: process.env.NODE_ENV === "production" ? "static" : "",

  devServer: {
    public: "localhost:8080",
    headers: {
      "Access-Control-Allow-Origin": "*",
    },
  },

  configureWebpack: {
    plugins: [new CleanWebpackPlugin()],
  },

  lintOnSave: false,
};
