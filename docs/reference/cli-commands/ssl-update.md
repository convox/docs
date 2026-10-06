---
title: "ssl update"
description: "Update the certificate for an App endpoint."
---

# ssl update

Update the certificate for an App endpoint (generation 1 only). Associates a certificate with a specific Service port. Use this after importing a new certificate to bind it to your Service's HTTPS endpoint. The certificate id must be one that [`convox certs`](/reference/cli-commands/certs) lists. On rack version 20261005214736 or newer, any other id returns `ERROR: certificate not found` and changes nothing.

> **Note:** Gen 2 Apps get certificates from each Service's `domain:` setting, or from `certificate:` on `nlb:` ports, in `convox.yml`. For a Gen 2 App this command returns `ERROR: command not valid for generation 2 applications`.

## Syntax

```bash
$ convox ssl update <process:port> <certificate>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

## Example Usage

```bash
$ convox ssl update web:443 acm-3f9d27c1b6a4 -a myapp
Updating certificate... OK
```

## Key Types

A Gen 1 Service gets a Classic Load Balancer by default, and a Classic Load Balancer takes only RSA 1024-bit and 2048-bit certificates from ACM. An ACM certificate with any other key type returns an error that names the key type, and the App is not updated:

```bash
$ convox ssl update web:443 acm-8b2f4c71d0e9 -a myapp
Updating certificate... ERROR: certificate acm-8b2f4c71d0e9 is EC_prime256v1, which a Classic Load Balancer does not accept from ACM
```

A Service with an Application Load Balancer also takes ECDSA and 3072-bit or 4096-bit RSA certificates from ACM. Applying an ECDSA or a 3072-bit or 4096-bit RSA certificate requires rack version 20261005214736 or newer.

## See Also

- [ssl](/reference/cli-commands/ssl)
- [certs import](/reference/cli-commands/certs-import)
- [certs](/reference/cli-commands/certs)
- [SSL](/deployment/ssl)
- [Custom Domains](/deployment/custom-domains)
