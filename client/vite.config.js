import { defineConfig } from 'vite'
import solidPlugin from 'vite-plugin-solid'
// import devtools from 'solid-devtools/vite';

export default defineConfig({
  plugins: [
    solidPlugin()
  ],
  server: {
    port: 3000,
    proxy: {
      '/api': {
        // string shorthand: http://localhost:5173/foo -> http://localhost:4567/foo
        target: 'http://app:8000',

        // proxy debug event listeners
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
