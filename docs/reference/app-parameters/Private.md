---
title: "Private"
description: "Controls whether the application's Fargate and Isolate tasks run in private subnets with no public IP."
---

# Private

Private subnet placement for the application's tasks that have their own network interface: Services on Fargate or with [Isolate](/reference/app-parameters/Isolate), and timers on Fargate. When set to `Yes`, these tasks run in private subnets with public IP assignment disabled. When set to `No`, they run in public subnets with public IPs assigned. Services on EC2 without `Isolate` run on the Rack's instances, which are placed by the Rack's [Private](/reference/rack-parameters/Private) parameter.

| Setting | Value |
|:--|:--|
| Default value  | `No` |
| Allowed values | `Yes`, `No` |

## Use Cases

- Set to `Yes` for applications that should not have direct internet access and must route through a NAT gateway
- Set to `Yes` for backend services or internal APIs that do not need public IP addresses
- Keep as `No` (default) for applications that need direct public internet connectivity

## Additional Information

Each deploy or promote of a Generation 2 App sets this parameter to the Rack's [Private](/reference/rack-parameters/Private) value, so a value set with `convox apps params set` lasts only until the next deploy. To run one App with a different value from the Rack, set it under `params:` in `convox.yml`. See [Parameters Set From the Rack](/reference/app-parameters#parameters-set-from-the-rack).

> This parameter controls task-level networking (subnet placement and public IP assignment). It does not affect load balancer scheme. To make a Service's load balancer internal, use the [`internal`](/application/services) attribute in `convox.yml`.

When `Private=Yes`, these tasks are placed in the Rack's private subnets, which route outbound traffic through NAT gateways. A Rack has private subnets only when its [Private](/reference/rack-parameters/Private) or [PrivateBuild](/reference/rack-parameters/PrivateBuild) parameter is `Yes`. Without them, setting `Private: "Yes"` makes the deploy of an App with Fargate Services or Fargate timers roll back.

The [Isolate](/reference/app-parameters/Isolate) parameter provides additional network isolation on top of Private mode.

For example, to run a Fargate App on a private Rack in public subnets:

```yaml
params:
  Private: "No"
```

## See Also

- [Isolate](/reference/app-parameters/Isolate)
- [PlaceLambdaInVpc](/reference/app-parameters/PlaceLambdaInVpc)
- [InternalDomains](/reference/app-parameters/InternalDomains)
- [Rack Parameter: Private](/reference/rack-parameters/Private)
- [Private Networking](/networking/private-networking)
- [Services](/application/services)
