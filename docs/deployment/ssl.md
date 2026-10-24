---
title: "SSL"
description: "Configure SSL certificates for Convox Services using AWS ACM, including generation, import, and local Rack setup."
---

# SSL

This page covers the V2 engine. V3 version: [SSL](https://docs.convox.com/deployment/ssl).

Convox will, if needed, automatically generate a valid SSL certificate for your Service via [AWS ACM](https://aws.amazon.com/certificate-manager/). Convox uses an existing ACM certificate for a Service when the certificate:

- is issued and unexpired
- is in the same region as the Rack
- has an RSA 1024-bit or 2048-bit key
- was not created by Convox for a Service `domain:`
- covers every domain in the Service's configuration

If no such certificate exists, the first deploy of a Service with a custom `domain:` requests a new certificate with email validation, and ACM sends a validation email for each domain. To use DNS validation instead, [pre-generate your certificate](#pre-generate-your-certificate) before the first deploy.

## Pre-generate your certificate

Convox allows you to generate your certificate ahead of time to ensure minimal delay before having your Service available during your first deploy. Pass every domain the Service lists. The certificate uses DNS validation, and the command prints the DNS records to add:

```sh
$ convox certs generate "*.example.org" "myapp.example.org"
Generating certificate... OK, acm-eeae31f242e9
ZONE         NAME                                                  TYPE   VALUE                                                              CHECK
example.org  _4123456789abcdef0123456789abcdef.example.org.        CNAME  _bbdcba9876543210fedcba9876543210.efghijklmn.acm-validations.aws.  missing
example.org  _5123456789abcdef0123456789abcdef.myapp.example.org.  CNAME  _aadcba9876543210fedcba9876543210.fghijklmno.acm-validations.aws.  missing
Deploy the domain once convox certs shows acm-eeae31f242e9 as dns
```

Add the records at your DNS provider and leave them in place, because ACM checks them again at every renewal. Once `convox certs` shows the certificate as `dns`, it is ready and you won't need to do anything further during your first deploy. See [Certificate DNS Validation](/deployment/certificate-dns-validation#adding-the-records) for the steps for each DNS provider.

DNS validation requires rack version 20261021234916 or newer, and CLI version 3.25.12 or newer, or 20261021234916 or newer. With an older CLI, or against an older Rack, the certificate uses email validation and ACM sends a validation email for each domain.

Each run requests a new certificate. To remove one you no longer need, pass its id to `convox certs delete`.

## Certificate management

To list your current certificates:

```sh
$ convox certs
ID                          DOMAIN                                                       EXPIRES             STATUS
acm-89ea927329d7            *.test-router-uactd9og6b40-1310739275.us-east-1.convox.site  10 months from now  dns
acm-eeae31f242e9            *.example.org                                                1 year from now     dns
acm-f1a2b3c4d5e6            api.example.org                                              2 months from now   email
cert-test-1580524125-66328  *.*.elb.amazonaws.com                                        10 months from now  imported

1 certificate uses email validation, which stops renewing on 2027-09-30: run convox certs dns
```

`STATUS` shows how ACM validates each certificate. See [certs](/reference/cli-commands/certs#status) for every status.

`convox certs` lists ACM certificates of every key type. A Service `domain:` uses only an RSA 1024-bit or 2048-bit certificate, so it never selects a listed ECDSA or 3072-bit or 4096-bit RSA certificate. Listing ECDSA and 3072-bit or 4096-bit RSA certificates requires rack version 20261005214736 or newer.

To delete an existing certificate, pass its full id:

```sh
$ convox certs delete acm-eeae31f242e9
Deleting certificate acm-eeae31f242e9... OK
```

To import an existing certificate:

```sh
$ convox certs import ~/.ssl/my_cert.pub ~/.ssl/my_key
Importing certificate... OK, acm-a89c0937f196
```

## Switching to DNS validation

ACM stops renewing email-validated certificates on 2027-09-30. On rack version 20261021234916 or newer, the Rack switches its certificates to DNS validation in place, without reissuing them: `convox.site` certificates automatically, and certificates for your own domains once you add one DNS record per domain. See [Certificate DNS Validation](/deployment/certificate-dns-validation).

## Certificates on NLB listeners

Services that expose a port through the [Network Load Balancer](/networking/nlb) with `protocol: tls` reference the certificate ARN directly in `convox.yml`:

```yaml
services:
  api:
    nlb:
      - port: 443
        protocol: tls
        containerPort: 3000
        scheme: public
        certificate: arn:aws:acm:us-east-1:123456789012:certificate/abcd1234-5678-90ab-cdef-1234567890ab
```

The ARN must be for an ACM certificate in the Rack's region and account (IAM server-certificate ARNs are also accepted). Cross-region and cross-account ARNs are rejected at release promote. `convox certs` lists certificates by Convox ID (`acm-` followed by the last 12 characters of the ARN) or IAM server-certificate name, not by full ARN, so retrieve the ARN from the AWS Console or `aws acm list-certificates --region <rack-region>` (add `--includes keyTypes=RSA_1024,RSA_2048,RSA_3072,RSA_4096,EC_prime256v1,EC_secp384r1,EC_secp521r1` to list certificates of every key type; by default it omits ECDSA and larger RSA certificates). Unlike ALB-routed Services, where Convox auto-provisions ACM certificates for the Service's domain, NLB listeners require the operator to pre-provision the certificate and paste the ARN.

An NLB TLS listener takes RSA certificates of up to 3072 bits and ECDSA P-256, P-384 and P-521 certificates. AWS rejects a 4096-bit RSA certificate from ACM on an NLB listener, and an IAM certificate with a 4096-bit RSA key puts the listener in a non-functional state.

## Local Rack

The local Rack will use DNS names `[process].[app].convox` which resolves to your local Rack. The local load balancer uses a certificate from a convox CA. On Firefox, you will need to set `security.enterprise_roots.enabled` to true in `about:config` or else you will not be able to confirm the security exception of the certificate.

## See Also

- [Certificate DNS Validation](/deployment/certificate-dns-validation)
- [Custom Domains](/deployment/custom-domains)
- [Load Balancing](/networking/load-balancing)
- [Network Load Balancing](/networking/nlb)
