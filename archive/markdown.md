[Markdown](https://en.wikipedia.org/wiki/Markdown) is a markup language that uses plain text formatting to write structured documents.
The [[second-brain]] uses the Markdown format to store information.
Markdown provides universal readability, regardless of platform or tooling.
Markdown also allows easy history tracking, collaboration and reviewing with [[Git]]. 
The Markdown format is very established and is not sensitive to major technological shifts.

## Syntax
Markdown uses reserved symbols to mark the styling of text or give text a special meaning.

### Headings
Markdown allows a non-hierarchical structure of headings.
It is recommended however to start with a level one heading and only increment it with a maximum of one when introducing a subheading.

Prefix a line with a hash `#` symbol to make it a heading.
The amount of hashes defines the heading level.
It is recommended to not number the headings manually:
```md
# Heading 1
## Heading 2
### Heading 3
#### Heading 4
##### Heading 5
###### Heading 6
```

It is good practice (but not required) to have an empty line before a heading.

### Text styling
* Surround text with single asterisks `*` to make text *italic*.
* Surround text with double asterisks `**` to make text **bold**.
* Surround text with double tildes `~~` to make text ~~strike-through~~.

### Lists
Prefix a list item with the asterisk `*` symbol to make it an unordered list:
```md
* item
* item
* item
```

Alternatively an unordered list can be made using the hyphen `-` or addition `+` symbol.
The asterisk `*` symbol is preferred because it looks the most like a real bullet point.
Using an addition `+` symbol is discouraged because it takes more effort to type on most keyboards and is not an intuitive symbol.

Prefix a list item with a number and a period `.` to make it an ordered list:
```md
1. item
2. item
3. item
```

It is also possible to create a numbered list with only the prefix `1.`:
```md
1. item
1. item
1. item
```

Most modern renderers will correctly number the list.
It is not recommended tho, because it makes it harder to read when editing the Markdown.

### Inline code and code blocks
To create a code block surround the code between two lines with three backticks.
Add a programming language name after the first three backticks to indicate to Markdown processors that syntax highlighting can be used for that code block:
```
{three backticks}{language}
{place code here}
{three backticks}
```

## YAML front matter
For the Second Brain it is allowed to use [[yaml-front-matter]].

## Exporting Markdown
Sometimes it is required to convert Markdown to other file formats like `.pdf` or `.docx`.
The tool [[Pandoc]] can be used to convert Markdown to a lot of other file formats.

## See also
* The original [Markdown specification](https://daringfireball.net/projects/markdown/).
* The new and improved [CommonMark specification](https://commonmark.org/).
* Good recommendations for Markdown [Basic Syntax](https://www.markdownguide.org/basic-syntax/).
