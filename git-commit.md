[git](git.md) commits are snapshots of a [git](git.md) repository.
Commits should be short.

## Commit message

A commit message can have a short summary, and an optional longer explanation after an empty line.
Do not use emoji's in commit messages like [gitmoji](https://gitmoji.dev/) advocates. 
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

The [commitlint](https://github.com/conventional-changelog/commitlint) project explains this convention in more detail.

> [!IMPORTANT]
> It is important to not create to many prefixes because too many categories defeat the purpose if categorizing commits.

## Reverting commits

It is possible to revert changes without modifying the [git](git.md) history.
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
