---
title: "DefaultAmi"
description: "Defines the default ECS-optimized AMI used for x86_64-based Convox Rack instances."
---

# DefaultAmi

Default Amazon Machine Image (AMI) for x86_64-based Rack instances. This allows racks to automatically use the latest recommended ECS-optimized AMI without manual intervention.

| Default value  | `/aws/service/ecs/optimized-ami/amazon-linux-2/recommended/image_id` |
| Allowed values | AWS SSM AMI path (e.g., `/aws/service/ecs/optimized-ami/amazon-linux-2023/recommended/image_id`) |

By default, Convox racks use the **Amazon Linux 2 (AL2) ECS-optimized AMI**. However, with AL2 nearing deprecation, you can switch to **Amazon Linux 2023 (AL2023)** by setting:

```bash
$ convox rack params set DefaultAmi="/aws/service/ecs/optimized-ami/amazon-linux-2023/recommended/image_id"
```

If the `Ami` rack parameter is set, `DefaultAmi` will be ignored, and the explicitly set `Ami` value will be used instead.

## Use Cases

- Upgrade to Amazon Linux 2023 for longer support lifecycle, improved security defaults, and updated packages.
- Pin a specific AMI path for testing before rolling it out across all racks.
- Roll back to a known-good AMI path if a new ECS-optimized AMI introduces regressions.

## Additional Information

The value is an AWS Systems Manager (SSM) parameter path, not a direct AMI ID. AWS automatically updates the AMI ID that this path resolves to whenever a new ECS-optimized image is published. This means your Rack instances will pick up the latest AMI on the next instance replacement (e.g., during a scaling event or parameter update) without requiring manual changes.

If you need to pin a specific AMI ID rather than using the auto-resolving SSM path, use the [Ami](/reference/rack-parameters/Ami) parameter instead.

## See Also

- [DefaultAmiArm](/reference/rack-parameters/DefaultAmiArm)
- [Ami](/reference/rack-parameters/Ami)
- [InstanceType](/reference/rack-parameters/InstanceType)
- [Rack Parameters](/reference/rack-parameters) for a full list of available parameters
