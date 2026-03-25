---
title: "Private"
description: "Controls whether the application's ECS tasks are deployed to private subnets with no public IP."
---

# Private

Private subnet placement for the application's ECS tasks. When set to `Yes`, services run in private subnets with public IP assignment disabled. When set to `No`, services run in public subnets with public IPs assigned.

| Default value  | `No` |
| Allowed values | `Yes`, `No` |

## Use Cases

- Set to `Yes` for applications that should not have direct internet access and must route through a NAT gateway
- Set to `Yes` for backend services or internal APIs that do not need public IP addresses
- Keep as `No` (default) for applications that need direct public internet connectivity

## Additional Information

This parameter is **independent of the Rack-level `Private` parameter**. Each App can override its networking posture individually -- for example, a single App can be set to `Private=Yes` even if the Rack has `Private=No`, and vice versa.

> This parameter controls task-level networking (subnet placement and public IP assignment). It does not affect load balancer scheme. To make a service's load balancer internal, use the [`internal`](/application/services) attribute in `convox.yml`.

When `Private=Yes`, ECS tasks are placed in the Rack's private subnets. These subnets typically route outbound traffic through a NAT gateway. Ensure your Rack has private subnets and NAT gateways configured before setting this parameter.

The [Isolate](/reference/app-parameters/Isolate) parameter provides additional network isolation on top of Private mode.

```bash
$ convox apps params set Private=Yes
```

## See Also

- [Isolate](/reference/app-parameters/Isolate)
- [PlaceLambdaInVpc](/reference/app-parameters/PlaceLambdaInVpc)
- [InternalDomains](/reference/app-parameters/InternalDomains)
- [Rack Parameter: Private](/reference/rack-parameters/Private)
- [Private Networking](/networking/private-networking)
- [Services](/application/services)
