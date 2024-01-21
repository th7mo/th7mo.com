[Stow](https://www.gnu.org/software/stow/) is a utility that simplifies the process of making [symbolic-link](symbolic-link.md)s.
It can be used to easily manage [dotfiles](dotfiles.md).

## Installation

Install Stow by executing the following command:

```sh
sudo apt install stow
```

## Usage

Mirror what the `~/` directory structure should be exactly inside the `~/dotfiles/` directory, because Stow will make [symbolic-link](symbolic-link.md)s in directories following the same hierarchy.
The name of the Stow directory can be anything, but in this note the name `dotfiles` will be used.

For example: [Helix](https://helix-editor.com/) wants a `config.toml` file located at `~/.config/helix/config.toml`.
If the `config.toml` file is located at `~/dotfiles/.config/helix/config.toml`, Stow will create a [symbolic-link](symbolic-link.md) to where Helix wants to read the configuration file.

To create [symbolic-link](symbolic-link.md)s for all the files in the `~/dotfiles/` directory to the `~/` directory, execute the following command:

```sh
stow ~/dotfiles/
```

Or in the repository **at root level**:

```sh
stow .
```

> [!WARNING]
> Careful with where to execute this command.
> Only execute `stow .` at the root of the repository, otherwise the [symbolic-link](symbolic-link.md)s will be invalid.
> Also be careful when removing [symbolic-link](symbolic-link.md) and make sure to have a backup of the original configuration files ([dotfiles](dotfiles.md)).
