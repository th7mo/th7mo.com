---
title:
---

Version software when publishing changes. Any modifications must be
released as a new version. A good convention for versioning software is
[semver](https://semver.org/). A semantic version number consists of
three components:

```
{MAJOR}.{MINOR}.{PATCH}
```

Increment the:

-   `{MAJOR}`{is:raw=""} version when an incompatible API change is
    made.
-   `{MINOR}`{is:raw=""} version when a backward compatible *new
    functionality* is added.
-   `{PATCH}`{is:raw=""} version when a backward compatible *bug fix* is
    made.

When the `{MINOR}`{is:raw=""} version is incremented, reset the
`{PATCH}`{is:raw=""} version to zero. When the `{MAJOR}`{is:raw=""}
version is incremented, reset both the `{MINOR}`{is:raw=""} and
`{PATCH}`{is:raw=""} versions to zero.

The `{MAJOR}`{is:raw=""} version zero is used when software is not used
in production yet.
