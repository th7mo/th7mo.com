Another option to add an [[ssh-key]] to the ssh-agent is to use the [Keychain](https://www.funtoo.org/Funtoo:Keychain) command.

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

This process can be automated by having this command in the [[bash]] or [[z-shell]] configuration file.

Keychain has the benefit of starting ssh-agent if it has not already been started, and can be used to add GPG keys.

## List loaded SSH keys
Use the `-l` flag to list all the loaded SSH keys:
```sh
keychain -l
```