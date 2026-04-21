---
title: "SSL"
description: "Configure SSL certificates for Convox Services using AWS ACM, including generation, import, and local Rack setup."
---

# SSL

Convox will, if needed, automatically generate a valid SSL certificate for your service via [AWS ACM](https://aws.amazon.com/certificate-manager/). If you _already_ have a single matching certificate in AWS ACM, in the same region as the Rack is installed, for the domain(s) in your service's configuration, Convox will use the existing certificate.

If you specify a custom `domain:` attribute for your service be on the lookout for a validation email that will come the first time you deploy.

## Pre-generate your certificate

Convox allows you to generate your certificate ahead of time to ensure minimal delay before having your service available during your first deploy.

```sh
$ convox certs generate "*.example.org" "myapp.example.org"
Generating certificate... OK, acm-eeae31f242e9
```

This will initiate the validation email process, so once you have validated the certificate, it will be ready and you won't need to do anything further during your first deploy.

## Certificate management

To list your current certificates:

```sh
$ convox certs
ID                          DOMAIN                                                       EXPIRES
acm-89ea927329d7            *.test-router-uactd9og6b40-1310739275.us-east-1.convox.site  10 months from now
acm-a911c40399a1            *.example.org                                                1 year from now
cert-test-1580524125-66328  *.*.elb.amazonaws.com                                        10 months from now
```

To delete an existing certificate:

```sh
$ convox certs delete acm-a89c0937f196
Deleting certificate acm-a89c0937f196... OK
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

The ARN must be for an ACM certificate in the Rack's region and account (IAM server-certificate ARNs are also accepted). Cross-region and cross-account ARNs are rejected at release promote. `convox certs` lists the Rack's certificates by Convox short ID (`acm-<hash>`) or IAM server-certificate name, not by full ARN — retrieve the ARN from the AWS Console or `aws acm list-certificates --region <rack-region>`. Unlike ALB-routed Services — where Convox auto-provisions ACM certificates for the Service's domain — NLB listeners require the operator to pre-provision the certificate and paste the ARN.

## Local Rack

The local rack will use DNS names `[process].[app].convox` which resolves to your local rack. The local load balancer uses a certificate from a convox CA. On Firefox, you will need to set `security.enterprise_roots.enabled` to true in `about:config` or else you will not be able to confirm the security exception of the certificate.

## See Also

- [Custom Domains](/deployment/custom-domains)
- [Load Balancing](/networking/load-balancing)
- [Network Load Balancing](/networking/nlb)
- [Security](/reference/security)
