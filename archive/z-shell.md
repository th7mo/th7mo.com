The Z shell (Zsh) is a Unix shell like [[bash]] or Fish.
It is based on the Bourne Shell (sh) and [[bash]].
It is an enhancement on top of those shells.
A popular framework used on top of the Z shell is [[oh-my-zsh]]

## Installation
Install the Z shell by executing the following command:
```sh
sudo apt install zsh
```

Set Z shell as the default shell by executing the following command:
```sh
chsh -s $(which zsh)
```
* `$(which zsh)` evaluates to the path where `zsh` is located.
  This will most likely be `/usr/bin/zsh`.

## Configuration
Z shell is configured via the `.zshrc` file.
It is recommended to put the `.zshrc` file inside the [[dotfiles]] repository to keep configuration backed up.