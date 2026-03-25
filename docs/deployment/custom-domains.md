---
title: "Custom Domains"
description: "How to configure custom domains for your Convox Services, including SSL certificates and DNS setup."
---

# Custom Domains

## Definition

You can specify that your Service should listen on a custom domain:

```yaml
services:
  web:
    domain: myapp.example.org
    port: 3000
```

You can also specify multiple domains:

```yaml
services:
  web:
    domain:
      - myapp.example.org
      - "*.example.org"
    port: 3000
```

> Using a custom domain requires a valid SSL certificate for the domains being specified. If a single certificate does not already exist in the same region as the Rack and matching all the domains you specify for your Service, one will be automatically created. The domain owner will receive a validation email with a link that must be clicked to complete the process. This happens the first time you deploy your Service with the custom domain configuration.

> You can [pre-generate your SSL certificate](/deployment/ssl#pre-generate-your-certificate) ahead of deploy time if you wish.

## Dynamic Configuration

You can use environment interpolation so that you don't have to hardcode the hostname in your `convox.yml`:

```yaml
services:
  web:
    domain: ${HOST}
    port: 3000
```

## Configuring DNS

Run `convox rack` and find the `Router` value. Configure your custom domain as a `CNAME` to this domain.

### Example

| Field | Value |
|-------|-------|
| Name  | `docs.convox.com` |
| Type  | `CNAME` |
| Value | `produ-Route-1ABCDEFGHIJK-01234569.us-east-1.elb.amazonaws.com` |
| TTL   | `60` |

> To set up DNS for a root domain, use an **Alias** record with Route 53 or the equivalent with your DNS provider.

## See Also

- [SSL](/deployment/ssl)
- [Services](/application/services)
- [Load Balancing](/networking/load-balancing)
- [Port](/application/port)
