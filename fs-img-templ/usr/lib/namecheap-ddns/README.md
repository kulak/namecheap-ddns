# README

Files in this directory provided trusted certificates for used websites:

- gdig2.pem is from Go Daddy and it is used to authenticate <https://myip.supermicro.com/> calls for external IP resolution.
- CloudflareIncECCCA-3.pem is from Cloudflare and it is used to authenticate Namecheap DDNS update calls.

## From CRT to PEM

Original certificate files were downloded with crt extension and had to be converted with commands:

```sh
openssl x509 -in gdig2.crt -out gdig2.pem
openssl x509 -in CloudflareIncECCCA-3.crt -out CloudflareIncECCCA-3.pem
```
