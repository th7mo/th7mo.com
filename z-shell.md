The Z shell (Zsh) is a Unix shell like Bash or Fish.
It is based on the Bourne Shell (sh) and Bash.
It is an enhancement on top of those shells.

## Installation
Install the Z shell by executing the following command:
```sh
sudo apt install zsh
```

Set Z shell as the default shell by executing the following command:
```sh
chsh -s $(which zsh)
```

## Configuration
Z shell is configured via the `.zshrc` file.

## Oh My Zsh
A popular framework for managing a Z shell configuration is [Oh My Zsh](https://ohmyz.sh/).
Install Oh My Zsh by executing the following command:
```sh
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
```

### Plugins
This sections explains how to install and enable various plugins for Oh My Zsh.
All the available plugins for Oh My Zsh can be found at their [GitHub repository](https://github.com/ohmyzsh/ohmyzsh/tree/master/plugins).

#### Autosuggestions
Install the [git](git.md) repository for [Z shell autosuggestions](https://github.com/zsh-users/zsh-autosuggestions):
```sh
git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
```

To enable autosuggestions, add it to the plugins list in `.zshrc`:
```sh
plugins=(
    # other plugins...
    zsh-autosuggestions
)
```
