# SSH-keygen

`ssh-keygen` is a command line tool to create a [ssh-key](ssh-key.md).

## Generating an SSH key

To generate a new [ssh-key](ssh-key.md) execute the following command:

```sh
ssh-keygen -t ed25519 -C {email}
```

- The `-t` flag specifies the type of key to generate.
- The `-C` adds a comment to the [ssh-hey](ssh-key.md).
	This flag is optional but recommended.
	It will help indicate which [ssh-key](ssh-key.md) belongs to which [email](email.md) when using `ssh-add -l` or `keychain -l`.
- `{email}` is the [email](email.md) used for GitHub or Bitbucket, configured as `user.email` in the [gitconfig](gitconfig.md) file.

When prompted to "Enter a file in which to save the key", press `Enter` to accept the default file location.
Otherwise, if a different name or located is desired, a custom name can be specified.
The default directory for storing [ssh-key](ssh-key.md)s is `~/.ssh/`.
Press `Enter` until the file is generated.
