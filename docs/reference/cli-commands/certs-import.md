---
title: "certs import"
description: "Import a certificate from public certificate and private key files."
---

# certs import

Import a certificate from public certificate and private key files. Optionally include an intermediate certificate chain for full chain validation. This is the recommended way to add production certificates to your Rack.

The certificate is imported into AWS Certificate Manager (ACM) in the Rack's region, and the command prints its id. A Gen 2 Service uses an ACM certificate with an RSA 1024-bit or 2048-bit key that `convox certs` lists when it covers every one of the Service's `domain:` values, from the App's next deploy. When more than one listed certificate matches, Convox uses the first match, so a previously imported certificate for the same domains can still be selected. For a Gen 1 App, apply it to an endpoint with [`convox ssl update`](/reference/cli-commands/ssl-update).

The command returns once ACM lists the new certificate, so the next command can use the printed id straight away. The wait adds a few seconds, about 13 seconds at most. Requires rack version 20261005214736 or newer.

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
Importing certificate... OK, acm-3f9d27c1b6a4
```

## Key Types

`convox certs` lists, and `convox certs delete` deletes, ACM certificates of every key type: RSA 1024, 2048, 3072 and 4096 bits, and ECDSA P-256, P-384 and P-521. Listing and deleting ECDSA and 3072-bit or 4096-bit RSA certificates requires rack version 20261005214736 or newer.

Where a certificate can be used depends on its key type:

| Use | Key types |
|:--|:--|
| Gen 2 Service `domain:` | RSA 1024, RSA 2048 |
| Gen 2 `nlb:` port with `protocol: tls`, referenced by full ARN | RSA 1024, 2048 and 3072; ECDSA P-256, P-384 and P-521 |
| Gen 1 [`convox ssl update`](/reference/cli-commands/ssl-update) on a Classic Load Balancer, the Gen 1 default | RSA 1024, RSA 2048 |
| Gen 1 `convox ssl update` on an Application Load Balancer | RSA 2048, 3072 and 4096; ECDSA P-256, P-384 and P-521 |

## See Also

- [certs generate](/reference/cli-commands/certs-generate)
- [certs](/reference/cli-commands/certs)
- [ssl update](/reference/cli-commands/ssl-update)
- [SSL](/deployment/ssl)
