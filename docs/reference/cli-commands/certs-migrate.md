---
title: "certs migrate"
description: "Switch the Rack's ACM certificates from email validation to DNS validation once their DNS records are in place."
---

# certs migrate

Switch the Rack's ACM certificates from email validation to DNS validation. Each certificate keeps its ARN and expiry and is not reissued, so Apps keep serving HTTPS throughout. A certificate is switched only once every DNS record it needs answers with the right value, checked the way [`convox certs dns`](/reference/cli-commands/certs-dns#check-values) checks it. A certificate whose records are not in place yet stays on email validation and keeps working, so the command is safe to run at any time.

The Rack also does this on its own every few minutes. Run `convox certs migrate` to do it at once and to see where each certificate stands. For `convox.site` certificates, the Rack first publishes their records through Convox. With [Route53CertificateValidation](/reference/rack-parameters/Route53CertificateValidation) set to `Yes`, it also writes records into Route 53 zones in the Rack's AWS account. See [Certificate DNS Validation](/deployment/certificate-dns-validation) for the full migration.

Without arguments, the command covers every ACM certificate the Rack uses, including its `convox.site` certificates. Pass one or more certificate ids, such as `acm-0a1b2c3d4e5f`, to act on only those. A domain name is not accepted in place of an id.

Requires rack version 20261021234916 or newer, and CLI version 3.25.12 or newer, or 20261021234916 or newer. Against an older Rack the command returns `ERROR: convox certs migrate requires a rack on the latest release`.

## Syntax

```bash
$ convox certs migrate [cert...]
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for the switches to finish |

## Example Usage

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

Switch one certificate:

```bash
$ convox certs migrate acm-1b2c3d4e5f6a
CERTIFICATE       DOMAIN            RESULT
acm-1b2c3d4e5f6a  shop.example.org  waiting for its DNS record (run convox certs dns)
```

## Results

| Result | Meaning |
|:-------|:--------|
| `dns` | The certificate uses DNS validation. Nothing to do |
| `switching` | Every record answers, and the switch to DNS validation has started. `convox certs` shows `migrating` until ACM finishes |
| `publishing` | The record was just published, by Convox for a `convox.site` certificate or by the Rack in Route 53, and does not answer yet. The switch follows within a few minutes |
| `waiting for its DNS record (run convox certs dns)` | A record is missing or has another value, and the Rack cannot add it. Add the record [`convox certs dns`](/reference/cli-commands/certs-dns) shows |
| `preparing` | The Rack has not read the certificate's records from ACM yet. Run the command again in a minute |
| `pending` | ACM has not issued the certificate yet |
| `failed` | ACM could not issue the certificate, or its validation timed out |
| `imported` | An imported certificate or an ACM private certificate. These do not use email or DNS validation |
| `error: <message>` | The step failed. See [Troubleshooting](/deployment/certificate-dns-validation#troubleshooting) |

Common error messages:

| Message | Meaning |
|:--------|:--------|
| `error: not used by this rack` | The named certificate is not one this Rack uses. Run the command against the Rack that uses it |
| `error: certificate not found` | No ACM certificate in the Rack's region has that id |
| `error: validation failed: CAA_ERROR` | A CAA record on the domain does not allow Amazon to issue certificates. See [Troubleshooting](/deployment/certificate-dns-validation#troubleshooting) |
| `error: registrar <status>: <message>` | Convox could not publish a `convox.site` record. Messages that end in `contact Convox support` need Convox to act |
| `error: Post "https://certbot.convox.com/validation": ...` | The Rack cannot reach Convox to publish its `convox.site` records. Allow outbound HTTPS from the Rack, or through [HttpProxy](/reference/rack-parameters/HttpProxy) if the Rack uses one |

## Waiting

With `--wait`, the command waits on the certificates whose result was `switching`, `publishing` or `preparing`. It checks every 10 seconds for up to 15 minutes, then prints their results:

| Outcome | Exit status |
|:--------|:------------|
| Every certificate it waited on reached `dns` | `0`, after `Waiting for switches to finish... OK` |
| No certificate needed waiting on | `0`, returns at once |
| Any other result, or 15 minutes passed | `1`, with `ERROR: certificates did not finish switching: <ids>` |

ACM can take up to 30 minutes to finish a switch. A switch still in progress when `--wait` gives up continues in the background, and [`convox certs`](/reference/cli-commands/certs) shows its status, `migrating` and then `dns`. Without `--wait`, the command prints the results and exits `0`.

## See Also

- [certs dns](/reference/cli-commands/certs-dns)
- [certs](/reference/cli-commands/certs)
- [Certificate DNS Validation](/deployment/certificate-dns-validation)
- [Route53CertificateValidation](/reference/rack-parameters/Route53CertificateValidation)
