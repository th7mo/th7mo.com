---
title: "NGINX"
---

Websites can be hosted with [NGINX](https://www.nginx.com/).

## Installation
NGINX can be installed by executing the following command:

```sh
sudo apt install nginx
```

## Configuration
NGINX configuration files live inside the `/etc/nginx/` directory.
The two main subdirectories in there are `/etc/nginx/sites-available/` and `/etc/nginx/sites-enabled/`.
The idea is that you can make a site configuration file in `sites-available` and when it's ready,
you make a [symbolic link](symbolic-link.html) to it in `sites-enabled/` which will activate it.

Make a configuration file for each domain.
It is recommended to name the configuration file `{domain-name-without-top-level-domain}.conf`.
For this website the configuration file could be `th7mo.conf`.
A minimal configuration looks like this:

```
server {
    server_name th7mo.com ;
    root /var/www/th7mo ;
    index index.html ;

    location / {
        try_files $uri $uri/ =404 ;
    }

    listen 80 ;
    listen [::]:80 ;
}
```

* `server_name` is the domain and base path of the website.
* `root` is the root of the project files. This directory should contain the website source code.
* `index` specifies the file that represents the landing page of the website (when navigating to the address specified in `server_name`)
* `location /` is a block that specifies NGINX how to look up files. The line below specifies to look in the `root` directory and throw a [404 Not Found](404-not-found.html) error if failing to find the default file specified at `index`.
* `listen 80 ;` specifies NGINX to listen for connections at port 80 for IPv4.
* `listen [::]:80 ;` specifies the same but for IPv6.

All these components are required for the most minimal configuration.
Without the `listen` for **both IP protocols** certbot **can not** generate SSL certificates.

### Redirecting subdomains
One of the most common use cases for redirecting subdomains is the `www` subdomain. Firstly make sure proper DNS configuration is set up for the subdomain. The `if` statement is used to check and to redirect subdomains:

```
server {
    server_name th7mo.com www.th7mo.com ;
    if ($host = www.th7mo.com) {
        return 301 https://th7mo.com$request_uri;
    }

    # all other configuration below ...
}
```

In the example above both `th7mo.com` and `www.th7mo.com` are specified in the `server_name` field.
In the `if` statement we check and redirect if the visitor of the website visits `www.th7mo.com` to `https://th7mo.com` with the path that comes after it.

### Link the configuration
When to configuration is complete, make a symbolic link to the `/etc/nginx/sites-enabled/` directory:

```sh
ln -s /etc/nginx/sites-available/{config-file} /etc/nginx/sites-enabled/
```

Only when a symbolic link is present in the `sites-enabled/` directory NGINX actually serves the website to users.

## Load changes
Every time the configuration for NGINX changes, NGINX needs to be restarted for changes to take effect. Run `nginx -t` to verify if the configuration file is correct according to NGINX. If the configuration is correct, execute the following command to restart NGINX:

```
sudo systemctl restart nginx
```

## See also
* Most of the information here I learned from Luke Smith.
  He explains how to set up NGINX at his [landchad.net](https://landchad.net/basic/nginx/) website.
