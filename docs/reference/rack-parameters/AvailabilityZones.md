---
title: "AvailabilityZones"
description: "Override the default AWS Availability Zones used by a Convox Rack."
---

# AvailabilityZones

Custom availability zone selection for the Rack. When blank, zones are selected automatically. Updating this parameter once a Rack is installed will require setting [MaxAvailabilityZones](/reference/rack-parameters/MaxAvailabilityZones) to the new AZs quantity you are choosing.

When left blank (the default), the Rack automatically selects availability zones in the region where it is installed.

| Default value  | "" |

## Use Cases

- Constraining a Rack to specific AZs that have capacity for the instance types you need
- Aligning Rack AZs with other infrastructure (e.g., databases, caches) to minimize cross-AZ data transfer costs
- Working around AZ-specific service limitations or outages in a region

## Additional Information

You must specify exactly 3 availability zones as a comma-delimited list (or 2 if you also set [MaxAvailabilityZones](/reference/rack-parameters/MaxAvailabilityZones) to `2`).

```bash
$ convox rack params set AvailabilityZones=us-east-1a,us-east-1b,us-east-1c
```

> **Warning:** Changing availability zones on an existing Rack can trigger the recreation of subnets and associated resources. This is a disruptive operation and should be planned carefully.

## See Also

- [MaxAvailabilityZones](/reference/rack-parameters/MaxAvailabilityZones)
- [HighAvailability](/reference/rack-parameters/HighAvailability)
- [Subnet0CIDR](/reference/rack-parameters/Subnet0CIDR)
- [SubnetPrivate0CIDR](/reference/rack-parameters/SubnetPrivate0CIDR)
- [ExistingVpc](/reference/rack-parameters/ExistingVpc)
