const path = require("path");
const merge = require("webpack-merge");
const base = require("./webpack.base.config");

let dev = {
  entry: {
    "attribution-web": path.resolve(__dirname, "../example/index.ts")
  },
  devtool: "inline-source-map",
  devServer: {
    contentBase: path.resolve(__dirname, "../example"),
    port: 8080,
    overlay: true
  },
  output: {
    filename: "[name].js",
    path: path.resolve(__dirname, "../example")
  },
  mode: "development"
};

module.exports = merge(base, dev);
