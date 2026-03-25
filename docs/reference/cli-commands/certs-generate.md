---
title: "certs generate"
description: "Generate a self-signed certificate for one or more domains."
---

# certs generate

Generate a self-signed certificate for one or more domains. Self-signed certificates are useful for development and testing. For production use, import a certificate from a trusted authority using `certs import`.

## Syntax

```bash
$ convox certs generate <domain> [domain...]
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--id` | | Send logs to stderr, certificate ID to stdout |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox certs generate example.com
Generating certificate... OK, cert-ab12cd
```

## See Also

- [certs import](/reference/cli-commands/certs-import)
- [certs](/reference/cli-commands/certs)
- [ssl](/reference/cli-commands/ssl)
- [SSL](/deployment/ssl)
