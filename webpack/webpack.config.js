const path = require('path');
const webpack = require('webpack');

module.exports = {
  mode: 'development',
  context: path.resolve(__dirname, '..'),
  entry: './assets/docs.js',
  output: {
    filename: 'docs.js',
    path: path.resolve(__dirname, '..', 'public', 'assets')
  },
  module: {
    rules: [
      {
        test: /\.css$/,
        use: [
          'style-loader',
          'css-loader'
        ]
      },
      {
        test: /\.js$/
      },
      {
        test: /\.scss$/,
        use: [
          { loader: 'style-loader' },
          { loader: 'css-loader' },
          {
            loader: 'postcss-loader',
            options: {
              plugins: function () {
                return [
                  require('precss'),
                  require('autoprefixer')
                ];
              }
            }
          },
          { loader: 'sass-loader' }
        ]
      },
      {
        test: /.(ttf|otf|eot|svg|woff(2)?)(\?[a-z0-9]+)?$/,
        exclude: /images/,
        use: [
          {
            loader: 'file-loader',
            options: {
              name: '[name].[ext]',
              outputPath: 'fonts/',
              publicPath: '/assets/fonts/'
            }
          }
        ]
      }
    ]
  },
  plugins: [
    new webpack.ProvidePlugin({ $: 'jquery', jquery: 'jquery', jQuery: 'jquery' })
  ],
  resolve: {
    modules: [ path.resolve(__dirname, 'node_modules') ],
  },
  resolveLoader: {
    modules: [ path.resolve(__dirname, 'node_modules') ],
  },
};
