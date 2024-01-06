# NetworkManager

[NetworkManager](https://networkmanager.dev) is a program for connecting to internet.
It is only needed for wireless connections.

## Installation

Debian probably has Networkmanager pre-installed.
A manual install can be done executing the following command:

```sh
sudo apt install network-manager
```

> [!TIP]
> It is recommended to use [iwd](https://wiki.gentoo.org/wiki/Iwd) as a backend for NetworkManager.
> Iwd is the modern replacement for wpa_supplicant.

## Usage

NetworkManager provides a command line (CLI) tool to connect with the internet called `nmcli` (NetworkManager CLI).
The graphical equivalent provided is `nmtui` (NetworkManager Terminal User Interface).

To list the available Wi-Fi networks execute:

```sh
nmcli device wifi list
```

To connect to a Wi-Fi network using `nmcli` execute:

```sh
nmcli device wifi connect {SSID} password {password}
```

- `{SSID}` is the name (unique ID) of the Wi-Fi network.
- `{password}` is the password of the Wi-Fi network.

### Eduroam

> [!NOTE]
> The references to external resources need to be replaced with a tested guide.

It is not possible to connect to the eduroam network using NetworkManager without custom configuration.
For more information reference about how to the following:

- The Arch Wiki section about [network configuration with eduroam](https://wiki.archlinux.org/title/Network_configuration/Wireless#eduroam).
- To generate configuration for NetworkManager use the [eduroam Configuration Assistant Tool](https://cat.eduroam.org/).

## References

- The Gentoo wiki article about [Networkmanager](https://wiki.gentoo.org/wiki/NetworkManager).
- The Arch Wiki article about [NetworkManager](https://wiki.archlinux.org/title/NetworkManager).
- The website of [NetworkManager](https://networkmanager.dev/docs/).
