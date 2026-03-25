---
title: "SpotFleetTargetType"
description: "Set the unit type used for Spot Fleet target capacity in a Convox Rack."
---

# SpotFleetTargetType

The unit type used for the Spot Fleet target capacity. This determines how the Spot Fleet measures whether it has reached its target capacity. This parameter can be used only when Spot Fleet is enabled by setting [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice).

| Default value  | `units`                           |
| Allowed values | `memory-mib`, `units`, `vcpu` |

## Use Cases

- Using `units` (default) to target a specific number of instances regardless of their size
- Switching to `vcpu` when your workloads are CPU-bound and you want to ensure a minimum total vCPU count across the fleet
- Using `memory-mib` for memory-intensive workloads where total fleet memory matters more than instance count

## Additional Information

The three target types work as follows:

- **units** -- Each instance counts as one unit. The target capacity equals the desired number of instances.
- **vcpu** -- Each instance contributes its vCPU count toward the target. A target of `8` could be fulfilled by one 8-vCPU instance or two 4-vCPU instances.
- **memory-mib** -- Each instance contributes its memory (in MiB) toward the target. This is useful when your primary resource constraint is memory.

This parameter only takes effect when Spot Fleet is enabled via [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice).

```bash
$ convox rack params set SpotFleetTargetType=vcpu
```

## See Also

- [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice)
- [SpotFleetAllocationStrategy](/reference/rack-parameters/SpotFleetAllocationStrategy)
- [SpotFleetAllowedInstanceTypes](/reference/rack-parameters/SpotFleetAllowedInstanceTypes)
- [SpotFleetExcludedInstanceTypes](/reference/rack-parameters/SpotFleetExcludedInstanceTypes)
- [SpotFleetMinMemoryMiB](/reference/rack-parameters/SpotFleetMinMemoryMiB)
- [SpotFleetMinVcpuCount](/reference/rack-parameters/SpotFleetMinVcpuCount)
- [SpotFleetMinOnDemandCount](/reference/rack-parameters/SpotFleetMinOnDemandCount)
- [SpotInstanceBid](/reference/rack-parameters/SpotInstanceBid)
