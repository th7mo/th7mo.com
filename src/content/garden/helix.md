---
title: "Helix"
description: "A post-modern modal text editor"
---

Helix is a text editor run in the terminal.
Helix also works in the [TTY](tty).
It is inspired by [Vim](https://www.vim.org/) and [Kakoune](https://kakoune.org/) text editors.
Helix can be run from the terminal with the `hx` command.
Some Linux distributions provide Helix with the `helix` command.
Helix's editing philosophy is `selection -> action`,
prioritizing safety through text verification before specifying an action.

## Installation
Because Helix is evolving quickly, the way Helix can be installed can evolve to.
To see how to install Helix, follow the [Helix documentation book article about installing Helix](https://docs.helix-editor.com/install.html).
Helix can be installed on Debian using an AppImage[^1]: Download the official Helix AppImage from the [Helix latest releases page](https://github.com/helix-editor/helix/releases).

[^1]: At the time of writing (January 2024) this is the best way to install Helix on Debian.

First move the Helix image to a directory inside the [`$PATH` variable](path-variable):

```sh
mv helix-*.AppImage /usr/local/bin/hx
```

Then change it to an executable using `chmod`:

```sh
chmod +x hx
```

Helix can be run now with the `hx` command
if the `/usr/local/bin/` directory is listed in the `$PATH` variable.

## Configuration
Helix relies on two primary configuration files:
`~/.config/helix/config.toml` for general configuration and
`~/.config/helix/languages.toml` for Language Server Protocol (LSP) configuration.
Due to built-in functionality, Helix's configurations files are typically short.
It is recommended to include the `config.toml` and `languages.toml` files inside the [dotfiles](dotfiles) repository for backup.

## Clipboard
Helix can copy to the system clipboard with <kbd>Space</kbd> + <kbd>y</kbd>.
When running Wayland as a desktop environment, make sure to install `wl-clipboard`:

```sh
sudo apt install wl-clipboard
```

Otherwise, when running Xorg as a desktop environment, make sure to install `xclip`:

```sh
sudo apt install xclip
```
