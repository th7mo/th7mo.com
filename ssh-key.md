# SSH key

> [!Note]
> This note is not finished yet.
> - this article has not been reviewed yet
> - add Gentoo package names too (and decide where)

A method to access and write data to GitHub, Gitlab and Bitbucket is the Secure Shell Protocol (SSH).
You can authenticate to those repository hosting services using an SSH key.

When you only need one SSH key, you can skip the section that covers [Multiple accounts](#multiple-accounts-on-the-same-repository-hosting-service).

## Generate a new SSH key

To generate a new SSH key execute the following command:

```sh
ssh-keygen -t ed25519 -C {email}
```

- `{email}` is the email you use for GitHub or Bitbucket, configured as `user.email` in the `.gitconfig` file.
- The `-t` flag specifies the type of key to generate.
- The `-C` is only a comment.
	This flag is optional.

When you are prompted to "Enter a file in which to save the key", you can press Enter to accept the default file location.
Otherwise if you want a different name for your SSH key, you can enter a custom name.
The default directory for storing SSH keys is `~/.ssh/`.
Press `Enter` until the file is generated.

## Adding the SSH key 

The previous command generated two files (assuming you did not change the name when prompted):

- `id_ed25519` (private key)
- `id_ed25519.pub` (public key)

Copy entire contents of the `id_ed25519.pub` file.
The following command demonstrates how to copy a file to the clipboard:

```sh
cat ~/.ssh/id_ed25519.pub | wl-copy
```

In this example the output of the `cat` command got redirected to `wl-copy`; a utility part of [wl-clipboard](https://github.com/bugaevc/wl-clipboard).
Add the key (that should be inside your clipboard now) to your {Github, Bitbucket, Gitlab} account.

## Adding the private SSH key to the ssh-agent

Choose one of the following options to add a SSH key to the ssh-agent.

### Option 1: use ssh-add manually

Add the private SSH key previously generated to the ssh-agent by executing the following command:

```sh
ssh-add ~/.ssh/id_ed25519
```

To automate this process put this command in your shell initialization file (`~/.bash_profile` or `~/.zshrc`).
 
### Option 2 (preferred): use Keychain

Another option to add a SSH key to the ssh-agent is to use the [Keychain](https://www.funtoo.org/Funtoo:Keychain) command.

```sh
keychain ~/.ssh/id_ed25519
```

To automate this process put this command in your shell initialization file (`~/.bash_profile` or `~/.zshrc`).

Keychain has benefit of starting ssh-agent if it has not already been started, and can be used to add GPG keys.

## Multiple accounts on the same repository hosting service

When you have a personal and work SSH key for the same repository hosting service, you probably need to take more steps to distinct the keys.

### Add both the SSH keys to the ssh-agent

Use one of the options explained in [the previous subsection](#adding-the-private-ssh-key-to-the-ssh-agent) for each SSH key.
Verify if all SSH keys are loaded by executing the command:

```sh
ssh-add -l
```

or the Keychain equivalent:

```sh
keychain -l
```

### Create or modify the .ssh/config file

To tell Git which SSH key to use when interacting with the remote repository, we need to configure the `~/.ssh/config` file. 
Make a host enty with the following settings for each key:

```sh
Host github # Use this name in the remote URL
	HostName github.com # domain name
	User your_user_name
	IdentityFile ~/.ssh/id_ed25519-personal # location of the private key
	IdentitiesOnly yes

Host bitbucket # Use this name in the remote URL
	HostName bitbucket.org #domain name
	User your_user_name
	IdentityFile ~/.ssh/id_ed25519-work # location of the private key
	IdentitiesOnly yes
```

> [!TIP]
> It is recommended to use a lowercase name for the `Host`, because you need to add this name to your remote url.

### Update the existing remotes

You need to modify the remotes to specify which host should be used:
Instead of the default domain name for the host, use the name specified in your `~/.ssh/config` file. 

```sh
git remote set-url {remote_name} git@{Host}:{workspace}/{repository}.git
```

Example for this repository:

```sh
git remote set-url origin git@github:th7mo/second-brain.git
```

Also make sure that your local `.gitconfig` has the correct `user.name` and `user.email` for authentication.

## IncludeIf configuration in gitconfig

> [!Note]
> This section needs to be expanded.

## Reference

- For a more detailed explanation reference the GitHub article [Generating a new SSH key and adding it to the ssh-agent](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent).
- For a more detailed explanation for multiple SSH keys reference:
	- [Managing multiple Bitbucket user SSH keys on one device](https://support.atlassian.com/bitbucket-cloud/docs/managing-multiple-bitbucket-user-ssh-keys-on-one-device/).
	- [Multiple SSH Keys settings for different Bitbucket Cloud Accounts](https://confluence.atlassian.com/bbkb/multiple-ssh-keys-settings-for-different-bitbucket-cloud-accounts-1168847503.html).
