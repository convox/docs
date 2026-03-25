---
title: "LoadBalancerAlgorithm"
description: "Sets the routing algorithm used by the application load balancer for distributing traffic to targets."
---

# LoadBalancerAlgorithm

Routing algorithm for the application load balancer.

| Default value  | `round_robin` |
| Allowed values | `round_robin`, `least_outstanding_requests` |

## Use Cases

- Use `least_outstanding_requests` when your services have uneven processing times and you want to route traffic to the least busy target
- Use `least_outstanding_requests` to improve latency distribution when some requests take significantly longer than others
- Keep as `round_robin` (default) for workloads with uniform request processing times

## Additional Information

The `round_robin` algorithm distributes requests evenly across all healthy targets in rotation. This works well when all requests take roughly the same amount of time to process.

The `least_outstanding_requests` algorithm routes each new request to the target with the fewest in-flight requests. This can significantly improve performance for applications where request processing times vary widely, as it prevents slow requests from causing a backlog on a single target while other targets sit idle.

This setting applies to all services in the application that use an Application Load Balancer (ALB) target group.

```bash
$ convox apps params set LoadBalancerAlgorithm=least_outstanding_requests
```

## See Also

- [LoadBalancerSuccessCodes](/reference/app-parameters/LoadBalancerSuccessCodes)
- [LoadBalancerGrpcSuccessCodes](/reference/app-parameters/LoadBalancerGrpcSuccessCodes)
- [SlowStartDuration](/reference/app-parameters/SlowStartDuration)
- [Rack Parameters](/reference/rack-parameters)
