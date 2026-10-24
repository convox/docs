---
title: "Autoscale"
description: "Enable or disable automatic scaling of Convox Rack EC2 instances."
---

# Autoscale

Automatic scaling of Rack cluster instances based on the CPU and memory reserved by the containers your Services are scaled to run. See the [Scaling documentation](/scaling/scaling#autoscaling) for more information.

When enabled, the Rack automatically adjusts the number of EC2 instances based on the resource requirements of the containers your Services are scaled to run. When disabled, you must manually manage the instance count.

| Setting | Value |
|:--|:--|
| Default value  | `Yes`       |
| Allowed values | `Yes`, `No` |

## Use Cases

- Keeping enabled (default) so the Rack automatically adds instances when deploying new Services or scaling up processes
- Disabling for Racks where you want full manual control over instance count, such as cost-sensitive environments with predictable workloads
- Disabling temporarily during maintenance operations to prevent unwanted scaling events

## Additional Information

The autoscaler runs every minute as a Lambda function in the Rack's stack. It sets [InstanceCount](/reference/rack-parameters/InstanceCount), or [NoHaInstanceCount](/reference/rack-parameters/NoHaInstanceCount) on a Rack without high availability, to the number of instances it calculates, so it replaces a count you set by hand on its next run. It does not run while the Rack's stack is updating, or after a Rack update that rolled back, until the next Rack update succeeds. Its output goes to its own CloudWatch log group, not to `convox rack logs`, and the group's name starts with `/aws/lambda/` and contains `InstancesAutoscaler`.

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
