# Git

[Git](https://git-scm.com/) is a version control system.
It is used for the Second Brain to keep track of changes and preserve history.
It also allows for easy editing on different devices, like phones.

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

Git is configured with `.gitconfig` files.
The main configuration file is located at `~/.gitconfig`.
The `.gitconfig` file stores all Git configuration like a `user.name` and `user.email` that is used for each commit.
It also stores custom aliases and branch handling preferences.

> [!TIP]
> It is recommended to put `.gitconfig` files in the [Dotfiles](/dotfiles.md).
> This will make sure the configuration is backed up.

## Commits

Commits should be short.
Commit messages can have a short summary, and an optional longer explanation after an empty line.
Do not use emoji's in commit messages like [gitmoji](https://gitmoji.dev/). 
Emojis do not provide any benefit to the commit message.
There are still a lot of tools that lack proper emoji support.
It is also way more difficult to filter or search commit messages when they use emojis for describing the commit.

Use the following prefixes for commit message titles:

- `feat: ` for new functionality.
- `fix: ` for reworking existing code to solve bugs.
- `refact: ` for refactoring code.
- `docs: ` for adding documentation.
- `test: ` for adding or fixing tests.
- `bump: ` for incrementing version numbers of dependencies.

It is important to not create to many prefixes because too many categories defeat the purpose if categorizing commits.

## Reverting commits

It is possible to revert changes without modifying the Git history.
This makes reverting commits a safer option than resetting commits.
To revert a commit execute the following command:

```sh
git revert {commit-hash}
```

Unlike `git reset`, the `git revert` command only reverts the changes of a specified commit, and will not revert the commits that came after the specified commit.

The `git revert` command also accepts a range of commits:

```sh
git revert --no-commit {commit-hash}..HEAD
git commit -m "revert: {reason for reverting}"
```

The `git revert` command above reverts all changes from the `{commit-hash}` up to and including where `HEAD` is.
It also adds the reverts to the staging area.

> [!TIP]
> For more safety use the `--no-commit` flag.
> It allows for reviewing the reverted changes before the commit is made.
> To abort the revert execute `git revert --abort`

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

## Rebase

A Git rebase adds commits from a branch on top of another branch.

### Usage

The most simple use case for a rebase is when a remote branch somebody is also locally working on has new changes.
If somebody is working on branch `main`, but `origin/main` has new commits, the command `git pull origin main` can be executed to make a merge commit which adds the new commits on `origin/main` to the local `main` branch.

To preserve a clean commit history, an alternative approach would use a rebase.

First, fetch the remote to download contents from the remote repository:

```sh
git fetch
```

If new changes have been made to the remote branch, include those remote changes by rebasing the commits made on top of the local branch as if they were made locally:

```sh
git rebase origin main
```

Now all the new commits made on the remote `main` branch will be applied to the local `main` branch.

A shorthand for this workflow is the `pull` command with the `--rebase` flag:

```sh
git pull --rebase {remote-name} {branch-name}
```

The `--rebase` flag can be omitted when the `pull.rebase` option is configured in the `.gitconfig`:

```ini
[pull]
    rebase = true
```

### Pull request / Merge request

The rebase option is also a good option for pull requests.

> [!TIP]
> Do not squash commits or merge a pull request using the `squash` option, unless the commit messages are not providing any information.
> It is always better to leave as much history as possible for later debugging of code.

## Worktrees

> [!NOTE]
> This section will be explained when I actually included work trees in my own workflow.

Git worktrees make branches directories at the root of the repository.
Instead of switching branches, just switch directories.

### Make a directory

```sh
mkdir {repository-name}
```

### Clone a repository without checking out a branch:

Clone the repository from inside the repository directory:

```sh
cd {repository-name}
```

After that clone the repository without worktrees: (The `--bare` flag specifies to only clone the essential Git files)

```sh
git clone --bare {git-remote-url}
```

- `{git-remote-url}` can be a `HTTPS` or `SSH` URL.

### Add a worktree

Only add worktrees from the bare repository, so navigate to the bare repository:

```sh
cd {repository-name}/.git/
```

Now add a worktree to the root of the repository:

```sh
git worktree add ../{branch-name}
```

Now the directory structure should look something like this: (as an example this repository has been used)

```
second-brain
├── second-brain/.git/
└── main/
```

To switch to branch `main`, simply `cd main`.

### Add a worktree with a new branch

Make sure to be in the bare repository when managing worktrees.

```sh
git worktree add -b {new-branch-name} ../{worktree-name}
```

- `{branch-name}` is the name of the new branch
- `{worktree-name}` is the name of directory (worktree)

### Add a worktree based on a remote branch

```sh
git worktree add {worktree-name} {remote-branch-name}
```

> [!TIP]
> When the worktree commit log does not line up with the remote commit log for that specific branch, it is probably because `{remote-branch-name}` was not the correct remote branch name when executing the `git worktree add` command.

### Remove a worktree

Remove worktrees from the bare repository:

```sh
cd {repository-name}/.git/
```

Remove a worktree by executing the following command:

```sh
git worktree remove {worktree-name}
```

For a full reference to worktrees see the official [Git worktree documentation](https://git-scm.com/docs/git-worktree).

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

## See also

- A [Reddit thread](https://www.reddit.com/r/git/comments/wwapum/comment/ilkdpzv/) about worktrees.
- The [Medium article](https://medium.com/ngconf/git-worktrees-in-use-f4e516512feb) about worktrees. 
- The [Git documentation](https://git-scm.com/docs/git-worktree) about worktrees.
