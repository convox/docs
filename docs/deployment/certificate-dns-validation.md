---
title: "Certificate DNS Validation"
description: "Switch a Rack's ACM certificates from email validation to DNS validation, with steps for Route 53, Cloudflare and other DNS providers."
---

# Certificate DNS Validation

AWS Certificate Manager (ACM) is ending email validation. On rack version 20261021234916 or newer, the Rack switches its ACM certificates to DNS validation in place. Each certificate keeps its ARN and expiry, nothing is reissued, and Apps keep serving HTTPS throughout. A certificate is switched only once its validation records answer in public DNS, so a certificate whose records are not in place yet stays on email validation and keeps working.

## AWS Timeline

| Date | Change |
|:-----|:-------|
| 2027-01-01 | Email validation is not offered in newly launched AWS Regions |
| 2027-03-31 | ACM stops accepting new email-validated certificate requests |
| 2027-09-30 | ACM stops renewing email-validated certificates |
| 2028-03-15 | Certificate authorities can no longer use email validation for any public certificate |

See the [AWS announcement](https://aws.amazon.com/blogs/security/aws-certificate-manager-will-discontinue-email-validation-to-prove-domain-validation-for-certificates/) for details.

## What Needs Action

| Certificate | Example | What you do |
|:------------|:--------|:------------|
| `convox.site` certificates for the Rack, its routers and `--wildcard-domain` Services | `*.myrack-router-123456789.us-east-1.convox.site` | Update the Rack. Convox publishes the records, and the Rack switches these certificates on its own a few minutes after the update |
| Certificates for your own domains: Service `domain:`, `convox certs generate`, Gen 1 `convox ssl update` | `www.example.com` | Update the Rack, add one DNS record per domain name, then the Rack switches the certificate |
| Certificates imported with `convox certs import`, and IAM server certificates | any | Nothing. They do not use email or DNS validation |
| ACM certificates in the region that the Rack does not use | any | Nothing on this Rack, which leaves them alone. `convox certs dns` names the ones still on email validation. Switch each one on the Rack that uses it, or in the [ACM console](#switching-in-the-acm-console) if something outside Convox uses it |

Certificates the Rack creates itself use email validation: the ones it creates when the Rack is installed, when [Internal](/reference/rack-parameters/Internal) is enabled, when a `--wildcard-domain` Service is created, and when a Service `domain:` that no issued certificate covers is deployed. The Rack switches each one to DNS validation once ACM issues it and its records answer: on its own for `convox.site` certificates, and once you add the records for a Service `domain:`, as in [Adding the Records](#adding-the-records). For a new custom domain, request a DNS-validated certificate before you deploy, as described in [Adding a New Custom Domain](#adding-a-new-custom-domain).

## Before You Start

| Requirement | Detail |
|:------------|:-------|
| Rack version | 20261021234916 or newer. Check with `convox rack`, update with `convox rack update` |
| CLI version | 3.25.12 or newer, or 20261021234916 or newer. Update with `sudo convox update`. Older CLIs do not have `convox certs dns` or `convox certs migrate` |
| HTTPS access to `certbot.convox.com` | The Rack publishes its `convox.site` records through Convox. If the Rack reaches the internet through [HttpProxy](/reference/rack-parameters/HttpProxy), the proxy must allow it |
| Public DNS lookups | The Rack checks each record against public DNS: `1.1.1.1`, `8.8.8.8` and the zone's nameservers, on port 53. If it cannot reach them, it uses the VPC's DNS resolver instead |
| IAM | The Rack's default IAM policies allow every call it makes. If an administrator restricts the Rack's `ApiRole` with a permissions boundary or a service control policy, it must allow `acm:UpdateCertificateOptions`, `acm:ListCertificateDomainValidations`, `acm:RequestCertificate`, `acm:AddTagsToCertificate` and `acm:DeleteCertificate`, plus the `route53` calls listed on [Route53CertificateValidation](/reference/rack-parameters/Route53CertificateValidation#permissions) if you enable that parameter |

## Migrating

### 1. Update the CLI and the Rack

```bash
$ sudo convox update
$ convox rack update --wait
```

A few minutes after the update finishes, the Rack switches its `convox.site` certificates to DNS validation. There is nothing else to do for them.

### 2. Check Certificate Status

```bash
$ convox certs
ID                                DOMAIN                                           EXPIRES            STATUS
acm-0a1b2c3d4e5f                  example.com                                      3 months from now  email
acm-1b2c3d4e5f6a                  shop.example.org                                 4 months from now  email
acm-89ea927329d7                  *.myrack-router-123456789.us-east-1.convox.site  5 months from now  dns

2 certificates use email validation, which stops renewing on 2027-09-30: run convox certs dns
```

`dns` means the certificate uses DNS validation, and ACM renews it while it is in use and its DNS record stays in place. `email` means it still needs a DNS record. See [certs](/reference/cli-commands/certs#status) for every status.

### 3. List the DNS Records

```bash
$ convox certs dns
ZONE         NAME                                                 TYPE   VALUE                                                              CHECK    CERTIFICATES
example.com  _0123456789abcdef0123456789abcdef.example.com.       CNAME  _fedcba9876543210fedcba9876543210.abcdefghij.acm-validations.aws.  missing  acm-0a1b2c3d4e5f
example.org  _1123456789abcdef0123456789abcdef.shop.example.org.  CNAME  _eedcba9876543210fedcba9876543210.bcdefghijk.acm-validations.aws.  missing  acm-1b2c3d4e5f6a

2 records: 2 missing
When every record shows ok: convox certs migrate --wait
```

Each record appears once, even when several certificates need it, and `example.com` and `*.example.com` share one record. Right after the Rack update, the command can print `N certificates are still being prepared, run again in a few minutes` while the Rack reads the records from ACM. `convox certs migrate` reads them at once.

### 4. Add the Records

Add each record at your DNS provider as a `CNAME`, with the name and value exactly as shown, and leave it in place permanently. ACM checks the record again at every renewal, and a certificate whose record is removed stops renewing. Any TTL works.

See [Adding the Records](#adding-the-records) for the steps for your DNS provider. Run `convox certs dns` again at any time: `CHECK` shows `ok` once a record answers.

### 5. Switch the Certificates

```bash
$ convox certs migrate --wait
CERTIFICATE       DOMAIN                                           RESULT
acm-0a1b2c3d4e5f  example.com                                      switching
acm-1b2c3d4e5f6a  shop.example.org                                 switching
acm-89ea927329d7  *.myrack-router-123456789.us-east-1.convox.site  dns
Waiting for switches to finish... OK
CERTIFICATE       DOMAIN            RESULT
acm-0a1b2c3d4e5f  example.com       dns
acm-1b2c3d4e5f6a  shop.example.org  dns
```

`convox certs migrate` switches only certificates whose records already answer, so it is safe to run at any time. A certificate still waiting on a record is listed as `waiting for its DNS record (run convox certs dns)` and stays as it was. If you skip this step, the Rack switches each certificate on its own within a few minutes of its record answering.

ACM can take up to 30 minutes to finish a switch, and `convox certs` shows the certificate as `migrating` until then. `--wait` waits up to 15 minutes. If it gives up with `ERROR: certificates did not finish switching`, the switch continues in the background; check `convox certs` again later. See [certs migrate](/reference/cli-commands/certs-migrate) for every result.

## Adding the Records

| Where your DNS is | Section |
|:------------------|:--------|
| Route 53, in the Rack's AWS account | [Route 53 in the Rack's AWS Account](#route-53-in-the-racks-aws-account) |
| Route 53, in another AWS account | [Route 53 in Another AWS Account](#route-53-in-another-aws-account) |
| Cloudflare | [Cloudflare](#cloudflare) |
| Any other provider | [Other DNS Providers](#other-dns-providers) |
| Managed by Terraform, external-dns or similar | [DNS Managed as Code](#dns-managed-as-code) |

To import records from a file, save them with `--zone-file`. It prints every record grouped by zone, with absolute names and a TTL of 300 seconds:

```bash
$ convox certs dns --zone-file > records.txt
$ cat records.txt
; zone example.com
_0123456789abcdef0123456789abcdef.example.com. 300 IN CNAME _fedcba9876543210fedcba9876543210.abcdefghij.acm-validations.aws.
; zone example.org
_1123456789abcdef0123456789abcdef.shop.example.org. 300 IN CNAME _eedcba9876543210fedcba9876543210.bcdefghijk.acm-validations.aws.
```

Import each zone's records into that zone only. To write one zone's records to their own file:

```bash
$ awk -v zone=example.org '/^; zone / { z = $3; next } z == zone' records.txt > example.org.txt
```

The file includes records that already show `ok`. Many DNS providers reject an import that contains a record the zone already has, so remove those lines first.

### Route 53 in the Rack's AWS Account

When `convox certs dns` prints `<zone> is a Route 53 zone in this AWS account`, the Rack can add the records itself:

```bash
$ convox rack params set Route53CertificateValidation=Yes --wait
$ convox certs migrate --wait
```

`convox certs migrate` writes the missing records and switches each certificate once its record answers. The first run usually shows `publishing` until Route 53 serves the new records, and the switch follows within a few minutes. Without the command, the Rack does the same in the background every few minutes. The parameter also covers certificates you request later with `convox certs generate`.

The Rack writes only into a public hosted zone in its own AWS account that public DNS delegates to. It only creates records, never changes a record that has another value, and never deletes a record. It writes records only for certificates not yet on DNS validation, so if a record for a `dns` certificate goes missing, add it back by hand. See [Route53CertificateValidation](/reference/rack-parameters/Route53CertificateValidation).

### Route 53 in Another AWS Account

The Rack cannot write to a hosted zone in another AWS account, so add the records with credentials for the account that holds the zone. In the Route 53 console, create each record in the hosted zone with **Record type** `CNAME`, **Record name** set to the part of `NAME` before the zone (for example `_1123456789abcdef0123456789abcdef.shop` in `example.org`), and **Value** set to `VALUE`.

With the AWS CLI, build a change batch from the zone file and apply it, once per zone. In this example `dns` is an AWS CLI profile for the account that holds the zone:

```bash
$ convox certs dns --zone-file > records.txt
$ aws route53 list-hosted-zones-by-name --dns-name example.org --query "HostedZones[?Name=='example.org.'].[Id,Config.PrivateZone]" --output text --profile dns
$ awk -v zone=example.org '
    BEGIN { printf "{\"Changes\":[" }
    /^; zone / { inzone = ($3 == zone); next }
    inzone { printf "%s{\"Action\":\"UPSERT\",\"ResourceRecordSet\":{\"Name\":\"%s\",\"Type\":\"CNAME\",\"TTL\":300,\"ResourceRecords\":[{\"Value\":\"%s\"}]}}", sep, $1, $5; sep = "," }
    END { print "]}" }' records.txt > example.org.json
$ aws route53 change-resource-record-sets --hosted-zone-id Z0123456789ABCDEFGHIJ --change-batch file://example.org.json --profile dns
```

Use the id of the zone whose second column is `False`, without the `/hostedzone/` prefix. A zone showing `True` is private, and ACM cannot see records in it. If more than one zone shows `False`, use the one whose name servers in `aws route53 get-hosted-zone --id <id> --profile dns` match `dig +short NS example.org`, because records in the other zone are never served. `UPSERT` creates each record or updates it to the same value, so the batch is safe to apply again.

### Cloudflare

Validation records must be **DNS only**. Cloudflare does not proxy CNAME records whose target is under `acm-validations.aws`, and a record that is not served as a plain CNAME never validates.

To import a zone's records:

1. Write the zone's records to their own file, as shown in [Adding the Records](#adding-the-records), and remove any line for a record that already shows `ok`.
2. In the Cloudflare dashboard, open the zone and go to **DNS** > **Records**.
3. Choose **Import and Export**, then **Import DNS records**, and select the file.
4. Clear **Proxy imported DNS records**, then import.

To add a record by hand, choose **Add record** with:

| Field | Value |
|:------|:------|
| Type | `CNAME` |
| Name | The part of `NAME` before the zone, for example `_1123456789abcdef0123456789abcdef.shop` in `example.org` |
| Target | `VALUE` |
| Proxy status | DNS only |
| TTL | Auto |

If **CNAME flattening** is on for every CNAME record in the zone, or **Flatten** is on for the record, Cloudflare stops returning the CNAME and `convox certs dns` shows `missing`. Turn flattening off for these records. On a partial (CNAME) setup, add the records at the zone's authoritative DNS provider instead of in Cloudflare.

### Other DNS Providers

The same rules apply at every provider:

| Field | Rule |
|:------|:-----|
| Type | `CNAME` |
| Name | Most dashboards add the zone to the name for you. Enter only the part of `NAME` before the zone, or the record ends up as `_0123456789abcdef0123456789abcdef.example.com.example.com` and shows `missing` |
| Value | `VALUE`, including its leading underscore. Remove the trailing dot only if the provider rejects it |
| TTL | Any. `300` matches the zone file |
| Proxying, forwarding, flattening | Off |

| Provider | Import from a zone file | Name field |
|:---------|:------------------------|:-----------|
| GoDaddy | **DNS** > **Actions** > **Import Zone File**. The import fails if a record already exists | Part before the zone |
| Namecheap | Not available. Add each record under **Advanced DNS** > **Host Records** | Part before the zone |
| Google Cloud DNS | `gcloud dns record-sets import example.org.txt --zone=<managed-zone> --zone-file-format`. Fails if a record already exists. Never add `--delete-all-existing`, which replaces every record in the zone | Full name with trailing dot in `gcloud` |
| Azure DNS | `az network dns zone import -g <resource-group> -n example.org -f example.org.txt` merges the records into the zone | Part before the zone, as in `az network dns record-set cname set-record -g <resource-group> -z example.org -n _1123456789abcdef0123456789abcdef.shop -c _eedcba9876543210fedcba9876543210.bcdefghijk.acm-validations.aws --ttl 300` |
| DigitalOcean | Not available. Add each record under **Networking** > **Domains** | Part before the zone |
| DNSimple | **DNS** > **Import**, then paste the file's lines | Part before the zone |

### DNS Managed as Code

If Terraform, external-dns, DNSControl or another tool manages the zone, add the records to that tool's configuration, or its next run can remove them and the certificates stop renewing.

### Delegated Subdomains and Private Zones

`ZONE` names the zone public DNS uses for the record. When a subdomain such as `shop.example.org` is delegated to its own nameservers, `ZONE` shows the subdomain, and the record belongs in that zone. ACM validates only through public DNS, so a record in a Route 53 private hosted zone or on an internal DNS server does not count. With split-horizon DNS, add the record to the public zone. If the Rack cannot reach public DNS on port 53 and the private zone is associated with the Rack's VPC, add the record to the private zone as well, because the Rack then checks records through the VPC's resolver.

## Checking a Record

`convox certs dns` asks the zone's authoritative nameservers, so it usually shows `ok` as soon as your DNS provider serves the record. See [certs dns](/reference/cli-commands/certs-dns#check-values) for when it checks through the VPC's resolver instead, and run `convox certs migrate` if a record you added still shows `missing`. To check from your own machine:

```bash
$ dig +short CNAME _1123456789abcdef0123456789abcdef.shop.example.org @1.1.1.1
_eedcba9876543210fedcba9876543210.bcdefghijk.acm-validations.aws.
```

On Windows, use `nslookup -type=CNAME <name> 1.1.1.1`. An empty answer means the record is not served yet, is proxied or flattened (both hide the CNAME), or a public resolver still holds an earlier "not found" answer for a while. If the record exists only at a name with the zone appended twice, such as `_1123456789abcdef0123456789abcdef.shop.example.org.example.org`, recreate it with the shorter name.

## Adding a New Custom Domain

A deploy uses an existing certificate only if ACM has issued it and it covers every domain the Service lists. Otherwise the deploy requests a new certificate with email validation. To use DNS validation for a new domain, request the certificate first:

1. Request a certificate with every domain the Service lists:

   ```bash
   $ convox certs generate example.com www.example.com
   Generating certificate... OK, acm-74ec8b405093
   ZONE         NAME                                                TYPE   VALUE                                                              CHECK
   example.com  _0123456789abcdef0123456789abcdef.example.com.      CNAME  _fedcba9876543210fedcba9876543210.abcdefghij.acm-validations.aws.  ok
   example.com  _3123456789abcdef0123456789abcdef.www.example.com.  CNAME  _ccdcba9876543210fedcba9876543210.defghijklm.acm-validations.aws.  missing
   Deploy the domain once convox certs shows acm-74ec8b405093 as dns
   ```

2. Add the records that show `missing` within 72 hours, or ACM stops waiting and the certificate shows `failed`. A record that shows `ok` is already in place from an earlier certificate. With [Route53CertificateValidation](/reference/rack-parameters/Route53CertificateValidation) set to `Yes`, the Rack adds them for zones in its AWS account.
3. Wait until `convox certs` shows the certificate as `dns`. ACM can take up to 30 minutes after the records answer.
4. Add the domains to the Service's `domain:` in `convox.yml` and deploy. The deploy uses the new certificate.

A domain that is already deployed does not need a new certificate. The migration switches the certificate it already uses.

## Several Racks and AWS Accounts

Run the same commands against each Rack with `-r <rack>`. A validation record belongs to one domain name in one AWS account and works in every region, so Racks in the same AWS account share records: once a record is in place, every certificate for that name in that account uses it. A Rack in another AWS account gets a different record for the same name, and both records can exist side by side.

## Gen 1 Apps

Certificates applied to Gen 1 Apps with `convox ssl update` are switched like any other certificate the Rack uses. For a new certificate, run `convox certs generate`, add the records it prints, wait until `convox certs` shows the certificate as `dns`, then apply it:

```bash
$ convox ssl update web:443 acm-74ec8b405093 -a myapp
Updating certificate... OK
```

`convox ssl update` accepts only an issued certificate. A certificate that is still `pending` returns `ERROR: certificate not found`.

## Logs

| Command | Line | When |
|:--------|:-----|:-----|
| `convox rack logs` | `N certificates need DNS records: run convox certs dns` | Every hour while certificates the Rack uses for your own domains are waiting for a record |
| `convox rack logs` | `certificate <id> (<domain>) switched to DNS validation` | When the Rack switches a certificate in the background. A switch made by `convox certs migrate` is not logged |
| `convox logs -a <app>` | `certificate <id> (<domain>) needs a DNS record before 2027-09-30: run convox certs dns` | Every hour, for a certificate Convox created for one of the App's Services that is waiting for a record |
| `convox logs -a <app>` | `certificate <id> (<domain>) switched to DNS validation` | When the Rack switches, in the background, a certificate Convox created for one of the App's Services |

## Troubleshooting

| Symptom | Cause and fix |
|:--------|:--------------|
| `CHECK` shows `missing` | The record is not served yet, is in another zone than `ZONE`, has the zone appended twice, is proxied or flattened, or exists only in a private zone. Check it as shown in [Checking a Record](#checking-a-record) |
| `CHECK` shows `wrong` | Another CNAME answers for the name. Usually a wildcard record on a parent name, such as `*.example.com`, answers because the validation record does not exist yet: add the validation record, which takes precedence over the wildcard. Otherwise the record has an old value, or its value lost the leading underscore. Correct it to `VALUE` |
| `CHECK` shows `unknown` | The lookup failed or did not finish. Run `convox certs dns` again. If it persists, the zone's nameservers are not answering the Rack |
| `N certificates are still being prepared` | The Rack has not read these certificates' records yet. Run `convox certs migrate`, or wait a few minutes |
| `<id>: <error>` lines in `convox certs dns` | The Rack could not read that certificate or its records. Run `convox certs migrate` to try again |
| A certificate from `convox certs generate` shows `failed` | ACM did not see its records within 72 hours, or a CAA record blocked it (next row). Add the records, run `convox certs generate` again, and remove the failed certificate with `convox certs delete` |
| `error: validation failed: CAA_ERROR` | A CAA record on the domain does not allow Amazon to issue certificates. Add a CAA record that allows `amazon.com`, such as `example.com. CAA 0 issue "amazon.com"`, and an `issuewild` record for wildcard certificates if the domain has `issuewild` records. Then run `convox certs migrate` |
| `convox certs` reports `convox.site` certificates still on email validation | Run `convox certs migrate` and read the result for those certificates. `error: Post "https://certbot.convox.com/validation": ...` means the Rack cannot reach Convox: allow outbound HTTPS from the Rack, or through its HttpProxy. A message that ends in `contact Convox support` needs Convox to act |
| `ERROR: certificates did not finish switching` from `--wait` | ACM is still working, or a certificate ended with another result. Check `convox certs` later, and run `convox certs migrate` to see any certificate that is not `dns` or `migrating` |
| `error: not used by this rack` | The named certificate is not one this Rack uses. Run the command against the Rack that uses it |
| `ERROR: convox certs dns requires a rack on the latest release` | The Rack is older than 20261021234916. Run `convox rack update` |
| Rack logs show `record <name> exists with another value, left alone` | With `Route53CertificateValidation=Yes`, a record with that name has another value. Correct it in Route 53 |
| Rack logs show `route53 write <name>: <error>` | The Rack could not write the record, for example because IAM denied the call. See [Route53CertificateValidation](/reference/rack-parameters/Route53CertificateValidation#permissions) |

## Switching in the ACM Console

The ACM console can also switch an email-validated certificate: open the certificate, choose **Update validation method**, and add the records it shows. The Rack recognizes a certificate switched this way and shows it as `migrating`, then `dns`. If the records are not in place within 72 hours, the switch lapses, and the certificate stays on email validation and keeps working.

## Downgrading

Certificates already switched stay on DNS validation and keep renewing on an older rack version. An older version does not declare [Route53CertificateValidation](/reference/rack-parameters/Route53CertificateValidation), so the downgrade drops it; set it again after you update back. Against an older Rack, `convox certs generate` requests email validation and prints no records.

## See Also

- [SSL](/deployment/ssl)
- [certs dns](/reference/cli-commands/certs-dns)
- [certs migrate](/reference/cli-commands/certs-migrate)
- [Route53CertificateValidation](/reference/rack-parameters/Route53CertificateValidation)
