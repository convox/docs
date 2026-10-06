---
title: "certs delete"
description: "Delete an ACM certificate or IAM server certificate by its id."
---

# certs delete

Delete an ACM certificate in the Rack's region or an IAM server certificate in the Rack's AWS account. Pass the full certificate id printed by `convox certs generate` or listed by `convox certs`. A certificate still waiting for validation is not listed; delete it with the id `convox certs generate` printed.

An id that starts with `acm-` deletes that ACM certificate, whatever its key type. An `acm-` id that matches no certificate, such as a partial id, `acm-` alone, or a certificate that was already deleted, returns `ERROR: certificate not found` and deletes nothing. Any other id, including a name such as `acme-2019`, deletes the IAM server certificate with that exact name. Deleting an ECDSA or a 3072-bit or 4096-bit RSA certificate requires rack version 20261005214736 or newer.

The certificate must not be in use by any App. For a Gen 1 App, apply a different certificate to the endpoint first with [`convox ssl update`](/reference/cli-commands/ssl-update). A Gen 2 App keeps using a certificate until the App is deployed without it, with no Service `domain:` selecting it and no `nlb:` port `certificate:` naming its ARN.

## Syntax

```bash
$ convox certs delete <cert>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox certs delete acm-74ec8b405093
Deleting certificate acm-74ec8b405093... OK
```

## See Also

- [certs](/reference/cli-commands/certs)
- [certs generate](/reference/cli-commands/certs-generate)
- [certs import](/reference/cli-commands/certs-import)
- [SSL](/deployment/ssl)
