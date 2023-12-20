# Email

I use my own custom domain `th7mo.com` for my emails.
I use the following emails:

- `dev@th7mo.com`
- `thimo@th7mo.com`, 
- `void@th7mo.com`, this is public and can be used for shopping and spam.

Legacy addresses:

- `work@th7mo.com`, currently used for only LinkedIn.
- `gaming@th7mo.com`, currently used for gaming platforms (League, Blizzard, Xbox, Parsec, Steam)
- `media@th7mo.com`, currently used for media platforms (Spotify, Tidal, Discord, Instagram, Snapchat, Reddit)

# New situation - Yet to implement

## Design philosophy

Excessive categorizing is doing me more harm than good.
Setting up the emails on a new email client takes a lot of time and I doubt a lot which email to use for some accounts.
I am using a unique password for each account, so separating really important stuff from regular stuff is secure enough for me.
I handle and archive each email on each account, so there is no real difference from separate emails.
There is no real difference from separate emails since I handled and archived each email on each account anyway. 

## root@th7mo.com 

For everything that involves money (except regular subscriptions) or other really important stuff. 
Contains the sensitive old `thimo@th7mo.com` accounts.

## thimo@th7mo.com

For public. 
Probably the most accounts will be using this email address.
Merges `media@th7mo.com`, `dev@th7mo.com`, `work@th7mo.com`, `gaming@th7mo.com` and old `thimo@th7mo.com` public stuff.

## thimoquinten@gmail.com

For throw-away stuff and spam I never use.
Migrates old `void@th7mo.com` accounts.
A fallback for when my other emails are compromised.

# Improvements

- implement the new situation
- remove the old situation
- cleanup references that hint that there was a previous implementation
- don't forget to update my email in `frontend-monorepo` from `dev@th7mo.com` to `thimo@th7mo.com`
- don't forget to use `thimo@th7mo.com` instead of `media@th7mo.com` for YouTube Premium
