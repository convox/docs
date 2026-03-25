---
title: "RedirectHttps"
description: "Controls whether HTTP requests are automatically redirected to HTTPS on the application load balancer."
---

# RedirectHttps

HTTP-to-HTTPS redirect control. When set to `Yes` (default), all HTTP requests are automatically redirected to HTTPS via a 301 response. Set to `No` to allow HTTP traffic.

| Default value  | `Yes` |
| Allowed values | `Yes`, `No` |

## Use Cases

- Keep as `Yes` (default) to enforce HTTPS for all traffic, which is recommended for production applications
- Set to `No` when your application is behind an upstream proxy or CDN that already handles TLS termination and sends traffic over HTTP
- Set to `No` for internal-only services where TLS is not required

## Additional Information

When set to `Yes`, the Rack's application load balancer returns an HTTP 301 redirect from port 80 to port 443 for all incoming requests to this app. This ensures all client traffic uses HTTPS.

When set to `No`, a port 80 forwarding rule is created for the app, allowing the load balancer to accept HTTP traffic without redirecting. The application receives the requests as-is over HTTP.

This parameter applies at the load balancer level. If your application needs to enforce HTTPS at the application layer as well, you should implement that separately in your application code.

```bash
$ convox apps params set RedirectHttps=No
```

## See Also

- [Private](/reference/app-parameters/Private)
- [InternalDomains](/reference/app-parameters/InternalDomains)
- [Rack Parameters](/reference/rack-parameters)
- [Services](/application/services)
