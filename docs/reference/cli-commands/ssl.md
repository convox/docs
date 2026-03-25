---
title: "ssl"
description: "List certificate associations for an App."
---

# ssl

List certificate associations for an App. Shows which certificates are bound to which Service ports, along with the associated domain and expiration date. Use this to verify that your endpoints are secured with the correct certificates.

## Syntax

```bash
$ convox ssl
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox ssl -a myapp
ENDPOINT  CERTIFICATE   DOMAIN           EXPIRES
web:443   cert-ab12cd   example.com      3 months from now
api:443   cert-ef34gh   api.example.com  5 days ago
```

## See Also

- [ssl update](/reference/cli-commands/ssl-update)
- [certs](/reference/cli-commands/certs)
- [SSL](/deployment/ssl)
- [Custom Domains](/deployment/custom-domains)
