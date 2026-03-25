---
title: "BuildMemory"
description: "Configure the memory allocated to Convox build containers."
---

# BuildMemory

Memory allocation (in MB) for build containers on each build.

| Default value  | `1000` |

## Use Cases

- Increasing memory for builds that compile large applications or use memory-intensive build tools
- Decreasing memory on Racks where builds are lightweight to allow more concurrent builds on the same instance
- Tuning alongside [BuildInstance](/reference/rack-parameters/BuildInstance) to optimize the cost-to-performance ratio for your build workloads

## Additional Information

> **Note:** Getting build errors like "Starting build... ERROR: not enough memory available to start process"? Either reduce this parameter, or change the [InstanceType](/reference/rack-parameters/InstanceType) parameter to an [instance type](https://aws.amazon.com/ec2/instance-types/) with more memory.

> **Warning:** If you set BuildMemory to an amount that is more than half of the total memory available to the build instance, you can only run one build at a time. If this value is too high, builds may fail.

If you are using Fargate builds (via [BuildMethod](/reference/rack-parameters/BuildMethod)), use [FargateBuildMemory](/reference/rack-parameters/FargateBuildMemory) instead.

```bash
$ convox rack params set BuildMemory=2048
```

## See Also

- [BuildCpu](/reference/rack-parameters/BuildCpu)
- [BuildInstance](/reference/rack-parameters/BuildInstance)
- [BuildMethod](/reference/rack-parameters/BuildMethod)
- [FargateBuildMemory](/reference/rack-parameters/FargateBuildMemory)
- [BuildVolumeSize](/reference/rack-parameters/BuildVolumeSize)
- [InstanceType](/reference/rack-parameters/InstanceType)
