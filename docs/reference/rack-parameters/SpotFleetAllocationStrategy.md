---
title: "SpotFleetAllocationStrategy"
description: "Set the allocation strategy for EC2 Spot Fleet instances in a Convox Rack."
---

# SpotFleetAllocationStrategy

The Spot Fleet allocation strategy. This parameter can be used only when Spot Fleet is enabled by setting [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice).

| Default value  | `lowestPrice`                                         |
| Allowed values | `lowestPrice`, `diversified`, `capacityOptimized` |

## Use Cases

- Using `lowestPrice` (default) to minimize costs by selecting the cheapest available Spot pools
- Switching to `capacityOptimized` for workloads that need higher availability by selecting pools with the most available capacity, reducing the chance of interruption
- Using `diversified` to spread instances across multiple Spot pools, reducing the impact of a single pool being reclaimed

## Additional Information

The three strategies behave as follows:

- **lowestPrice** -- Spot instances are launched from the pool with the lowest price. This maximizes cost savings but may lead to higher interruption rates if the cheapest pool runs low on capacity.
- **diversified** -- Spot instances are distributed across all available pools. This reduces the risk of all instances being interrupted simultaneously.
- **capacityOptimized** -- Spot instances are launched from pools with the most available capacity. This is the best choice for workloads that have a high cost of interruption.

This parameter only takes effect when Spot Fleet is enabled via [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice).

```bash
$ convox rack params set SpotFleetAllocationStrategy=capacityOptimized
```

## See Also

- [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice)
- [SpotFleetAllowedInstanceTypes](/reference/rack-parameters/SpotFleetAllowedInstanceTypes)
- [SpotFleetExcludedInstanceTypes](/reference/rack-parameters/SpotFleetExcludedInstanceTypes)
- [SpotFleetMinMemoryMiB](/reference/rack-parameters/SpotFleetMinMemoryMiB)
- [SpotFleetMinVcpuCount](/reference/rack-parameters/SpotFleetMinVcpuCount)
- [SpotFleetMinOnDemandCount](/reference/rack-parameters/SpotFleetMinOnDemandCount)
- [SpotFleetTargetType](/reference/rack-parameters/SpotFleetTargetType)
- [SpotInstanceBid](/reference/rack-parameters/SpotInstanceBid)
