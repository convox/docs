---
title: "RouterInternalSecurityGroup"
description: "Assign custom security groups to the internal Convox Rack router for controlling VPC-internal traffic."
---

# RouterInternalSecurityGroup

Comma-delimited list of security group IDs to assign to the internal Rack router. If blank, the Rack creates its own internal router security group that allows TCP traffic on ports 80 and 443 from within the VPC CIDR range.

| Default value  | "" |

## Use Cases

- Restricting which VPC-internal services or subnets can reach your internal applications
- Applying organization-mandated security groups that enforce specific firewall rules or logging
- Using a shared security group across multiple Racks for consistent internal network policies

## Additional Information

The internal router is only created when [Internal](/reference/rack-parameters/Internal) is set to `Yes`. It handles traffic for services configured with `internal: true` in their `convox.yml`, making them accessible only from within the VPC.

When this parameter is blank, the Rack creates a default security group that permits inbound TCP on ports 80 and 443 from the [VPCCIDR](/reference/rack-parameters/VPCCIDR) range. If you provide custom security group IDs, you are fully responsible for ensuring the correct ingress rules are in place.

```bash
$ convox rack params set RouterInternalSecurityGroup=sg-0abc123def456
```

## See Also

- [RouterSecurityGroup](/reference/rack-parameters/RouterSecurityGroup)
- [Internal](/reference/rack-parameters/Internal)
- [InternalOnly](/reference/rack-parameters/InternalOnly)
- [InternalRouterSuffix](/reference/rack-parameters/InternalRouterSuffix)
- [VPCCIDR](/reference/rack-parameters/VPCCIDR)
