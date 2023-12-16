import { defineConfig } from 'vite';
import solidPlugin from 'vite-plugin-solid';

export default defineConfig({
  plugins: [solidPlugin()],
  server: {
    port: 3000,
    watch: {
      usePolling: true
    },
    proxy: {
      '/api': {
        target: 'http://back-end:5000',
        changeOrigin: true,
        secure: false,
      }
    },
  },
  build: {
    target: 'esnext',
  },
});
