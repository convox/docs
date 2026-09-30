---
title: "certs import"
description: "Import a certificate from public certificate and private key files."
---

# certs import

Import a certificate from public certificate and private key files. Optionally include an intermediate certificate chain for full chain validation. This is the recommended way to add production certificates to your Rack.

The certificate is imported into AWS Certificate Manager (ACM) in the Rack's region, and the command prints its id. A Gen 2 Service uses an ACM certificate that `convox certs` lists when it covers every one of the Service's `domain:` values, from the App's next deploy. When more than one listed certificate matches, Convox uses the first match, so a previously imported certificate for the same domains can still be selected. For a Gen 1 App, apply it to an endpoint with [`convox ssl update`](/reference/cli-commands/ssl-update).

A certificate with an ECDSA key, or with a 3072-bit or 4096-bit RSA key, imports and prints an id, but `convox certs` does not list it and `convox certs delete` returns `ERROR: certificate not found` for it. List and delete those certificates in ACM directly. A Service `domain:` never selects one and `convox ssl update` cannot apply one, but an `nlb:` port with `protocol: tls` can reference one by its full ARN.

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

## See Also

- [certs generate](/reference/cli-commands/certs-generate)
- [certs](/reference/cli-commands/certs)
- [ssl update](/reference/cli-commands/ssl-update)
- [SSL](/deployment/ssl)
