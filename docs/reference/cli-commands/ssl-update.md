---
title: "ssl update"
description: "Update the certificate for an App endpoint."
---

# ssl update

Update the certificate for an App endpoint (generation 1 only). Associates a certificate with a specific Service port. Use this after importing a new certificate to bind it to your Service's HTTPS endpoint.

> **Note:** Gen 2 Apps manage SSL automatically via the `domain` and `tls` settings in `convox.yml`. This command is only needed for Gen 1 Apps.

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
$ convox ssl update web:443 cert-ef34gh -a myapp
Updating certificate... OK
```

## See Also

- [ssl](/reference/cli-commands/ssl)
- [certs import](/reference/cli-commands/certs-import)
- [certs](/reference/cli-commands/certs)
- [SSL](/deployment/ssl)
- [Custom Domains](/deployment/custom-domains)
