[Git](https://git-scm.com/) is a version control system.
It is used for the [[second-brain]] to keep track of changes and preserve history.
It also allows for easy editing on different devices, like phones.

<!-- TODO -->
> [!NOTE]
> This note is not complete yet.
> The Git workflow (branch strategy) can have a better explanation.
> 
> The terms used in the worktree section is inconsistent, and there should be a warning about only executing commands in the bare workspace `{repository-name}/.git/`
> 
> There is also a way to do not rely on GNU Stow but use a bare repository to manage dotfiles.
> This needs to be researched first.

## Installation
Install Git by executing the following command:
```sh
sudo apt install git
```

## Configuration
All the configuration for Git is stored in [[gitconfig]] files.

## Usage
A [[git-commit]] creates a snapshot of the current state of a Git repository.
Git uses [[git-branch]]es for collaborate work and diverging from the default environment.

## Pull request / Merge request
The [[git-rebase]] option is also a good option for pull requests.

> [!TIP]
> Do not squash commits or merge a pull request using the `squash` option, unless the commit messages are not providing any usefull information.
> It is always better to leave as much history as possible for later debugging of code.

## Worktrees
[[Git-worktree]]s can be used to switch between branches without committing or stashing changes. 

## Multiple Git identities
See [[multiple-git-identities]] for how to configure multiple Git identities.

## See also
* A [Reddit thread](https://www.reddit.com/r/git/comments/wwapum/comment/ilkdpzv/) about worktrees.
* The [Medium article](https://medium.com/ngconf/git-worktrees-in-use-f4e516512feb) about worktrees. 
* The [Git documentation](https://git-scm.com/docs/git-worktree) about worktrees.
