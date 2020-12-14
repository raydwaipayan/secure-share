module.exports = {
  "transpileDependencies": [
    "vuetify"
  ],
  pages: {
    'index': {
      entry: './src/main.js',
      template: 'public/index.html',
      title: 'Home',
      chunks: ['chunk-vendors', 'chunk-common', 'index']
    }
  }
}