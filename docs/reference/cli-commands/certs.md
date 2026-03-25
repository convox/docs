---
title: "certs"
description: "List certificates on the Rack."
---

# certs

List certificates on the Rack. Shows each certificate's ID, domain, and expiration date. Use this to audit which certificates are available before associating them with App endpoints.

## Syntax

```bash
$ convox certs
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox certs
ID            DOMAIN              EXPIRES
cert-ab12cd   example.com         3 months from now
cert-ef34gh   *.example.com       6 months from now
cert-ij56kl   api.example.com     5 days ago
```

## See Also

- [certs generate](/reference/cli-commands/certs-generate)
- [certs import](/reference/cli-commands/certs-import)
- [certs delete](/reference/cli-commands/certs-delete)
- [ssl](/reference/cli-commands/ssl)
- [SSL](/deployment/ssl)
