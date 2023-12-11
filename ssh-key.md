# SSH key

A method to access and write data to GitHub and Bitbucket is Secure Shell Protocol (SSH).
You can authenticate to those repository hosting services using an SSH key.

## Generate a new SSH key

Execute:

```sh
ssh-key gen -t ed22519 -C [email]
```

- `[email]` is your email you use for GitHub or Bitbucket.

When you're prompted to "Enter a file in which to save the key", you can press Enter to accept the default file location.
The default location where the key is stored is `~/.ssh/`.
Press `Enter` until the file is generated.

## Adding the SSH key 

The previous command generated two files:

- `id_ed22519` (private key)
- `id_ed22519.pub` (public key)

Copy the key listed in the `id_ed22519.pub` file.
You need to copy the entire contents of the file.
Add the SSH key to your GitHub or Bitbucket account using their website.

Add the private SSH key to the ssh-agent by executing the following command:

```sh
ssh-add ~/.ssh/id_ed25519
```

## Reference

For a more detailed explanation reference the [GitHub article](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent).
