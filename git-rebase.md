A [[Git]] rebase adds one or multiple [[git-commit]]s from a [[git-branch]] on top of another branch.

The most simple use case for a rebase is when a [[git-remote]] branch somebody is also locally working on has new changes.
If somebody is working on branch `main`, but `origin/main` has new commits, the command `git pull origin main` can be executed to make a merge commits which adds the new commits on `origin/main` to the local `main` branch.

To preserve a clean commit history, an alternative approach would be a `git rebase`.

First, fetch the remote to download contents from the remote repository:
```sh
git fetch
```

If new changes have been made to the remote branch, include those remote changes by rebasing the commit made on top of the local branch as if they were made locally:
```sh
git rebase origin main
```

Now all the new commits made on the remote `main` branch will be applied to the local `main` branch.

A shorthand for this workflow is the `pull` command with the `--rebase` flag:
```sh
git pull --rebase {remote-name} {branch-name}
```

The `--rebase` flag can be omitted when the `pull.rebase` option is configured in the [[gitconfig]]:
```ini
[pull]
    rebase = true
```
