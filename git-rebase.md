# Git rebase

A [git](/git.md) rebase adds one or multiple [git-commit](/git-commit.md)s from a [git-branch](/git-branch.md) on top of another [git-branch](/git-branch.md).

### Usage

The most simple use case for a rebase is when a remote [git-branch](/git-branch.md) somebody is also locally working on has new changes.
If somebody is working on [git-branch](/git-branch.md) `main`, but `origin/main` has new [git-commit](/git-commit.md)s, the command `git pull origin main` can be executed to make a merge [git-commit](/git-commit.md) which adds the new [git-commit](/git-commit.md)s on `origin/main` to the local `main` branch.

To preserve a clean [git-commit](/git-commit.md) history, an alternative approach would use a rebase.

First, fetch the remote to download contents from the remote repository:

```sh
git fetch
```

If new changes have been made to the remote [git-branch](/git-branch.md), include those remote changes by rebasing the [git-commit](/git-commit.md) made on top of the local [git-branch](/git-branch.md) as if they were made locally:

```sh
git rebase origin main
```

Now all the new [git-commit](/git-commit.md)s made on the remote `main` [git-branch](/git-branch.md) will be applied to the local `main` [git-branch](/git-branch.md).

A shorthand for this workflow is the `pull` command with the `--rebase` flag:

```sh
git pull --rebase {remote-name} {branch-name}
```

The `--rebase` flag can be omitted when the `pull.rebase` option is configured in the [gitconfig.md](/gitconfig.md):

```ini
[pull]
    rebase = true
```
