---
title: "BuildInstance"
description: "Set the EC2 instance type used as the Convox Rack dedicated build instance."
---

# BuildInstance

EC2 instance type to create and use as the Rack's [dedicated build instance](/deployment/builds).

The build instance will also use the [InstanceBootCommand](/reference/rack-parameters/InstanceBootCommand) and [InstanceRunCommand](/reference/rack-parameters/InstanceRunCommand) Rack params, if defined.

| Default value  | `t3.small`                                                       |
| Allowed values | [EC2 Instance Types](https://aws.amazon.com/ec2/instance-types/) |

## Use Cases

- Upgrading to a larger instance type (e.g., `c5.xlarge`) for faster builds on projects with large Docker images or complex compilation steps
- Using a memory-optimized instance type for builds that require large amounts of RAM during compilation
- Choosing an ARM-based instance type (e.g., `t4g.small`) if your application targets ARM architecture

## Additional Information

The build instance runs separately from the runtime cluster. It is dedicated to running builds so that build activity does not compete with your application workloads for resources.

If you are getting build errors like "not enough memory available to start process", consider upgrading this to an instance type with more memory, or adjust [BuildMemory](/reference/rack-parameters/BuildMemory) to fit within the current instance's capacity.

```bash
$ convox rack params set BuildInstance=c5.xlarge
```

## See Also

- [InstanceType](/reference/rack-parameters/InstanceType)
- [BuildMemory](/reference/rack-parameters/BuildMemory)
- [BuildCpu](/reference/rack-parameters/BuildCpu)
- [BuildVolumeSize](/reference/rack-parameters/BuildVolumeSize)
- [BuildMethod](/reference/rack-parameters/BuildMethod)
- [BuildInstancePolicy](/reference/rack-parameters/BuildInstancePolicy)
- [BuildInstanceSecurityGroup](/reference/rack-parameters/BuildInstanceSecurityGroup)
- [Deployment: Builds](/deployment/builds)
