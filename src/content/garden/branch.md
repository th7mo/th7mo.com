---
title: "Branch"
---

A [Git](git) branch is a path of one or more [commits](commit).
Branches allow developers to work on different features without affecting the main codebase.

When working on a project alone, working with a single branch is sufficient.
Create a new feature branch when it is desired to test in isolation
or when doing an entire project refactoring.

Use kebab-case for branch names as it is the most readable convention.
Suffix the name with an underscore `_` and an issue identifier when working on issues:

```
feat/add-primary-navigation-component_#29
```

or for [Jira](https://www.atlassian.com/software/jira) tickets:

```
fix/navigation-overflow-on-mobile_APP-43
```

Make branch names as descriptive as possible.
Long branch names shouldn't be avoided since
[Bash](bash), [Zsh](zsh), and other shells have branch
autocompletion for Git branches.
