> [!NOTE]
> The implementation for symlinking might be replaced with a `--bare` repository using a worktree implementation.

Dotfiles are files prefixed with a dot `.`
They are hidden by default on [[Unix]]-based operating systems.
Dotfiles are used to store any type of configuration.
All the user configuration is stored in dotfiles somewhere in or under the [[home-directory]].
A lot of applications store their configuration in the `~/.config/` directory.

It is a good idea to make a [[git]] repository to back up and sync configuration files.
However, it is not recommended to make a Git repository directly in the home directory.
A better alternative is to move configuration files into a separate directory like `~/dotfiles/`.
To make sure installed programs still have access to the configuration files, [[symbolic-link]]s to the configuration files can be made using [[stow]].

## History
That dotfiles are invisible was an accident[^1].
A long time ago Unix decided to hide the `.` and `..` directories because they exist in every directory.
They only checked if the first character of the file (or directory) name was a dot `.`, and hid it when that was the case:
```go
if fileName[0] == "." {
    hideFile()
}
```

This results in every file starting with a dot to be hidden.
What they should have done is check if the filename is `.` or `..` instead:
```go
if fileName == "." || fileName == ".." {
    hideFile()
}
```

[^1]: http://xahlee.info/UnixResource_dir/writ/unix_origin_of_dot_filename.html
