---
title: "Private"
description: "Create Rack resources in private subnets with no direct public internet access."
---

# Private

Private subnet placement for all Rack instances. When set to `Yes`, instances have no direct internet access and all outbound traffic passes through a NAT gateway. See the [Private Networking](/networking/private-networking) documentation for more information.

When set to `Yes`, EC2 instances, ECS tasks, and other Rack resources are placed in private subnets that do not have direct routes to the internet. Outbound internet access is provided through NAT Gateways.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Use Cases

- Running workloads in a private network to meet security and compliance requirements
- Preventing Rack instances from being directly accessible from the internet
- Deploying applications that should only be reachable through a load balancer or VPN

## Additional Information

When `Private` is set to `Yes`, the Rack creates NAT Gateways for outbound internet connectivity. NAT Gateways incur additional AWS charges (per-hour and per-GB data processing fees).

Setting `Private=Yes` also makes the [PrivateBuild](/reference/rack-parameters/PrivateBuild) parameter unnecessary, since all instances including build instances will be placed in private subnets.

If you also want the Rack API itself to be unreachable from the public internet, set [PrivateApi](/reference/rack-parameters/PrivateApi) to `Yes` as well.

```bash
$ convox rack params set Private=Yes
```

## See Also

- [PrivateApi](/reference/rack-parameters/PrivateApi)
- [PrivateApiSecurityGroup](/reference/rack-parameters/PrivateApiSecurityGroup)
- [PrivateBuild](/reference/rack-parameters/PrivateBuild)
- [PlaceLambdaInVpc](/reference/rack-parameters/PlaceLambdaInVpc)
- [Internal](/reference/rack-parameters/Internal)
- [InternalOnly](/reference/rack-parameters/InternalOnly)
- [Private Networking](/networking/private-networking)
