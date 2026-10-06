---
title: "InstanceSecurityGroup"
description: "Assign a custom security group to Convox Rack ECS instances for fine-grained network access control."
---

# InstanceSecurityGroup

The security group to assign to the ECS instances. If blank, the Rack creates a security group open to all IPs in your VPC.

| Setting | Value |
|:--|:--|
| Default value  | "" |

## Use Cases

- Restricting inbound traffic to instances from specific CIDR ranges or security groups
- Applying a pre-existing security group that meets your organization's network security policies
- Locking down instance access in environments where the default open-VPC security group is too permissive

## Additional Information

The value should be a valid AWS Security Group ID, for example:

```bash
$ convox rack params set InstanceSecurityGroup=sg-0abc1234def56789a
```

When this parameter is blank, Convox creates a default security group that allows all TCP and UDP traffic from the VPC CIDR, plus all traffic from the NLB security groups while an NLB is enabled. If you provide a custom security group, ensure it permits the necessary traffic for the ECS agent, load balancers, and inter-container communication.

If you also need to customize the security group for build instances specifically, see [BuildInstanceSecurityGroup](/reference/rack-parameters/BuildInstanceSecurityGroup).

On rack version 20261005214736 or newer, if the Rack uses an NLB with client IP preservation ([NLBPreserveClientIP](/reference/rack-parameters/NLBPreserveClientIP), [NLBInternalPreserveClientIP](/reference/rack-parameters/NLBInternalPreserveClientIP), or `preserve_client_ip: true` on a deployed `nlb:` port), the new group needs an ingress rule allowing all traffic from the NLB security group (`${Rack}:NLBSecurityGroup` or `${Rack}:NLBInternalSecurityGroup`), and the Rack refuses the change until the rule exists. Remove that rule before disabling the NLB or uninstalling the Rack. See [Custom InstanceSecurityGroup](/reference/rack-parameters/NLBPreserveClientIP#custom-instancesecuritygroup).

## See Also

- [BuildInstanceSecurityGroup](/reference/rack-parameters/BuildInstanceSecurityGroup)
- [InstancePolicy](/reference/rack-parameters/InstancePolicy)
- [Private](/reference/rack-parameters/Private)
- [WhiteList](/reference/rack-parameters/WhiteList)
