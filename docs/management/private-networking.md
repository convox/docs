---
title: "Private Networking"
---

Convox Racks are installed using best practices for privacy and security. In cases with exceptional needs for isolation, Convox offers a "private" Rack mode that is easy to configure.

## Public vs. Private Racks

Convox Rack creates a cluster of server instances into which your apps are deployed. In the default "public" mode, these instances are connected to a public subnet. Security Groups are used to keep out unwanted inbound traffic, and outbound traffic to the Internet is unrestricted.

By contrast, server instances in a private Rack have no direct connection to the Internet. Inbound traffic is routed from the load balancer to a private subnet, and outbound traffic passes through a NAT gateway.

## Installing a Private Rack

A Rack can be installed as private by selecting the **Private** checkbox in the Rack installation page in the Convox console.

We recommend [installing via the Console](/docs/installing-a-rack/), but it's also possible using the Convox CLI:

```
convox rack install Private=Yes
```

Note that Racks installed via the CLI will need to be added manually to the Console.

## Converting an Existing Rack to Private

An existing Rack can be flipped to private mode:

```
convox rack params set Private=Yes
```

## Why Use a Private Rack?

The benefits of running a Rack in private mode include:

- **Network Isolation**: Private mode adds an extra layer of network security. The private subnet means all inbound traffic must pass through an ELB that explicitly listens only on known ports. The NAT gateway allows outbound traffic from the private subnet.

- **Static Outbound IP**: The NAT gateway created to handle outbound traffic has a static IP address. This can be useful in cases where your applications need to make requests to services that require IP whitelisting.

## Cost

Private networking mode requires the provisioning of up to three NAT gateways (one for each availability zone) which are charged based on time and traffic. This will increase the monthly cost of your rack. More pricing details can be found in the [AWS docs](https://aws.amazon.com/vpc/pricing/).

## See also

* [VPC Configurations](/docs/vpc-configurations)
