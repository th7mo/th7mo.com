# To Do

This is a unordered, unstructured collection of things I still need to do.
This To Do list can never contain everything I need to do. 
I depend on [Git](https://git-scm.com/) for version control.

- Install [NVM](https://github.com/nvm-sh/nvm) (don't use Gentoo package repository for this)
- Use NPM to install LSP servers for the languages I use (check existing `languages.toml` on Github for used languages).
- Swapping the stuff from my closet and my bed-drawers.
- Write why I still use Xorg (nvidia) and explain how to set it up.
  - Explain that you NEVER use xrandr but use `.Xresources` to scale.
  - When specifing a cursor and size you get consistant cursor.
  - Don't forget to source the `.Xresources` file in the `.xinitrc` file.
- Remove dependency on Google Photos for my media.
  I want a platform independent solution to store my media.
  A Network Attached Storage (NAS) or home server solution should be future proof.
- Figure out how to make the text while logging in my laptop and computer high resolution and also readable.
  At the moment my computer font is way to big and blurry, and my laptop font size is way to small to read anything.
- Write a guide for my Gentoo setup (including WM).
  - Explain that for Vulkan needs to be used for rendering Sway for NVIDIA card.
- Explain that I need to use Vulkan for Sway using NVIDIA.
- Explain that I need `iwd` as [Networkmanager](/networkmanager.md) backend instead of `wpa_supplicant`.
- Backup my Gentoo portage configuration settings
- Backup my Linux kernel `.config` file.
- Make a seperate branch for pc and laptop for my `dotfiles` repository.
- Explain how to make audio work before you forget it
  - global pulseaudio flag
  - install pipewire and wireplumber
  - start your windowmanager using dbus-run-session or dbus-launch (otherwise your audio will probably not work)
  - use the gentoo-pipewire-launcher (see i3 config) on startup
  - use pavucontrol for easy managing of audio (it is not worth it for the terminal)
 
