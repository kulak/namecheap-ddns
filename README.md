# README

## Overview

Calls `https://myip.supermicro.com/` to detect external IP address and then calls Namecheap Dynamic DNS REST API to update DNS record.

Released image is based on `portablectl` image. See references section below for details.

## Usage

### Application Syntax

    namecheap-ddns [config_file]

- config_file is a YAML configuration file.   Default value is `namecheap-ddns.yaml`.  See, [example](https://github.com/Kulak/namecheap-ddns/blob/master/namecheap-ddns.yaml)

### Portable Image Deployment Example

All of the following commands require `sudo` prepended if running from non-root account.

```sh
mkdir -p /usr/local/portable-images/
cd /usr/local/portable-images
curl -LO https://github.com/kulak/namecheap-ddns/releases/download/1.2/namecheap-ddns.raw
mkdir /etc/namecheap-ddns

dd of=/etc/namecheap-ddns/namecheap-ddns.yaml
# paste your file configuration
hosts: 
  - hostname1
  - hostname2
domain: domain.com
password: your_password
# Use Ctrl-D to end input stream

portablectl attach /usr/local/portable-images/namecheap-ddns.raw

# run once
systemctl start namecheap-ddns.service

# enable and run timer 
systemctl enable --now namecheap-ddns.timer

# monitor service
journalctl -u namecheap-ddns
```

## Build

The build file uses [taskfile.dev](https://taskfile.dev/).

## Refernces

### Namecheap Documentation

1. [How do I enable Dynamic DNS for a domain?](https://www.namecheap.com/support/knowledgebase/article.aspx/595/11/how-do-i-enable-dynamic-dns-for-a-domain)
2. [How do I set up a Host for Dynamic DNS?](https://www.namecheap.com/support/knowledgebase/article.aspx/43/11/how-do-i-set-up-a-host-for-dynamic-dns)
3. [How do I use a browser to dynamically update the host's IP?](https://www.namecheap.com/support/knowledgebase/article.aspx/29/11/how-do-i-use-a-browser-to-dynamically-update-the-hosts-ip)

### Portable Image

1. [Portable Services Introduction](https://systemd.io/PORTABLE_SERVICES/)
2. [portablectl manpage](https://www.freedesktop.org/software/systemd/man/portablectl.html)
