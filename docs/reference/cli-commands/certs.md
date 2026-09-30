---
title: "certs"
description: "List the certificates available to the Rack."
---

# certs

List the certificates available to the Rack. Shows each certificate's ID, domain, and expiration date. Use this to audit which certificates are available before associating them with App endpoints.

The list covers IAM server certificates in the Rack's AWS account, listed by name, and ACM certificates in the Rack's region, listed as `acm-` followed by the last 12 characters of the certificate ARN. It is not limited to certificates this Rack created. Only ACM certificates with the status `ISSUED` and an RSA 1024-bit or 2048-bit key are listed, so a certificate still waiting for validation does not appear until ACM issues it, an expired ACM certificate is not shown, and a certificate with an ECDSA or larger RSA key is never shown. Certificates Convox creates for Service domains are not listed.

An IAM certificate named `cert-<rack>-<unix time>-<number>` is the self-signed certificate the Rack uploads for Gen 1 `https` and `tls` ports that have no certificate applied.

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
ID                                DOMAIN                 EXPIRES
acm-3f9d27c1b6a4                  *.example.com          6 months from now
acm-74ec8b405093                  example.com            1 year from now
cert-production-1758802512-04821  *.*.elb.amazonaws.com  5 days ago
```

## See Also

- [certs generate](/reference/cli-commands/certs-generate)
- [certs import](/reference/cli-commands/certs-import)
- [certs delete](/reference/cli-commands/certs-delete)
- [ssl](/reference/cli-commands/ssl)
- [SSL](/deployment/ssl)
