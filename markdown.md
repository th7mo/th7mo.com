# Markdown

[Markdown](https://en.wikipedia.org/wiki/Markdown) is a markup language that uses plain text formatting to write structured documents.
The Second Brain uses the Markdown format to store information.
Markdown provides universal readability, regardless of platform or tooling.
Markdown also allows easy history tracking, collaboration and reviewing with [Git](/git.md). 
The Markdown format is very established and is not sensitive to major technological shifts.

## Syntax

Markdown uses reserved symbols to mark the styling of text or give text a special meaning.

### Headings

Markdown allows a non-hierarchical structure of headings.
It is recommended however to start with a level one heading and only increment it with a maximum of one when introducing a subheading.

Prefix a line with a hash (`#`) symbol to make it a heading.
The amount of hashes defines the heading level.
It is recommended to not number the headings manually.

```md
# Heading 1
## Heading 2
### Heading 3
#### Heading 4
##### Heading 5
###### Heading 6
```

### Text styling

- Surround text with single asterisks (`*`) to make text *italic*.
- Surround text with double asterisks (`**`) to make text **bold**.
- Surround text with double tildes (`~~`) to make text ~~strike-through~~.


### Lists

Prefix a list item with the hyphen (`-`) symbol to make it an unordered list:

```md
- item
- item
- item
```

Alternatively an unordered list can be made using the asterisk (`*`) or addition (`+`) symbol.
Using an asterisk (`*`) symbol is discouraged because it is also used for making text bold or italic.
Using an addition (`+`) symbol is discouraged because it takes more effort to type on most keyboards and is not an intuitive symbol.

Prefix a list item with a number and a period (`.`) to make it an ordered list:

```md
1. item
2. item
3. item
```

### Inline code and code blocks

To create a code block surround the code between two lines with three backticks.
Add a programming language name after the first three backticks to indicate to Markdown processors that syntax highlighting can be used for that code block:

```
{three backticks}{language}
{place code here}
{three backticks}
```

## See also

- The original [Markdown specification](https://daringfireball.net/projects/markdown/).
- The new and improved [CommonMark specification](https://commonmark.org/).
- Good recommendations for Markdown [Basic Syntax](https://www.markdownguide.org/basic-syntax/).
