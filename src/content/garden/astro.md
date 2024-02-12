---
title: "Astro"
description: "The web framework for content-driven websites"
---

[Astro](https://astro.build/) is a frontend framework for content-first websites.
Astro primarily supports content written in [Markdown](markdown).
Additional metadata about the content is provided in [YAML front matter](yaml-front-matter).
The philosophy of Astro is opting in for JavaScript only when it's needed.
Astro in its most simple setup is a static site generator
like [Hugo](https://gohugo.io/) or [Jekyll](https://jekyllrb.com/).
Without JavaScript, Astro only generates HTML and [CSS](css) files.

The most popular method to create a new Astro project is with `npm`:

```sh
npm create astro@latest
```

The `scripts` object inside the `package.json` file created will show
how to run, build and preview the Astro project.

## Project structure

* `src/pages/`: All the pages of the website.
* `src/layouts/`: Astro templates.
* `src/content/`: Content collections.
* `public/`: Non-code, like favicons, images and fonts.
* `dist/`: Generated Astro build.

## Footnotes

Astro is a tool that converts [footnotes](footnote) written
in Markdown to HTML links displayed in a superscript.
The references for footnotes have classes
that are prefixed with `user-content-fn`.
The actual footnotes are listed at the end of the generated HTML file
under a level 2 [heading](markdown#headings) named 'Footnotes'.
The classes for the footnotes themselves are prefixed with `user-content-fnref`.
This allows for custom styling of both the footnotes and their references:

```css
/* Only target footnote reference */
a[href^="user-content-fn"]:not(href^="user-content-fnref") {
    text-decoration: underline;
}
```
