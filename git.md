# Git

[Git](https://git-scm.com/) is a version control system.
It is used for the Second Brain to keep track of changes and preserve history.
It also allows for easy editing on different devices, like phones.

## Installation

Install Git by executing the following command:

```sh
sudo apt install git
```

## Configuration

Git is configured with `.gitconfig` files.
The main configuration file is located at `~/.gitconfig`.
The `.gitconfig` file stores all Git configuration like a `user.name` and `user.email` that is used for each commit.
It also stores custom aliases and branch handling preferences.

> [!TIP]
> It is recommended to put `.gitconfig` files in the [Dotfiles](/dotfiles.md).
> This will make sure the configuration is backed up.

## Multiple Git identities

Sometimes it is desired to overwrite the default configuration specified in `~/.gitconfig`.
This can be done by creating another `.gitconfig` file anywhere, and referencing it in the main `~/.gitconfig` file.
The `includeIf` option is used to specify when a `.gitconfig` file should be active.

For example:

```sh
[includeIf "gitdir:~/work/"]
    path = ~/work/.gitconfig
```

This specifies that the `~/work/.gitconfig` configuration should be used when inside a Git repository anywhere under the `~/work/` folder.
The `~/work/.gitconfig` can have a different `user.name` and `user.email` for professional use.

> [!IMPORTANT]
> Always end the `gitdir` path with a slash (`/`), otherwise it will not work!

## Commits

Commits should be short.
Commit messages can have a short summary, and an optional longer explanation after an empty line.
Do not use emoji's in commit messages like [gitmoji](https://gitmoji.dev/). 
It is not a benefit.
There are still a lot of tools that cannot handle emoji's.
It is also way more difficult to filter or search commit messages when they use emoji's for describing the commit.

Use the following prefixes for commit message titles:

`feat: ` for new functionality.

`fix: ` for reworking existing code to solve bugs.

`refact: ` for refactoring code.

`docs: ` for adding documentation.

`test: ` for adding or fixing tests.

`bump: ` for incrementing version numbers of dependencies.

It is important to not create to many prefixes because too many categories defeat the purpose if categorizing commits.

## Branches

When working alone on a small project a single `main` branch works perfect most of the time.
Add a new `feat/` branch when it is desired to test in isolation or when doing an entire overhaul/refactor of a project.

Use `lower-kabab-case` for naming branches, and use an underscore `_` for including an issue in the branch name:

```
feat/add-primary-navigation-component_#29
```

or for [Jira](https://www.atlassian.com/software/jira) tickets:

```
fix/navigation-overflow-on-mobile_APP-43
```

A single feature branch should ideally only solve one issue, so adding an issue identifier should work most of the time.

It is perfectly fine to have a long branch name.
A more descriptive name is always better.
[Bash](https://www.gnu.org/software/bash/) and other shells have branch autocompletion for Git branches and for graphical interfaces the branch name length does not matter either.

## Worktrees

> [!NOTE]
> This section will be explained when I actually included work trees in my own workflow.

Git Worktrees make branches root-level directories in the repository.
Instead of switching branches, just switch directories.

### Make a directory

```sh
mkdir {name-of-repository}
```

### Clone a repository without checking out a branch:

Clone the repository from inside the repository directory:

```sh
cd {name-of-repository}
```

After that clone the repository without Worktrees (`--bare` only clones the essential Git files).

```sh
git clone --bare {git-remote-url}
```

### Add a Worktree

Only add Worktrees from the bare repository, so navigate to the bare repository:

```sh
cd {name-of-repository}.git/
```

Now add a Worktree to the root directory of the repository:

```sh
git worktree add ../{branch-name}
```

Now the directory structure should look something like this: (as an example this repository has been used)

```
second-brain
├── second-brain.git/
└── main/
```

To switch to branch `main`, simply `cd main`.

### Add a Worktree with a new branch

```sh
git worktree add -b {new-branch-name} {worktree-name}
```

- `{branch-name}` is the name of the new branch
- `{worktree-name}` is the name of directory (Worktree)

### Add a Worktree based on a remote branch

```sh
git worktree add {worktree-name} {remote-branch-name}
```

### Remove a Worktree

```sh
git worktree remove {worktree-name}
```

For a full reference to Worktrees see the official [Git Worktree documentation](https://git-scm.com/docs/git-worktree).

