---
title: "SpotFleetExcludedInstanceTypes"
description: "Exclude specific EC2 instance types from the Convox Rack Spot Fleet."
---

# SpotFleetExcludedInstanceTypes

Comma-separated list of excluded instance types in the Spot Fleet. This parameter cannot be used together with [SpotFleetAllowedInstanceTypes](/reference/rack-parameters/SpotFleetAllowedInstanceTypes). This parameter can be used only when Spot Fleet is enabled by setting [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice).

| Default value  | "" |

Wildcard patterns are supported. Examples: `m5.8xlarge`, `c5*.*`, `m5a.*`, `r*`, `*3*`.

## Use Cases

- Excluding older-generation instance types that may have lower performance characteristics
- Blocking specific instance families known to have high interruption rates in your region
- Removing instance types that are incompatible with your application's resource profile

## Additional Information

When this parameter is blank (the default), no instance types are excluded from the Spot Fleet. The fleet considers all types that meet the minimum resource requirements defined by [SpotFleetMinMemoryMiB](/reference/rack-parameters/SpotFleetMinMemoryMiB) and [SpotFleetMinVcpuCount](/reference/rack-parameters/SpotFleetMinVcpuCount).

If you need to define an explicit allow list instead of an exclusion list, use [SpotFleetAllowedInstanceTypes](/reference/rack-parameters/SpotFleetAllowedInstanceTypes). The two parameters are mutually exclusive -- if both are set, `SpotFleetAllowedInstanceTypes` takes precedence.

This parameter only takes effect when Spot Fleet is enabled via [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice).

```bash
$ convox rack params set SpotFleetExcludedInstanceTypes="t2.nano,t2.micro"
```

## See Also

- [SpotFleetAllowedInstanceTypes](/reference/rack-parameters/SpotFleetAllowedInstanceTypes)
- [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice)
- [SpotFleetAllocationStrategy](/reference/rack-parameters/SpotFleetAllocationStrategy)
- [SpotFleetMinMemoryMiB](/reference/rack-parameters/SpotFleetMinMemoryMiB)
- [SpotFleetMinVcpuCount](/reference/rack-parameters/SpotFleetMinVcpuCount)
- [SpotFleetMinOnDemandCount](/reference/rack-parameters/SpotFleetMinOnDemandCount)
- [SpotFleetTargetType](/reference/rack-parameters/SpotFleetTargetType)
- [SpotInstanceBid](/reference/rack-parameters/SpotInstanceBid)
