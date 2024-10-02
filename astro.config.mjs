import { defineConfig } from "astro/config";

// https://astro.build/config
export default defineConfig({
  site: 'https://th7mo.com',
  markdown: {
    shikiConfig: {
      theme: 'css-variables',
      wrap: true
    }
  }
});
