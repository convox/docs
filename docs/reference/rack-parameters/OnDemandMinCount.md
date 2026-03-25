---
title: "OnDemandMinCount"
description: "Set the minimum number of on-demand EC2 instances in the Rack cluster when using spot instances."
---

# OnDemandMinCount

Minimum number of on-demand instances to maintain when using spot instances. This should be set to a value that will guarantee the minimum acceptable service availability. You must set it even if you are using [HighAvailability](/reference/rack-parameters/HighAvailability) as `false`, as this will be used to create the minimum on-demand instances.

| Default value  | `3` |
| Minimum value | `1` |

## Use Cases

- Ensuring a baseline of 3 on-demand instances for production workloads while using spot instances for cost savings on additional capacity
- Reducing to `1` in non-HA development Racks that use spot instances to minimize costs
- Increasing above `3` for mission-critical workloads where spot instance interruptions must not impact availability

## Additional Information

This parameter is primarily relevant when using spot instances. On-demand instances provide guaranteed availability, while spot instances can be interrupted when AWS needs the capacity back. By setting an appropriate `OnDemandMinCount`, you ensure that your Rack always has a stable base of instances that cannot be interrupted.

If you are not using spot instances (i.e., [SpotInstanceBid](/reference/rack-parameters/SpotInstanceBid) is not set), this parameter still defines the minimum on-demand count in the Auto Scaling group.

```bash
$ convox rack params set OnDemandMinCount=2
```

## See Also

- [SpotInstanceBid](/reference/rack-parameters/SpotInstanceBid)
- [InstanceCount](/reference/rack-parameters/InstanceCount)
- [HighAvailability](/reference/rack-parameters/HighAvailability)
- [NoHaInstanceCount](/reference/rack-parameters/NoHaInstanceCount)
