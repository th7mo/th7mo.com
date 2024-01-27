`ssh-keygen` is a command line tool to create an [[ssh-key]].

## Generating an SSH key
To generate a new SSH key execute the following command:
```sh
ssh-keygen -t ed25519 -C {email}
```
* The `-t` flag specifies the type of key to generate.
* The `-C` flag adds a comment to the SSH key.
	This flag is optional but recommended.
	It will help indicate which SSH key belongs to which [[email-address]] when using `ssh-add -l` or [[Keychain#List loaded ssh keys]].
* `{email}` is the email address used for GitHub or Bitbucket, configured as `user.email` in the [[gitconfig]] file.

When prompted to "Enter a file in which to save the key", press `Enter` to accept the default file location.
Otherwise, if a different name or located is desired, a custom name can be specified.
The default directory for storing SSH keys is `~/.ssh/`.
Press `Enter` until the file is generated.
