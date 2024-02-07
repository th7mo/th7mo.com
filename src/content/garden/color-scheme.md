---
title: "Color scheme"
---

A CSS color scheme for websites should be stored in CSS variables.
Usually the variables for the color scheme are stored in the `:root` pseudo-class:

```css
:root {
    --color-text: black;
    --color-base: white;
}
```

Now everywhere in the stylesheet the `--color-text` can be accessed with `var(--color-text)`:

```css
p {
    color: var(--color-text);
}
```

## Color scheme based on browser preference
Users express their color scheme preference by toggling light or dark mode (or 'automatic') in their browser settings.
The `prefers-color-scheme` CSS media feature detects if a light or dark color scheme is preferred[^1].
The media feature can be leveraged to overwrite the CSS variables that store the color scheme colors[^2]:

```css
@media (prefers-color-scheme:dark) {
    :root {
        --color-text: white;
        --color-base: black;
    }
}
```

[^1]: This is explained in detail at the [MDN web docs](https://developer.mozilla.org/en-US/docs/Web/CSS/@media/prefers-color-scheme).
[^2]: This also works inside `<style>` tags in SVG files.

The same media feature can be used inside HTML to load a dark or light favicon based on the color scheme preference:

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

The black favicon is used when light mode is preferred and the white favicon is used when dark mode is preferred.
