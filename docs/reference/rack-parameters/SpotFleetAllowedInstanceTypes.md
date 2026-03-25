---
title: "SpotFleetAllowedInstanceTypes"
description: "Restrict which EC2 instance types are allowed in the Convox Rack Spot Fleet."
---

# SpotFleetAllowedInstanceTypes

Comma-separated list of allowed instance types in the Spot Fleet. This parameter cannot be used together with [SpotFleetExcludedInstanceTypes](/reference/rack-parameters/SpotFleetExcludedInstanceTypes) -- if both are set, `SpotFleetAllowedInstanceTypes` takes precedence. This parameter can be used only when Spot Fleet is enabled by setting [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice).

| Default value  | "" |

Wildcard patterns are supported. Examples: `m5.8xlarge`, `c5*.*`, `m5a.*`, `r*`, `*3*`.

## Use Cases

- Restricting the Spot Fleet to specific instance families (e.g., `m5.*`, `c5.*`) that are known to work well with your workloads
- Limiting to instances with specific hardware capabilities such as enhanced networking or NVMe storage
- Targeting only current-generation instance types to ensure consistent performance

## Additional Information

When this parameter is blank (the default), the Spot Fleet considers all instance types that meet the minimum resource requirements defined by [SpotFleetMinMemoryMiB](/reference/rack-parameters/SpotFleetMinMemoryMiB) and [SpotFleetMinVcpuCount](/reference/rack-parameters/SpotFleetMinVcpuCount).

If you need to exclude specific instance types rather than defining an allow list, use [SpotFleetExcludedInstanceTypes](/reference/rack-parameters/SpotFleetExcludedInstanceTypes) instead. The two parameters are mutually exclusive.

This parameter only takes effect when Spot Fleet is enabled via [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice).

```bash
$ convox rack params set SpotFleetAllowedInstanceTypes="m5.large,m5.xlarge,c5.large"
```

## See Also

- [SpotFleetExcludedInstanceTypes](/reference/rack-parameters/SpotFleetExcludedInstanceTypes)
- [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice)
- [SpotFleetAllocationStrategy](/reference/rack-parameters/SpotFleetAllocationStrategy)
- [SpotFleetMinMemoryMiB](/reference/rack-parameters/SpotFleetMinMemoryMiB)
- [SpotFleetMinVcpuCount](/reference/rack-parameters/SpotFleetMinVcpuCount)
- [SpotFleetMinOnDemandCount](/reference/rack-parameters/SpotFleetMinOnDemandCount)
- [SpotFleetTargetType](/reference/rack-parameters/SpotFleetTargetType)
- [SpotInstanceBid](/reference/rack-parameters/SpotInstanceBid)
