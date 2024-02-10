---
title: "Squash"
description: "Merging multiple commits"
---

Squashing [commits](commit) is a technique used in [Git](git)
to merge multiple commits into a single commit.
This is helpful in keeping a clean Git history.
Services like [GitHub](https://github.com/) can squash pull request,
which combines all the commits from that pull request into a single commit.
To squash commits, the [rebase](rebase) command is used with the `-i` flag
followed by the number of commits to be squashed:

```sh
git rebase -i HEAD~{number-of-commits}
```

Since squashing commits modifies the Git history,
the `-i` flag needs to be used for an 'interactive' rebase
which allows the user to edit the commit history.
