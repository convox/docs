---
title: "FargateServices"
description: "Run all services for this application on AWS Fargate or Fargate Spot instead of EC2."
---

# FargateServices

Launch type override for application Services. Set to `Yes` to run all Services in [Fargate](https://aws.amazon.com/fargate/), or `Spot` for [Fargate Spot](https://aws.amazon.com/blogs/aws/aws-fargate-spot-now-generally-available/).

| Setting | Value |
|:--|:--|
| Default value  | `No` |
| Allowed values | `Yes`, `Spot`, `No` |

## Use Cases

- Set to `Yes` when you want serverless compute without managing EC2 instances
- Set to `Spot` to reduce costs for fault-tolerant or stateless workloads that can handle interruptions
- Keep as `No` when you need EC2-specific features such as GPU instances, or when you want to use existing Reserved Instances

## Additional Information

When set to `Yes`, all Services in the application run on AWS Fargate with on-demand pricing. When set to `Spot`, Services run on Fargate Spot, which can provide up to 70% cost savings but may be interrupted when AWS reclaims capacity.

Individual Services can also run on Fargate. Add `FARGATE` or `FARGATE_SPOT` as a fourth value to the Service's `<Service>Formation` parameter, after the count, CPU and memory that `convox apps params` shows, for example `WorkerFormation=2,256,512,FARGATE` for a Service named `worker`. The app-level `FargateServices` parameter acts as a blanket override for all Services.

Fargate Services do not require the Rack to have EC2 instances available. This can simplify capacity planning and eliminate the need to manage instance types or autoscaling groups.

```bash
$ convox apps params set FargateServices=Yes
```

### Moving a Deployed App

An App that already runs on EC2 can move to Fargate with `FargateServices` or a Service's `Formation`. The Services keep their names, endpoints and Release. A Service with `port:` or `nlb:` ports can fail some requests for a short time while it moves. Setting `FargateServices=No`, or removing the launch type from `Formation`, moves the Services back to EC2.

The move requires rack version 20261103122112 or newer. On an App last deployed or promoted before the Rack moved to that version, deploy or promote the App once first, or the parameter change rolls back and the App stays on EC2:

```bash
$ convox releases promote RABCDEFGHIJ -a myapp --wait
$ convox apps params set FargateServices=Yes -a myapp --wait
```

Use the App's current Release, shown by `convox apps info -a myapp`.

Agent Services (`agent: true`) cannot run on Fargate. On an App with an agent Service, keep `FargateServices=No` and set `FARGATE` or `FARGATE_SPOT` in the other Services' `Formation` instead.

### Moving the Rack to an Older Version

Before moving the Rack to a version older than 20261103122112, move any `internalAndExternal` Service that runs on Fargate back to EC2. Otherwise the App's first deploy or promote on the older version removes the external router's access to that Service's processes. Its external domain stops answering, and the deploy does not finish for hours unless you cancel it with `convox apps cancel -a myapp`.

## See Also

- [FargateTimers](/reference/app-parameters/FargateTimers)
- [Private](/reference/app-parameters/Private)
- [Isolate](/reference/app-parameters/Isolate)
- [Rack Parameters](/reference/rack-parameters)
