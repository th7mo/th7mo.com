When having a personal and a work identity for [git](git.md), there is additional configuration to be made to let [git](git.md) know when to use which identity.

Sometimes it is desired to overwrite the default [gitconfig](gitconfig.md).
This can be done by creating another [gitconfig](gitconfig.md) anywhere, and referencing it in the main [gitconfig](gitconfig.md) file.
The `includeIf` option is used to specify when a [gitconfig](gitconfig.md) file should be active.

For example:
```ini
[includeIf "gitdir:~/work/"]
    path = ~/.gitconfig-work
```

This specifies that the `~/.gitconfig-work` configuration should be used when inside a [git](git.md) repository anywhere under the `~/work/` directory.
The `~/.gitconfig-work` can have a different `user.name` and `user.email` for professional use.

> [!IMPORTANT]
> Always end the `gitdir` path with a forward slash `/`, otherwise it will not work!
