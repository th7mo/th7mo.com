---
title: "Paragraph width"
---

Paragraphs should be in between 50 and 70 characters of
width[^1]. Too wide paragraphs will make readers take big horizontal movements when starting a new line.
Too small paragraphs break the flow of reading text.

[^1]: According to the [Baymard Institute](https://baymard.com/blog/line-length-readability).

When working with [CSS](css) the `ch` unit can be used to achieve
this. The `ch` unit [is defined as](https://drafts.csswg.org/css-values-3/#ch):

> Equal to the used advance measure of the "0" (ZERO, U+0030) glyph
> found in the font used to render it.

In proportional typefaces, `1ch` is *usually* around 20% wider than the
average character width[^2].
This makes the optimal paragraph width in CSS between `40ch` and `60ch`:

```css
p {
    max-width: 55ch;
}
```

[^2]: Discovered by Eric A. Meyer explained in [this blog post](https://meyerweb.com/eric/thoughts/2018/06/28/what-is-the-css-ch-unit/) about the `ch` unit.

Fonts can however vary in their width. It is best to experiment width
the value on each font to determine the best paragraph width.
