---
title: "SubnetPrivate1CIDR"
description: "Set the CIDR block for Private Subnet 1 in the Convox Rack VPC."
---

# SubnetPrivate1CIDR

Private Subnet 1 CIDR Block. This defines the IP address range for the second private subnet in the Rack's VPC. The CIDR block must fall within the [VPCCIDR](/reference/rack-parameters/VPCCIDR) range.

| Default value  | `10.0.5.0/24` |

## Use Cases

- Customizing subnet ranges when installing into an existing VPC to avoid overlap with pre-existing subnets
- Adjusting the CIDR size to allocate more or fewer IP addresses based on expected workload density
- Aligning subnet CIDRs with your organization's IP address management (IPAM) scheme

## Additional Information

This subnet is placed in the second Availability Zone. Private subnets are only created when the [Private](/reference/rack-parameters/Private) parameter is set to `Yes`. Resources in private subnets do not have direct internet access and route outbound traffic through a NAT gateway. The default `/24` CIDR provides 254 usable IP addresses.

When installing a Rack into an [ExistingVpc](/reference/rack-parameters/ExistingVpc), ensure the CIDR blocks do not overlap with existing subnets in that VPC. Subnet CIDR values should typically only be customized at install time or when using [ExistingVpc](/reference/rack-parameters/ExistingVpc).

```bash
$ convox rack params set SubnetPrivate1CIDR=10.1.5.0/24
```

## See Also

- [SubnetPrivate0CIDR](/reference/rack-parameters/SubnetPrivate0CIDR)
- [SubnetPrivate2CIDR](/reference/rack-parameters/SubnetPrivate2CIDR)
- [Subnet1CIDR](/reference/rack-parameters/Subnet1CIDR)
- [VPCCIDR](/reference/rack-parameters/VPCCIDR)
- [Private](/reference/rack-parameters/Private)
- [ExistingVpc](/reference/rack-parameters/ExistingVpc)
