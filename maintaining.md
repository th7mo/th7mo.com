This [[note]] provides rules to make the writing of the entire [[second-brain]] consistent.

<!-- TODO: Complete rewrite using Wikipedia, ArchWiki and Gentoo wiki as inspiration. -->

## Formatting
* Use the rules described in the [[Markdown]] note.
* Every Markdown file name must be written in `lowercase-kabab-case`.
* Try to make the name of the note singular (for example: note instead of notes).
* Capitalize the first character of the note title.
* Only one heading 1 is allowed in the entire note.
* When writing commands using the triple backtick syntax, never prefix the command with `$` or `#`.
  Write the command without any prefixes, so it can easily be copied to the terminal.
* Use the `<!-- TODO -->` mark before a `> [!NOTE]` GitHub Flavored Markdown syntax for explaining that a section is work in progress.
  It will render the blockquote as a proper note, without the HTML comment before it.
  The prefixing comment can be used to quickly search all things that need revision.
* Use the work in progress indicator as close to the source as possible, so the reader knows that a part of the note is not finished or needs adjustment.

## Language
Try to avoid first and second person writing.
This Second Brain should be mostly informative.
This improves the ability to make more notes publicly available.
It is allowed to use first or second person when writing personal notes.

Use English (not British) for writing notes.
The only exception are words that do not have a proper English equivalent.

## Terminology
* **Note**:
  Every Markdown file in the Second Brain is a note.
  Use the term 'note' instead of page, article or document.
* **Execute**:
  When writing commands with the triple backtick syntax, use the term execute instead of run.
  For example: "Execute the following command:" instead of "Run the following command:".
* **Directory**:
  Use the term directory instead of folder.

## Special blockquotes
> [!NOTE]  
> Highlights information that users should take into account, even when skimming.

> [!TIP]
> Optional information to help a user be more successful.

> [!IMPORTANT]  
> Crucial information necessary for users to succeed.

> [!WARNING]  
> Critical content demanding immediate user attention due to potential risks.

> [!CAUTION]
> Negative potential consequences of an action.

## See also
* The Second Brain setup of David Luhr.
  He also uses [[Git]] and Markdown to store his Second Brain.
  He explains his setup on his [personal website](https://luhr.co/blog/2023/04/21/my-custom-second-brain-setup-part-2-how-it-works/).
* There is also a [GitHub repository](https://github.com/KasperZutterman/Second-Brain) that contains a list of people their public Second Brain.
