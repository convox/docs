---
title: "ApiRouter"
description: "Legacy load balancer type label for the Convox Rack API."
---

# ApiRouter

A legacy label for the Rack API load balancer type. The Rack API always uses an Application Load Balancer (ALBv2) regardless of this setting.

| Default value  | `ELB`        |
| Allowed values | `ALB`, `ELB` |

## Use Cases

- This parameter exists for backward compatibility. The Rack API load balancer is always an ALBv2 regardless of the value set here.

## Additional Information

Despite the `ALB`/`ELB` naming, the Rack API load balancer is always provisioned as an `AWS::ElasticLoadBalancingV2::LoadBalancer` (Application Load Balancer). The `ELB` default is a legacy label and does not create a Classic Load Balancer. Changing this parameter has no operational effect on the load balancer type.

```bash
$ convox rack params set ApiRouter=ALB
```

## See Also

- [ApiCount](/reference/rack-parameters/ApiCount)
- [ApiCpu](/reference/rack-parameters/ApiCpu)
- [ApiWebMemory](/reference/rack-parameters/ApiWebMemory)
- [SslPolicy](/reference/rack-parameters/SslPolicy)
