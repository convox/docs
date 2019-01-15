---
title: "SSL"
---

Convox will, if needed, automatically generate a valid SSL certificate for your service via [AWS ACM](https://aws.amazon.com/certificate-manager/). If you _already_ have a matching certificate in AWS ACM for the domain(s) in question Convox will use the existing certificate.

If you specify a custom `domain:` attribute for your service be on the lookout for a validation email that will come the first time you deploy.
