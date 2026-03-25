---
title: "FargateServices"
description: "Run all services for this application on AWS Fargate or Fargate Spot instead of EC2."
---

# FargateServices

Launch type override for application services. Set to `Yes` to run all services in [Fargate](https://aws.amazon.com/fargate/), or `Spot` for [Fargate Spot](https://aws.amazon.com/blogs/aws/aws-fargate-spot-now-generally-available/).

| Default value  | `No` |
| Allowed values | `Yes`, `Spot`, `No` |

## Use Cases

- Set to `Yes` when you want serverless compute without managing EC2 instances
- Set to `Spot` to reduce costs for fault-tolerant or stateless workloads that can handle interruptions
- Keep as `No` when you need EC2-specific features such as GPU instances, or when you want to use existing Reserved Instances

## Additional Information

When set to `Yes`, all services in the application run on AWS Fargate with on-demand pricing. When set to `Spot`, services run on Fargate Spot, which can provide up to 70% cost savings but may be interrupted when AWS reclaims capacity.

Individual services can also be configured to run on Fargate at the service level by setting the launch type in the service's `Formation` parameter to `FARGATE` or `FARGATE_SPOT`. The app-level `FargateServices` parameter acts as a blanket override for all services.

Fargate services do not require the Rack to have EC2 instances available. This can simplify capacity planning and eliminate the need to manage instance types or autoscaling groups.

```bash
$ convox apps params set FargateServices=Yes
```

## See Also

- [FargateTimers](/reference/app-parameters/FargateTimers)
- [Private](/reference/app-parameters/Private)
- [Isolate](/reference/app-parameters/Isolate)
- [Rack Parameters](/reference/rack-parameters)
