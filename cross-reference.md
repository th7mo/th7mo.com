# Cross-reference

A cross-reference is a connection to another part of the content.
For the Second Brain a cross-reference is a connection between two [note](note.md)s.

## Cross-referencing other [note](/note.md)s

When cross-referencing another [note](note.md), make the link name the same as the file name of the [note](note.md).
For this [note](note.md) the file name is "cross-reference".
The link itself should start with a slash (`/`).
Because the [file-structure](/file-structure.md) has no directories, every [note](note.md) can be cross-referenced using the syntax `/{file-name}`.

For example: the complete syntax to cross-reference this [note](note.md) from anywhere in the Second Brain is: 

```
[cross-reference](/cross-reference.md)
```

Prefer internal cross-references above external cross-references:

```
cross-reference.md -> another-note.md -> https://git-scm.com/
```

is better than:

```
cross-reference.md -> https://git-scm.com/ <- another-note.md
```

### Cross-referencing sections of other [note](/note.md)s

For a more specific cross-reference it is also possible to cross-reference headings of other [note](note.md)s.
This can be achieved by specifying the heading in `lower-kabab-case` after a `#` as divider:

```
[cross-reference#Cross-referencing sections of other notes](/cross-reference.md#cross-referencing-sections-of-other-notes)
```

The `#` hash character can indicate that only a section of a [note](note.md) is cross-referenced.

