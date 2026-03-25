---
title: "ExistingVpc"
description: "Specifies an existing AWS VPC ID to use instead of creating a new one for the Convox Rack."
---

# ExistingVpc

Existing VPC ID from AWS. If blank, a VPC will be created. The additional parameter [InternetGateway](/reference/rack-parameters/InternetGateway) must be set to use **ExistingVpc**.

| Default value  | "" |
| Allowed values | VPC ID    |

## Use Cases

- Install a Convox Rack into an existing VPC that contains other infrastructure, databases, or services that your applications need to communicate with via private networking.
- Share a VPC across multiple Racks or AWS services to reduce the number of VPCs in your account and simplify network management.
- Use a pre-configured VPC with specific CIDR ranges, route tables, or VPN/Direct Connect attachments that cannot be replicated by the Rack's auto-created VPC.

## Additional Information

When specifying an existing VPC, you must also provide the [InternetGateway](/reference/rack-parameters/InternetGateway) parameter with the ID of the Internet Gateway attached to that VPC. Without it, the Rack's subnets will not have internet access and the installation will fail.

You may also want to configure the subnet CIDR parameters to avoid conflicts with existing subnets in the VPC:

- [Subnet0CIDR](/reference/rack-parameters/Subnet0CIDR), [Subnet1CIDR](/reference/rack-parameters/Subnet1CIDR), [Subnet2CIDR](/reference/rack-parameters/Subnet2CIDR)
- [SubnetPrivate0CIDR](/reference/rack-parameters/SubnetPrivate0CIDR), [SubnetPrivate1CIDR](/reference/rack-parameters/SubnetPrivate1CIDR), [SubnetPrivate2CIDR](/reference/rack-parameters/SubnetPrivate2CIDR)

> **Note:** This parameter should be set at Rack installation time. Changing it on an existing Rack is not supported and may result in infrastructure failures.

```bash
$ convox rack params set ExistingVpc=vpc-0abc123def456
```

> **Note:** This parameter is typically set at install time rather than changed on a running Rack.

## See Also

- [InternetGateway](/reference/rack-parameters/InternetGateway)
- [VPCCIDR](/reference/rack-parameters/VPCCIDR)
- [Private](/reference/rack-parameters/Private)
- [AvailabilityZones](/reference/rack-parameters/AvailabilityZones)
- [Rack Parameters](/reference/rack-parameters) for a full list of available parameters
