A [markdown](markdown.md) file can have YAML front matter.
YAML front matter is indicated with the following syntax:
```md
---
{YAML-here}
---
```

In the YAML front matter metadata about the document can be stored.
Information about the author, title, language and date of the document are the most used pieces of metadata:
```md
---
author: "Thimo van Velzen"
title: "Intership Report Pre-final"
lang: NL
---

# Introduction
```

The `lang: NL` field is used by the `ltex-ls` language server to check the language of the current buffer.
