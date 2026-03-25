---
title: "BuildCpu"
description: "Configure the CPU units allocated to Convox build containers."
---

# BuildCpu

How much CPU should be allocated to builds. The value is specified in ECS CPU units, where 1024 units equals one full vCPU.

| Default value  | `256` |

## Use Cases

- Increasing CPU allocation for large or complex builds that are CPU-bound (e.g., compiling native extensions)
- Reducing CPU allocation on Racks where builds are I/O-bound rather than CPU-bound to leave more capacity for running services

## Additional Information

This value is an ECS CPU reservation for the build container. It works alongside [BuildMemory](/reference/rack-parameters/BuildMemory) to define the resources available to each build. If you are using Fargate builds (via [BuildMethod](/reference/rack-parameters/BuildMethod)), use [FargateBuildCpu](/reference/rack-parameters/FargateBuildCpu) instead.

```bash
$ convox rack params set BuildCpu=512
```

## See Also

- [BuildMemory](/reference/rack-parameters/BuildMemory)
- [BuildInstance](/reference/rack-parameters/BuildInstance)
- [BuildMethod](/reference/rack-parameters/BuildMethod)
- [FargateBuildCpu](/reference/rack-parameters/FargateBuildCpu)
- [BuildVolumeSize](/reference/rack-parameters/BuildVolumeSize)
