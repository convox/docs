---
title: "VolumeSize"
description: "Set the default EBS volume size in GB for Convox Rack EC2 instances."
---

# VolumeSize

Default disk size (in gibibytes) of the EBS volume attached to each EC2 instance in the cluster. This is the root volume that stores the operating system, Docker images, and container data for running applications.

| Default value  | `50` |

## Use Cases

- Increasing volume size for Racks running applications with large Docker images or many concurrent containers
- Reducing volume size for cost optimization on Racks with lightweight workloads
- Increasing when instances report "no space left on device" errors during normal application operation (as opposed to build-time errors, which are controlled by [BuildVolumeSize](/reference/rack-parameters/BuildVolumeSize))

## Additional Information

This parameter controls the size of the EBS volume on runtime instances. For build instances, use [BuildVolumeSize](/reference/rack-parameters/BuildVolumeSize) instead.

If you are running out of disk space on runtime instances, you can increase this value:

```bash
$ convox rack params set VolumeSize=100
```

The change takes effect as instances are replaced (e.g., during scaling events or parameter updates). Existing instances will retain their current volume size until replaced.

## See Also

- [BuildVolumeSize](/reference/rack-parameters/BuildVolumeSize)
- [SwapSize](/reference/rack-parameters/SwapSize)
- [InstanceType](/reference/rack-parameters/InstanceType)
- [EncryptEbs](/reference/rack-parameters/EncryptEbs)
