---
title: "DefaultAmiArm"
description: "Defines the default ECS-optimized AMI used for ARM64-based Convox Rack instances."
---

# DefaultAmiArm

Default Amazon Machine Image (AMI) for ARM64-based Rack instances. This ensures ARM-based racks always use the latest ECS-optimized AMI unless manually overridden.

| Default value  | `/aws/service/ecs/optimized-ami/amazon-linux-2/arm64/recommended/image_id` |
| Allowed values | AWS SSM AMI path (e.g., `/aws/service/ecs/optimized-ami/amazon-linux-2023/arm64/recommended/image_id`) |

If your rack runs on **ARM64 architecture**, it will use the **Amazon Linux 2 (AL2) ARM64 ECS-optimized AMI** by default. You can switch to **Amazon Linux 2023 for ARM** by setting:

```bash
$ convox rack params set DefaultAmiArm="/aws/service/ecs/optimized-ami/amazon-linux-2023/arm64/recommended/image_id"
```

If the `Ami` rack parameter is set, `DefaultAmiArm` will be ignored, and the explicitly set `Ami` value will be used instead.

## Use Cases

- Upgrade ARM64 racks to Amazon Linux 2023 for extended support and updated packages.
- Run cost-efficient Graviton-based instances (e.g., `t4g`, `m6g`, `c7g`) with the correct ARM-optimized AMI.
- Test a new ARM64 AMI path before applying it across production racks.

## Additional Information

Like [DefaultAmi](/reference/rack-parameters/DefaultAmi), this value is an AWS Systems Manager (SSM) parameter path that auto-resolves to the latest recommended AMI ID. Your Rack instances will receive updated AMIs on the next instance replacement without manual intervention.

ARM64 instances such as the AWS Graviton family offer improved price-performance for many workloads. Ensure your application containers are built for the `arm64` architecture when using ARM-based instance types.

## See Also

- [DefaultAmi](/reference/rack-parameters/DefaultAmi)
- [Ami](/reference/rack-parameters/Ami)
- [InstanceType](/reference/rack-parameters/InstanceType)
- [Rack Parameters](/reference/rack-parameters) for a full list of available parameters
