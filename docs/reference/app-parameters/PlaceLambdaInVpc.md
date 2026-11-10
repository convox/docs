---
title: "PlaceLambdaInVpc"
description: "Place the App's timer launcher Lambda function inside the VPC when the App's Private parameter is Yes."
---

# PlaceLambdaInVpc

VPC placement control for the App's timer launcher Lambda function. When set to `Yes` and [Private](/reference/app-parameters/Private) is enabled, the function is placed inside the VPC.

| Setting | Value |
|:--|:--|
| Default value  | `No` |
| Allowed values | `Yes`, `No` |

## Use Cases

- Enable when your Private Rack has no internet access and Lambda functions need to reach VPC resources
- Enable for compliance requirements that mandate all compute resources run within the VPC
- Keep disabled when Lambda functions do not need VPC access, as VPC-attached Lambdas have slower cold start times

## Additional Information

This parameter only takes effect when the App's [Private](/reference/app-parameters/Private) parameter is `Yes`.

The only Lambda function in an App's stack is the timer launcher. It exists only when the App has timers and [FargateTimers](/reference/app-parameters/FargateTimers) is `Yes` or `Spot`, so this parameter has no effect on an App whose timers run on EC2.

When enabled, the timer launcher is configured with VPC networking, placing it in the Rack's private subnets with the instances security group. This allows it to access resources within the VPC but requires a NAT gateway or VPC endpoints for any external internet access.

VPC-attached Lambda functions experience longer cold start times compared to non-VPC Lambdas due to the overhead of attaching an elastic network interface (ENI). Consider this trade-off when enabling this parameter.

Each deploy or promote sets this parameter to the Rack's [PlaceLambdaInVpc](/reference/rack-parameters/PlaceLambdaInVpc) value. To keep a different value for one App, set it under `params:` in `convox.yml`. See [Parameters Set From the Rack](/reference/app-parameters#parameters-set-from-the-rack).

```yaml
params:
  PlaceLambdaInVpc: "Yes"
```

## See Also

- [Private](/reference/app-parameters/Private)
- [FargateTimers](/reference/app-parameters/FargateTimers)
- [Rack Parameters](/reference/rack-parameters)
- [Private Networking](/networking/private-networking)
