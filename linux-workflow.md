This [note](note.md) describes a solid Linux workflow stack, using reliable, stable and future-proof software.

## Philosophy
An operating system should serve the user, not the other way around.
The user should mostly use tools that are available to anyone.
A good operating system should be stable by nature, but should offer the option to use unstable software when needed.
Although it is really fun to optimize a system for hardware resources, the optimization is not noticeable on modern hardware.
A good operating system should have a solid foundation and be mature.

It is tempting to use the latest software, but it comes at the cost of trust.
Trust that your perfect system never fails.

## Operating system
The perfect operating system to fulfill these requirements is [Debian](https://www.debian.org).
Debian is one of the most stable distributions for Linux.
Stability is also the reason why Debian has for more forks than any other distribution.

## Desktop environment
A desktop environment (or window manager) is mainly personal preference.
However, it is recommended to pick a desktop environment that is popular and supported.
A solid desktop environment is [GNOME](https://www.gnome.org/).
It is also the default desktop environment for Debian.
It is popular, minimalistic (from a design perspective), very polished, and has little customization options (for the better).

## Text editor
The [Helix](https://helix-editor.com/) text editor is pretty new, but is one of the few text editors that just works out of the box.
File navigation, syntax highlighting and LSP support are built-in, and almost nobody their configuration file is longer than 30 lines.
It is used in combination with [tmux](tmux.md) for multiple terminal instances.

## Applications
Prefer applications available in the Debian repository over anything else.
When the Debian repository does not provide a specific application, a [flatpak](flatpak.md) packaged application is a good alternative.

## Shell
I use [z-shell](z-shell.md) as my shell.
It makes it easier than Bash to add custom plugins using Oh My Zsh.
