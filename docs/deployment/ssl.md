---
title: "SSL"
---

Convox will, if needed, automatically generate a valid SSL certificate for your service via [AWS ACM](https://aws.amazon.com/certificate-manager/). If you _already_ have a matching certificate in AWS ACM for the domain(s) in question Convox will use the existing certificate.

If you specify a custom `domain:` attribute for your service be on the lookout for a validation email that will come the first time you deploy.

## Local Rack

The local rack will use DNS names `[process].[app].convox` which resolves to your local rack. The local load balancer uses a certificate from a convox CA. On Firefox, you will need to set `security.enterprise_roots.enabled` to true in `about:config` or else you will not be able to confirm the security exception of the certificate.
