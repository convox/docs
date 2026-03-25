---
title: "LoadBalancerSuccessCodes"
description: "Specifies the HTTP status codes that healthy targets must return when responding to an ALB health check."
---

# LoadBalancerSuccessCodes

HTTP success codes for load balancer health checks.

| Default value  | `200-399,401` |
| Allowed values | `200-499` |

You can specify multiple values (e.g., `200,202`) or a range (e.g., `200-299`).

## Use Cases

- Narrow to `200` or `200-299` when your health check endpoint should only return success codes
- Add `401` or other codes when your health check path is behind authentication and returns those codes when healthy
- Widen the range when your application returns redirect (3xx) responses on the health check path and those should be treated as healthy

## Additional Information

The default value of `200-399,401` is deliberately permissive. It treats any 2xx, 3xx, or 401 response as healthy. This works well for most applications, including those where the health check path may redirect or require authentication.

If your application has a dedicated health check endpoint (e.g., `/health`) that returns a clean `200 OK`, consider narrowing this to `200` for more accurate health detection.

```bash
$ convox apps params set LoadBalancerSuccessCodes=200
```

## See Also

- [LoadBalancerGrpcSuccessCodes](/reference/app-parameters/LoadBalancerGrpcSuccessCodes)
- [LoadBalancerAlgorithm](/reference/app-parameters/LoadBalancerAlgorithm)
- [SlowStartDuration](/reference/app-parameters/SlowStartDuration)
- [Rack Parameters](/reference/rack-parameters)
