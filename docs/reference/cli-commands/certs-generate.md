---
title: "certs generate"
description: "Request an AWS Certificate Manager (ACM) certificate for one or more domains."
---

# certs generate

Request a certificate from AWS Certificate Manager (ACM) in the Rack's region. The first domain is the certificate's domain name and any other domains are added as subject alternative names. The command prints the new certificate's id.

The certificate uses DNS validation. After the id, the command prints the DNS records the certificate needs and the line `Deploy the domain once convox certs shows <id> as dns`. No validation email is sent. Add each record at your DNS provider and leave it in place, because ACM checks the records again at every renewal. A record that already shows `ok`, such as one added for an earlier certificate for the same domain, needs nothing more. ACM issues the certificate once every record answers, which can take up to 30 minutes, and [`convox certs`](/reference/cli-commands/certs) then shows it as `dns`. If the records are not available yet, the command prints `Run convox certs dns in a minute to see the records to add`. With [Route53CertificateValidation](/reference/rack-parameters/Route53CertificateValidation) set to `Yes`, the Rack adds the records itself for domains whose hosted zone is in its AWS account.

DNS validation requires rack version 20261021234916 or newer, and CLI version 3.25.12 or newer, or 20261021234916 or newer. With an older CLI, or against an older Rack, the certificate uses email validation: ACM sends a validation email for each domain, and the command prints only the id.

To use the certificate for a Service, pass every domain the Service lists, add the records, wait until `convox certs` shows the certificate as `dns`, then deploy. A deploy uses only an issued certificate that covers all of the Service's domains. If none exists, the deploy requests a new certificate with email validation. See [Certificate DNS Validation](/deployment/certificate-dns-validation#adding-a-new-custom-domain).

Each run requests a new certificate, and each one must be validated separately. `convox certs` lists the new certificate with the status `pending`, and [`convox certs delete`](/reference/cli-commands/certs-delete) removes a certificate you no longer need. On a rack version older than 20261021234916, `convox certs` does not list the certificate until ACM issues it, so keep the id this command prints.

The command returns once ACM lists the new request, so `convox certs delete` accepts the printed id straight away. Requires rack version 20261005214736 or newer. Waiting for the listing adds up to about 13 seconds. With DNS validation, the Rack also waits up to 20 seconds for ACM to produce the records, and the CLI then reads them the way `convox certs dns` does, which takes longer in an account with many certificates.

## Syntax

```bash
$ convox certs generate <domain> [domain...]
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--id` | | Send logs and the DNS records to stderr, certificate ID to stdout |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox certs generate example.com www.example.com
Generating certificate... OK, acm-74ec8b405093
ZONE         NAME                                                TYPE   VALUE                                                              CHECK
example.com  _0123456789abcdef0123456789abcdef.example.com.      CNAME  _fedcba9876543210fedcba9876543210.abcdefghij.acm-validations.aws.  ok
example.com  _3123456789abcdef0123456789abcdef.www.example.com.  CNAME  _ccdcba9876543210fedcba9876543210.defghijklm.acm-validations.aws.  missing
Deploy the domain once convox certs shows acm-74ec8b405093 as dns
```

## See Also

- [certs dns](/reference/cli-commands/certs-dns)
- [certs import](/reference/cli-commands/certs-import)
- [certs](/reference/cli-commands/certs)
- [SSL](/deployment/ssl)
