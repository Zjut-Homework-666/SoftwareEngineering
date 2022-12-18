const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true
})
module.exports = {
    publicPath: './',
    configureWebpack: {
          resolve: {
            extensions: [".ts", ".tsx", ".js", ".json"],
             alias: {}
          },
          module: {
            rules: [
              {
                  test: /\.tsx?$/,
                  use: [
                      {
                          loader: 'thread-loader'
                      },
                      {
                          loader: 'ts-loader',
                          options: {
                              happyPackMode: true,
                              appendTsSuffixTo: [/\.vue$/],
                          }
                      }
                  ],
              }
            ]
          }
    }
}
