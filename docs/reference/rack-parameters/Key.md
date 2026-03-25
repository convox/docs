---
title: "Key"
description: "SSH key name for accessing Convox Rack EC2 cluster instances."
---

# Key

SSH key name for access to cluster instances. Set this to the name of an existing EC2 Key Pair in your AWS account to enable SSH access to the Rack's EC2 instances.

When left blank (the default), SSH access to cluster instances is disabled. This is the recommended setting for production environments where instance access should be managed through other mechanisms such as ECS Exec or SSM Session Manager.

| Default value  | "" |

## Use Cases

- Debugging container or instance-level issues that require direct SSH access to the host
- Inspecting ECS agent logs or Docker state on a specific instance
- Performing emergency maintenance on cluster instances when other access methods are unavailable

## Additional Information

The value must match the name of an EC2 Key Pair that already exists in the same AWS region as your Rack. You can create a key pair in the AWS Console under **EC2 > Key Pairs** or via the AWS CLI:

```bash
$ aws ec2 create-key-pair --key-name my-convox-key --query 'KeyMaterial' --output text > my-convox-key.pem
$ convox rack params set Key=my-convox-key
```

Setting this parameter will trigger a rolling update of your instances. The key will apply to both runtime and build instances.

For security best practices, consider leaving this blank and using [AWS Systems Manager Session Manager](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager.html) for instance access instead.

## See Also

- [InstanceType](/reference/rack-parameters/InstanceType)
- [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup)
- [InstanceBootCommand](/reference/rack-parameters/InstanceBootCommand)
- [InstanceRunCommand](/reference/rack-parameters/InstanceRunCommand)
