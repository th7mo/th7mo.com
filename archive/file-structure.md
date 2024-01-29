The [[second-brain]] has the most simple file structure possible.
It does not contain any directories and only uses [[Markdown]] as a file format.

Every [[note]] is located at the root of the repository.
The filenames of notes have a naming convention mentioned in [[maintaining#Formatting]].

The notes in the Second Brain have a complex relation to each other.
There is not a simple file structure solution because the notes are very different in nature.
It would be very difficult to create a file structure that does not require a lot of maintenance.

Although it is not intuitive, it is fine to have a thousand notes in a single directory.
Also, there is no real benefit to categorizing the notes.
It would result in more overhead and complexity and could never properly represent the coherency of the notes.

When searching for a note, just search for the filename of that note.
Because the Second Brain uses just plain text in all notes, searching for a word in all the notes is also a good option.

The benefit for not using directories is that it makes a plain Markdown links to others note easy.
The relative link to another note is always `./{name-of-note}`.

> [!NOTE]
> Although the Second Brain uses [[wikilink]]s that do not require a path, it is still important have a flat file structure.
> This preserve compatibility when converting wikilinks to plain Markdown links.
