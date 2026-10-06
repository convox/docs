---
title: "certs generate"
description: "Request an AWS Certificate Manager (ACM) certificate for one or more domains."
---

# certs generate

Request a certificate from AWS Certificate Manager (ACM) in the Rack's region. The first domain is the certificate's domain name and any other domains are added as subject alternative names. The command prints the new certificate's id.

Each run requests a new certificate, and each one must be validated separately. Keep the id this command prints: `convox certs` does not list the certificate until ACM issues it, and [`convox certs delete`](/reference/cli-commands/certs-delete) takes that id to remove a certificate you no longer need.

The command returns once ACM lists the new request, so `convox certs delete` accepts the printed id straight away. The wait adds a few seconds, about 13 seconds at most. Requires rack version 20261005214736 or newer.

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
Generating certificate... OK, acm-74ec8b405093
```

## See Also

- [certs import](/reference/cli-commands/certs-import)
- [certs](/reference/cli-commands/certs)
- [ssl](/reference/cli-commands/ssl)
- [SSL](/deployment/ssl)
