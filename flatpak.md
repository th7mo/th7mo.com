# Flatpak

Some proprietary applications or applications are not provided by the Debian repository.
In those cases it could be better to install the [Flatpak](https://flatpak.org/) packaged version of that application.
[Steam](https://store.steampowered.com/) and [Parsec](https://parsec.app/) are examples of applications that are ugly to install on Debian and should be isolated using Flatpak.

## Installation

The Debian package for Flatpak is `flatpak`.
Install Flatpak with the following command:

```sh
sudo apt install flatpak
```

## Usage

Refer to the [Flatpak website](https://flatpak.org/) for how to install applications using Flatpak.

## Make application launchers find Flatpak applications

When the Flatpak applications are installed, most application launchers are not able to find them.
This is because the Flatpak applications are not in a `$PATH` directory.
To make those application launchers find the Flatpak applications, it is possible to create [symbolic-link](symbolic-link.md)s to the `/usr/bin/` directory (or any other directory included in the `$PATH` variable).

```sh
ln -s /var/lib/flatpak/exports/bin/{installed-flatpak-name} /usr/bin/{custom-name}
```

For example for [Steam](https://flathub.org/apps/com.valvesoftware.Steam):

```sh
ln -s /var/lib/flatpak/exports/bin/com.valvesoftware.Steam /usr/bin/steam
```

Now application launchers (and the terminal) are able to find the Flatpak application, because it is added to the `$PATH` variable.

## See also

- The Gentoo article about [Flatpak](https://wiki.gentoo.org/wiki/Flatpak).
- I learned about [symbolic-link](symbolic-link.md)ing of Flatpak installed applications from this from a very specific [GitHub Gist](https://gist.github.com/curioswati/668e9e120ddd4b6f8d07dc28b5780d22).
