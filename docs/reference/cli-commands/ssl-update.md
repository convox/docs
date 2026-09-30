---
title: "ssl update"
description: "Update the certificate for an App endpoint."
---

# ssl update

Update the certificate for an App endpoint (generation 1 only). Associates a certificate with a specific Service port. Use this after importing a new certificate to bind it to your Service's HTTPS endpoint. The certificate id must be one that [`convox certs`](/reference/cli-commands/certs) lists.

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

## See Also

- [ssl](/reference/cli-commands/ssl)
- [certs import](/reference/cli-commands/certs-import)
- [certs](/reference/cli-commands/certs)
- [SSL](/deployment/ssl)
- [Custom Domains](/deployment/custom-domains)
