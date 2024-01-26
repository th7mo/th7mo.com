A cross-reference is a connection to another resource.
For the [second-brain](second-brain.md) a cross-reference is a reference to another [note](note.md).

## Cross-referencing another note
There are two popular methods to cross-reference a note.
The first one is the [[markdown]] syntax using `[{title}]({file-name})`.
The second one is the wiki links syntax using `[[{file-name}]]`.
Do not use wiki links.
They are more elegant, but not compatible with most applications because they are not an official Markdown feature.

When cross-referencing another note, make the link name the same as the file name of the note.
For this note the file name is "cross-reference".
Because the [file-structure](file-structure.md) has no directories, every note can be cross-referenced just by referencing the file name.

For example: the complete syntax to cross-reference this note from anywhere in the Second Brain is: 
```
[cross-reference](cross-reference.md)
```

To check if cross-references are formatted properly, use the following [regular-expression](regular-expression.md):
```regex
\(/
```

If this search shows results, there are cross-references using an undesired forward slash.

## Duplicate cross-referencing notes
A cross-reference should appear only *once* in a note.``

## Cross-referencing a heading of another note
For a more specific cross-reference it is also possible to cross-reference a heading of other note.
This can be achieved by specifying the heading in `lower-kabab-case` after a `#` as divider:
```
[cross-reference#Cross-referencing a heading of another note](cross-reference.md#cross-referencing-a-heading-of-another-note)
```

The `#` hash character can indicate that only a section of a note is cross-referenced.

> [!IMPORTANT]
> Do not cross-reference another note.md inside a heading.
> This is because a cross-reference to a heading that also has a cross-reference, is ugly to handle.

## Cross-referencing a future note
It is recommended to create cross-references to a note that does not exist yet but is desired to be written in the future.
[obsidian](obsidian.md) and other applications warn about the non-existing note.
With this workflow it is easy to fill in gaps of knowledge.
