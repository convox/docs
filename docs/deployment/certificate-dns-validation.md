---
title: "Certificate DNS Validation"
description: "Switch a Rack's ACM certificates from email validation to DNS validation, with steps for Route 53, Cloudflare and other DNS providers."
---

# Certificate DNS Validation

AWS Certificate Manager (ACM) is ending email validation, where ACM sends a validation email for each domain in a certificate. A certificate that still uses email validation is not renewed after September 30, 2027, and expires at the end of the term it holds on that date. Its replacement is DNS validation: each domain name gets one DNS record, and ACM checks that record every time it renews the certificate.

Rack version 20261021234916 and newer switch your Rack's ACM certificates to DNS validation without replacing them. Each certificate keeps its ARN and expiry date, and your Apps keep serving HTTPS during the switch. The Rack switches a certificate only after its DNS record is in place, so a certificate whose record you have not added yet keeps working on email validation.

In short:

1. Update the CLI and the Rack. Certificates for `convox.site` hostnames, which Convox assigns, need nothing more.
2. Run `convox certs` to see which certificates still use email validation.
3. Run `convox certs dns` to get the DNS records.
4. Add the records at your DNS provider.
5. Run `convox certs migrate --wait` to switch the certificates now. This step is optional: without it, the Rack switches each certificate within a few minutes of its record being in place.

## AWS Timeline

| Date | Change |
|:-----|:-------|
| January 1, 2027 | Email validation is not offered in newly launched AWS Regions |
| March 31, 2027 | ACM stops accepting new email-validated certificate requests |
| September 30, 2027 | ACM stops renewing email-validated certificates |
| March 15, 2028 | Certificate authorities can no longer use email validation for any public certificate |

See the [AWS announcement](https://aws.amazon.com/blogs/security/aws-certificate-manager-will-discontinue-email-validation-to-prove-domain-validation-for-certificates/) for details.

## What Needs Action

| Certificate | Example | What you do |
|:------------|:--------|:------------|
| `convox.site` certificates for the Rack, its routers and `--wildcard-domain` Services | `*.myrack-router-123456789.us-east-1.convox.site` | Update the Rack. Convox publishes the DNS records, and the Rack switches these certificates on its own a few minutes after the update |
| Certificates for your own domains: created for a Service `domain:`, requested with `convox certs generate`, or applied to Gen 1 Apps with `convox ssl update` | `www.example.com` | Update the Rack and add one DNS record per domain name. The Rack then switches the certificate |
| Certificates imported with `convox certs import`, and IAM server certificates | any | Nothing. They do not use email or DNS validation |
| ACM certificates in the Rack's region that this Rack does not use | any | Switch each one from the Rack that uses it, or in the [ACM console](#switching-in-the-acm-console) if something outside Convox uses it. This Rack leaves them alone, but `convox certs dns` lists the ones still on email validation |

## Before You Start

| Requirement | Detail |
|:------------|:-------|
| Rack version | 20261021234916 or newer. Check with `convox rack`, and update with `convox rack update` |
| CLI version | A 3.x CLI needs 3.25.12 or newer, and a date-versioned CLI needs 20261021234916 or newer. Update with `sudo convox update`. Older CLIs do not have `convox certs dns` or `convox certs migrate` |

Most Racks need nothing else. If your Rack's outbound traffic or IAM permissions are restricted, it also needs:

| Requirement | Detail |
|:------------|:-------|
| HTTPS access to `certbot.convox.com` | The Rack publishes its `convox.site` records through Convox. If the Rack reaches the internet through [HttpProxy](/reference/rack-parameters/HttpProxy), the proxy must allow this host |
| DNS lookups on port 53 | The Rack checks each record against public DNS: `1.1.1.1`, `8.8.8.8` and the zone's nameservers. Without this access, it checks records through the VPC's DNS resolver instead. In that case, see [Delegated Subdomains and Private Zones](#delegated-subdomains-and-private-zones) if a private hosted zone for the domain is associated with the Rack's VPC |
| IAM permissions | The Rack's default IAM policies already allow every call it makes. If a permissions boundary or a service control policy restricts the Rack's `ApiRole`, it must allow `acm:UpdateCertificateOptions`, `acm:ListCertificateDomainValidations`, `acm:RequestCertificate`, `acm:AddTagsToCertificate` and `acm:DeleteCertificate`. With [Route53CertificateValidation](/reference/rack-parameters/Route53CertificateValidation#permissions) enabled, it must also allow the `route53` calls listed there |

## Migrating

### 1. Update the CLI and the Rack

```bash
$ sudo convox update
$ convox rack update --wait
```

A few minutes after the update finishes, the Rack switches its `convox.site` certificates to DNS validation. You do not need to do anything else for them.

### 2. Check Certificate Status

```bash
$ convox certs
ID                                DOMAIN                                           EXPIRES            STATUS
acm-0a1b2c3d4e5f                  example.com                                      3 months from now  email
acm-1b2c3d4e5f6a                  shop.example.org                                 4 months from now  email
acm-89ea927329d7                  *.myrack-router-123456789.us-east-1.convox.site  5 months from now  dns

2 certificates use email validation, which stops renewing on 2027-09-30: run convox certs dns
```

The `STATUS` column shows how each certificate is validated:

- `dns`: the certificate uses DNS validation. ACM renews it while it is in use, as long as its DNS record stays in place.
- `email`: the certificate still uses email validation and needs a DNS record.

See [certs](/reference/cli-commands/certs#status) for every status.

### 3. List the DNS Records

```bash
$ convox certs dns
ZONE         NAME                                                 TYPE   VALUE                                                              CHECK    CERTIFICATES
example.com  _0123456789abcdef0123456789abcdef.example.com.       CNAME  _fedcba9876543210fedcba9876543210.abcdefghij.acm-validations.aws.  missing  acm-0a1b2c3d4e5f
example.org  _1123456789abcdef0123456789abcdef.shop.example.org.  CNAME  _eedcba9876543210fedcba9876543210.bcdefghijk.acm-validations.aws.  missing  acm-1b2c3d4e5f6a

2 records: 2 missing
When every record shows ok: convox certs migrate --wait
```

Each row is one DNS record the Rack's certificates need:

| Column | Meaning |
|:-------|:--------|
| `ZONE` | The DNS zone the record goes in. This is usually your domain, such as `example.com`. A delegated subdomain shows as its own zone |
| `NAME`, `TYPE`, `VALUE` | The record to create |
| `CHECK` | Whether the record is in place: `missing` until you add it, then `ok`. See [Troubleshooting](#troubleshooting) for `wrong` and `unknown` |
| `CERTIFICATES` | The certificates that need the record |

When several certificates need the same record, it is listed once, and `CERTIFICATES` names all of them. `example.com` and `*.example.com` also share one record.

Right after the Rack update, the command can print `N certificates are still being prepared, run again in a few minutes`. The Rack is still reading those certificates' records from ACM. Wait a few minutes, or run `convox certs migrate`, which reads them right away. You can run it before adding any records, because it switches only certificates whose records are in place.

### 4. Add the Records

Add each record at your DNS provider as a `CNAME`. Most DNS dashboards add the zone to the name for you, so enter only the part of `NAME` before the zone, for example `_1123456789abcdef0123456789abcdef.shop` for the `example.org` zone. Use `VALUE` as the value. Any TTL works. [Adding the Records](#adding-the-records) has the steps for Route 53, Cloudflare and other providers.

Leave the records in place permanently. ACM checks them again at every renewal, and a certificate whose record is removed stops renewing.

Run `convox certs dns` again at any time. A record is in place once its `CHECK` shows `ok`.

### 5. Switch the Certificates

This step is optional. Without it, the Rack switches each certificate on its own within a few minutes of its record being in place.

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

`convox certs migrate` switches only the certificates whose records are in place, so you can run it at any time. A certificate that is still waiting for its record shows `waiting for its DNS record (run convox certs dns)` and stays as it was.

`--wait` waits up to 15 minutes, but ACM can take up to 30 minutes to finish a switch. Until it finishes, `convox certs` shows the certificate as `migrating`. If `--wait` stops with `ERROR: certificates did not finish switching`, the switch keeps going in the background. Check `convox certs` again later. See [certs migrate](/reference/cli-commands/certs-migrate) for every result.

## Adding the Records

| Where your DNS is hosted | Steps |
|:-------------------------|:------|
| Route 53, in the Rack's AWS account | [Route 53 in the Rack's AWS Account](#route-53-in-the-racks-aws-account) |
| Route 53, in another AWS account | [Route 53 in Another AWS Account](#route-53-in-another-aws-account) |
| Cloudflare | [Cloudflare](#cloudflare) |
| Any other provider | [Other DNS Providers](#other-dns-providers) |
| Managed by Terraform, external-dns or a similar tool | [DNS Managed as Code](#dns-managed-as-code) |

### Saving the Records to a File

Many DNS providers can import records from a zone file. `--zone-file` prints every record in that format, grouped by zone, with full names and a TTL of 300 seconds:

```bash
$ convox certs dns --zone-file > records.txt
$ cat records.txt
; zone example.com
_0123456789abcdef0123456789abcdef.example.com. 300 IN CNAME _fedcba9876543210fedcba9876543210.abcdefghij.acm-validations.aws.
; zone example.org
_1123456789abcdef0123456789abcdef.shop.example.org. 300 IN CNAME _eedcba9876543210fedcba9876543210.bcdefghijk.acm-validations.aws.
```

The file also includes records that already show `ok`. Many DNS providers reject an import that contains a record the zone already has, so remove those lines before you import.

Import each zone's records into that zone only. To write one zone's records to their own file:

```bash
$ awk -v zone=example.org '/^; zone / { z = $3; next } z == zone' records.txt > example.org.txt
```

### Route 53 in the Rack's AWS Account

If `convox certs dns` prints `<zone> is a Route 53 zone in this AWS account`, the Rack can write the records for you. Turn on [Route53CertificateValidation](/reference/rack-parameters/Route53CertificateValidation) and run the migration:

```bash
$ convox rack params set Route53CertificateValidation=Yes --wait
$ convox certs migrate --wait
```

`convox certs migrate` writes the missing records and switches each certificate once its record is in place. The first run usually shows `publishing` until Route 53 serves the new records, and the switch follows within a few minutes. If you do not run the command, the Rack does the same in the background every few minutes. The parameter also covers certificates you request later with `convox certs generate`.

The Rack follows these rules when it writes records:

- It writes only into a public hosted zone in its own AWS account, and only when that zone is the one the domain's public DNS is delegated to.
- It only creates records. It never changes a record that has another value, and never deletes a record.
- It writes records only for certificates not yet on DNS validation. If a record for a `dns` certificate goes missing, add it back yourself.

See [Route53CertificateValidation](/reference/rack-parameters/Route53CertificateValidation) for details.

### Route 53 in Another AWS Account

The Rack cannot write to a hosted zone in another AWS account, so add the records yourself with credentials for the account that holds the zone.

In the Route 53 console, open the public hosted zone named in `ZONE`. ACM cannot see records in a private hosted zone. If the account has more than one public hosted zone with that name, use the one whose name servers match the output of `dig +short NS example.org`. Records in the other zone are never served. Create one record for each row with that `ZONE` whose `CHECK` is not `ok`:

| Field | Value |
|:------|:------|
| Record name | The part of `NAME` before the zone, for example `_1123456789abcdef0123456789abcdef.shop` in the `example.org` zone |
| Record type | `CNAME` |
| Value | `VALUE` |

With many records, the AWS CLI is faster. In the commands below, `dns` is an AWS CLI profile for the account that holds the zone. Repeat steps 2 to 4 for each zone.

1. Save the records:

   ```bash
   $ convox certs dns --zone-file > records.txt
   ```

2. Find the hosted zone's id:

   ```bash
   $ aws route53 list-hosted-zones-by-name --dns-name example.org --query "HostedZones[?Name=='example.org.'].[Id,Config.PrivateZone]" --output text --profile dns
   ```

   Use the zone whose second column is `False`, which is the public zone, and drop the `/hostedzone/` prefix from its id. If more than one zone shows `False`, compare each one's name servers from `aws route53 get-hosted-zone --id <id> --profile dns` with `dig +short NS example.org`, as above.

3. Build a change batch from the zone's records:

   ```bash
   $ awk -v zone=example.org '
       BEGIN { printf "{\"Changes\":[" }
       /^; zone / { inzone = ($3 == zone); next }
       inzone { printf "%s{\"Action\":\"UPSERT\",\"ResourceRecordSet\":{\"Name\":\"%s\",\"Type\":\"CNAME\",\"TTL\":300,\"ResourceRecords\":[{\"Value\":\"%s\"}]}}", sep, $1, $5; sep = "," }
       END { print "]}" }' records.txt > example.org.json
   ```

4. Apply it to the zone:

   ```bash
   $ aws route53 change-resource-record-sets --hosted-zone-id Z0123456789ABCDEFGHIJ --change-batch file://example.org.json --profile dns
   ```

`UPSERT` creates each record, or sets it to the value in the batch if it already exists, so you can apply the same batch again safely.

### Cloudflare

Validation records must be **DNS only**. A proxied record is not served as a plain CNAME, and never validates.

To import a zone's records:

1. Write the zone's records to their own file, as shown in [Saving the Records to a File](#saving-the-records-to-a-file), and remove any line for a record that already shows `ok`.
2. In the Cloudflare dashboard, open the zone and go to **DNS** > **Records**.
3. Choose **Import and Export**, then **Import DNS records**, and select the file.
4. Clear **Proxy imported DNS records**, then import.

To add a record by hand, choose **Add record** and fill in:

| Field | Value |
|:------|:------|
| Type | `CNAME` |
| Name | The part of `NAME` before the zone, for example `_1123456789abcdef0123456789abcdef.shop` in `example.org` |
| Target | `VALUE` |
| Proxy status | DNS only |
| TTL | Auto |

Turn off CNAME flattening for these records. A flattened record stops returning the CNAME, so `convox certs dns` shows it as `missing`. Flattening can be on for every CNAME record in the zone (the zone's **CNAME flattening** setting) or for one record (the record's **Flatten** option), so check both.

If the zone uses a Cloudflare partial (CNAME) setup, add the records at the DNS provider that hosts the zone instead of in Cloudflare.

### Other DNS Providers

The same rules apply at every provider:

| Field | Rule |
|:------|:-----|
| Type | `CNAME` |
| Name | Enter only the part of `NAME` before the zone. Most dashboards add the zone for you, so entering the full name creates a record such as `_0123456789abcdef0123456789abcdef.example.com.example.com`, which shows `missing` |
| Value | `VALUE`, including its leading underscore. Remove the trailing dot only if the provider rejects it |
| TTL | Any. `300` matches the zone file |
| Proxying, forwarding, flattening | Off |

| Provider | Importing a zone file | Name field |
|:---------|:----------------------|:-----------|
| GoDaddy | **DNS** > **Actions** > **Import Zone File**. The import fails if a record already exists | Part before the zone |
| Namecheap | Not available. Add each record under **Advanced DNS** > **Host Records** | Part before the zone |
| Google Cloud DNS | `gcloud dns record-sets import example.org.txt --zone=<managed-zone> --zone-file-format`. The import fails if a record already exists. Never add `--delete-all-existing`, which replaces every record in the zone | Full name with trailing dot in `gcloud` |
| Azure DNS | `az network dns zone import -g <resource-group> -n example.org -f example.org.txt` merges the records into the zone | Part before the zone (`-n` in the command below) |
| DigitalOcean | Not available. Add each record under **Networking** > **Domains** | Part before the zone |
| DNSimple | **DNS** > **Import**, then paste the file's lines | Part before the zone |

To add one record with the Azure CLI:

```bash
$ az network dns record-set cname set-record -g <resource-group> -z example.org -n _1123456789abcdef0123456789abcdef.shop -c _eedcba9876543210fedcba9876543210.bcdefghijk.acm-validations.aws --ttl 300
```

### DNS Managed as Code

If Terraform, external-dns, DNSControl or another tool manages the zone, add the records to that tool's configuration. Otherwise its next run can remove them, and the certificates stop renewing.

### Delegated Subdomains and Private Zones

If a subdomain such as `shop.example.org` is delegated to its own nameservers, `ZONE` shows the subdomain, and the record goes in that zone.

ACM checks records only through public DNS. A record in a Route 53 private hosted zone or on an internal DNS server does not count. With split-horizon DNS, add the record to the public zone.

If the Rack cannot reach public DNS on port 53, it checks records through the VPC's DNS resolver instead. In that case, if a private hosted zone for the domain is associated with the Rack's VPC, add the record to the private zone as well, so the Rack can see it.

## Checking a Record

`convox certs dns` usually asks the zone's own nameservers, so a record shows `ok` as soon as your DNS provider serves it. In some cases it checks through the VPC's resolver instead; see [certs dns](/reference/cli-commands/certs-dns#check-values) for when. `convox certs migrate` always asks the nameservers when it can reach them, so if a record you added still shows `missing`, run `convox certs migrate`.

To check a record from your own machine:

```bash
$ dig +short CNAME _1123456789abcdef0123456789abcdef.shop.example.org @1.1.1.1
_eedcba9876543210fedcba9876543210.bcdefghijk.acm-validations.aws.
```

On Windows, use `nslookup -type=CNAME <name> 1.1.1.1`.

An empty answer means one of these:

- The record is not served yet.
- The record is proxied or flattened, which hides the CNAME.
- A public resolver still holds an earlier "not found" answer for a while.
- The record was created with the zone added twice, such as `_1123456789abcdef0123456789abcdef.shop.example.org.example.org`. Recreate it with the shorter name.

## New Certificates Still Use Email Validation

The Rack still requests a new certificate with email validation when:

- the Rack is installed
- [Internal](/reference/rack-parameters/Internal) is enabled
- a `--wildcard-domain` Service is created
- you deploy a Service `domain:` that no issued certificate covers

Once ACM issues one of these certificates, the Rack switches it to DNS validation: a `convox.site` certificate on its own, and a certificate for your own domain once you add its DNS record (see [Adding the Records](#adding-the-records)).

To skip email validation for a new custom domain, see [Adding a New Custom Domain](#adding-a-new-custom-domain).

## Adding a New Custom Domain

A domain that is already deployed does not need a new certificate. The migration switches the certificate it already uses.

A deploy uses an existing certificate only if ACM has issued it and it covers every domain the Service lists. Otherwise the deploy requests a new certificate with email validation. To use DNS validation for a new domain, request the certificate before you deploy:

1. Request a certificate with every domain the Service lists:

   ```bash
   $ convox certs generate example.com www.example.com
   Generating certificate... OK, acm-74ec8b405093
   ZONE         NAME                                                TYPE   VALUE                                                              CHECK
   example.com  _0123456789abcdef0123456789abcdef.example.com.      CNAME  _fedcba9876543210fedcba9876543210.abcdefghij.acm-validations.aws.  ok
   example.com  _3123456789abcdef0123456789abcdef.www.example.com.  CNAME  _ccdcba9876543210fedcba9876543210.defghijklm.acm-validations.aws.  missing
   Deploy the domain once convox certs shows acm-74ec8b405093 as dns
   ```

2. Add the records that show `missing` within 72 hours. After that, ACM stops waiting and the certificate shows `failed`. A record that shows `ok` is already in place from an earlier certificate. With [Route53CertificateValidation](/reference/rack-parameters/Route53CertificateValidation) set to `Yes`, the Rack adds the records for zones in its AWS account.
3. Wait until `convox certs` shows the certificate as `dns`. ACM can take up to 30 minutes after the records are in place.
4. Add the domains to the Service's `domain:` in `convox.yml` and deploy. The deploy uses the new certificate.

## Several Racks and AWS Accounts

Run the same commands against each Rack with `-r <rack>`.

A validation record belongs to one domain name and to the AWS account the Rack runs in, not the account that hosts your DNS. It works in every region. Racks in the same AWS account share records: once a record is in place, every certificate for that name in that account uses it. A Rack in another AWS account gets a different record for the same name, and both records can exist side by side.

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
| `convox rack logs` | `N certificates need DNS records: run convox certs dns` | Every hour, while certificates the Rack uses for your own domains are waiting for a record |
| `convox rack logs` | `certificate <id> (<domain>) switched to DNS validation` | When the Rack switches a certificate in the background. A switch made by `convox certs migrate` is not logged |
| `convox logs -a <app>` | `certificate <id> (<domain>) needs a DNS record before 2027-09-30: run convox certs dns` | Every hour, for a certificate Convox created for one of the App's Services that is waiting for a record |
| `convox logs -a <app>` | `certificate <id> (<domain>) switched to DNS validation` | When the Rack switches, in the background, a certificate Convox created for one of the App's Services |

## Troubleshooting

| Symptom | Cause | Fix |
|:--------|:------|:----|
| `CHECK` shows `missing` | The record is not served yet, is in a different zone than `ZONE`, has the zone added twice, is proxied or flattened, or exists only in a private zone | Check the record as shown in [Checking a Record](#checking-a-record) |
| `CHECK` shows `wrong`, and a parent name has a wildcard record such as `*.example.com` | The wildcard answers for the name because the validation record does not exist yet | Add the validation record, which takes precedence over the wildcard |
| `CHECK` shows `wrong`, with no wildcard involved | The record has an old value, or its value lost the leading underscore | Correct its value to `VALUE` |
| `CHECK` shows `unknown` | The lookup failed or did not finish | Run `convox certs dns` again. If it keeps happening, the zone's nameservers are not answering the Rack |
| `N certificates are still being prepared` | The Rack has not read these certificates' records from ACM yet | Run `convox certs migrate`, or wait a few minutes |
| `<id>: <error>` lines in `convox certs dns` | The Rack could not read that certificate or its records | Run `convox certs migrate` to try again |
| A certificate from `convox certs generate` shows `failed` | ACM did not see its records within 72 hours, or a CAA record blocked it (next row) | Add the records, run `convox certs generate` again, and remove the failed certificate with `convox certs delete` |
| `error: validation failed: CAA_ERROR` | A CAA record on the domain does not allow Amazon to issue certificates | Add a CAA record that allows `amazon.com`, such as `example.com. CAA 0 issue "amazon.com"`. If the domain has `issuewild` records, add an `issuewild` record for wildcard certificates too. Then run `convox certs migrate` |
| `convox certs` shows `convox.site` certificates still on email validation | The Rack has not published or checked their records yet, publishing failed, or the switch failed, for example because IAM denied `acm:UpdateCertificateOptions` | Run `convox certs migrate` and read the result for those certificates. `error: Post "https://certbot.convox.com/validation": ...` means the Rack cannot reach Convox: allow outbound HTTPS from the Rack, or through its HttpProxy. A message that ends in `contact Convox support` needs Convox to act |
| `ERROR: certificates did not finish switching` from `--wait` | ACM is still working, or a certificate ended with another result | Check `convox certs` later, and run `convox certs migrate` to see any certificate that is not `dns` or `migrating` |
| `error: not used by this rack` | The named certificate is not one this Rack uses | Run the command against the Rack that uses it |
| `ERROR: convox certs dns requires a rack on the latest release` | The Rack is older than 20261021234916 | Run `convox rack update` |
| Rack logs show `record <name> exists with another value, left alone` | With `Route53CertificateValidation=Yes`, a record with that name already exists with another value | Correct the record in Route 53 |
| Rack logs show `route53 write <name>: <error>` | The Rack could not write the record, for example because IAM denied the call | See [Route53CertificateValidation](/reference/rack-parameters/Route53CertificateValidation#permissions) |

## Switching in the ACM Console

You can also switch an email-validated certificate in the ACM console. Open the certificate, choose **Update validation method**, and add the records it shows. The Rack recognizes a certificate switched this way and shows it as `migrating`, then `dns`. If the records are not in place within 72 hours, the switch lapses, and the certificate stays on email validation and keeps working.

## Downgrading

Certificates already switched stay on DNS validation and keep renewing on an older Rack version. Versions older than 20261021234916 do not have the [Route53CertificateValidation](/reference/rack-parameters/Route53CertificateValidation) parameter, so a downgrade removes it. Set it again after you update back. Against an older Rack, `convox certs generate` requests email validation and prints no records.

## See Also

- [SSL](/deployment/ssl)
- [certs dns](/reference/cli-commands/certs-dns)
- [certs migrate](/reference/cli-commands/certs-migrate)
- [Route53CertificateValidation](/reference/rack-parameters/Route53CertificateValidation)
