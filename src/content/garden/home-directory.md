---
title: "Home directory"
description: "The directory on a file-system for a specific user"
---

The home directory of a regular user is found at `/home/{username}/`,
while the root user's home directory is located at `/root/`.
In most operating systems, when the root user creates a new user,
a home directory is automatically generated for them.
This home directory typically comes equipped with default directories:

* `Desktop/`: files in this directory will be shown on the desktop environment.
* `Documents/`: the default location for storing files for word processors.
* `Downloads/`: the default location for downloading files.
* `Music/`: for audio.
* `Pictures/`: for images.
* `Public/`: a place where files get placed for remote access.
* `Templates/`: for templates like templates for word processors.
* `Videos/`: for videos.

I always create the following additional directories on a fresh installation.
These directories are in lowercase letters to distinguish them from the default directories:

* `dev/`: for all personal software development related files.
* `sandbox/`: a sandbox environment for testing purposes.
* `work/`: for all professional work related files.

## Shell handling
On Unix based systems the `$HOME` variable will evaluate to the
home directory of the active user.
In the shell the tilde `~` character is an alias for `/home/{active-user}`.
A user can navigate to the home folder no matter the current directory with:

```sh
cd ~
```

Executing the `cd` command without any arguments will also navigate to the active user's home directory:

```sh
cd
```
