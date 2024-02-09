---
title: "Zsh"
description: "An extended Bourne shell with many improvements"
---

The Z shell, also known as Zsh,
is a Unix shell that is similar to [Bash](bash) or [Fish](https://fishshell.com/).
It is built upon the Bourne Shell (sh) and Bash,
and provides additional features compared to those shells.
[Oh My Zsh](oh-my-zsh) is a commonly used framework that can be used with Zsh.

## Installation
Zsh can be installed with `apt`:

```sh
sudo apt install zsh
```

Set Zsh as the default shell with the `chsh` (change shell) command:

```sh
chsh -s $(which zsh)
```

`$(which zsh)` evaluates to the path where Zsh is located.
On most systems this path will be `/usr/bin/zsh`.

## Configuration
Z shell is configured via the `.zshrc` file. It is recommended to put
the `.zshrc` file inside the [dotfiles](dotfiles) repository to keep
configuration backed up.
