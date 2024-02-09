---
title: "SSH key"
description: "Managing authentication with SSH keys"
---

A method to access and write data to GitHub, GitLab and Bitbucket is the Secure Shell Protocol (SSH).
SSH keys can be used to authenticate to those repository hosting services.

When only one SSH key is needed to configure, the section that covers [Managing multiple SSH keys](#managing-multiple-ssh-keys) can be ignored.

## Generate a new SSH key
Create a new SSH key using the [ssh-keygen](ssh-keygen) tool.

## Adding the SSH key 
The previous command generated two files (unless another file name or path was specified at the [Generate a new SSH key](#generate-a-new-ssh-key) section):
* `~/.ssh/id_ed25519` (private key)
* `~/.ssh/id_ed25519.pub` (public key)

Copy entire contents of the `~/.ssh/id_ed25519.pub` file.
The following command demonstrates how to copy a file to the clipboard:
```sh
cat ~/.ssh/id_ed25519.pub | wl-copy
```

In this example the output of the `cat` command got redirected to `wl-copy`;
a utility part of [`wl-clipboard`](https://github.com/bugaevc/wl-clipboard).
Add the key (that should be inside the clipboard now) to a {GitHub, Bitbucket, GitLab} account.

## Adding the private SSH key to the ssh-agent
Choose one of the following options to add an SSH key to the ssh-agent.

### Option 1: use ssh-add manually
Add the private SSH key previously generated to the ssh-agent by executing the following command:
```sh
ssh-add ~/.ssh/id_ed25519
```

This process can be automated by having this command in the [Bash](bash) or [Zsh](zsh) configuration file.

### Option 2 (preferred): use Keychain
The [Keychain](keychain) tool has better tooling for adding SSH keys.
Refer to [Add an SSH key to Keychain](keychain#add-an-ssh-key-to-keychain) for how to add an SSH key to Keychain.

## Managing multiple SSH keys
When two or more SSH keys are used, more steps are required to let SSH know which key to use when.

### Add both the SSH keys to the ssh-agent
Use one of the options explained in [Adding the private SSH key to the ssh-agent](#adding-the-private-ssh-key-to-the-ssh-agent) for each SSH key.
Verify that all SSH keys are loaded by executing the command:
```sh
ssh-add -l
```

if Keychain is used reference [List loaded SSH keys](keychain#list-loaded-ssh-keys).

### Create or modify the SSH configuration
To tell [Git](git) which SSH key to use when interacting with the remote repository, we need to configure the `~/.ssh/config` file. 
Make a host entry with the following settings for each key:

```sh
Host github.com
    HostName github.com
    User user_name
    IdentityFile ~/.ssh/id_ed25519-personal
    IdentitiesOnly yes

Host bitbucket.com
    HostName bitbucket.org
    User user_name
    IdentityFile ~/.ssh/id_ed25519-work
    IdentitiesOnly yes
```

Make the `Host` name the same has the `HostName`.
When those are the same there is no need to update any existing remotes because the URL stays the same.

When using multiple Git identities, also follow [Git identities](git-identities).

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

Make sure that the local [gitconfig](gitconfig) has the correct `user.name` and `user.email` for authentication.

## See also
* For a more detailed explanation reference the GitHub article [Generating a new SSH key and adding it to the ssh-agent](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent).
* For a more detailed explanation for multiple SSH keys reference:
    * [Managing multiple Bitbucket user SSH keys on one device](https://support.atlassian.com/bitbucket-cloud/docs/managing-multiple-bitbucket-user-ssh-keys-on-one-device/).
    * [Multiple SSH Keys settings for different Bitbucket Cloud Accounts](https://confluence.atlassian.com/bbkb/multiple-ssh-keys-settings-for-different-bitbucket-cloud-accounts-1168847503.html).
