---
title: "NoHAAutoscaleExtra"
description: "Set the number of extra autoscale capacity instances for non-high-availability Racks."
---

# NoHAAutoscaleExtra

Extra instance capacity for autoscaling in non-HA Rack clusters. This only applies when [HighAvailability](/reference/rack-parameters/HighAvailability) is set to `false`. This functions similarly to the [AutoscaleExtra](/reference/rack-parameters/AutoscaleExtra) parameter but is used exclusively in non-HA configurations.

| Default value  | `0` |

## Use Cases

- Adding a buffer of one or more spare instances in non-HA Racks to absorb traffic spikes without waiting for new instances to launch
- Setting to `1` to ensure at least one instance is always available beyond the minimum needed for running workloads
- Keeping at `0` in development or staging environments where cost is a priority over burst headroom

## Additional Information

When [HighAvailability](/reference/rack-parameters/HighAvailability) is `true`, the [AutoscaleExtra](/reference/rack-parameters/AutoscaleExtra) parameter is used instead. `NoHAAutoscaleExtra` is ignored in HA configurations.

The autoscaler uses this value to determine how many additional instances to keep running beyond what is required to schedule current workloads. A value of `0` means autoscale will only provision exactly the capacity needed.

```bash
$ convox rack params set NoHAAutoscaleExtra=1
```

## See Also

- [AutoscaleExtra](/reference/rack-parameters/AutoscaleExtra)
- [Autoscale](/reference/rack-parameters/Autoscale)
- [HighAvailability](/reference/rack-parameters/HighAvailability)
- [NoHaInstanceCount](/reference/rack-parameters/NoHaInstanceCount)
