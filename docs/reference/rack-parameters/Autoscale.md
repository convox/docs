---
title: "Autoscale"
description: "Enable or disable automatic scaling of Convox Rack EC2 instances."
---

# Autoscale

Automatic scaling of Rack cluster instances based on resource utilization. See the [Scaling documentation](/scaling/scaling#autoscaling) for more information.

When enabled, the Rack automatically adjusts the number of EC2 instances based on the resource requirements of running containers. When disabled, you must manually manage the instance count.

| Default value  | `Yes`       |
| Allowed values | `Yes`, `No` |

## Use Cases

- Keeping enabled (default) so the Rack automatically adds instances when deploying new services or scaling up processes
- Disabling for Racks where you want full manual control over instance count, such as cost-sensitive environments with predictable workloads
- Disabling temporarily during maintenance operations to prevent unwanted scaling events

## Additional Information

When autoscale is enabled, use [AutoscaleExtra](/reference/rack-parameters/AutoscaleExtra) to control how many extra instances of headroom the autoscaler maintains. For non-HA Racks, use [NoHAAutoscaleExtra](/reference/rack-parameters/NoHAAutoscaleExtra) instead.

```bash
$ convox rack params set Autoscale=No
```

## See Also

- [AutoscaleExtra](/reference/rack-parameters/AutoscaleExtra)
- [NoHAAutoscaleExtra](/reference/rack-parameters/NoHAAutoscaleExtra)
- [InstanceCount](/reference/rack-parameters/InstanceCount)
- [InstanceType](/reference/rack-parameters/InstanceType)
- [Scaling](/scaling/scaling)
