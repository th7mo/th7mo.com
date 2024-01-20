# Multiple Git identities

When having a personal and a work identity for [git](/git.md), there is additional configuration to be made to let [git](/git.md) know when to use which identity.

Sometimes it is desired to overwrite the default [gitconfig](/gitconfig.md).
This can be done by creating another [gitconfig](/gitconfig.md) anywhere, and referencing it in the main [gitconfig](/gitconfig.md) file.
The `includeIf` option is used to specify when a [gitconfig](/gitconfig.md) file should be active.

For example:

```sh
[includeIf "gitdir:~/work/"]
    path = ~/work/.gitconfig
```

This specifies that the `~/work/.gitconfig` configuration should be used when inside a [git](/git.md) repository anywhere under the `~/work/` directory.
The `~/work/.gitconfig` can have a different `user.name` and `user.email` for professional use.

> [!IMPORTANT]
> Always end the `gitdir` path with a slash (`/`), otherwise it will not work!
