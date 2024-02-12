---
title: "Network Manager"
description: "Managing wireless internet connections on Linux"
---

[NetworkManager](https://networkmanager.dev) is a program for connecting
to internet. It is only needed for wireless connections.

## Installation
Debian probably has NetworkManager pre-installed. A manual install can
be done by executing the following command:

```sh
sudo apt install network-manager
```

It is recommended to use [iwd](https://wiki.gentoo.org/wiki/Iwd) as a
backend for NetworkManager. Iwd is the modern replacement for
`wpa_supplicant`.

## Usage

NetworkManager provides a command line (CLI) tool to connect with the
internet called `nmcli` (NetworkManager CLI). The graphical equivalent
provided is `nmtui` (NetworkManager Terminal User Interface).

To list the available Wi-Fi networks execute:

```sh
nmcli device wifi list
```

To connect to a Wi-Fi network using `nmcli` execute:

```sh
nmcli device wifi connect {SSID} password {password}
```

* `{password}` is the password of the Wi-Fi network.

## Eduroam
It is not possible to connect to the eduroam network using
NetworkManager without custom configuration. For more information
reference about how to the following:

* The Arch Wiki section about [network configuration with eduroam](https://wiki.archlinux.org/title/Network_configuration/Wireless#eduroam).
* To generate configuration for NetworkManager use the [eduroam Configuration Assistant Tool](https://cat.eduroam.org/).

## See also
* The [Gentoo wiki article](https://wiki.gentoo.org/wiki/NetworkManager) about NetworkManager.
* The [Arch Wiki article](https://wiki.archlinux.org/title/NetworkManager) about NetworkManager.
* The website of [NetworkManager](https://networkmanager.dev/docs/).
