Version software when publishing changes.
Any modifications must be released as a new version.
A good convention for versioning software is [semver](https://semver.org/).
A version number consists of three components:
```
{MAJOR}.{MINOR}.{PATCH}
```

Increase the:
* `{MAJOR}` version when an incompatible API change is made.
* `{MINOR}` version when a backward compatible *new functionality* is added.
* `{PATCH}` version when a backward compatible *bug fix* is made.

The `{MAJOR}` version 0 can be used when in development.
