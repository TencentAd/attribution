const path = require("path");
const merge = require("webpack-merge");
const UglifyJsPlugin = require("uglifyjs-webpack-plugin");
const base = require("./webpack.base.config");
const CleanWebpackPlugin = require("clean-webpack-plugin");

let prod = {
  entry: {
    "attribution-miniprogram": path.resolve(__dirname, "../src/index.ts")
  },
  plugins: [
    new CleanWebpackPlugin(),
    new UglifyJsPlugin({
      test: /\.ts$/i,
      uglifyOptions: {
        compress: {
          pure_funcs: ["console.log", "alert"] // https://webpack.js.org/plugins/uglifyjs-webpack-plugin/
        }
      }
    })
  ],
  output: {
    filename: `[name].min.js`,
    path: path.resolve(__dirname, "../dist"),
    publicPath: "//i.gtimg.cn/ams-web/attribution-miniprogram/"
  },
  mode: "production"
};

module.exports = merge(base, prod);
