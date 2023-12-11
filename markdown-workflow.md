# Markdown

Markdown is a markup language that uses plain text formatting to write structured documents.

## Syntax

Markdown uses reserved symbols to mark the styling of text or give text a special meaning.

### Headings

Markdown allows a non-hierarchical structure of headings.
It is recommended however to start with a level one heading and only increment it with a maximum of one when introducing a subheading.

Prefix a line with a hash (`#`) symbol to make it a heading.
The amount of hashes defines the heading level.

```md
# Heading 1
## Heading 2
### Heading 3
#### Heading 4
##### Heading 5
###### Heading 6
```


### Lists

Prefix a list item with the (`-`) hyphen symbol to make it an unordered list.

```md
- item
- item
- item
```

Alternatively an unordered list can be made using the asterisk (`*`) or addition (`+`) symbol.
Using an asterisk (`*`) symbol is discouraged because it is also used for making text bold or italic.
Using an addition (`+`) symbol is discouraged because it takes more effort to type on most keyboards and is not an intuitive symbol.


[WIP: split the rest of the article into separate subjects]

I strongly believe that simple is better.
That applies to technology as well.
To keep my document workflow simple I want to minimize dependencies my documents rely on.

## Requirements

- The workflow should ensure that my documents are robust and are not sensitive to technological shifts.
- The workflow should not lock in on proprietary programs.
- The workflow should suit all my document needs.
- The workflow should provide universal readability, regardless of platform or tooling.
- The workflow should allow history tracking.
- The workflow should allow collaboration and reviewing of the documents.
- The workflow should handle large documents
- The workflow should dictate that the source files are independent of any tooling being used.

## Decision

Use the [Markdown](https://en.wikipedia.org/wiki/Markdown) markup language for documents and [Git](https://git-scm.com/) for version control.
Markdown and Git are the only two dependencies that can not be swapped by other implementations.
Tools to convert Markdown to other formats can change over time.

I utilize [Pandoc](https://pandoc.org/) to convert these documents into `.pdf` format.
Notably, no adjustments are required in the Markdown file themselves.
A `.tex` template file and Pandoc serve solely to transform the content into the required `.pdf` format.
As a result, the source files remain free of dependencies.

### Markdown styling

Write strict Markdown without flavors or extensions like GitHub Flavored Markdown (GFM).
Only two exceptions are allowed:

- referencing headers with the non-official syntax: `[](#header-name)` (does not break compatibility in any way, it is still valid link syntax)
- plain tables using the `|` pipe syntax (the table convention is standardized enough to depend on)

Each line always contains one full sentence.
When the sentence extends the screen, the sentence should probably be rewritten because it is too long for the reader anyway.

The maximum amount of consecutive blank lines is one. 

Headings (`#` ... `#######`) are not numbered.
Converter tools like Pandoc handle the presentation of the document.

When making an unordered list the dash (`-`) symbol is used.
Prefer a dash over an asterisk (`*`) symbol because the asterisk is also used for **bold** and *italics* markup.
Prefer a dash over a plus (`+`) symbol because a plus is more difficult to write on most keyboards (`Shift` + `=`) than a dash.

The official Markdown specification I follow can be found at [CommonMark](https://commonmark.org/).
The original but incomplete Markdown specification can be found [here](https://daringfireball.net/projects/markdown/).

When I want to submit a single document, all the contents of that document lives in a single Markdown file.
Editors and tools can handle large Markdown files.
The benefits of splitting (each section) into different Markdown files do not outweigh the downsides of being dependent on a tool that merges those files back.

The naming convention for each Markdown file is `name-in-kabab-case.md`.
The most popular file extension name for Markdown is `.md`, so that is the one being used.

The indicator `[WIP]` is used to indicate that a section is not finished.
More information can be provided using a colon: `[WIP: add more documentation]`.
For comments the `[COMMENT: the comment itself]` indicator is used. 

I write the Markdown files using an editor that identifies spelling and grammar issues.

## Consequences

Most people want to receive a `.pdf` formatted document, so before submitting documents a conversion needs to be made.

Pandoc ignores all HTML inside Markdown documents when converting using LaTeX, so any HTML in Markdown files will be lost when generating a `.pdf` document.
This makes it difficult to rely on HTML tables instead of tables with the pipe syntax.
