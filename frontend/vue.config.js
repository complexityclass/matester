// vue.config.js
module.exports = {
  configureWebpack: () => {
    if (process.env.NODE_ENV === 'production') {
      // изменение конфигурации для production...
    } else {
      // изменения для разработки...
    }
  },
  runtimeCompiler: false,
  devServer: {
    proxy: 'https://matester23.herokuapp.com',
    contentBase: './dist',
    host: 'localhost',
    disableHostCheck: true,
    hot: false,
    liveReload: false,
    overlay: {
      warnings: true,
      errors: true
    }
  }
}
