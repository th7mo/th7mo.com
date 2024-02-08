---
title: "Footnote"
---

Footnotes references can be created in HTML with the `<a>` tag wrapped inside the `<sup>` tag to make the footnote float at the top of the line:

```html
<p>
    some words<sup><a href="#footnote-1">[1]</a></sup>  
</p>
```

In this example the footnote with id `footnote-1` is linked.
It is recommended to come up with a standard prefix for footnotes to distinct footnotes from regular links in CSS:

```css
a[href^="footnote-"] {
    /* styling for footnotes only */
}
```

Below the main content the footnotes can be listed in an ordered list:

```html
    <!-- main content -->
    <hr>
    <ol>
        <li id="footnote-1">A footnote explanation.</li>
        <li id="footnote-2">Another footnote explanation.</li>
    </ol>
```
