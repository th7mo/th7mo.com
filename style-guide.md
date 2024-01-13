# Style guide

This note provides rules to make the writing of the entire Second Brain consistent.

## Terminology

**Second Brain**:
The Second Brain is a digital collection of all thoughts, documentation and ideas.
It is stored in the form of a repository, but that is only an implementation detail.
Use the term Second Brain instead of repository.

**Note**:
Every [Markdown](/markdown-workflow.md) file in the Second Brain is a note.
A note contains information about a single topic.
Use the term note instead of page, article or document.

**Execute**:
When writing commands with the triple backtick syntax, use the term execute instead of run.
For example: "Execute the following command:" instead of "Run the following command:".

**Directory**:
Use the term directory instead of folder.

## Formatting

Use the rules described in the [Markdown Workflow](/markdown-workflow.md).

Every Markdown file name must be written in `lowercase-kabab-case`.

Start the document with a heading 1 (`#`) with the name of the Markdown file.
Use spaces instead of slashes and capitalize the first character.
Only one heading 1 is allowed in the entire Markdown file.

When writing commands using the triple backtick syntax, never prefix the command with `$` or `#`.
Write the command without any prefixes, so it can easily be copied to the terminal.

Use the `> [!NOTE]` GitHub syntax for explaining that a section is work in progress.
Use the work in progress indicator as close to the source as possible, so the reader knows that a part of the note is not finished or needs adjustments.
If this rule is followed than there should not be anything work in progress related in the [To-Do](/to-do.md) list.

When referencing another note, make the link name the same as the title of the note.
For this note the title is "Style guide".
The link itself should start with a slash (`/`).
Because there are no directories inside the Second Brain (the [File structure](/file-structure.md) note explains why), every note can be referenced using the syntax `/{file-name}`.
For example: the complete syntax to reference this note from anywhere in the Second Brain is: `[Style guide](/style-guide.md)`.

## Language

Try to avoid first and second person writing.
This Second Brain should be mostly informative.
This improves the ability to make more notes publicly available.
It is allowed to use first or second person when writing personal notes.

Use English for writing notes.
The only exception is naming products that are not known in English.
