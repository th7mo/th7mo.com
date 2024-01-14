# Dotfiles

It is a good idea to make a Git repository to back up and sync configuration files.
However, it is not recommended to make a Git repository directly into the `~/` home directory.
A better alternative is to move configuration files into a separate directory like `~/dotfiles/`.
To make sure installed programs still have access to the configuration files, symbolic links to the configuration files can be made.
[Stow](https://www.gnu.org/software/stow/) is a utility that simplifies the process of making symbolic links.

> [!NOTE]
> This might be replaced with a Git `--bare` repository using a worktree implementation.

## Installation

Install Stow by executing the following command:

```sh
sudo apt install stow
```

## Usage

Mirror what the `~/` directory structure should be exactly inside the `~/dotfiles/` directory, because Stow will make symlinks in directories following the same hierarchy.
The name of the Stow directory can be anything, but in this note the name `dotfiles` will be used.

For example: [Helix](https://helix-editor.com/) wants a `config.toml` file located at `~/.config/helix/config.toml`.
If the `config.toml` file is located at `~/dotfiles/.config/helix/config.toml`, Stow will create a symbolic link to where Helix wants to read the configuration file.

To create symbolic links for all the files in the `~/dotfiles/` directory to the `~/` directory, execute the following command:

```sh
stow ~/dotfiles/
```

Or in the repository **at root level**:

```sh
stow .
```

> [!WARNING]
> Be careful with where to execute this command.
> Only execute `stow .` at the root of the repository, otherwise the symbolic links will be invalid.
> Also be careful when removing symlinks and make sure to have a backup of the original configuration files (dotfiles).
