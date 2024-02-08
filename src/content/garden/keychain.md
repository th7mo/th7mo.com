---
title: "Keychain"
description: "A practical manager for ssh-agent"
---

[Keychain](https://www.funtoo.org/Funtoo:Keychain) can be used to add [SSH keys](ssh-key) to the ssh-agent.

## Installation
Keychain does not come pre-installed on most systems, so first install Keychain:

```sh
sudo apt install keychain
```

## Add an SSH key to Keychain
To add a private SSH key to Keychain, execute:

```sh
keychain ~/.ssh/id_ed25519
```

This process can be automated by having this command in the [Bash](bash) or [Zsh](zsh) configuration file.

Keychain has the benefit of starting ssh-agent if it has not already been started,
and can be used to add GPG keys.

## List loaded SSH keys
Use the `-l` flag to list all the loaded SSH keys:

```sh
keychain -l
```
