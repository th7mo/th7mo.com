---
title: "Wikilink"
description: "A bidirectional link between two resources"
---

A wikilink is a [bidirectional link](https://maggieappleton.com/bidirectionals)
from one article to another article.
Wikilinks are the default method to link two resources (articles)
 in [Obsidian](obsidian) and [Wikipedia](https://www.wikipedia.org/).

## Link to an article
A wikilink is created by enclosing the name of an article by double
square brackets:

```md
[[{name-of-article}]]
```

Only the first character of the wikilink name is *case-insensitive*.
This allows for proper capitalizing the first character of a sentence.
It also allows to properly capitalize wikilinks when the link text
represents something that should be capitalized by default.

## Link to an article heading
To link to a heading in another article, append a `#` and the heading
name to the article name:

```md
[[{name-of-article}#{name-of-heading}]]
```

The name of the heading is *case-sensitive*.

To link to a heading in the same article, emit the name of the article:

```md
[[#Link to a article heading]]
```

## Amount of wikilinks
A wikilink for a article should appear only *once* in a article. It is
allowed however to have multiple wikilinks to the same article if the
wikilinks link to different headings. This makes the article more
readable and makes for a simple rule to follow. It is also easier to
write, especially in technical articles where there can be a lot of
mentions of other articles.

A wikilink should be placed at the first mention of the concept it will
link to.

## Compatibility
Wikilinks are not part of the official [Markdown](markdown)
specification. They are also not well-supported either. However, they
are more pleasant to read, write and maintain than regular Markdown
links. There are tools available that convert wikilinks into regular
Markdown links.

## See also
-   The [Obsidian
    help](https://help.obsidian.md/Linking+notes+and+files/Internal+links)
    about internal links using wikilinks.
-   The [Wikipedia article](https://en.wikipedia.org/wiki/Help:Link)
    about wikilinks.
