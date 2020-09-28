const path = require("path");

module.exports = {
  resolve: {
    extensions: [".ts", ".json", ".js"]
  },
  output: {
    library: "AttributionWeb",
    libraryExport: "default",
    libraryTarget: "umd"
  },
  plugins: [],
  module: {
    rules: [
      {
        test: /\.js$/,
        loader: "source-map-loader",
        enforce: "pre"
      },
      {
        test: /\.ts$/,
        exclude: /node_modules/,
        use: [
          {
            loader: "babel-loader"
          },
          {
            loader: "tslint-loader",
            options: {
              configFile: path.resolve(__dirname, "../tslint.json"),
              emitErrors: true,
              fix: false,
              tsConfigFile: path.resolve(__dirname, "../tsconfig.json")
            }
          }
        ]
      }
    ]
  }
};
