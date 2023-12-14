# SSH key

A method to access and write data to GitHub and Bitbucket is Secure Shell Protocol (SSH).
You can authenticate to those repository hosting services using an SSH key.

## Generate a new SSH key

Execute:

```sh
ssh-key gen -t ed22519 -C {email}
```

- `{email}` is your email you use for GitHub or Bitbucket.

When you're prompted to "Enter a file in which to save the key", you can press Enter to accept the default file location.
The default location where the key is stored is `~/.ssh/`.
Press `Enter` until the file is generated.

## Adding the SSH key 

The previous command generated two files:

- `id_ed22519` (private key)
- `id_ed22519.pub` (public key)

Copy entire contents of the `id_ed22519.pub` file.
Add the SSH key to your GitHub or Bitbucket account using their website.

Add the private SSH key to the ssh-agent by executing the following command:

```sh
ssh-add ~/.ssh/id_ed25519
```

## Multiple accounts on the same repository hosting service

When you have a personal and work ssh key for the same repository hosting service, you probably need to take more steps to distinct the keys.

### Add both the ssh keys to the ssh-agent

Use the `ssh-add {ssh-key-path}` command explained in [the previous subsection](#adding-the-ssh-key) for each ssh key.
Verify if all ssh keys are loaded by using the command:

```sh
ssh-add -l
```

### Create or modify the .ssh/config file

Make a host with these settings for each key:

```sh
Host bitbucket.org-personal # Use this name in the remote URL
	HostName bitbucket.org
	User your_user_name
	IdentityFile ~/.ssh/id_ed25519-personal
	IdentitiesOnly yes

Host bitbucket.org-work # Use this name in the remote URL
	HostName bitbucket.org
	User your_user_name
	IdentityFile ~/.ssh/id_ed25519-work
	IdentitiesOnly yes
```

### Update the existing remotes

You need to modify the remotes to specify which host will be used:
Instead of the default name for the host, use the name specified in your `~/.ssh/config` file. 

```sh
git remote set-url {remote_name} git@{Host}:{workspace}/{repository}.git
```

Example for this repository:

```sh
git remote set-url origin git@bitbucket.org-personal:m7to/wiki.git
```

Also make sure that your local `.gitconfig` has the correct `user.name` and `user.email` for authentication.

## Reference

- For a more detailed explanation reference the GitHub article [Generating a new SSH key and adding it to the ssh-agent](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent).
- For a more detailed explanation for multiple ssh keys reference:
	- [Managing multiple Bitbucket user SSH keys on one device](https://support.atlassian.com/bitbucket-cloud/docs/managing-multiple-bitbucket-user-ssh-keys-on-one-device/).
	- [Multiple SSH Keys settings for different Bitbucket Cloud Accounts](https://confluence.atlassian.com/bbkb/multiple-ssh-keys-settings-for-different-bitbucket-cloud-accounts-1168847503.html).
