module.exports = {
  entry: [
    './src/old.thing/resources/assets/js/reactjs/index.js'
    // './src/old.thing/resources/assets/sass/app.scss'
  ],
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        use: ['babel-loader']
      },
      {
        test: /\.scss$/,
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
      }
    ]
  },
  resolve: {
    extensions: ['*', '.js', '.jsx']
  },
  output: {
    path: __dirname + '/src/old.thing/dist/bundle',
    publicPath: '/',
    filename: '[name].js'
  }
};
