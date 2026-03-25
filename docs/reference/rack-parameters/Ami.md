---
title: "Ami"
description: "Override the default Amazon Machine Image (AMI) used for Convox Rack EC2 instances."
---

# Ami

Custom Amazon Machine Image (AMI) for the Rack's EC2 instances. When blank, the latest ECS-optimized AMI is used.

When left blank (the default), the Rack automatically uses the latest ECS-optimized AMI via AWS SSM parameter lookup based on your instance architecture.

| Default value  | "" |

## Use Cases

- Pinning a specific AMI version when you need to control exactly which OS image your instances run
- Using a custom AMI that includes pre-installed monitoring agents, security tools, or other software
- Testing a new ECS-optimized AMI before rolling it out across all Racks

## Additional Information

The AMI you specify must be ECS-optimized (i.e., it must include the ECS container agent and Docker runtime). Using a non-ECS-optimized AMI will prevent instances from joining the ECS cluster.

If you need to change the default AMI path rather than hardcoding a specific image ID, use [DefaultAmi](/reference/rack-parameters/DefaultAmi) or [DefaultAmiArm](/reference/rack-parameters/DefaultAmiArm) instead.

```bash
$ convox rack params set Ami=ami-0abcdef1234567890
```

## See Also

- [DefaultAmi](/reference/rack-parameters/DefaultAmi)
- [DefaultAmiArm](/reference/rack-parameters/DefaultAmiArm)
- [InstanceType](/reference/rack-parameters/InstanceType)
- [AWS ECS-Optimized AMI Documentation](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html)
