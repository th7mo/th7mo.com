---
title: "Semantic versioning"
description: "How to version API's and public software"
---

Version software when publishing changes.
Any modifications must be released as a new version.
A good convention for versioning software is [semver](https://semver.org/).
A semantic version number consists of three components:

```
{MAJOR}.{MINOR}.{PATCH}
```

Increment the:

* `{MAJOR}` version when an incompatible API change is made.
* `{MINOR}` version when a backward compatible *new functionality* is added.
* `{PATCH}` version when a backward compatible *bug fix* is made.

When the `{MINOR}` version is incremented,
the `{PATCH}` version should be reset to zero.
Similarly, when the `{MAJOR}` version is incremented,
both the `{MINOR}` and `{PATCH}` versions should be reset to zero.

The `{MAJOR}` version zero is typically assigned to software that has not yet been released for production use.
