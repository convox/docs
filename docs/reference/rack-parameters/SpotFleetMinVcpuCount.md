---
title: "SpotFleetMinVcpuCount"
description: "Set the minimum vCPU count for instance types in the Convox Rack Spot Fleet."
---

# SpotFleetMinVcpuCount

Spot Fleet's minimum vCPU count. Instance types with fewer vCPUs than this value are excluded from the Spot Fleet's instance selection. This parameter can be used only when Spot Fleet is enabled by setting [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice).

| Default value  | `0` |

## Use Cases

- Increasing the minimum vCPU count to ensure Spot instances can handle CPU-intensive workloads
- Setting a higher threshold to exclude small instance types (e.g., `t3.nano`, `t3.micro`) that would be too small for your containers
- Keeping the default of `0` to allow the widest possible range of instance types for maximum Spot availability

## Additional Information

This parameter works with [SpotFleetMinMemoryMiB](/reference/rack-parameters/SpotFleetMinMemoryMiB) to define the minimum hardware requirements for instances selected by the Spot Fleet. Together, these parameters filter out instance types that would be too small to run your workloads effectively.

A value of `0` means no vCPU minimum is enforced, allowing the Spot Fleet to consider all instance types that meet other criteria.

You can further refine instance selection using [SpotFleetAllowedInstanceTypes](/reference/rack-parameters/SpotFleetAllowedInstanceTypes) or [SpotFleetExcludedInstanceTypes](/reference/rack-parameters/SpotFleetExcludedInstanceTypes).

This parameter only takes effect when Spot Fleet is enabled via [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice).

```bash
$ convox rack params set SpotFleetMinVcpuCount=2
```

## See Also

- [SpotFleetMinMemoryMiB](/reference/rack-parameters/SpotFleetMinMemoryMiB)
- [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice)
- [SpotFleetAllowedInstanceTypes](/reference/rack-parameters/SpotFleetAllowedInstanceTypes)
- [SpotFleetExcludedInstanceTypes](/reference/rack-parameters/SpotFleetExcludedInstanceTypes)
- [SpotFleetAllocationStrategy](/reference/rack-parameters/SpotFleetAllocationStrategy)
- [SpotFleetMinOnDemandCount](/reference/rack-parameters/SpotFleetMinOnDemandCount)
- [SpotFleetTargetType](/reference/rack-parameters/SpotFleetTargetType)
- [SpotInstanceBid](/reference/rack-parameters/SpotInstanceBid)
