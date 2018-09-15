---
title: "Security"
---

Convox always has your security in mind. Below are some of the ways your apps and resources are protected.

## AWS Isolation

Your Convox Rack is installed in your own AWS account. Unlike a multi-tenant PaaS, you never have to worry about other users gaining access to your infrastructure.

[Convox Console](https://console.convox.com) is a multi-tenant management layer that proxies API requests to your Rack and provides team management and workflow tools. Many security measures are baked in such as rollable API keys, programmatic [deploy keys](/docs/deploy-keys), and user roles. However, if you don't want to use the multi-tenant Console you can run your own! [Contact us](mailto:support@convox.com) for information about plans and pricing.

## AWS Permissions

Convox strives to limit the scope of AWS permissions needed to manage a Rack. When you install a Rack, an IAM "KernelUser" is created and granted only the permissions required. Additionally, many of the permissions are scoped down further to only apply to the Rack's resources. As with security in general, this is an ongoing process. Future Rack updates will continue to limit these permissions as much as possible.

Some examples of limited permissions include:
- Access to IAM resources (roles, users, policies, etc) that belong to the `/convox/` path
- Access to DynamoDB tables that only begin with the name of the Rack as a prefix
- Access to RDS instances that only begin with the name of the Rack as a prefix

## VPC Isolation

All of the infrastructure that Convox creates for you runs inside a Virtual Private Cloud ([VPC](https://aws.amazon.com/vpc/)). This provides additional isolation at the networking layer. By default, all resources such as datastores are created in such a way that they can only be accessed from inside the VPC.

## Load Balancers

Convox uses AWS Load Balancers to route traffic to your application. Load balancers and their Security Groups are set up to only listen on the ports you specify and only route requests to the relevant application services.

## Private Networking

If you'd like to take network isolation one step further you can run your Rack in private networking mode, where the Rack instances run in private subnets that access the Internet through NATs. These instances are not routable via the public Internet. Read more on the private networking [doc](/docs/private-networking/).

## Dedicated Instances

If you would like to ensure hardware single tenancy all the way down to the AWS infrastructure level you can do so by passing `Tenancy=dedicated` option to the `convox rack install` command when setting up your Rack.
