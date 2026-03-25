---
title: "BuildInstanceSecurityGroup"
description: "Assign a custom security group to the Convox Rack build instances."
---

# BuildInstanceSecurityGroup

The security group to assign to build instances. If blank, the Rack uses the same security group as the runtime instances (or the value of [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup) if set).

| Default value  | "" |

## Use Cases

- Isolating build instance network access from runtime instances for security compliance
- Granting build instances access to specific internal resources (e.g., private artifact repositories, license servers) that runtime instances should not reach
- Restricting outbound internet access on build instances in regulated environments

## Additional Information

The security group must exist in the same VPC as the Rack. When specifying a custom security group, ensure it allows the necessary outbound access for pulling Docker images, accessing ECR, and communicating with the ECS control plane.

```bash
$ convox rack params set BuildInstanceSecurityGroup=sg-0abcdef1234567890
```

## See Also

- [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup)
- [BuildInstance](/reference/rack-parameters/BuildInstance)
- [BuildInstancePolicy](/reference/rack-parameters/BuildInstancePolicy)
- [PrivateBuild](/reference/rack-parameters/PrivateBuild)
