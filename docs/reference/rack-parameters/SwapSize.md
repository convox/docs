---
title: "SwapSize"
description: "Set the swap volume size in GB for Convox Rack EC2 instances."
---

# SwapSize

Default swap volume size in GB. Set this value to `0` to disable swap.

| Default value  | `5` |

## Use Cases

- Increasing swap size for Racks running memory-intensive workloads that occasionally exceed physical memory
- Setting to `0` to disable swap entirely for latency-sensitive applications where swap thrashing would be unacceptable
- Reducing swap size on instances with large amounts of RAM where swap is unlikely to be needed

## Additional Information

Swap provides a safety net when containers approach their memory limits, allowing the operating system to temporarily page memory to disk rather than immediately killing processes. However, swap is significantly slower than RAM, so relying heavily on swap indicates that instances need more memory.

If your applications are frequently using swap, consider increasing the [InstanceType](/reference/rack-parameters/InstanceType) to an instance with more memory, or increasing the memory reservations for individual services.

```bash
$ convox rack params set SwapSize=10
```

## See Also

- [VolumeSize](/reference/rack-parameters/VolumeSize)
- [InstanceType](/reference/rack-parameters/InstanceType)
