---
title: "Semantic versioning"
---

Version software when publishing changes. Any modifications must be
released as a new version. A good convention for versioning software is
[semver](https://semver.org/). A semantic version number consists of
three components:

```
{MAJOR}.{MINOR}.{PATCH}
```

Increment the:

-   `{MAJOR}` version when an incompatible API change is
    made.
-   `{MINOR}` version when a backward compatible *new
    functionality* is added.
-   `{PATCH}` version when a backward compatible *bug fix* is
    made.

When the `{MINOR}` version is incremented, reset the
`{PATCH}` version to zero. When the `{MAJOR}`
version is incremented, reset both the `{MINOR}` and
`{PATCH}` versions to zero.

The `{MAJOR}` version zero is used when software is not used
in production yet.
