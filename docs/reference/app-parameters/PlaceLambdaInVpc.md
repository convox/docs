---
title: "PlaceLambdaInVpc"
description: "Place Convox-related Lambda functions inside the VPC when the Rack is running in Private mode."
---

# PlaceLambdaInVpc

VPC placement control for Convox-related Lambda functions. When set to `Yes` and [Private](/reference/app-parameters/Private) is enabled, Lambda functions are placed inside the VPC.

| Default value  | `No` |
| Allowed values | `Yes`, `No` |

## Use Cases

- Enable when your Private Rack has no internet access and Lambda functions need to reach VPC resources
- Enable for compliance requirements that mandate all compute resources run within the VPC
- Keep disabled when Lambda functions do not need VPC access, as VPC-attached Lambdas have slower cold start times

## Additional Information

This parameter only takes effect when the [Private](/reference/app-parameters/Private) app parameter (or the Rack-level Private parameter) is set to `Yes`. If the Rack is not private, this setting has no effect.

When enabled, the Convox Lambda functions (such as the timer launcher) are configured with VPC networking, placing them in the Rack's private subnets with the instances security group. This allows the Lambda functions to access resources within the VPC but requires a NAT gateway or VPC endpoints for any external internet access.

VPC-attached Lambda functions experience longer cold start times compared to non-VPC Lambdas due to the overhead of attaching an elastic network interface (ENI). Consider this trade-off when enabling this parameter.

```bash
$ convox apps params set PlaceLambdaInVpc=Yes
```

## See Also

- [Private](/reference/app-parameters/Private)
- [FargateTimers](/reference/app-parameters/FargateTimers)
- [Rack Parameters](/reference/rack-parameters)
- [Private Networking](/networking/private-networking)
