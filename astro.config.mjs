import { defineConfig } from "astro/config";

// https://astro.build/config
export default defineConfig({
  site: "https://th7mo.com",
  build: {
    format: "file"
  },
  markdown: {
    syntaxHighlight: false
  }
});
