const CopyWebpackPlugin = require('copy-webpack-plugin')
const MiniCssExtractPlugin = require("mini-css-extract-plugin")

module.exports = {
  entry: [
    './src/old.thing/resources/assets/reactjs/index.js',
    './src/old.thing/resources/assets/sass/app.scss',
  ],
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        use: ['babel-loader']
      },
      {
        test: /\.(scss|css)$/,
        use: [
          {
            loader: MiniCssExtractPlugin.loader,
            options: {
              // you can specify a publicPath here
              // by default it use publicPath in webpackOptions.output
            }
          },
            "css-loader", // translates CSS into CommonJS
            "sass-loader" // compiles Sass to CSS, using Node Sass by default
        ]
      },
      {
				test: /\.(ttf|eot|svg|gif|woff)(\?v=[0-9]\.[0-9]\.[0-9])?$/,
				loader: 'file-loader',
				query: {
					name: 'images/[name].[hash:8].[ext]'
				}
			},
    ]
  },
  resolve: {
    extensions: ['*', '.js', '.jsx']
  },
  output: {
    path: __dirname + '/src/old.thing/dist',
    publicPath: '/',
    filename: '[name].js'
  },
  plugins: [
    new CopyWebpackPlugin([

    ]),
    new MiniCssExtractPlugin({
      // Options similar to the same options in webpackOptions.output
      // both options are optional
      filename: '[name].css',
      chunkFilename: '[id].css',
    })
  ],
};
