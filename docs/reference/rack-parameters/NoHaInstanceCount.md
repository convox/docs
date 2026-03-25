---
title: "NoHaInstanceCount"
description: "Set the number of EC2 instances for non-high-availability Convox Rack clusters."
---

# NoHaInstanceCount

The number of EC2 instances in your non-High Availability Rack cluster. This parameter is only used when [HighAvailability](/reference/rack-parameters/HighAvailability) is set to `false`. For HA Racks, see [InstanceCount](/reference/rack-parameters/InstanceCount).

| Default value  | `1` |
| Minimum value | `1` |

## Use Cases

- Running a single-instance development or staging Rack to minimize costs
- Increasing to 2 or more instances in non-HA Racks that still need some level of workload distribution
- Scaling up a non-HA Rack to handle higher application loads without switching to full HA mode

## Additional Information

When [HighAvailability](/reference/rack-parameters/HighAvailability) is `true`, the [InstanceCount](/reference/rack-parameters/InstanceCount) parameter is used instead and `NoHaInstanceCount` is ignored.

Non-HA Racks run with a minimum of 1 instance. While you can increase this value, keep in mind that non-HA mode does not provide the same redundancy guarantees as HA mode. For production workloads that require fault tolerance, consider enabling [HighAvailability](/reference/rack-parameters/HighAvailability) instead.

```bash
$ convox rack params set NoHaInstanceCount=2
```

## See Also

- [InstanceCount](/reference/rack-parameters/InstanceCount)
- [HighAvailability](/reference/rack-parameters/HighAvailability)
- [NoHAAutoscaleExtra](/reference/rack-parameters/NoHAAutoscaleExtra)
- [InstanceType](/reference/rack-parameters/InstanceType)
