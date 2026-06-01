---
title: "LoadBalancerGrpcSuccessCodes"
description: "Specifies the gRPC status codes that healthy targets must return when responding to a gRPC health check."
---

# LoadBalancerGrpcSuccessCodes

GRPC success codes for load balancer health checks.

| Default value  | `12` |
| Allowed values | `0-99` |

You can specify multiple values (e.g., `12,13`) or a range (e.g., `10-99`).

## Use Cases

- Adjust when your gRPC Services return non-default status codes for healthy responses
- Set to `0` if your gRPC health check endpoint returns `OK` (status code 0) to indicate health
- Set to a range or comma-separated list when your Service may return multiple valid status codes

## Additional Information

The default value of `12` corresponds to the gRPC `UNIMPLEMENTED` status code. This is the default because many services do not implement the gRPC health checking protocol, and the ALB treats `UNIMPLEMENTED` as a healthy response in that case.

If your Service implements the [gRPC Health Checking Protocol](https://github.com/grpc/grpc/blob/master/doc/health-checking.md), you should set this to `0` (the `OK` status code) to ensure the health check correctly validates your Service's health endpoint.

```bash
$ convox apps params set LoadBalancerGrpcSuccessCodes=0
```

## See Also

- [LoadBalancerSuccessCodes](/reference/app-parameters/LoadBalancerSuccessCodes)
- [LoadBalancerAlgorithm](/reference/app-parameters/LoadBalancerAlgorithm)
- [Rack Parameters](/reference/rack-parameters)
