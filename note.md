The [second-brain](second-brain.md) contains notes.
A note represents a *single* concept.
They should be as atomic as possible.
Notes can [wikilink](wikilink.md) each other.
Notes should be small, and split up if it has multiple responsibilities.

## Renaming
Sometimes it is necessary to rename a note, or a heading inside a note.

> [!WARNING]
> Never just rename a note or heading, because all [wikilink](wikilink.md)s that pointed to that note or heading will be broken.
>
> Instead, always refactor a heading or note name using [LSP](https://microsoft.github.io/language-server-protocol/)'s or indexed projects.
> This will make sure all [wikilink](wikilink.md)s to that note or heading will also be renamed appropriately.
