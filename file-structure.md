The [second-brain](second-brain.md) has the most simple file structure possible.
It does not contain any directories and only uses [markdown](markdown.md) as a file format.

Every [note](note.md) is located at the root of the repository.
The filenames of [note](note.md)s have a naming convention mentioned in [maintaining#Formatting](maintaining.md#formatting).

## Why
The [note](note.md)s in the [second-brain](second-brain.md) have a complex relation to each other.
There is not a simple file structure solution because the [note](note.md)s are very different in nature.
It would be very difficult to create a file structure that does not require a lot of maintenance.

Although it is not intuitive, it is fine to have a thousand [note](note.md) in a single directory.
Also, there is no real benefit to categorizing the [note](note.md)s.
It would result in more overhead and complexity.
Moving [note](note.md) to different directories would invalidate every [wikilink](wikilink.md) to that [note](note.md).

When searching for a [note](note.md), just search for the filename of that [note](note.md).
Because the [second-brain](second-brain.md) uses just plain text in all [note](note.md)s, searching for a word in all the [note](note.md)s is also a good option.

The benefit for not using directories is that it makes a [wikilink](wikilink.md) to any other [note](note.md) easy.
The path to another [note](note.md) is always `/{name-of-note}`.
