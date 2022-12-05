const { defineConfig } = require('@vue/cli-service')
// module.exports = defineConfig({
//   transpileDependencies: true
// })
module.exports = {
  configureWebpack:{
    resolve:{
      alias:{
        'assets':'@/assets',
        'pages':'@/pages',
        'components':'@/components',
        'network':'@/network'
      }
    }
  }
}
