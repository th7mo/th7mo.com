---
title: "Revert"
---

It is possible to revert changes without modifying the [Git](git)
history. This makes reverting [commits](commit) a safer option than
[resetting](reset) commits. To revert a commit execute the
following command:

```sh
git revert {commit-hash}
```

Unlike `git reset`, the `git revert` command only reverts the changes of
a specified commit, and will not revert the commit that came after the
specified commit.

The `git revert` command also accepts a range of commits:

```sh
git revert --no-commit {commit-hash}..HEAD
```

```sh
git commit -m "revert: {reason for reverting}"
```

The `git revert` command above reverts all changes from the
`{commit-hash}` up to and including where `HEAD` is. It also
adds the files that have reverts to the [staging area](staging-area).
For more safety use the `--no-commit` flag.
It allows for reviewing the reverted changes before the commit is made.
To abort the revert execute:

```sh
git revert --abort
```
