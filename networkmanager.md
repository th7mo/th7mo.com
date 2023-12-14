# NetworkManager

NetworkManager is a program for connecting to internet.

## Wi-Fi

NetworkManager provides a command line (CLI) tool to connect with the internet called `nmcli`.
The graphical equivalent provided is `nmtui`.

To list the available Wi-Fi networks execute:

```sh
nmcli device wifi list
```

To connect to Wi-Fi using `nmcli` execute:

```sh
nmcli device wifi connect {SSID} password {password}
```

- `{SSID}` is the name (unique ID) of the Wi-Fi network.
- `{password}` is the password of the Wi-Fi network.

## References

- The website of [NetworkManager](https://networkmanager.dev/docs/).
- The Arch Wiki article about [NetworkManager](https://wiki.archlinux.org/title/NetworkManager).
