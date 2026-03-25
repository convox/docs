---
title: "SpotFleetMaxPrice"
description: "Set the maximum hourly price for EC2 Spot Fleet instances and enable Spot Fleet mode in a Convox Rack."
---

# SpotFleetMaxPrice

The maximum price for instances in the Spot Fleet per hour. Setting this parameter enables Spot Fleet mode, which uses AWS EC2 Spot Fleet requests to fulfill instance demand instead of an Auto Scaling group. The fleet will try to launch instances until it crosses the price threshold, even if the target [InstanceCount](/reference/rack-parameters/InstanceCount) or [NoHaInstanceCount](/reference/rack-parameters/NoHaInstanceCount) is not fulfilled.

| Default value  | "" |

## Use Cases

- Enabling Spot Fleet to significantly reduce EC2 costs for non-critical or fault-tolerant workloads
- Setting a budget ceiling for Spot instance spending while allowing AWS to select the most cost-effective instance types
- Running development or staging environments at a fraction of on-demand pricing

## Additional Information

Setting this parameter enables Spot Fleet, which fundamentally changes how the Rack manages instances. Instead of an Auto Scaling group, the Rack uses an AWS Spot Fleet request to fulfill instance demand. The fleet is managed by the Spot request rather than the Auto Scaling group.

> **Note:** Currently Spot Fleet has only single availability zone support, even if you set [HighAvailability](/reference/rack-parameters/HighAvailability) to `true`.

[SpotFleetMinOnDemandCount](/reference/rack-parameters/SpotFleetMinOnDemandCount) is used to launch on-demand instances alongside Spot instances, ensuring a baseline of guaranteed capacity. The instance type for on-demand instances is taken from the [InstanceType](/reference/rack-parameters/InstanceType) parameter.

When Spot Fleet is enabled, you can fine-tune instance selection using [SpotFleetAllowedInstanceTypes](/reference/rack-parameters/SpotFleetAllowedInstanceTypes), [SpotFleetExcludedInstanceTypes](/reference/rack-parameters/SpotFleetExcludedInstanceTypes), [SpotFleetMinMemoryMiB](/reference/rack-parameters/SpotFleetMinMemoryMiB), and [SpotFleetMinVcpuCount](/reference/rack-parameters/SpotFleetMinVcpuCount).

```bash
$ convox rack params set SpotFleetMaxPrice=0.05
```

## See Also

- [SpotFleetAllocationStrategy](/reference/rack-parameters/SpotFleetAllocationStrategy)
- [SpotFleetAllowedInstanceTypes](/reference/rack-parameters/SpotFleetAllowedInstanceTypes)
- [SpotFleetExcludedInstanceTypes](/reference/rack-parameters/SpotFleetExcludedInstanceTypes)
- [SpotFleetMinMemoryMiB](/reference/rack-parameters/SpotFleetMinMemoryMiB)
- [SpotFleetMinVcpuCount](/reference/rack-parameters/SpotFleetMinVcpuCount)
- [SpotFleetMinOnDemandCount](/reference/rack-parameters/SpotFleetMinOnDemandCount)
- [SpotFleetTargetType](/reference/rack-parameters/SpotFleetTargetType)
- [SpotInstanceBid](/reference/rack-parameters/SpotInstanceBid)
- [InstanceType](/reference/rack-parameters/InstanceType)
