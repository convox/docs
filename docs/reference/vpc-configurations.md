---
title: "VPC Configurations"
---

When you install a Rack, most of the AWS resources used by Convox are launched inside of a new [VPC](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_Introduction.html). That default behavior might not be sufficient in all cases, so we've covered a few of the common non-default VPC configurations below.

1. [Installing into an existing VPC](#installing-into-an-existing-vpc)
1. [Installing a private Rack into an existing VPC](#installing-a-private-rack-into-an-existing-vpc)
1. [Peering VPCs in the same region](#peering-vpcs-in-the-same-region)

## Installing into an existing VPC

By default, Convox Rack installations create a new VPC with subnets in two or three (when available) Availability Zones in your chosen AWS Region. If you'd like to install a Convox Rack into an existing VPC, we recommend setting aside a /24 block in each of three Availability Zones.

To install a Rack into an existing VPC, you'll need to provide:

* the VPC ID
* the VPC CIDR
* the CIDRs of the blocks into which the Convox subnets will be installed
* the Internet Gateway ID

Also, make sure your VPC:

* Has an Internet Gateway attached
* Has `DNS resolution` and `DNS hostnames` options enabled

See [Rack Parameters](/reference/rack-parameters) for the specific parameter names you will need to configure during Rack installation.

### Choosing suitable CIDR blocks

Your existing VPC has a CIDR block, and each of your existing subnets has its own CIDR block within that larger VPC block. From the remaining addresses in your VPC CIDR block, Convox will create an additional subnet in each Availability Zone in which you'd like to run Convox instances. Convox recommends three subnets with /24 CIDR blocks to give your Convox installation 254 addresses per subnet.

## Installing a private Rack into an existing VPC

Installing a private Rack into an existing VPC requires specifying a few more options in the **Advanced Network and Compute Options** section of the Convox web installer:

- `ExistingVpc=vpc-id`
- `InternetGateway=igw-id`
- `Subnet0CIDR=172.0.1.0/24`
- `Subnet1CIDR=172.0.2.0/24`
- `Subnet2CIDR=172.0.3.0/24`
- `SubnetPrivate0=172.0.4.0/24`
- `SubnetPrivate1=172.0.5.0/24`
- `SubnetPrivate2=172.0.6.0/24`
- `VPCCIDR=172.0.0.0/16` (make sure to create this block on your custom VPC first)

Even when running in non-private mode (`Private=No`) we still use private subnets for creating endpoints with some AWS services, so don't forget to specify them.

If you want a Private rack, also add:

- `Private=Yes`

## Peering VPCs

An alternative to installing a Convox Rack into an existing VPC is to install the Rack into its own isolated VPC and then peer that VPC with another containing your non-Convox infrastructure.

> A VPC peering connection allows you to route traffic between the peer VPCs using private IP addresses; as if they are part of the same network.

 The above excerpt comes from the AWS [Peering Guide](http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide/Welcome.html), a great place to learn more about this technique.

If you are ready to set up a peering connection between two VPCs, the [Working with VPC Peering Connections](http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide/working-with-vpc-peering.html) section of that guide walks you through the steps, which include the following and more:

* Creating a VPC Peering Connection
* Accepting a VPC Peering Connection
* Updating Route Tables for Your VPC Peering Connection
* Updating Your Security Groups to Reference Peered VPC Security Groups

Keep in mind that VPC Peering has a number of [limitations](http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide/vpc-peering-overview.html#vpc-peering-limitations) that can complicate its setup. For example, you cannot create a VPC peering connection between VPCs that have matching or overlapping CIDR blocks.

## See also

* [Rack Parameters](/docs/rack-parameters)
