---
title: "HighAvailability"
description: "Configure whether a Convox Rack runs in high-availability mode with redundant instances across multiple availability zones."
---

# HighAvailability

> **Warning:** This parameter cannot be changed after the Rack is created.

High Availability mode for the Rack, choosing between failure resiliency and cost efficiency. This ensures proper resource redundancy to mitigate system failures.

If HighAvailability is set to `true`, the [InstanceCount](/reference/rack-parameters/InstanceCount) parameter is used as the initial cluster size. If set to `false`, the [NoHaInstanceCount](/reference/rack-parameters/NoHaInstanceCount) is used as the initial cluster size. Both can be scaled up to 1000 instances.

| Default value  | `true`          |
| Allowed values | `true`, `false` |

## Use Cases

- Setting to `false` for development or staging environments to reduce costs by running fewer instances
- Keeping as `true` for production workloads where uptime and fault tolerance are critical
- Using `false` for short-lived test Racks that do not require multi-AZ redundancy

## Additional Information

In high-availability mode, Convox distributes instances across multiple AWS Availability Zones and runs a minimum of 3 instances, providing resilience against individual instance or AZ failures. When disabled, the Rack can operate with as few as 1 instance, significantly reducing AWS infrastructure costs at the expense of redundancy.

Because this parameter cannot be changed after Rack creation, plan your availability requirements before installing the Rack. If you need to switch modes, you must create a new Rack.

This parameter can only be set during Rack installation and cannot be changed afterward. You can view the current value with:

```bash
$ convox rack params
```

## See Also

- [InstanceCount](/reference/rack-parameters/InstanceCount)
- [NoHaInstanceCount](/reference/rack-parameters/NoHaInstanceCount)
- [Autoscale](/reference/rack-parameters/Autoscale)
- [AutoscaleExtra](/reference/rack-parameters/AutoscaleExtra)
- [Private](/reference/rack-parameters/Private)
