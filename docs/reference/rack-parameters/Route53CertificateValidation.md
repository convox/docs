---
title: "Route53CertificateValidation"
description: "Let the Rack write ACM DNS validation records into Route 53 hosted zones in its own AWS account."
---

# Route53CertificateValidation

Write the ACM DNS validation records for your own domains' certificates into Route 53. With `Yes`, the Rack adds the records that a certificate on email validation, or waiting for DNS validation, needs when the domain's public hosted zone is in the Rack's AWS account, then switches the certificate to DNS validation once the records answer. See [Certificate DNS Validation](/deployment/certificate-dns-validation).

| Setting | Value |
|:--|:--|
| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

Requires rack version 20261021234916 or newer.

## Use Cases

- Switching certificates for your own domains to DNS validation without adding records by hand, when the domains' DNS is in Route 53 in the same AWS account as the Rack
- Having the records for new certificates from `convox certs generate` added for you

## Additional Information

```bash
$ convox rack params set Route53CertificateValidation=Yes --wait
```

While the parameter is `No`, [`convox certs dns`](/reference/cli-commands/certs-dns) prints a line for each zone this parameter would cover:

```text
example.org is a Route 53 zone in this AWS account. To have the rack add its records: convox rack params set Route53CertificateValidation=Yes
```

### Which Records the Rack Writes

The Rack writes a record when all of these hold:

| Condition | Detail |
|:--|:--|
| The Rack uses the certificate | The certificates [`convox certs dns`](/reference/cli-commands/certs-dns) lists. `convox.site` records are published by Convox, not written to your zones |
| The record does not already answer with the right value | Records that show `ok` are left as they are |
| The certificate is not yet on DNS validation | It uses email validation, is switching to DNS validation, or is waiting for DNS validation. For a certificate already on DNS validation the Rack writes nothing, so add any of its records that `convox certs dns` shows as not `ok` by hand |
| The zone is the one public DNS uses for the name | A public hosted zone in the Rack's AWS account, with the same name and nameservers as the zone public DNS delegates to. A private hosted zone, an undelegated copy of the zone, or a zone in another AWS account is never written |

The Rack writes records in the background every few minutes, and [`convox certs migrate`](/reference/cli-commands/certs-migrate) writes them at once. `convox certs dns` and `convox certs generate` do not write records.

Each record is a CNAME with a TTL of 300 seconds. The Rack only creates records. When a record with the same name already exists with another value, the Rack leaves it alone, `convox certs migrate` shows the certificate as `waiting for its DNS record`, and `convox rack logs` shows `certificate <id> (<domain>) record <name> exists with another value, left alone`. Correct the record in Route 53 and the Rack switches the certificate on its next pass.

The Rack never deletes the records it wrote, including after the parameter is set back to `No` or the certificate is deleted. ACM needs a certificate's records to stay in place to renew it.

### Zones Managed by Other Tools

The parameter is off by default because it writes into your own DNS zones. If Terraform, external-dns or another tool manages a zone, add the records `convox certs dns` lists to that tool's configuration instead, or its next run can remove them and the certificates stop renewing.

### Permissions

The Rack's default IAM policies allow the calls this parameter makes. If an administrator restricts the Rack's `ApiRole` with a permissions boundary or a service control policy, it must allow `route53:ListHostedZonesByName`, `route53:GetHostedZone`, `route53:ListResourceRecordSets` and `route53:ChangeResourceRecordSets`. A denied zone lookup appears in `convox rack logs` as `route53 zone <zone>: <error>`, and a denied write as `certificate <id> (<domain>) route53 write <name>: <error>`.

### Disabling

```bash
$ convox rack params set Route53CertificateValidation=No --wait
```

The Rack stops writing records. Records it already wrote stay in place, and certificates already on DNS validation keep renewing.

### Downgrading

A rack version older than 20261021234916 does not declare this parameter, so a downgrade drops it. Certificates already on DNS validation keep renewing. After updating back to 20261021234916 or newer, set the parameter again.

## See Also

- [Certificate DNS Validation](/deployment/certificate-dns-validation)
- [certs dns](/reference/cli-commands/certs-dns)
- [certs migrate](/reference/cli-commands/certs-migrate)
- [Rack Parameters](/reference/rack-parameters) for a full list of available parameters
