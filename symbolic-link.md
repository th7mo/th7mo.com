A symbolic link, often called a 'symlink', is a file that serves as a reference or pointer to another file or directory.

With Unix based systems the `ln` command with the `-s` flag can be used to create symbolic link:
```sh
ln -s {target_path} {link_path}
```
* `{target_path}` is the path to the file or directory to link to
* `{link_path}` the desired path for the (new) symbolic link
