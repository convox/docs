---
title: "Private Networking"
description: "Configure Convox Racks for private networking with isolated subnets, NAT gateways, and static outbound IPs."
---

# Private Networking

Convox Racks are installed using best practices for privacy and security. In cases with exceptional needs for isolation, Convox offers a "private" Rack mode that is easy to configure.

## Public vs. Private Racks

Convox Rack creates a cluster of server instances into which your apps are deployed. In the default "public" mode, these instances are connected to a public subnet. Security Groups are used to keep out unwanted inbound traffic, and outbound traffic to the Internet is unrestricted.

By contrast, server instances in a private Rack have no direct connection to the Internet. Inbound traffic is routed from the load balancer to a private subnet, and outbound traffic passes through a NAT gateway.

## Installing a Private Rack

A Rack can be installed as private by setting the parameter `Private` to `Yes` in Console during Rack installation.

## Converting an Existing Rack to Private

It's not currently possible to toggle a rack between Public and Private due to how EFS mount targets work (one subnet per AZ at a time, so it's not possible to switch from public to private subnets, or vice versa).

## Privacy Parameters

Convox provides three separate privacy controls that can be combined for different levels of isolation:

### Private

The [Private](/reference/rack-parameters/Private) parameter places all Rack instances (both workload and build) into private subnets. This is the broadest privacy control. When set to `Yes`, instances have no direct internet access — all outbound traffic passes through a NAT gateway.

### PrivateBuild

The [PrivateBuild](/reference/rack-parameters/PrivateBuild) parameter places only the build instances into private subnets while keeping workload instances in public subnets. This is useful when you want to reduce the attack surface for build infrastructure without requiring a fully private Rack. This parameter has no effect if `Private` is already set to `Yes`.

### PrivateApi

The [PrivateApi](/reference/rack-parameters/PrivateApi) parameter makes the Rack API load balancer internal, so it is only accessible from within the VPC or through a VPN/peering connection. This is independent of `Private` — you can have a private API with public instances, or vice versa. When enabled, you will need VPN, Direct Connect, or VPC peering to manage the Rack via the CLI or CI/CD.

## Why Use a Private Rack?

The benefits of running a Rack in private mode include:

- **Network Isolation**: Private mode adds an extra layer of network security. The private subnet means all inbound traffic must pass through an ALB that explicitly listens only on known ports. The NAT gateway allows outbound traffic from the private subnet.

- **Static Outbound IP**: The NAT gateway created to handle outbound traffic has a static IP address. This can be useful in cases where your applications need to make requests to services that require IP whitelisting.

## Cost

Private networking mode requires the provisioning of up to three NAT gateways (one for each availability zone) which are charged based on time and traffic. This will increase the monthly cost of your rack. More pricing details can be found in the [AWS docs](https://aws.amazon.com/vpc/pricing/).

## See Also

- [Private](/reference/rack-parameters/Private)
- [PrivateApi](/reference/rack-parameters/PrivateApi)
- [PrivateBuild](/reference/rack-parameters/PrivateBuild)
- [PrivateApiSecurityGroup](/reference/rack-parameters/PrivateApiSecurityGroup)
- [VPC Configurations](/networking/vpc-configurations)
- [Internal Services](/networking/internal-services)
