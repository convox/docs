---
title: "MaxAvailabilityZones"
description: "Set the maximum number of AWS Availability Zones used by the Convox Rack."
---

# MaxAvailabilityZones

The maximum number of Availability Zones that the Rack cluster should use. This controls how many AZs the Rack's subnets and resources are distributed across.

| Default value  | `3`      |
| Allowed values | `2`, `3` |

## Use Cases

- Setting to `2` in AWS regions that only have two Availability Zones
- Reducing to `2` to lower cross-AZ data transfer costs in cost-sensitive environments
- Updating this value when changing [AvailabilityZones](/reference/rack-parameters/AvailabilityZones) to match the new number of AZs specified

## Additional Information

When updating the [AvailabilityZones](/reference/rack-parameters/AvailabilityZones) parameter to a different set of AZs, you must also set `MaxAvailabilityZones` to match the quantity of AZs you are selecting.

Using 3 Availability Zones (the default) provides better fault tolerance since your workloads are distributed across more isolated data centers. Only reduce to 2 when necessary due to region limitations or specific cost constraints.

```bash
$ convox rack params set MaxAvailabilityZones=2
```

## See Also

- [AvailabilityZones](/reference/rack-parameters/AvailabilityZones)
- [HighAvailability](/reference/rack-parameters/HighAvailability)
- [InstanceCount](/reference/rack-parameters/InstanceCount)
