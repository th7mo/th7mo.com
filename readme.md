This repository is the website infrastructure for [th7mo.com](https://th7mo.com).
The content is stored as a submodule at [github.com/th7mo/notes](https://github.com/th7mo/notes).
The website is built with HTML and CSS. Intentionally shipped without any
JavaScript. The website uses [Astro](https://astro.build/) as a static site
generator to add the boilerplate HTML to every content file.

## Astro Commands
All commands are run from the root of the project, from a terminal:

| Command                   | Action                                           |
| :------------------------ | :----------------------------------------------- |
| `npm install`             | Installs dependencies                            |
| `npm run dev`             | Starts local dev server at `localhost:4321`      |
| `npm run build`           | Build your production site to `./dist/`          |
| `npm run preview`         | Preview your build locally, before deploying     |
| `npm run update` (custom) | Rebase changes in content repo and push commit   |
| `npm run astro ...`       | Run CLI commands like `astro add`, `astro check` |
| `npm run astro -- --help` | Get help using the Astro CLI                     |
