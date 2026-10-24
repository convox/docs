---
title: "certs"
description: "List the certificates available to the Rack, with each certificate's validation status."
---

# certs

List the certificates available to the Rack. Shows each certificate's ID, domain, expiration date and status. Use this to audit which certificates are available before associating them with App endpoints, and to find certificates that still use email validation.

The list covers IAM server certificates in the Rack's AWS account, listed by name, and ACM certificates in the Rack's region, listed as `acm-` followed by the last 12 characters of the certificate ARN. It is not limited to certificates this Rack created. ACM certificates that are issued, waiting for validation, failed, or whose validation timed out are listed, including certificates Convox created for Service domains. Expired, revoked and inactive ACM certificates are not shown. `EXPIRES` is blank for a certificate ACM has not issued. On a rack version older than 20261021234916, only issued ACM certificates are listed, and certificates Convox created for Service domains are not.

ACM certificates of every key type are listed: RSA 1024, 2048, 3072 and 4096 bits, and ECDSA P-256, P-384 and P-521. Listing ECDSA and 3072-bit or 4096-bit RSA certificates requires rack version 20261005214736 or newer; older racks list only RSA 1024-bit and 2048-bit ACM certificates. A Service `domain:` uses only RSA 1024-bit and 2048-bit certificates; see [certs import](/reference/cli-commands/certs-import#key-types) for the key types each load balancer accepts.

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
ID                                DOMAIN                                           EXPIRES            STATUS
acm-0a1b2c3d4e5f                  example.com                                      3 months from now  email
acm-1b2c3d4e5f6a                  shop.example.org                                 4 months from now  email
acm-89ea927329d7                  *.myrack-router-123456789.us-east-1.convox.site  5 months from now  dns
cert-production-1758802512-04821  *.*.elb.amazonaws.com                            5 days ago         imported

2 certificates use email validation, which stops renewing on 2027-09-30: run convox certs dns
```

## Status

| Status | Meaning |
|:-------|:--------|
| `dns` | An ACM certificate that uses DNS validation. ACM renews it while it is in use and its DNS records stay in place |
| `email` | An ACM certificate that uses email validation. ACM stops renewing email-validated certificates on 2027-09-30. See [Certificate DNS Validation](/deployment/certificate-dns-validation) |
| `migrating` | Switching from email to DNS validation. The certificate keeps working while ACM finishes |
| `pending` | Waiting for validation. ACM has not issued it yet |
| `failed` | ACM could not issue the certificate, or its validation timed out |
| `imported` | An IAM server certificate, a certificate imported into ACM, or an ACM private certificate. These do not use email or DNS validation, and ACM does not renew IAM or imported certificates |

When certificates the Rack uses are still on email validation, lines after the table name the next command:

```text
2 certificates use email validation, which stops renewing on 2027-09-30: run convox certs dns
1 convox.site certificate is still on email validation: run convox certs migrate to see why
```

These lines count only certificates the Rack uses. [`convox certs dns`](/reference/cli-commands/certs-dns) names email-validated certificates in the region that the Rack does not use.

The status values require rack version 20261021234916 or newer. A date-versioned CLI shows the `STATUS` column from version 20261021234916 and leaves it empty against an older Rack. A 3.x CLI labels the column `Status` and shows `Ready` for every certificate against an older Rack. The lines after the table require CLI version 3.25.12 or newer, or 20261021234916 or newer.

## See Also

- [certs dns](/reference/cli-commands/certs-dns)
- [certs migrate](/reference/cli-commands/certs-migrate)
- [certs generate](/reference/cli-commands/certs-generate)
- [Certificate DNS Validation](/deployment/certificate-dns-validation)
- [SSL](/deployment/ssl)
