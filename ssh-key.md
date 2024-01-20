# SSH key

A method to access and write data to GitHub, GitLab and Bitbucket is the Secure Shell Protocol (SSH).
SSH keys can be used to authenticate to those repository hosting services.

When only one SSH key is needed to configure, the section that covers [Multiple SSH keys](#managing-multiple-ssh-keys) can be ignored.

## Generate a new SSH key

To generate a new SSH key execute the following command:

```sh
ssh-keygen -t ed25519 -C {email}
```

- The `-t` flag specifies the type of key to generate.
- The `-C` adds a comment to the SSH key.
	This flag is optional but recommended.
	It will help indicate which key belongs to which [email](/email.md) when using `ssh-add -l` or `keychain -l`.
- `{email}` is the [email](/email.md) used for GitHub or Bitbucket, configured as `user.email` in the [gitconfig](/gitconfig.md) file.

When prompted to "Enter a file in which to save the key", press `Enter` to accept the default file location.
Otherwise, if a different name or located is desired, a custom name can be specified.
The default directory for storing SSH keys is `~/.ssh/`.
Press `Enter` until the file is generated.

## Adding the SSH key 

The previous command generated two files (unless another file name or path was specified at the [Generate a new SSH key](#generate-a-new-ssh-key) section):

- `~/.ssh/id_ed25519` (private key)
- `~/.ssh/id_ed25519.pub` (public key)

Copy entire contents of the `~/.ssh/id_ed25519.pub` file.
The following command demonstrates how to copy a file to the clipboard:

```sh
cat ~/.ssh/id_ed25519.pub | wl-copy
```

In this example the output of the `cat` command got redirected to `wl-copy`; a utility part of [`wl-clipboard`](https://github.com/bugaevc/wl-clipboard).
Add the key (that should be inside the clipboard now) to a {GitHub, Bitbucket, GitLab} account.

## Adding the private SSH key to the ssh-agent

Choose one of the following options to add an SSH key to the ssh-agent.

### Option 1: use ssh-add manually

Add the private SSH key previously generated to the ssh-agent by executing the following command:

```sh
ssh-add ~/.ssh/id_ed25519
```

This process can be automated by having this command in the shell initialization file (`~/.bash_profile`, `.profile` or `~/.zshrc`).
 
### Option 2 (preferred): use Keychain

Another option to add an SSH key to the ssh-agent is to use the [Keychain](https://www.funtoo.org/Funtoo:Keychain) command.

Keychain does not come pre-installed on most systems, so first install Keychain:

```sh
sudo apt install keychain
```

To add a private SSH key to Keychain, execute:

```sh
keychain ~/.ssh/id_ed25519
```

This process can be automated by having this command in the shell initialization file (`~/.bash_profile`, `.profile` or `~/.zshrc`).

Keychain has the benefit of starting ssh-agent if it has not already been started, and can be used to add GPG keys.

## Managing multiple SSH keys

When two or more SSH keys are used, more steps are required to let SSH know which key to use when.

### Add both the SSH keys to the ssh-agent

Use one of the options explained in [the previous subsection](#adding-the-private-ssh-key-to-the-ssh-agent) for each SSH key.
Verify that all SSH keys are loaded by executing the command:

```sh
ssh-add -l
```

or the Keychain equivalent:

```sh
keychain -l
```

### Create or modify the .ssh/config file

To tell [git](/git.md) which SSH key to use when interacting with the remote repository, we need to configure the `~/.ssh/config` file. 
Make a host entry with the following settings for each key:

```sh
Host github.com                           # Use this name in the remote URL
	HostName github.com                     # domain name
	User user_name                          # put this in double quotes when the User name has spaces
	IdentityFile ~/.ssh/id_ed25519-personal # location of the private key
	IdentitiesOnly yes

Host bitbucket.com                        # Use this name in the remote URL
	HostName bitbucket.org                  # domain name
	User user_name                          # put this in double quotes when the User name has spaces
	IdentityFile ~/.ssh/id_ed25519-work     # location of the private key
	IdentitiesOnly yes
```

> [!TIP]
> Make the `Host` name the same has the `HostName`.
> When those are the same there is no need to update any existing remotes because the URL stays the same.

When using multiple [git](/git.md) identities, also follow the section [git#Multiple-git-identities](/git.md#multiple-git-identities).

### Update the existing remotes

When the `Host` name has the same name as the `HostName`, this section can be skipped.

When the `Host` name is not the same has the `HostName` need, modify the remotes to specify which host should be used:
Instead of the default domain name for the host, the `Host` name specified in the `~/.ssh/config` file must be used.

```sh
git remote set-url {remote_name} git@{Host}:{workspace}/{repository}.git
```

Example for this repository when `Host` is named `github` instead of `github.com`:

```sh
git remote set-url origin git@github:th7mo/second-brain.git
```

Make sure that the local [gitconfig](/gitconfig.md) has the correct `user.name` and `user.email` for authentication.

## See also

- For a more detailed explanation reference the GitHub article [Generating a new SSH key and adding it to the ssh-agent](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent).
- For a more detailed explanation for multiple SSH keys reference:
	- [Managing multiple Bitbucket user SSH keys on one device](https://support.atlassian.com/bitbucket-cloud/docs/managing-multiple-bitbucket-user-ssh-keys-on-one-device/).
	- [Multiple SSH Keys settings for different Bitbucket Cloud Accounts](https://confluence.atlassian.com/bbkb/multiple-ssh-keys-settings-for-different-bitbucket-cloud-accounts-1168847503.html).
