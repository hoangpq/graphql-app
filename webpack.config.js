const path = require('path');
const webpack = require('webpack');

module.exports = {
  // cheap-module-source-map
  devtool: 'cheap-module-source-map',
  entry: [
    './app/index',
  ],
  output: {
    path: path.join(__dirname, 'dist'),
    filename: 'bundle.js',
    publicPath: '/static/',
  },
  module: {
    rules: [
      {
        test: /\.js/,
        exclude: /(node_modules|bower_components)/,
        include: [
          path.join(__dirname, 'app'),
        ],
        use: [
          {
            loader: 'babel-loader',
          }
        ],
      },
      {
        test: /\.css/,
        use: [
          {
            loader: 'style-loader',
          },
          {
            loader: 'css-loader',
          }
        ]
      }
    ]
  },
  plugins: [
    new webpack.DefinePlugin({
      'process.env': {
        'NODE_ENV': JSON.stringify('production'),
      },
    }),
    new webpack.optimize.UglifyJsPlugin({
      // compress production
      sourceMap: true,
      parallel: true,
      exclude: /node_modules/,
      compress: true,
    })
  ],
};
