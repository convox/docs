---
title: "Subnet2CIDR"
description: "Set the CIDR block for Public Subnet 2 in the Convox Rack VPC."
---

# Subnet2CIDR

Public Subnet 2 CIDR Block. This defines the IP address range for the third public subnet in the Rack's VPC. The CIDR block must fall within the [VPCCIDR](/reference/rack-parameters/VPCCIDR) range.

| Default value  | `10.0.3.0/24` |

## Use Cases

- Customizing subnet ranges when installing into an existing VPC to avoid overlap with pre-existing subnets
- Adjusting the CIDR size to allocate more or fewer IP addresses based on expected workload density
- Aligning subnet CIDRs with your organization's IP address management (IPAM) scheme

## Additional Information

This subnet is placed in the third Availability Zone and is only created when [HighAvailability](/reference/rack-parameters/HighAvailability) is `true` and [MaxAvailabilityZones](/reference/rack-parameters/MaxAvailabilityZones) is `3`. Public subnets are used for resources that need direct internet access, such as load balancers and NAT gateways. The default `/24` CIDR provides 254 usable IP addresses.

When installing a Rack into an [ExistingVpc](/reference/rack-parameters/ExistingVpc), ensure the CIDR blocks do not overlap with existing subnets in that VPC. Subnet CIDR values should typically only be customized at install time or when using [ExistingVpc](/reference/rack-parameters/ExistingVpc).

```bash
$ convox rack params set Subnet2CIDR=10.1.3.0/24
```

## See Also

- [Subnet0CIDR](/reference/rack-parameters/Subnet0CIDR)
- [Subnet1CIDR](/reference/rack-parameters/Subnet1CIDR)
- [SubnetPrivate2CIDR](/reference/rack-parameters/SubnetPrivate2CIDR)
- [VPCCIDR](/reference/rack-parameters/VPCCIDR)
- [ExistingVpc](/reference/rack-parameters/ExistingVpc)
- [MaxAvailabilityZones](/reference/rack-parameters/MaxAvailabilityZones)
