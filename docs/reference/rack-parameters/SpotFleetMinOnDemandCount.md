---
title: "SpotFleetMinOnDemandCount"
description: "Set the minimum number of on-demand instances to maintain alongside Spot instances in a Convox Rack Spot Fleet."
---

# SpotFleetMinOnDemandCount

Spot Fleet's minimum on-demand instance count. This ensures a baseline of guaranteed, non-interruptible capacity alongside your Spot instances. The instance type for on-demand instances is taken from the [InstanceType](/reference/rack-parameters/InstanceType) parameter. This parameter can be used only when Spot Fleet is enabled by setting [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice).

| Default value (if HighAvailability disabled) | `1` |
| Default value (if HighAvailability enabled)  | `2` |

The CloudFormation default is `1`. When HighAvailability is enabled, the Rack automatically adjusts this to `2` to ensure redundancy.

## Use Cases

- Increasing the on-demand count for production workloads to guarantee a minimum level of availability even if all Spot instances are reclaimed
- Setting to `1` for development environments where brief interruptions are acceptable
- Setting higher than the default to ensure critical services always have dedicated capacity

## Additional Information

On-demand instances provide guaranteed capacity that cannot be interrupted by AWS, unlike Spot instances which can be reclaimed with a two-minute warning. By maintaining a minimum number of on-demand instances, you ensure that your Rack always has baseline capacity to run essential services.

This parameter only takes effect when Spot Fleet is enabled via [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice).

```bash
$ convox rack params set SpotFleetMinOnDemandCount=3
```

## See Also

- [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice)
- [SpotFleetAllocationStrategy](/reference/rack-parameters/SpotFleetAllocationStrategy)
- [SpotFleetAllowedInstanceTypes](/reference/rack-parameters/SpotFleetAllowedInstanceTypes)
- [SpotFleetExcludedInstanceTypes](/reference/rack-parameters/SpotFleetExcludedInstanceTypes)
- [SpotFleetMinMemoryMiB](/reference/rack-parameters/SpotFleetMinMemoryMiB)
- [SpotFleetMinVcpuCount](/reference/rack-parameters/SpotFleetMinVcpuCount)
- [SpotFleetTargetType](/reference/rack-parameters/SpotFleetTargetType)
- [SpotInstanceBid](/reference/rack-parameters/SpotInstanceBid)
- [InstanceType](/reference/rack-parameters/InstanceType)
- [OnDemandMinCount](/reference/rack-parameters/OnDemandMinCount)
