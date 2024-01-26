A wikilink is a link from one [[note]] to another note in the [[second-brain]].

## Syntax
When cross-referencing another note, make the link name the same as the file name of the note.
For this note the file name is "cross-reference".
Because the [[file-structure]] has no directories, every note can be cross-referenced just by referencing the file name.

For example: the complete syntax to cross-reference this note from anywhere in the Second Brain is: 
```
[cross-reference](cross-reference.md)
```

To check if cross-references are formatted properly, use the following [[regular-expression]]:
```regex
\(/
```

If this search shows results, there are cross-references using an undesired forward slash.

## Duplicate cross-referencing notes
A cross-reference should appear only *once* in a note.

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

## Cross-referencing a future note
It is recommended to create cross-references to a note that does not exist yet but is desired to be written in the future.
[[obsidian]] and other applications warn about the non-existing note.
With this workflow it is easy to fill in gaps of knowledge.
