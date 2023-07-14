# README

## Overview

Calls `https://myip.supermicro.com/` to detect external IP address and then calls Namecheap Dynamic DNS REST API to update DNS record.

## Syntax

    namecheap-ddns [config_file]

- config_file is a YAML configuration file.   Default value is `namecheap-ddns.yaml`.  See, [example](https://github.com/Kulak/namecheap-ddns/blob/master/namecheap-ddns.yaml)

## Release

Release contains `portablectl` image.

For details see [portablectl documentation](https://www.freedesktop.org/software/systemd/man/portablectl.html)

### Image Deployment Example

```sh
sudo mkdir -p /usr/local/portable-images/
# download
```

## Build

The build file uses [taskfile.dev](https://taskfile.dev/).

## Namecheap Documentation

1. [How do I enable Dynamic DNS for a domain?](https://www.namecheap.com/support/knowledgebase/article.aspx/595/11/how-do-i-enable-dynamic-dns-for-a-domain)
2. [How do I set up a Host for Dynamic DNS?](https://www.namecheap.com/support/knowledgebase/article.aspx/43/11/how-do-i-set-up-a-host-for-dynamic-dns)
3. [How do I use a browser to dynamically update the host's IP?](https://www.namecheap.com/support/knowledgebase/article.aspx/29/11/how-do-i-use-a-browser-to-dynamically-update-the-hosts-ip)
