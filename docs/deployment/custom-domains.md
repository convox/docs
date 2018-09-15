---
title: "Custom Domains"
order: 800
---

You can specify that your service should listen on a custom domain:

```yaml
services:
  web:
    domain: myapp.example.org
    port: 3000
```

You can also specify multiple domains using this syntax:

```yaml
services:
  web:
    domain:
      - myapp.example.org
      - "*.example.net"
    port: 3000
```

You can use environment interpolation so that you don't have to hardcode the hostname in your `convox.yml`:

```yaml
services:
  web:
    domain: ${HOST}
    port: 3000
```

## Configuring DNS

Run `convox rack` and find the `Domain` value. Configure your custom domain as a `CNAME` or `ALIAS` pointer to this domain.
