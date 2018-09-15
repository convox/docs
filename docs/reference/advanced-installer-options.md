---
title: "Advanced Rack Installer"
---

When installing a Rack via Console, click the **Advanced Network and Compute Options** link to display available advanced installation options.

### Network options

![Advanced network installation options](/assets/images/docs/advanced-rack-install-network.png) _Advanced network installation options_

#### Use an existing VPC

By default, private Racks are installed into new VPCs created by the Convox installer for that purpose.

If instead you would like to install the Rack into a VPC that already exists, select it from the **Use an existing VPC** dropdown menu, which is prepopulated with the existing VPCs found in your AWS account (within the region selected above).

For more detailed information, see [VPC Configurations](/docs/vpc-configurations).

#### Use an existing Internet Gateway

If you've selected an existing VPC, any existing Internet Gateways associated with it will be available in the **Use an existing Internet Gateway** dropdown.

#### VPC CIDR

If you haven't selected an existing VPC from the **Use an existing VPC** dropdown menu, you can specify a new VPC CIDR to be created in the **VPC CIDR** field.

#### Subnet CIDRs

In the **Subnet CIDRs** fields, you can provide the CIDRs of the subnets into which Convox should be installed. For more information, see [Choosing suitable CIDR blocks](/docs/vpc-configurations#choosing-suitable-cidr-blocks).

#### Private

Select the **Private** checkbox to install the Rack in [private networking mode](/docs/private-networking/) which will use private subnets and NAT gateways to shield instances.

This can be changed later by setting the [`Private` Rack parameter](/docs/rack-parameters/#private).

### Compute options

![Advanced compute installation options](/assets/images/docs/advanced-rack-install-compute.png) _Advanced compute installation options_

#### Instance type

In the **Instance type** dropdown, you can select the type of EC2 instance to be created for your Rack.

This can be changed later by setting the [`InstanceType` Rack parameter](/docs/rack-parameters/#instancetype).

#### Autoscale

Select **Yes** from the **Autoscale** dropdown menu to make the Rack scale its own instance count based on the needs of the containers it provisions.

This can be changed later by setting the [`Autoscale` Rack parameter](/docs/rack-parameters/#autoscale).

For more information, see [Autoscaling](/docs/scaling#autoscale).

#### Instance count

If the **Autoscale** option is disabled, you can specify how many instances you want to be provisioned with this Rack by modifying the number in the **Instance count** field.

This can be changed later using the [`convox rack scale` command](/docs/scaling/#scaling-the-rack) (or enabling Autoscaling).

#### Dedicated instances

Select the **Dedicated instances** checkbox to create EC2 instances on dedicated hardware.

## See also

* [Rack Parameters](/docs/rack-parameters)

