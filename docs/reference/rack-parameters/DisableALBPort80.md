---
title: "DisableALBPort80"
description: "Disables port 80 (HTTP) on the Application Load Balancer for a Convox Rack."
---

# DisableALBPort80

Port 80 (HTTP) listener control for the Application Load Balancer. When set to `Yes`, the ALB will no longer listen on port 80 (HTTP), and only HTTPS traffic on port 443 will be accepted.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Use Cases

- Enforce HTTPS-only traffic by removing the HTTP listener entirely, rather than relying on HTTP-to-HTTPS redirects.
- Meet compliance or security policies that require no unencrypted HTTP endpoints to be exposed.
- Simplify security audits by eliminating the HTTP port from the load balancer configuration.

## Additional Information

By default, Convox configures the ALB to listen on both port 80 (HTTP) and port 443 (HTTPS). Port 80 traffic is always redirected to HTTPS via a 301 redirect. Setting this parameter to `Yes` removes the port 80 listener completely.

> **Note:** This parameter only takes effect when [InternalOnly](/reference/rack-parameters/InternalOnly) is set to `No` (the default). When `InternalOnly` is `Yes`, there is no public router and therefore no port 80 listener — this parameter has no effect in that configuration.

If your applications rely on HTTP-to-HTTPS redirects at the load balancer level, disabling port 80 will prevent those redirects from working. Ensure that all clients are configured to connect directly via HTTPS before enabling this parameter.

```bash
$ convox rack params set DisableALBPort80=Yes
```

## See Also

- [SslPolicy](/reference/rack-parameters/SslPolicy)
- [Internal](/reference/rack-parameters/Internal)
- [InternalOnly](/reference/rack-parameters/InternalOnly)
- [Rack Parameters](/reference/rack-parameters) for a full list of available parameters
