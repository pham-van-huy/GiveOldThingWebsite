const CopyWebpackPlugin = require('copy-webpack-plugin')

module.exports = {
  entry: [
    './src/old.thing/resources/assets/reactjs/index.js',
    './src/old.thing/resources/assets/sass/app.scss'
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
              loader: "style-loader"
            },
            {
              loader: 'file-loader',
              options: {
                name: '/[name].css',
              }
            },
            {
              loader: 'sass-loader'
            }
        ]
      },
      {
				test: /\.(ttf|eot|svg)(\?v=[0-9]\.[0-9]\.[0-9])?$/,
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

    ])
  ],
};
