---
title: "IgnoreManualScaleCount"
description: "When autoscaling is enabled, ignore manually set desired counts so the autoscaler controls scaling exclusively."
---

# IgnoreManualScaleCount

Manual scaling suppression when autoscaling is enabled. When set to `Yes`, manually set desired counts are ignored in favor of autoscaler-managed scaling.

| Default value  | `No` |
| Allowed values | `Yes`, `No` |

If set to `Yes`, any manually set service counts using `convox scale` will be ignored, and the autoscaler will dynamically manage the service count based on the targets defined in `convox.yml`.
Additionally, during deployments, the **active** instance count at the time of deployment will be used as the new desired count, ensuring the autoscaler fully controls scaling behavior without manual overrides.

## Use Cases

- Enable when you use autoscaling and want to prevent manual `convox scale` commands from overriding the autoscaler
- Enable to ensure deployments preserve the current running count rather than resetting to a hard-coded desired count
- Keep disabled when you occasionally need to manually override scale counts for debugging or load testing

## Additional Information

Without this parameter, a `convox scale` command or a deployment that specifies a desired count can override the autoscaler's current target. This can cause unexpected scaling behavior, such as scaling down from a high count that the autoscaler had set in response to load.

With `IgnoreManualScaleCount=Yes`, the system uses the actual running task count at deployment time as the baseline, allowing the autoscaler to remain in control throughout the deployment lifecycle.

For more details on how this parameter interacts with scaling, see [Scaling](/scaling/scaling).

```bash
$ convox apps params set IgnoreManualScaleCount=Yes
```

## See Also

- [Scaling](/scaling/scaling)
- [Services](/application/services)
- [Rack Parameters](/reference/rack-parameters)
