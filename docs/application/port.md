---
title: "Port"
---

The Convox router listens on ports 80 and 443 and handles SSL termination for you. You will need to define the port on which your application is listening.

## Definition

```yaml
services:
  web:
    port: 3000
```

Your service should listen for HTTP requests on this port.

#### End-to-End Encryption

You can specify that your application is listening on HTTPS:

```yaml
services:
  web:
    port: https:3001
```

This syntax implies that your container is listening on HTTPS for requests between the router and the container.

## See also

- [Custom Domains](/docs/custom-domains)
- [SSL](/docs/ssl)
