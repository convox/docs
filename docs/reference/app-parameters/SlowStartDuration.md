---
title: "SlowStartDuration"
description: "Sets the ramp-up period during which a newly deployed service gradually receives increasing traffic."
---

# SlowStartDuration

Traffic ramp-up period (in seconds) for newly deployed services.

| Default value  | `0` |
| Allowed values | `0`, `30-900` seconds |

Once the duration expires, the full share of traffic will be directed to the new service.

## Use Cases

- Set to `30-120` seconds for applications that need to warm up caches, establish database connection pools, or compile JIT code before handling full traffic
- Set for JVM-based applications (Java, Scala, Kotlin) that benefit from JIT compilation warm-up
- Keep at `0` (default) for applications that are ready to handle full traffic immediately after passing health checks

## Additional Information

The slow start duration is configured on the ALB target group. During the slow start period, the load balancer linearly increases the share of traffic sent to newly registered targets. After the duration expires, the target receives its full share of traffic according to the configured [LoadBalancerAlgorithm](/reference/app-parameters/LoadBalancerAlgorithm).

Valid values are `0` (disabled) or any integer from `30` to `900`. Values between `1` and `29` are not allowed by AWS.

```bash
$ convox apps params set SlowStartDuration=60
```

## See Also

- [LoadBalancerAlgorithm](/reference/app-parameters/LoadBalancerAlgorithm)
- [LoadBalancerSuccessCodes](/reference/app-parameters/LoadBalancerSuccessCodes)
- [CircuitBreaker](/reference/app-parameters/CircuitBreaker)
- [Rack Parameters](/reference/rack-parameters)
