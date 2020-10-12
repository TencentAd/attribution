const devConfig = require("./webpack/webpack.dev.config");
const prodConfig = require("./webpack/webpack.prod.config");

module.exports = process.env.NODE_ENV === "production" ? prodConfig : devConfig;
