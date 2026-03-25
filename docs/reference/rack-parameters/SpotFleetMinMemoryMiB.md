---
title: "SpotFleetMinMemoryMiB"
description: "Set the minimum memory requirement in MiB for instance types in the Convox Rack Spot Fleet."
---

# SpotFleetMinMemoryMiB

Spot Fleet's minimum memory in MiB. Instance types with less memory than this value are excluded from the Spot Fleet's instance selection. This parameter can be used only when Spot Fleet is enabled by setting [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice).

| Default value  | `1000` |

## Use Cases

- Increasing the minimum memory to ensure Spot instances can handle memory-intensive workloads
- Setting a higher threshold to exclude small instance types that would not have enough memory for your application containers
- Keeping the default for general-purpose workloads where the Spot Fleet should consider a wide range of instance types

## Additional Information

This parameter works with [SpotFleetMinVcpuCount](/reference/rack-parameters/SpotFleetMinVcpuCount) to define the minimum hardware requirements for instances selected by the Spot Fleet. Together, these parameters filter out instance types that would be too small to run your workloads effectively.

You can further refine instance selection using [SpotFleetAllowedInstanceTypes](/reference/rack-parameters/SpotFleetAllowedInstanceTypes) or [SpotFleetExcludedInstanceTypes](/reference/rack-parameters/SpotFleetExcludedInstanceTypes).

This parameter only takes effect when Spot Fleet is enabled via [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice).

```bash
$ convox rack params set SpotFleetMinMemoryMiB=2048
```

## See Also

- [SpotFleetMinVcpuCount](/reference/rack-parameters/SpotFleetMinVcpuCount)
- [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice)
- [SpotFleetAllowedInstanceTypes](/reference/rack-parameters/SpotFleetAllowedInstanceTypes)
- [SpotFleetExcludedInstanceTypes](/reference/rack-parameters/SpotFleetExcludedInstanceTypes)
- [SpotFleetAllocationStrategy](/reference/rack-parameters/SpotFleetAllocationStrategy)
- [SpotFleetMinOnDemandCount](/reference/rack-parameters/SpotFleetMinOnDemandCount)
- [SpotFleetTargetType](/reference/rack-parameters/SpotFleetTargetType)
- [SpotInstanceBid](/reference/rack-parameters/SpotInstanceBid)
