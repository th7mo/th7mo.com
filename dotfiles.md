# Dotfiles

> [!NOTE]
> This might be replaced with a Git `--bare` repository using a worktree implementation.

Dotfiles are files prefixed with a dot (`.`).
They are hidden by default on UNIX-based operating systems.
Dotfiles are used to store any type of configuration.
All the user configuration is stored in dotfiles somewhere in or under the `~/` home directory.
A lot of applications store their configuration in the `~/.config/` directory.

It is a good idea to make a [git](/git.md) repository to back up and sync configuration files.
However, it is not recommended to make a [git](/git.md) repository directly into the `~/` home directory.
A better alternative is to move configuration files into a separate directory like `~/dotfiles/`.
To make sure installed programs still have access to the configuration files, symbolic links to the configuration files can be made using [stow](/stow.md).

## History

That dotfiles are invisible was an accident, because a long time ago UNIX decided to hide the `.` and `..` directories because they exist in every directory.
They only checked if the first character of the file (or directory) name was a dot (`.`), and hid it when that was the case. 
