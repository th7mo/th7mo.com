[tmux](https://github.com/tmux/tmux) is a terminal multiplexer.
It is used to manage terminal windows and sessions.

Tmux can be configured in the `~/.tmux.conf` file.
Tmux can be launched when starting a terminal application by `exec tmux` inside the `~/.bashrc` file.

The tmux configuration can be reloaded by executing the following command inside tmux:

```sh
:source-file ~/.tmux.conf
```
