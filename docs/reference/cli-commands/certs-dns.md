---
title: "certs dns"
description: "List the DNS validation records the Rack's ACM certificates need, with a live check of each record."
---

# certs dns

List the DNS records that the Rack's ACM certificates need for DNS validation, one row per record, with a live check of each one. Add these records at your DNS provider and leave them in place, because ACM checks them again at every renewal. See [Certificate DNS Validation](/deployment/certificate-dns-validation) for the full migration.

The list covers certificates for your own domains that the Rack uses: certificates Convox created for Service domains, certificates `convox certs generate` requested with DNS validation, certificates on the Rack's load balancers, and certificates applied to Gen 1 Apps with `convox ssl update`. It includes certificates on email validation, certificates switching to DNS validation, and certificates already on DNS validation, so records that are already in place stay listed. Each record appears once, even when several certificates need it, and `example.com` and `*.example.com` share one record.

Records for `convox.site` certificates are not listed, because Convox publishes them. Imported certificates, failed certificates and certificates still waiting for email validation have no records.

Requires rack version 20261021234916 or newer, and CLI version 3.25.12 or newer, or 20261021234916 or newer. Against an older Rack the command returns `ERROR: convox certs dns requires a rack on the latest release`.

## Syntax

```bash
$ convox certs dns
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |
| `--zone-file` | | Print the records as zone-file lines |

## Example Usage

```bash
$ convox certs dns
ZONE         NAME                                                 TYPE   VALUE                                                              CHECK    CERTIFICATES
example.com  _0123456789abcdef0123456789abcdef.example.com.       CNAME  _fedcba9876543210fedcba9876543210.abcdefghij.acm-validations.aws.  ok       acm-0a1b2c3d4e5f
example.org  _1123456789abcdef0123456789abcdef.shop.example.org.  CNAME  _eedcba9876543210fedcba9876543210.bcdefghijk.acm-validations.aws.  missing  acm-1b2c3d4e5f6a

2 records: 1 ok, 1 missing
example.org is a Route 53 zone in this AWS account. To have the rack add its records: convox rack params set Route53CertificateValidation=Yes
When every record shows ok: convox certs migrate --wait
```

## Columns

| Column | Content |
|:-------|:--------|
| `ZONE` | The DNS zone the record goes in: the closest parent of the record name that has its own NS records in public DNS. A delegated subdomain shows as its own zone |
| `NAME` | The record name, in full, with a trailing dot |
| `TYPE` | `CNAME` |
| `VALUE` | The record value |
| `CHECK` | The result of a live lookup of the record |
| `CERTIFICATES` | The certificates that use the record |

## Check Values

| Value | Meaning |
|:------|:--------|
| `ok` | The record answers with the expected value |
| `missing` | Nothing answers for the name, or the name has no CNAME record |
| `wrong` | A CNAME answers with a different value. A wildcard record on a parent name, such as `*.example.com`, often causes this |
| `unknown` | The lookup failed or did not finish in time. Run the command again |

The Rack asks the zone's authoritative nameservers in public DNS, so a record usually shows `ok` as soon as your DNS provider serves it, without waiting for resolver caches. When those nameservers answer, a record that only a private hosted zone or an internal DNS server answers stays `missing`, because ACM validates through public DNS. When the Rack cannot reach public DNS on port 53, when no nameserver answers, or once the command has run for about 25 seconds (as it can in a region with many ACM certificates), it checks the remaining records through the VPC's resolver instead. A record that only a private hosted zone or an internal DNS server answers can then show `ok`, and a record you just added can show `missing` until the resolver's cached answer expires. `convox certs migrate` always asks the nameservers when it can reach them.

## Other Output

After the table, a summary counts the records by check value. These lines can follow it:

| Line | Meaning |
|:-----|:--------|
| `<zone> is a Route 53 zone in this AWS account. To have the rack add its records: convox rack params set Route53CertificateValidation=Yes` | The zone is a public hosted zone in the Rack's AWS account. See [Route53CertificateValidation](/reference/rack-parameters/Route53CertificateValidation). Not printed once the parameter is `Yes` |
| `N certificates are still being prepared, run again in a few minutes` | The Rack has not read these certificates' records from ACM yet. It does this in the background every few minutes, and [`convox certs migrate`](/reference/cli-commands/certs-migrate) does it at once |
| `<id>: <error>` | The records for that certificate could not be read. `convox certs migrate` tries again |
| `<id> uses email validation but is not used by this rack, so it leaves it alone` | An email-validated certificate in the Rack's region that this Rack does not use. The Rack does not switch it. Run the command against the Rack that uses it, or delete it if nothing does |
| `When every record shows ok: convox certs migrate --wait` | Printed while any record is not `ok` |

## Zone File

`--zone-file` prints every record, including records that already show `ok`, as zone-file lines grouped by zone. Names and values are absolute, with a trailing dot, and every record has a TTL of 300 seconds. Lines about certificates that are still being prepared, or whose records could not be read, go to stderr, so redirecting the output captures only the records:

```bash
$ convox certs dns --zone-file > records.txt
$ cat records.txt
; zone example.com
_0123456789abcdef0123456789abcdef.example.com. 300 IN CNAME _fedcba9876543210fedcba9876543210.abcdefghij.acm-validations.aws.
; zone example.org
_1123456789abcdef0123456789abcdef.shop.example.org. 300 IN CNAME _eedcba9876543210fedcba9876543210.bcdefghijk.acm-validations.aws.
```

Import each `; zone` block into the zone it names. Many DNS providers reject an import that contains a record the zone already has, so remove the lines for records that already show `ok`. See [Adding the Records](/deployment/certificate-dns-validation#adding-the-records) for provider-specific steps.

## See Also

- [certs migrate](/reference/cli-commands/certs-migrate)
- [certs](/reference/cli-commands/certs)
- [Certificate DNS Validation](/deployment/certificate-dns-validation)
- [Route53CertificateValidation](/reference/rack-parameters/Route53CertificateValidation)
