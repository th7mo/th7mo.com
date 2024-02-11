---
title: "Color scheme"
description: "Managing color schemes on websites using HTML and CSS"
---

A website's color scheme should be stored in [CSS](css) custom properties.
It is common practice to define custom properties on the `:root` pseudo-class for global reference:

```css
:root {
    --color-text: black;
    --color-base: white;
}
```

With the code example above the text color can be referenced with `var(--color-text)`:

```css
p {
    color: var(--color-text);
}
```

## Browser preference
Users express their color scheme preference by choosing light mode, dark mode or 'automatic' in their browser settings.
The `prefers-color-scheme` CSS media feature indicates if a light or dark color scheme is preferred[^1].
This media feature can be utilized to overwrite the CSS variables that store the color scheme colors[^2]:

```css
@media (prefers-color-scheme:dark) {
    :root {
        --color-text: white;
        --color-base: black;
    }
}
```

[^1]: This is explained in detail in the [MDN web docs](https://developer.mozilla.org/en-US/docs/Web/CSS/@media/prefers-color-scheme).
[^2]: This also works inside `<style>` tags in SVG files.

The same media feature can be used inside HTML to show a light or dark favicon based on the color scheme preference:

```html
<head>
    <link
        rel="icon"
        href="favicon-black.ico"
        media="(prefers-color-scheme: light)"
    >
    <link
        rel="icon"
        href="favicon-white.ico"
        media="(prefers-color-scheme: dark)"
    >
</head>
```

The black favicon is shown when light mode is selected, while the white favicon shown in dark mode.
