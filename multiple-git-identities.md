When having a personal and a work identity for [[Git]] there is additional configuration to be made to let Git know when to use which identity.

Sometimes it is desired to overwrite the default [[gitconfig]].
This can be done by creating another `.gitconfig` configuration anywhere, and referencing it in the main `~/.gitconfig` file.
The `includeIf` option is used to specify when a `.gitconfig` file should be active.

For example:
```ini
[includeIf "gitdir:~/work/"]
    path = ~/.gitconfig-work
```

This specifies that the `~/.gitconfig-work` configuration should be used when inside a Git repository anywhere under the `~/work/` directory.
The `~/.gitconfig-work` can have a different name (`user.name`) and [[email-address]] (`user.email`) for professional use.

> [!IMPORTANT]
> Always end the `gitdir` path with a forward slash `/`, otherwise it will not work!
