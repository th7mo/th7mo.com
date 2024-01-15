# File structure

The Second Brain has the most simple file structure possible.
It does not contain any directories and only uses plain text files (using the [Markdown workflow](/markdown-workflow.md)).

Every note is located at the root of the repository.
The filenames of notes have a naming convention mentioned in the [Style guide](/style-guide.md).

## Why

The notes in the Second Brain have a complex relation to each other.
There is not a simple file structure solution because the notes are very different in nature.
It would be very difficult to create a file structure that does not require a lot of maintenance.

Although it is not intuitive, it is fine to have a thousand notes in a single directory just like it is fine to have a thousand words in a single note.
Also, there is no real benefit to categorizing the notes.
It would result in more overhead and complexity.
Moving notes to different directories would invalidate every cross-reference to that note.

When searching for a note, just search for the filename of that note.
Because the Second Brain uses just plain text in all notes, searching for a word in all the notes is also a good option.

The benefit for not using directories is that it makes referencing other notes easy.
The path to another note is always `/{name-of-note}`.
