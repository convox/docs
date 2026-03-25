---
title: "InternalOnly"
description: "Restrict a Convox Rack to only support VPC-internal applications with no public-facing router."
---

# InternalOnly

VPC-only mode for the Rack. When set to `Yes`, no public router is created and applications are only accessible from within the VPC.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Use Cases

- Deploying a Rack in a private network where no services should be publicly accessible
- Meeting compliance requirements that prohibit public-facing endpoints
- Running backend infrastructure Racks that are accessed exclusively through VPN or VPC peering

## Additional Information

When `InternalOnly` is set to `Yes`, the Rack only creates an internal Application Load Balancer. No internet-facing ALB is provisioned, so no services on this Rack can receive traffic from the public internet.

This differs from the [Internal](/reference/rack-parameters/Internal) parameter, which adds an internal ALB alongside the public one. `InternalOnly` removes the public ALB entirely.

If you need some services to be internal and others to be public, use [Internal](/reference/rack-parameters/Internal) instead.

```bash
$ convox rack params set InternalOnly=Yes
```

## See Also

- [Internal](/reference/rack-parameters/Internal)
- [InternalRouterSuffix](/reference/rack-parameters/InternalRouterSuffix)
- [Private](/reference/rack-parameters/Private)
