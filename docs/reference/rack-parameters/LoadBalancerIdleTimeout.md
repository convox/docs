---
title: "LoadBalancerIdleTimeout"
description: "Configure the idle timeout for the Rack's Application Load Balancer in seconds."
---

# LoadBalancerIdleTimeout

The idle timeout value for the Rack's Application Load Balancer (ALB), in seconds. This controls how long the ALB waits before closing idle connections. The valid range is 1-4000 seconds.

| Default value  | `3600` |
| Allowed values | `1` - `4000` |

## Use Cases

- Increasing the timeout for applications that use long-lived connections such as WebSockets or server-sent events
- Reducing the timeout to free up ALB resources faster in high-traffic environments with short-lived requests
- Matching the timeout to your application's keep-alive settings to prevent premature connection drops

## Additional Information

The default value of 3600 seconds (1 hour) is appropriate for most workloads. If your application uses WebSockets or long-polling, you may need to keep this at or above 3600 to avoid connection interruptions.

For applications that only serve short HTTP request/response cycles, lowering this value can help the ALB release idle connections more quickly and improve resource utilization.

```bash
$ convox rack params set LoadBalancerIdleTimeout=300
```

The idle timeout applies to the front-facing ALB that routes traffic to your application services. Make sure your application's own keep-alive timeout is set higher than this value to prevent the application from closing the connection before the load balancer does.

## See Also

- [SslPolicy](/reference/rack-parameters/SslPolicy)
- [Internal](/reference/rack-parameters/Internal)
- [InternalOnly](/reference/rack-parameters/InternalOnly)
- [Private](/reference/rack-parameters/Private)
