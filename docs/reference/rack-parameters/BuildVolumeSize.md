---
title: "BuildVolumeSize"
description: "Set the default build disk size in GB for the Convox Rack build instance."
---

# BuildVolumeSize

Build disk size in GB for the dedicated build instance's EBS volume.

| Default value  | `100` |

## Use Cases

- Increasing disk size for builds that produce large Docker images or pull many large base images
- Increasing when builds fail with "No space left on device" errors
- Reducing on Racks with small, lightweight builds to lower EBS storage costs

## Additional Information

> **Note:** Getting errors like "No space left on device" on your builds (not your running applications)? Extend the space on the device by increasing this parameter.

This parameter controls the EBS volume size attached to the build instance. It does not affect the volume size of runtime instances (see [VolumeSize](/reference/rack-parameters/VolumeSize) for that).

```bash
$ convox rack params set BuildVolumeSize=200
```

## See Also

- [VolumeSize](/reference/rack-parameters/VolumeSize)
- [BuildInstance](/reference/rack-parameters/BuildInstance)
- [BuildMemory](/reference/rack-parameters/BuildMemory)
- [BuildCpu](/reference/rack-parameters/BuildCpu)
- [BuildMethod](/reference/rack-parameters/BuildMethod)
