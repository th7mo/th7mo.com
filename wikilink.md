A wikilink is a link from one [[note]] to another note in the [[second-brain]].

> [!IMPORTANT]
> [[Markdown#Headings]] are not allowed to have wikilinks.
> This is because handling wikilinks to headings that have wikilinks is difficult.

## Link to a note
A wikilink is created by enclosing the name of a note by double square brackets:
```md
[[{name-of-note}]]
```

The following code example shows the syntax for a wikilink to this note:
```md
[[wikilink]]
```

Only the first character of the wikilink name is *case-insensitive*.
This allows for proper capitalizing the first character of a sentence.

## Link to a note heading
To link to a heading in another note, append a `#` and the heading name to the note name:
```md
[[{name-of-note}#{name-of-heading}]]
```

The name of the heading is *case-sensitive*.

The following code example shows the syntax for a wikilink to this heading of this note:
```md
[[wikilink#Link to a note heading]]
```

To link to a heading in the same note, emit the name of the note:
```md
[[#Link to a note heading]]
```

## Amount of wikilinks
A wikilink for a note should appear only *once* in a note.
It is allowed however to have multiple wikilinks to the same note if the wikilinks link to different headings.
This makes the note more readable and makes for a simple rule to follow.
It is also easier to write, especially in technical notes where there can be a lot of mentions of other notes.

A wikilink should be placed at the first mention of the concept it will link to.
An except to this rule are markers that indicate that a note needs adjustments.
This is due their temporary nature.

## Compatibility
Wikilinks are not part of the official [[Markdown]] specification.
They are also not well-supported either.
However, they are more pleasant to read, write and maintain than regular Markdown links.
There are tools available that convert wikilinks into regular Markdown links.
If the Second Brain ever gets ported into a different (file) format, a converter can be used to preserve compatibility.

## Linking future notes
It is recommended to create wikilinks to notes that do not exist yet but are desired to have in the future.
[[Obsidian]] and other applications warn about the non-existing note.
Linking future notes make it easy to fill in gaps of knowledge over time.

## See also
* The [Obsidian help](https://help.obsidian.md/Linking+notes+and+files/Internal+links) about internal links using wikilinks.
* The [Wikipedia article](https://en.wikipedia.org/wiki/Help:Link) about wikilinks.
