import { defineConfig } from "astro/config";

import svelte from "@astrojs/svelte";

// https://astro.build/config
export default defineConfig({
  site: "https://th7mo.com",

  build: {
    format: "file"
  },

  markdown: {
    shikiConfig: {
      themes: {
        light: "github-light",
        dark: "github-dark"
      },
      transformers: [{
        preprocess(code) {
          if (code.endsWith('\n')) {
            code = code.slice(0, -1);
          }
          return code;
        },
      }]
    }
  },

  integrations: [svelte()]
});
