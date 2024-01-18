# Style guide

This note provides rules to make the writing of the entire Second Brain consistent.

> [!NOTE]
> This note should be split or be renamed to 'maintainer guide' or something like that.
> Now it does more than only document the styling.

## Cross-referencing other notes

When cross-referencing another note, make the link name the same as the title of the note.
For this note the title is "Style guide".
The link itself should start with a slash (`/`).
Because the [File structure](/file-structure.md) has no directories, every note can be cross-referenced using the syntax `/{file-name}`.

For example: the complete syntax to cross-reference this note from anywhere in the Second Brain is: 

```
[Style guide](/style-guide.md)
```

Prefer internal cross-references above external cross-references:

```
style-guide.md -> git.md -> https://git-scm.com/
```

is better than:

```
style-guide.md -> https://git-scm.com/ <- git.md
```

### Cross-referencing sections of other notes

For a more specific cross-reference it is also possible to cross-reference headings of other notes.
This can be achieved by specifying the heading in `lower-kabab-case` after a `#` as divider:

```
[Style guide | Cross-referencing sections of other notes](/style-guide.md#cross-referencing-sections-of-other-notes)
```

The `|` pipe character can indicate that a section of a note is cross-referenced.

## Formatting

- Use the rules described in the [Markdown](/markdown.md) note.
- Every Markdown file name must be written in `lowercase-kabab-case`.
- Start the document with a heading 1 (`#`) with the name of the Markdown file.
- Capitalize the first character of the note title.
- Only one heading 1 is allowed in the entire note.
- When writing commands using the triple backtick syntax, never prefix the command with `$` or `#`.
  Write the command without any prefixes, so it can easily be copied to the terminal.
- Use the `> [!NOTE]` GitHub Flavored Markdown syntax for explaining that a section is work in progress.
- Use the work in progress indicator as close to the source as possible, so the reader knows that a part of the note is not finished or needs adjustment.

## Renaming

Sometimes it is necessary to rename a note, or a heading inside a note.

> [!WARNING]
> Never just rename a note or heading, because all cross-references that pointed to that note or heading will be broken.

Instead, always refactor a heading or note name using [LSP](https://microsoft.github.io/language-server-protocol/)'s or indexed projects.
This will make sure all cross-references to that note or heading will also be renamed appropriately.

## Language

Try to avoid first and second person writing.
This Second Brain should be mostly informative.
This improves the ability to make more notes publicly available.
It is allowed to use first or second person when writing personal notes.

Use English (not British) for writing notes.
The only exception are words that do not have a proper English equivalent.

## Terminology

**Second Brain**:
The Second Brain is a digital collection of all thoughts, documentation and ideas.
It is stored in the form of a repository, but that is only an implementation detail.
Use the term Second Brain instead of repository.
Second Brain is a name, and should be written with capital letters.

**Note**:
Every Markdown file in the Second Brain is a note.
A note contains information about a single topic.
Use the term note instead of page, article or document.

**Execute**:
When writing commands with the triple backtick syntax, use the term execute instead of run.
For example: "Execute the following command:" instead of "Run the following command:".

**Directory**:
Use the term directory instead of folder.

## See also

- The Second Brain setup of David Luhr.
  He also uses [Git](/git.md) and [Markdown](/markdown.md) to store his Second Brain.
  He explains his setup on his [personal website](https://luhr.co/blog/2023/04/21/my-custom-second-brain-setup-part-2-how-it-works/).
- There is also a [GitHub repository](https://github.com/KasperZutterman/Second-Brain) that contains a list of people their public Second Brains.
