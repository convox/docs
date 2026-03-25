---
title: "InstanceCount"
description: "Set the number of EC2 instances in a high-availability Convox Rack cluster."
---

# InstanceCount

The number of EC2 instances in your Rack cluster. This parameter is only used for clusters with [HighAvailability](/reference/rack-parameters/HighAvailability) set to `true`.

| Default value  | `3` |
| Minimum value | `3` |

## Use Cases

- Increasing instance count to handle higher application workloads or run more containers
- Setting to a higher value when deploying multiple services that require dedicated instance capacity
- Adjusting alongside [Autoscale](/reference/rack-parameters/Autoscale) to set a minimum floor for the cluster size

## Additional Information

When [HighAvailability](/reference/rack-parameters/HighAvailability) is set to `false`, this parameter is not used. Instead, the `NoHaInstanceCount` parameter controls the cluster size, which has a minimum value of `1`.

If [Autoscale](/reference/rack-parameters/Autoscale) is enabled, the cluster may scale beyond this value based on workload demand. The `InstanceCount` effectively serves as the baseline number of instances.

```bash
$ convox rack params set InstanceCount=5
```

## See Also

- [HighAvailability](/reference/rack-parameters/HighAvailability)
- [InstanceType](/reference/rack-parameters/InstanceType)
- [Autoscale](/reference/rack-parameters/Autoscale)
- [AutoscaleExtra](/reference/rack-parameters/AutoscaleExtra)
- [InstanceUpdateBatchSize](/reference/rack-parameters/InstanceUpdateBatchSize)
