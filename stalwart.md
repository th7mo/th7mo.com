[Stalwart](https://stalw.art/) is a mail server written in Rust.
Stalwart can be used for hosting [email](email.md) on an own server.

## Installation
How to install Stalwart can be found at their [documentation](https://stalw.art/docs/install/linux).
After running the installation script the only thing left to configure is the user accounts in the SQLite database.
How to add users ([email](email.md) addresses) can be found [here](https://stalw.art/docs/directory/types/sql).
After configuring everything restart the service using:
```sh
systemctl restart stalwart-mail
```

## Ports
For [email](email.md) to work it is required to open ports that are used by [email](email.md) protocols.
Services that rent servers usually disable [email](email.md) ports by default.
It is most likely required to open a ticket to let support open those ports on an own server.
Following ports are required to be open to let Stalwart receive and send [email](email.md): 
```
443
465
993
587
25 
143
```

Open a port by executing the following `ufw` command:
```sh
ufw enable {port-number}
```

Repeat this command for every port listed above.

## DNS
Stalwart generates keys needed to add to the DNS configuration.

* `@` | `MX` | `mx.th7mo.com.` (trailing dot `.`)
* `stalwart._domainkey` | `TXT` | `{generated-DKIM1-key}` (including quotes)
* `_dmarc` | `TXT` | `{generated-DMARC1-key}` (including quotes) 
* `@` | `TXT` | `{generated-spf1-key}` (something like ` "v=spf1 a:mail.th7mo.com mx -all ra=postmaster" `)

## Add account to email client
### Gmail (mobile)
1. In the first screen `Add your email address` enter the [email](email.md) address.
2. When choosing which type of account this is, choose `Personal (IMAP)`.
3. Enter the password.
4. For `Incoming server settings`, adjust the `Server` from `th7mo.com` to `mail.th7mo.com`.
5. Also change the `Username` that defaults to the [email](email.md) address, to only the prefix.
   For example: when the [email](email.md) is `thimo@th7mo.com`, make the `Username` only `thimo` (without the domain name).
6. Apply step 5 again for `Outgoing server settings`.
7. Now the account is set up, a `Account name` and `Your name` can be chosen (optional).

### Manual setup 
This configuration is for other [email](email.md) clients.
It is possible that not every listed setting is required by the chosen [email](email.md) client.
The individual setting names can vary too.

#### Incoming server settings (IMAP)
* **Username**: [email](email.md) without domain (`thimo` for `thimo@th7mo.com`)
* **Server**: `mail.th7mo.com`
* **Port**: `993`
* **Connection security**: `SSL/TLS`
* **Authentication method**: `normal password`

#### Outgoing server settings (SMTP)
* **Username**: [email](email.md) without domain (`thimo` for `thimo@th7mo.com`)
* **Server**: `mail.th7mo.com`
* **Port**: `465`
* **Connection security**: `SSL/TLS`
* **Authentication method**: `normal password`
