---
title: "FargateTimers"
description: "Run all timers for this application on AWS Fargate or Fargate Spot instead of EC2."
---

# FargateTimers

Launch type override for application timers. Set to `Yes` to run all timers in [Fargate](https://aws.amazon.com/fargate/), or `Spot` for [Fargate Spot](https://aws.amazon.com/blogs/aws/aws-fargate-spot-now-generally-available/).

| Default value  | `No` |
| Allowed values | `Yes`, `Spot`, `No` |

## Use Cases

- Set to `Yes` when timer tasks should run on serverless compute without depending on EC2 instance availability
- Set to `Spot` to reduce costs for timer workloads that can tolerate occasional interruptions
- Keep as `No` when timers should run on existing EC2 instances to avoid Fargate per-task pricing

## Additional Information

When set to `Yes`, all timers in the application are launched as Fargate tasks. When set to `Spot`, timers run on Fargate Spot for cost savings. Timer tasks are short-lived by nature, making them good candidates for Fargate Spot since interruptions are less impactful on short-running workloads.

This parameter is independent of [FargateServices](/reference/app-parameters/FargateServices). You can run services on EC2 while running timers on Fargate, or vice versa.

```bash
$ convox apps params set FargateTimers=Yes
```

## See Also

- [FargateServices](/reference/app-parameters/FargateServices)
- [Private](/reference/app-parameters/Private)
- [PlaceLambdaInVpc](/reference/app-parameters/PlaceLambdaInVpc)
- [Rack Parameters](/reference/rack-parameters)
