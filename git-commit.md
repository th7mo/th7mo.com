Commits are snapshots of a [[Git]] repository.
Commits should be short, representing an atomic change.

## Commit message
A commit message can have a short summary, and an optional longer explanation after an empty line.
Do not use emoji's in commit messages like [gitmoji](https://gitmoji.dev/) advocates. 
Emojis do not provide any benefit to the commit message.
There are still a lot of tools that lack proper emoji support.
It is also way more difficult to filter or search commit messages when they use emojis for describing the commit.

Use the following prefixes for commit message titles:
* `feat: ` for new functionality.
* `fix: ` for reworking existing code to solve bugs.
* `refact: ` for refactoring code.
* `docs: ` for adding documentation.
* `test: ` for adding or fixing tests.
* `bump: ` for incrementing version numbers of dependencies.

A good convention to use is [conventional commits](https://www.conventionalcommits.org).
The [commitlint](https://github.com/conventional-changelog/commitlint) project expands on this and adds more types.

> [!IMPORTANT]
> Do not introduce to many prefix categories because it defeats the purpose if categorizing commits.

## Reverting commits
See [[git-revert]] for how to reverse changes stored in a commit.

## See also
* A [Medium article](https://medium.com/neudesic-innovation/conventional-commits-a-better-way-78d6785c2e08) that explains the most common commit prefixes.
* A [GitHub Gist](https://gist.github.com/qoomon/5dfcdf8eec66a051ecd85625518cfd13) about conventional commit messages.
