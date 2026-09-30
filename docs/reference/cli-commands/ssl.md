---
title: "ssl"
description: "List certificate associations for an App."
---

# ssl

List certificate associations for an App. Shows which certificates are bound to which Service ports, along with the associated domain and expiration date. For a Gen 2 App, only port 443 of each Service is shown, and a Service that uses a certificate Convox created for its `domain:` has no row, because `convox certs` does not list those certificates. Use this to verify that your endpoints are secured with the correct certificates.

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
ENDPOINT  CERTIFICATE                       DOMAIN                 EXPIRES
web:443   acm-3f9d27c1b6a4                  *.example.com          6 months from now
api:443   cert-production-1758802512-04821  *.*.elb.amazonaws.com  5 days ago
```

## See Also

- [ssl update](/reference/cli-commands/ssl-update)
- [certs](/reference/cli-commands/certs)
- [SSL](/deployment/ssl)
- [Custom Domains](/deployment/custom-domains)
