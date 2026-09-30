---
title: "Custom Domains"
description: "Gen 1 (End of Life): How to map custom domains to Gen 1 Convox applications using CNAME records and DNS configuration."
---

# Custom Domains

> **This page documents Generation 1, which has reached End of Life.** Gen 1 apps use `docker-compose.yml`. For current documentation, see [Custom Domains](/deployment/custom-domains).

You can easily map a custom domain to a Convox application by creating a `CNAME` to your load balancer hostname.

## Balancer Hostname

You can find the load balancer hostname(s) for your application using `convox services`:

```bash
$ convox services
SERVICE  DOMAIN                                                  PORTS
web      docs-web-R72RMTP-326048479.us-east-1.elb.amazonaws.com  80
```

## Configuring DNS

Create an appropriate DNS entry to map your desired custom domain to your Convox app. In the example above one might create the following DNS entry:

| Field | Value |
|:--|:--|
| Name | `docs.convox.com` |
| Type | `CNAME` |
| Value | `docs-web-R72RMTP-326048479.us-east-1.elb.amazonaws.com` |
| TTL | `60` |

> **Note:** A root domain cannot use a `CNAME` record. Use an alias record in Route 53, or the equivalent from your DNS provider, pointing at the load balancer hostname.

## See Also

- [Custom Domains (Gen 2)](/deployment/custom-domains)
- [SSL](/gen1/ssl)
- [Load Balancers](/gen1/load-balancers)
