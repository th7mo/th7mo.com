# To-Do

This is an unordered, unstructured collection of things I still need to do.
This To-Do list can never contain everything I need to do. 
I depend on [Git](https://git-scm.com/) for version control.

- Install [NVM](https://github.com/nvm-sh/nvm) (never use Debian package repository for managing Node versions)
- Use NPM to install LSP servers for the languages I use (check existing `languages.toml` on GitHub for used languages).
- Swapping the stuff from my closet and my bed-drawers.
- Write why I still use X.org (NVIDIA) and explain how to set it up.
  - Explain that you NEVER use `xrandr` but use `.Xresources` to scale.
  - When specifying a cursor and size you get consistent cursor.
  - Don't forget to source the `.Xresources` file in the `.xinitrc` file.
- Remove dependency on Google Photos for my media.
  I want a platform independent solution to store my media.
  A Network Attached Storage (NAS) or home server solution should be future-proof.
- Make a separate branch for pc and laptop for my `dotfiles` repository.
- Explain how to make audio work before you forget it
  - global pulseaudio flag
  - install pipewire and wireplumber
  - start your windowmanager using dbus-run-session or dbus-launch (otherwise your audio will probably not work)
  - use pavucontrol for easy managing of audio (it is not worth it for the terminal)
