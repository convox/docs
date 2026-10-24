---
title: "InstanceCount"
description: "Set the number of EC2 instances in a high-availability Convox Rack cluster."
---

# InstanceCount

The number of EC2 instances in your Rack cluster. This parameter is only used for clusters with [HighAvailability](/reference/rack-parameters/HighAvailability) set to `true`.

| Setting | Value |
|:--|:--|
| Default value  | `3` |
| Minimum value | `3` |

## Use Cases

- Increasing instance count to handle higher application workloads or run more containers
- Setting to a higher value when deploying multiple Services that require dedicated instance capacity
- Setting the cluster size by hand on a Rack with [Autoscale](/reference/rack-parameters/Autoscale) set to `No`

## Additional Information

When [HighAvailability](/reference/rack-parameters/HighAvailability) is set to `false`, this parameter is not used. Instead, the `NoHaInstanceCount` parameter controls the cluster size, which has a minimum value of `1`.

If [Autoscale](/reference/rack-parameters/Autoscale) is enabled, the autoscaler sets this parameter itself based on workload demand, with a minimum of `3`, and replaces a value you set on its next run. Set `Autoscale=No` to control the instance count by hand.

```bash
$ convox rack params set InstanceCount=5
```

## See Also

- [HighAvailability](/reference/rack-parameters/HighAvailability)
- [InstanceType](/reference/rack-parameters/InstanceType)
- [Autoscale](/reference/rack-parameters/Autoscale)
- [AutoscaleExtra](/reference/rack-parameters/AutoscaleExtra)
- [InstanceUpdateBatchSize](/reference/rack-parameters/InstanceUpdateBatchSize)
