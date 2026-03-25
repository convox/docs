---
title: "InstanceType"
description: "Set the EC2 instance type for runtime instances in a Convox Rack cluster."
---

# InstanceType

The type of EC2 instance to run in your Rack cluster.

| Default value  | `t3.small`                                                       |
| Allowed values | [EC2 Instance Types](https://aws.amazon.com/ec2/instance-types/) |

## Use Cases

- Upgrading to a larger instance type when your applications require more memory or CPU
- Switching to a compute-optimized instance family (e.g., `c5`) for CPU-intensive workloads
- Using ARM-based instances (e.g., `t4g.small`) for cost savings on compatible workloads
- Downsizing for development or staging Racks where resource demands are lower

## Additional Information

The instance type determines the CPU, memory, and networking capacity available to each node in your cluster. When choosing an instance type, consider the total resource requirements of all containers that will run on each instance.

If your builds require a different instance type than your runtime workloads, use [BuildInstance](/reference/rack-parameters/BuildInstance) to configure a dedicated build instance separately.

See also the [BuildInstance](/reference/rack-parameters/BuildInstance) Rack parameter.

```bash
$ convox rack params set InstanceType=t3.medium
```

## See Also

- [BuildInstance](/reference/rack-parameters/BuildInstance)
- [InstanceCount](/reference/rack-parameters/InstanceCount)
- [CpuCredits](/reference/rack-parameters/CpuCredits)
- [Ami](/reference/rack-parameters/Ami)
