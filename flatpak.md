# Flatpak

For some proprietary applications or applications not covered by the Gentoo repository, it could be better install a [Flatpak](https://flatpak.org/) package for that application.
[Steam](https://store.steampowered.com/) and [Parsec](https://parsec.app/) are examples of applications that are really ugly to install on Gentoo and should be isolated using Flatpak.

## Installation

The Gentoo package for Flatpak is [`sys-apps/flatpak`](https://packages.gentoo.org/packages/sys-apps/flatpak).
Install Flatpak with the following command:

```sh
emerge --ask sys-apps/flatpak
```

## Usage

Refer to the [Flatpak website](https://flatpak.org/) for how to install applications using Flatpak.

## Creating symbolic links

When the Flatpak applications are installed, most application launcher are not able to find them.
To make those launchers find the applications, it is possible to create symbolic links to the `/usr/bin/` directory (or any other directory included in the `$PATH` variable).

```sh
ln -s /var/lib/flatpak/exports/bin/{your-flatpak-name} /usr/bin/{custom-name}
```

For example for [Steam](https://flathub.org/apps/com.valvesoftware.Steam):

```sh
ln -s /var/lib/flatpak/exports/bin/com.valvesoftware.Steam /usr/bin/steam
```

Now application launchers (and your terminal) are able to find the application, because it is added to the `$PATH` variable.

## References

- The Gentoo article about [Flatpak](https://wiki.gentoo.org/wiki/Flatpak).
- I learned about symbolic linking of Flatpak installed application from this from a very specific [Github Gist](https://gist.github.com/curioswati/668e9e120ddd4b6f8d07dc28b5780d22).
