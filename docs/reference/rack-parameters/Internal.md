---
title: "Internal"
description: "Enable an internal load balancer on a Convox Rack for VPC-only accessible applications."
---

# Internal

Internal load balancer for the Rack, enabling private service routing within the VPC. See [Internal Services](/networking/internal-services) for more information.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Use Cases

- Enabling internal routing for Services that should only be accessible within the VPC, such as backend APIs or microservices
- Supporting a mix of public-facing and internal-only applications on the same Rack
- Restricting access to sensitive Services while keeping other Services publicly routable

## Additional Information

When set to `Yes`, Convox creates an additional internal (non-internet-facing) Application Load Balancer. Services can then be configured as internal in `convox.yml` to route traffic only through this internal ALB.

This is different from [InternalOnly](/reference/rack-parameters/InternalOnly), which removes the public-facing router entirely. With `Internal` set to `Yes`, the Rack supports both public and internal Services simultaneously.

```bash
$ convox rack params set Internal=Yes
```

## See Also

- [InternalOnly](/reference/rack-parameters/InternalOnly)
- [InternalRouterSuffix](/reference/rack-parameters/InternalRouterSuffix)
- [Private](/reference/rack-parameters/Private)
- [Internal Services](/networking/internal-services)
