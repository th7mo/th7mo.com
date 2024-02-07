---
title: "Symbolic link"
---

A symbolic link, often called a 'symlink', is a file that serves as a
reference/pointer to another file or directory.

With [Unix](unix)-based systems the `ln` (link) command with the
`-s` (soft) flag can be used to create symbolic link:

```sh
ln -s {target_path} {link_path}
```

-   `{target_path}` is the path to the file or directory to
    link to.
-   `{link_path}` the desired path for the (new) symbolic
    link.
