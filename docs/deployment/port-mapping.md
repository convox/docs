---
title: "Port Mapping"
order: 700
---

You can define a port on which your service will listen.

```yaml
services:
  web:
    port: 3000
```

Your service should listen for HTTP requests on this port.

You can specify that your application creates an HTTPS listener using this syntax:

```yaml
services:
  web:
    port: https:3001
```

Convox will automatically configure the Rack's load balancer to route external ports 80 and 443 to your application appropriately.

## See also

- [Custom Domains](/docs/custom-domains)
- [SSL](/docs/ssl)
