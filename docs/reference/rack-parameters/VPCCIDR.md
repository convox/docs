---
title: "VPCCIDR"
description: "Set the CIDR block for the VPC created by the Convox Rack."
---

# VPCCIDR

VPC CIDR Block. This defines the overall IP address range for the Rack's VPC. All subnet CIDR blocks must fall within this range.

| Default value  | `10.0.0.0/16` |

## Use Cases

- Changing the VPC CIDR during initial Rack installation to avoid conflicts with existing VPCs when setting up VPC peering
- Using a smaller CIDR (e.g., `/20`) when your organization has limited IP address space
- Using a different CIDR range to comply with corporate network addressing policies

## Additional Information

> **Note:** Changing this parameter after the Rack is created has no effect, since VPC CIDR ranges cannot be changed after creation. This value must be set at Rack installation time.

The default `10.0.0.0/16` provides 65,536 IP addresses, which is sufficient for most deployments. The subnet CIDR parameters ([Subnet0CIDR](/reference/rack-parameters/Subnet0CIDR), [Subnet1CIDR](/reference/rack-parameters/Subnet1CIDR), [Subnet2CIDR](/reference/rack-parameters/Subnet2CIDR), [SubnetPrivate0CIDR](/reference/rack-parameters/SubnetPrivate0CIDR), [SubnetPrivate1CIDR](/reference/rack-parameters/SubnetPrivate1CIDR), [SubnetPrivate2CIDR](/reference/rack-parameters/SubnetPrivate2CIDR)) must all fall within this VPC CIDR range.

If installing into an existing VPC using [ExistingVpc](/reference/rack-parameters/ExistingVpc), this parameter should match the CIDR of that VPC.

```bash
$ convox rack params set VPCCIDR=10.1.0.0/16
```

## See Also

- [Subnet0CIDR](/reference/rack-parameters/Subnet0CIDR)
- [Subnet1CIDR](/reference/rack-parameters/Subnet1CIDR)
- [Subnet2CIDR](/reference/rack-parameters/Subnet2CIDR)
- [SubnetPrivate0CIDR](/reference/rack-parameters/SubnetPrivate0CIDR)
- [SubnetPrivate1CIDR](/reference/rack-parameters/SubnetPrivate1CIDR)
- [SubnetPrivate2CIDR](/reference/rack-parameters/SubnetPrivate2CIDR)
- [ExistingVpc](/reference/rack-parameters/ExistingVpc)
