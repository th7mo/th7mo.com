# Stalwart

[Stalwart](https://stalw.art/) is a mail server written in Rust.
I use Stalwart for hosting my email on my own server.

## Installation

How to install Stalwart can be found at their [docs](https://stalw.art/docs/install/linux).
After running the installation script the only thing left to configure is the user accounts in the SQLite database.
How to add users (email addresses) can be found [here](https://stalw.art/docs/directory/types/sql).
After configuring everything restart the service using:

```sh
systemctl restart stalwart-mail
```

## Ports

For email to even work on the server you need to open ports email protocols use.
Services that let you rent a server usually disable email ports by default.
You probably need to open a ticket to let support open those ports on your server.
I don't know which ports are strictly needed for Stalwart to work, but I have a working version with the following ports open:

```
443                        ALLOW IN    Anywhere
465                        ALLOW IN    Anywhere
993                        ALLOW IN    Anywhere
587                        ALLOW IN    Anywhere
25                         ALLOW IN    Anywhere
143                        ALLOW IN    Anywhere
```

Open a port by using the following `ufw` command:

```sh
ufw enable {port-number}
```

## DNS

Stalwart generates keys you need to add to your DNS configuration.

- `@` | `MX` | `mx.th7mo.com.` (trailing dot (`.`))
- `stalwart._domainkey` | `TXT` | `{your-DKIM1-key}` (including quotes)
- `_dmarc` | `TXT` | `{your-DMARC1-key}` (including quotes) 
- `@` | `TXT` | `{your-spf1-key}` (something like ` "v=spf1 a:mail.th7mo.com mx -all ra=postmaster" `)

## Add account to email client

### Gmail (mobile)

1. In the first screen `Add your email address` enter your email address.
2. When choosing which type of account this is, choose `Personal (IMAP)`.
3. Enter your password.
4. For `Incoming server settings`, adjust the `Server` from `th7mo.com` to `mail.th7mo.com`.
5. Also change your `Username` that defaults to your email address, to only the prefix.
   For example: when your email is `thimo@th7mo.com`, make your `Username` only `thimo` (without the domain name).
6. Apply step 5 again for `Outgoing server settings`.
7. Now your account is set up, you can choose to full in a `Account name` and `Your name` (optional).

### Manual setup 

This configuration is for other email clients.
It is possible that not every listed setting is required by your email client.
The individual setting names can vary too.

#### Incoming server settings (IMAP)

**Username**: email without domain (`thimo` for `thimo@th7mo.com`)

**Server**: `mail.th7mo.com`

**Port**: `993`

**Connection security**: `SSL/TLS`

**Authentication method**: `normal password`

#### Outgoing server settings (SMTP)

**Username**: email without domain (`thimo` for `thimo@th7mo.com`)

**Server**: `mail.th7mo.com`

**Port**: `465`

**Connection security**: `SSL/TLS`

**Authentication method**: `normal password`
