---
title: "Commit"
description: "A snapshot of the repository's currently staged changes"
---

A commit is a snapshot of a [Git](git) repository. Commits should be short, ideally representing an [atomic](atomic) change.

## Commit message
A commit message can have a short summary, and an optional longer explanation after an empty line. Do not use emoji's in commit messages like [the gitmoji project](https://gitmoji.dev/) tries to achieve.
Emojis do not provide any benefit to the commit message.
There are still a lot of tools that lack proper emoji support.
It is also way more difficult to filter or search commit messages when they use emojis for categorizing commits.

Use the following prefixes for commit message titles:

* `feat:` for new functionality.
* `fix:` for reworking existing code to solve bugs.
* `refact:` for refactoring code.
* `docs:` for adding documentation.
* `test:` for adding or fixing tests.
* `bump:` for incrementing version numbers of dependencies.

A good convention to use is [conventional commits](https://www.conventionalcommits.org).
The [commitlint project](https://github.com/conventional-changelog/commitlint) expands on this and adds more categories.

It is important to not introduce to many prefix categories.
This will defeat the purpose if categorizing commits.

## Undo commits
Sometimes it is desired to undo changes made by one or multiple commits.
Prefer [reverting](revert) over [resetting](reset) commits to preserve the git history.

## See also
* A [Medium article](https://medium.com/neudesic-innovation/conventional-commits-a-better-way-78d6785c2e08) that explains the most common commit prefixes.
* A [GitHub gist](https://gist.github.com/qoomon/5dfcdf8eec66a051ecd85625518cfd13) about conventional commit messages.
