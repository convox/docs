---
title: "BuildVolumeSize"
description: "Set the default build disk size in GB for the Convox Rack build instance."
---

# BuildVolumeSize

Build disk size in GB for the dedicated build instance's EBS volume.

| Setting | Value |
|:--|:--|
| Default value  | `100` |

## Use Cases

- Increasing disk size for builds that produce large Docker images or pull many large base images
- Increasing when builds fail with "No space left on device" errors
- Reducing on Racks with small, lightweight builds to lower EBS storage costs

## Additional Information

> **Note:** Getting errors like "No space left on device" on your builds (not your running applications)? On a Rack with [BuildMethod](/reference/rack-parameters/BuildMethod)=`ec2`, extend the space on the device by increasing this parameter. On a Rack with `BuildMethod`=`fargate`, increase [FargateBuildVolumeSize](/reference/rack-parameters/FargateBuildVolumeSize) instead.

This parameter controls the EBS volume size attached to the build instance. It does not affect the volume size of runtime instances (see [VolumeSize](/reference/rack-parameters/VolumeSize) for that).

A Rack that builds on Fargate runs no build instance, so a change to this parameter is accepted and applied to the stack but never reaches a running builder. Use [FargateBuildVolumeSize](/reference/rack-parameters/FargateBuildVolumeSize) to size the Fargate builder's disk instead.

```bash
$ convox rack params set BuildVolumeSize=200
```

## See Also

- [VolumeSize](/reference/rack-parameters/VolumeSize)
- [FargateBuildVolumeSize](/reference/rack-parameters/FargateBuildVolumeSize)
- [BuildInstance](/reference/rack-parameters/BuildInstance)
- [BuildMemory](/reference/rack-parameters/BuildMemory)
- [BuildCpu](/reference/rack-parameters/BuildCpu)
- [BuildMethod](/reference/rack-parameters/BuildMethod)
