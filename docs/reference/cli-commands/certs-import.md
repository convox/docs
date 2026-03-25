---
title: "certs import"
description: "Import a certificate from public certificate and private key files."
---

# certs import

Import a certificate from public certificate and private key files. Optionally include an intermediate certificate chain for full chain validation. This is the recommended way to add production certificates to your Rack.

## Syntax

```bash
$ convox certs import <pub> <key>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--chain` | | Intermediate certificate chain file |
| `--id` | | Send logs to stderr, certificate ID to stdout |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox certs import cert.pem key.pem --chain chain.pem
Importing certificate... OK, cert-ef34gh
```

## See Also

- [certs generate](/reference/cli-commands/certs-generate)
- [certs](/reference/cli-commands/certs)
- [ssl update](/reference/cli-commands/ssl-update)
- [SSL](/deployment/ssl)
