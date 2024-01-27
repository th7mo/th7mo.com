It is possible to revert changes without modifying the [[Git]] history.
This makes reverting [[git-commit]]s a safer option than [[git-reset]]ting commits.
To revert a commit execute the following command:
```sh
git revert {commit-hash}
```

Unlike `git reset`, the `git revert` command only reverts the changes of a specified commit, and will not revert the commit that came after the specified commit.

The `git revert` command also accepts a range of commits:
```sh
git revert --no-commit {commit-hash}..HEAD
git commit -m "revert: {reason for reverting}"
```

The `git revert` command above reverts all changes from the `{commit-hash}` up to and including where `HEAD` is.
It also adds the files that have reverts to the [[git-staging-area]].

> [!TIP]
> For more safety use the `--no-commit` flag.
> It allows for reviewing the reverted changes before the commit is made.
> To abort the revert execute `git revert --abort`
