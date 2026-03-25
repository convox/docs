---
title: "SpotInstanceBid"
description: "Set a bid price to use EC2 Spot Instances for Convox Rack cluster instances."
---

# SpotInstanceBid

A value, in dollars, that you want to pay for spot instances. If spot instances are available for the bid price, the Rack instances will use spot instances instead of on-demand instances, resulting in significant cost savings. If the parameter is empty, spot instances will not be utilized.

This must be used with the [OnDemandMinCount](/reference/rack-parameters/OnDemandMinCount) parameter to guarantee some on-demand instances are running if spot instances are not available. You must set `OnDemandMinCount` even if [HighAvailability](/reference/rack-parameters/HighAvailability) is `false` -- if not set, the default value will be used.

| Default value  | "" |

## Use Cases

- Setting a bid price for development or staging Racks to reduce costs by up to 60-90% compared to on-demand pricing
- Using spot instances for stateless workloads that can tolerate brief interruptions
- Combining with [OnDemandMinCount](/reference/rack-parameters/OnDemandMinCount) to maintain a baseline of guaranteed capacity while using spot instances for additional scaling

## Additional Information

When a bid price is set, the Rack's Auto Scaling group uses a mixed instances policy that combines on-demand and spot instances. The [OnDemandMinCount](/reference/rack-parameters/OnDemandMinCount) parameter determines the minimum number of on-demand instances that are always running, providing a safety net when spot capacity is unavailable.

For more advanced spot instance management with instance type flexibility, consider using [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice) instead, which enables AWS Spot Fleet with support for multiple instance types and allocation strategies.

```bash
$ convox rack params set SpotInstanceBid=0.10
```

## See Also

- [OnDemandMinCount](/reference/rack-parameters/OnDemandMinCount)
- [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice)
- [SpotFleetAllocationStrategy](/reference/rack-parameters/SpotFleetAllocationStrategy)
- [InstanceType](/reference/rack-parameters/InstanceType)
- [InstanceCount](/reference/rack-parameters/InstanceCount)
