# NetworkManager

NetworkManager is a program for connecting to internet.

## Wi-Fi

A command line tool to connect with the internet is called `nmcli`.
The graphical equivalent provided with NetworkManager is `nmtui`.

To list the available Wi-Fi networks execute:

```sh
nmcli device wifi list
```

To connect to Wi-Fi using `nmcli` execute:

```sh
nmcli device wifi connect [SSID] password [password]
```

- `[SSID]` is the name (unique ID) of the Wi-Fi network.
- `[password]` is the password of the Wi-Fi network.
