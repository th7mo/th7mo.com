Another option to add an [ssh-key](ssh-key.md) to the ssh-agent is to use the [Keychain](https://www.funtoo.org/Funtoo:Keychain) command.

## Installation
Keychain does not come pre-installed on most systems, so first install Keychain:
```sh
sudo apt install keychain
```

## Add an SSH key to Keychain
To add a private [ssh-key](ssh-key.md) to Keychain, execute:
```sh
keychain ~/.ssh/id_ed25519
```

This process can be automated by having this command in the [bash](bash.md) or [z-shell](z-shell.md) configuration file.

Keychain has the benefit of starting ssh-agent if it has not already been started, and can be used to add GPG keys.

## List loaded SSH keys
Use the `-l` flag to list all the loaded [ssh-key](ssh-key.md)s:
```sh
keychain -l
```
