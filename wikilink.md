A wikilink is a link from one [[note]]s to another note in the [[second-brain]].

## Link a note
A wikilink is the name of a note surrounded by double square brackets:
```md
[[{name-of-note}]]
```

The following code example shows the syntax for a wikilink to this note:
```md
[[wikilink]]
```

## Amount of wikilinks
A wikilink for a note should appear only *once* in a note.
This makes the note more readable and makes for a simple rule to follow.
It is also easier to write, especially in technical notes where there can be a lot of mentions of another note.

## Cross-referencing a heading of another note
For a more specific cross-reference it is also possible to cross-reference a heading of other note.
This can be achieved by specifying the heading in `lower-kabab-case` after a `#` as divider:
```
[cross-reference#Cross-referencing a heading of another note](cross-reference.md#cross-referencing-a-heading-of-another-note)
```

The `#` hash character can indicate that only a section of a note is cross-referenced.

> [!IMPORTANT]
> Do not cross-reference another note inside a heading.
> This is because a cross-reference to a heading that also has a cross-reference, is ugly to handle.

## Compatibility
Wikilinks are not part of the official Markdown specification.
They are also not well-supported either.
However, they are more pleasant to read, write and maintain.
There are tools available that convert wikilinks into regular Markdown links.
If the Second Brain ever gets ported into a different (file) format, a converter can be used to preserve compatibility.

## Linking future notes
It is recommended to create cross-references to a note that does not exist yet but is desired to be written in the future.
[[obsidian]] and other applications warn about the non-existing note.
With this workflow it is easy to fill in gaps of knowledge.

## See also
* The [Wikipedia article](https://en.wikipedia.org/wiki/Help:Link) about wikilinks.
