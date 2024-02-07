---
title: "YAML front matter"
---

A [Markdown](markdown) file can have YAML front matter. YAML front
matter is indicated with the following syntax:

```yaml
---
{YAML-here}
---
```

In the YAML front matter metadata about the document can be stored.
Information about the author, title, language and date of the document
are the most used pieces of metadata:

```yaml
---
author: "Thimo van Velzen"
title: "Intership Report Pre-final"
lang: NL
---

# Introduction
```

[Pandoc](pandoc) uses the YAML front matter to insert various
pieces of metadata into the generated document.

The `lang:` field is used by the `ltex-ls` language server to check the
language of the current buffer.
