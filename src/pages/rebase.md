---
title:
---

A [Git](git.html) rebase adds one or multiple [commits](commit.html)
from a [branch](branch.html) on top of another branch.

The most simple use case for a rebase is when a [remote](remote.html)
branch somebody is also locally working on has new changes. First, fetch
the remote to download contents from the remote repository:

``` lang-sh
git fetch
```

If new changes have been made to the remote branch, include those remote
changes by rebasing the commit made on top of the local branch as if
they were made locally:

``` lang-sh
git rebase origin main
```

Now all the new commits made on the remote `origin/main` branch will be
applied to the local `main` branch.

A shorthand for this workflow is the `pull` command with the `--rebase`
flag:

```sh
git pull --rebase {remote-name} {branch-name}
```

The `--rebase` flag can be omitted when the `pull.rebase` option is
configured in the [gitconfig](gitconfig.html):

    [pull]
        rebase = true
