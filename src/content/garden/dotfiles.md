---
title: "Dotfiles"
description: "The hidden configuration files in Linux"
---

Files prefixed with a dot `.` are called dotfiles.
They are hidden by default on [Unix](unix)-based operating systems.
Dotfiles are used to store any type of configuration.
All the user configuration is stored in dotfiles somewhere in or under the [home directory](home-directory).
A lot of applications store their configuration in the `~/.config/` directory.

It is a good idea to make a [Git](git) repository to back up and sync configuration files.
However, it is not recommended to make a Git repository directly in the home directory.
A better alternative is to move configuration files into a separate directory like `~/dotfiles/`. To make sure installed programs still have access to the configuration files,
[symbolic-links](symbolic-link) to the configuration files can be made using [Stow](stow).

## History
That dotfiles are invisible was an accident[^1].
A long time ago Unix decided to hide the `.` and `..` directories because they exist in every directory.
They only checked if the first character of a file or directory name was a dot `.`,
and hid it when that was the case:

```go
if fileName[0] == "." {
    hideFile()
}
```

[^1]: [Rob Pike on the Origin of Unix Dot File Names](http://xahlee.info/UnixResource_dir/writ/unix_origin_of_dot_filename.html)

This results in every file and directory starting with a dot to be hidden.
What they should have done is check if the filename is `.` or `..` instead:

```go
if fileName == "." || fileName == ".." {
    hideFile()
}
```
