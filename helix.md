[Helix](https://helix-editor.com/) is a text editor.
It is inspired by [Vim](https://www.vim.org/) and [Kakoune](https://kakoune.org/).
Helix can be run from the terminal with the `hx` command.
Some Linux distributions provide Helix with the `helix` command.

Helix's editing philosophy is `selection -> action`, prioritizing safety through text verification before specifying an action.

## Installation
Because Helix is evolving quickly, the way Helix can be installed can evolve to.
To see how to install Helix, follow the [Helix documentation book article about installing Helix](https://docs.helix-editor.com/install.html).

At the time of writing (January 2024) Helix can be installed on Debian using an AppImage:
Download the official Helix AppImage from the [latest releases](https://github.com/helix-editor/helix/releases/) page.
```sh
mv helix-*.AppImage /usr/local/bin/hx       # rename to hx
chmod +x hx                                 # change permission for executable mode
```

Helix can be run now with the `hx` command (if `/usr/local/bin/` directory is `$PATH`).

## Configuration
Helix relies on two primary configuration files: `~/.config/helix/config.toml` for general settings and `~/.config/helix/languages.toml` for Language Server Protocol (LSP) configuration.
Due to built-in functionality, Helix configurations are typically compact.
It is recommended to include the `config.toml` and `languages.toml` files inside the [dotfiles](dotfiles.md) repository for backup.

## Clipboard
Helix can copy to the system clipboard with `Space+y`.

### Wayland
When running [Wayland](https://wayland.freedesktop.org/) as a desktop environment, make sure to install `wl-clipboard`:
```sh
sudo apt install wl-clipboard
```

### Xorg
When running [Xorg](https://www.x.org/wiki/) as a desktop environment, make sure to install `xclip`:
```sh
sudo apt install xclip
```
