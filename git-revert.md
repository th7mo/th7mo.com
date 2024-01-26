It is possible to revert changes without modifying the [git](git.md) history.
This makes reverting [git-commit](git-commit.md)s a safer option than [git-reset](git-reset.md)ting [git-commit](git-commit.md)s.
To revert a [git-commit](git-commit.md) execute the following command:
```sh
git revert {commit-hash}
```

Unlike `git reset`, the `git revert` command only reverts the changes of a specified [git-commit](git-commit.md), and will not revert the [git-commit](git-commit.md) that came after the specified [git-commit](git-commit.md).

The `git revert` command also accepts a range of [git-commit](git-commit.md)s:
```sh
git revert --no-commit {commit-hash}..HEAD
git commit -m "revert: {reason for reverting}"
```

The `git revert` command above reverts all changes from the `{commit-hash}` up to and including where `HEAD` is.
It also adds the files that have reverts to the staging area.

> [!TIP]
> For more safety use the `--no-commit` flag.
> It allows for reviewing the reverted changes before the [git-commit](git-commit.md) is made.
> To abort the revert execute `git revert --abort`
