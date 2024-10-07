import { defineConfig } from "astro/config";

// https://astro.build/config
export default defineConfig({
  site: "https://th7mo.com",
  trailingSlash: "always",
  markdown: {
    syntaxHighlight: false
  }
});
