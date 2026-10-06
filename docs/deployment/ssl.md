---
title: "SSL"
description: "Configure SSL certificates for Convox Services using AWS ACM, including generation, import, and local Rack setup."
---

# SSL

Convox will, if needed, automatically generate a valid SSL certificate for your Service via [AWS ACM](https://aws.amazon.com/certificate-manager/). If you _already_ have an issued certificate in AWS ACM, in the same region as the Rack is installed, that `convox certs` lists, that has an RSA 1024-bit or 2048-bit key, and that covers every domain in your Service's configuration, Convox will use the existing certificate.

If you specify a custom `domain:` attribute for your Service be on the lookout for a validation email that will come the first time you deploy.

## Pre-generate your certificate

Convox allows you to generate your certificate ahead of time to ensure minimal delay before having your Service available during your first deploy.

```sh
$ convox certs generate "*.example.org" "myapp.example.org"
Generating certificate... OK, acm-eeae31f242e9
```

Once you validate the certificate and ACM issues it, it is ready and you won't need to do anything further during your first deploy.

Each run requests a new certificate. `convox certs` does not list a certificate until ACM issues it, so keep the id the command prints if you need to delete the request with `convox certs delete`.

## Certificate management

To list your current certificates:

```sh
$ convox certs
ID                          DOMAIN                                                       EXPIRES
acm-89ea927329d7            *.test-router-uactd9og6b40-1310739275.us-east-1.convox.site  10 months from now
acm-eeae31f242e9            *.example.org                                                1 year from now
cert-test-1580524125-66328  *.*.elb.amazonaws.com                                        10 months from now
```

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

- [Custom Domains](/deployment/custom-domains)
- [Load Balancing](/networking/load-balancing)
- [Network Load Balancing](/networking/nlb)
- [Security](/reference/security)
