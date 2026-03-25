---
title: "BuildMethod"
description: "Choose between EC2 and Fargate for the Convox Rack build process."
---

# BuildMethod

Build process type for the Rack. Controls whether builds run on dedicated EC2 instances or Fargate. `ec2` uses a dedicated build instance; `fargate` uses a Fargate task. When using `fargate`, set [FargateBuildCpu](/reference/rack-parameters/FargateBuildCpu) and [FargateBuildMemory](/reference/rack-parameters/FargateBuildMemory).

| Default value  | `ec2`            |
| Allowed values | `ec2`, `fargate` |

## Use Cases

- Using `ec2` (default) for builds that need large disk space, fast I/O, or Docker layer caching across builds
- Switching to `fargate` to eliminate the cost of a continuously running build instance when builds are infrequent
- Using `fargate` when you want builds to run in an isolated, ephemeral environment for security reasons

## Additional Information

When using `ec2`, build resource allocation is controlled by [BuildCpu](/reference/rack-parameters/BuildCpu), [BuildMemory](/reference/rack-parameters/BuildMemory), and [BuildInstance](/reference/rack-parameters/BuildInstance).

When using `fargate`, resource allocation is controlled by [FargateBuildCpu](/reference/rack-parameters/FargateBuildCpu) and [FargateBuildMemory](/reference/rack-parameters/FargateBuildMemory) instead. Note that Fargate builds do not benefit from Docker layer caching, which may result in longer build times.

```bash
$ convox rack params set BuildMethod=fargate
```

## See Also

- [BuildInstance](/reference/rack-parameters/BuildInstance)
- [BuildCpu](/reference/rack-parameters/BuildCpu)
- [BuildMemory](/reference/rack-parameters/BuildMemory)
- [FargateBuildCpu](/reference/rack-parameters/FargateBuildCpu)
- [FargateBuildMemory](/reference/rack-parameters/FargateBuildMemory)
- [BuildVolumeSize](/reference/rack-parameters/BuildVolumeSize)
