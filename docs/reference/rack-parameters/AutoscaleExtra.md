---
title: "AutoscaleExtra"
description: "Set the number of extra capacity instances that the Rack autoscaler keeps running."
---

# AutoscaleExtra

The number of instances of extra capacity that autoscale should keep running. This provides headroom so that new deployments or scaling events do not have to wait for new instances to launch.

| Default value  | `1` |

## Use Cases

- Increasing to 2 or more on production Racks to ensure there is always spare capacity for rapid deployments
- Setting to 0 on cost-sensitive Racks where you are willing to accept longer scaling times in exchange for lower costs
- Increasing during expected traffic spikes to pre-provision additional headroom

## Additional Information

This parameter only takes effect when [Autoscale](/reference/rack-parameters/Autoscale) is set to `Yes`. For non-HA Racks (where [HighAvailability](/reference/rack-parameters/HighAvailability) is `false`), use [NoHAAutoscaleExtra](/reference/rack-parameters/NoHAAutoscaleExtra) instead.

```bash
$ convox rack params set AutoscaleExtra=2
```

## See Also

- [Autoscale](/reference/rack-parameters/Autoscale)
- [NoHAAutoscaleExtra](/reference/rack-parameters/NoHAAutoscaleExtra)
- [InstanceCount](/reference/rack-parameters/InstanceCount)
- [InstanceType](/reference/rack-parameters/InstanceType)
- [Scaling](/scaling/scaling)
