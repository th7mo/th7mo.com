A cross-reference is a connection to another part of the content.
For the [second-brain](second-brain.md) a cross-reference is a connection between two [note](note.md)s.

# Cross-referencing another note

When cross-referencing another [note](note.md), make the link name the same as the file name of the [note](note.md).
For this [note](note.md) the file name is "cross-reference".
Because the [file-structure](file-structure.md) has no directories, every [note](note.md) can be cross-referenced using the syntax `{file-name}`.

For example: the complete syntax to cross-reference this [note](note.md) from anywhere in the [second-brain](second-brain.md) is: 

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

# Cross-referencing a heading of another note

For a more specific cross-reference it is also possible to cross-reference a heading of other [note](note.md).
This can be achieved by specifying the heading in `lower-kabab-case` after a `#` as divider:

```
[cross-reference#Cross-referencing a heading of another note](/cross-reference.md#cross-referencing-a-heading-of-another-note)
```

The `#` hash character can indicate that only a section of a [note](note.md) is cross-referenced.

> [!IMPORTANT]
> Do not cross-reference another [note](note.md) inside a heading.
> This is because a cross-reference to a heading that also has a cross-reference, is ugly to handle.
