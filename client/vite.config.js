import { defineConfig } from 'vite'
import solidPlugin from 'vite-plugin-solid'
import eslint from 'vite-plugin-eslint'
// import devtools from 'solid-devtools/vite';

export default defineConfig({
  plugins: [
    solidPlugin(),
    eslint(),
    'solid'
  ],
  extends: [
    'eslint:recommended',
    'plugin:solid/recommended'
  ],
  server: {
    port: 3000,
    watch: {
      usePolling: true
    },
    proxy: {
      '/api': {
        target: 'http://app:8000',
        changeOrigin: true,
        secure: false,
        configure: (proxy, _options) => {
          proxy.on('error', (err, _req, _res) => {
            console.log('proxy error', err)
          })
          proxy.on('proxyReq', (proxyReq, req, _res) => {
            console.log('Sending Request to the Target:', req.method, req.url)
          })
          proxy.on('proxyRes', (proxyRes, req, _res) => {
            console.log('Received Response from the Target:', proxyRes.statusCode, req.url)
          })
        }
      }
    }
  },
  build: {
    target: 'esnext'
  }
})
