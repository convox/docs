---
title: "PlaceLambdaInVpc"
description: "Place Convox-managed Lambda functions inside the Rack VPC for private network access."
---

# PlaceLambdaInVpc

VPC placement for Convox-related Lambda functions. This is useful when the Rack is configured as [Private](/reference/rack-parameters/Private) and the Lambda functions need access to resources within the VPC.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Use Cases

- Enabling VPC placement for Lambdas when the Rack is private and Lambdas need to communicate with VPC resources
- Meeting security requirements that mandate all compute resources run within the VPC
- Allowing Lambda functions to access private RDS instances, ElastiCache clusters, or other VPC-only resources used by the Rack

## Additional Information

When set to `Yes`, Convox-managed Lambda functions are deployed into the Rack's VPC subnets and associated with the Rack's security groups. This allows the Lambdas to access resources within the VPC but also means they are subject to VPC networking rules.

Lambda functions placed in a VPC do not have internet access by default unless the VPC has a NAT Gateway configured. If your Rack is private, this is typically already the case.

Enabling this parameter increases Lambda cold start times slightly because ENI (Elastic Network Interface) attachment is required.

```bash
$ convox rack params set PlaceLambdaInVpc=Yes
```

## See Also

- [Private](/reference/rack-parameters/Private)
- [PrivateApi](/reference/rack-parameters/PrivateApi)
- [ExistingVpc](/reference/rack-parameters/ExistingVpc)
