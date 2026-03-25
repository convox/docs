---
title: "Tenancy"
description: "Configure whether Convox Rack EC2 instances run on shared or dedicated hardware."
---

# Tenancy

Dedicated hardware. This controls the tenancy of EC2 instances and the VPC created by the Rack.

| Default value  | `default`                  |
| Allowed values | `default`, `dedicated` |

## Use Cases

- Using `dedicated` tenancy for regulatory or compliance requirements that mandate physical isolation from other AWS customers
- Keeping `default` (shared) tenancy for most workloads to benefit from lower costs
- Enabling dedicated tenancy when running security-sensitive workloads that require hardware-level isolation

## Additional Information

When set to `dedicated`, all EC2 instances in the Rack run on single-tenant hardware that is physically isolated from instances belonging to other AWS accounts. This also sets the VPC instance tenancy, meaning all instances launched in the VPC will use dedicated hardware.

> **Warning:** Dedicated tenancy has significantly higher costs than shared tenancy. Not all instance types support dedicated tenancy. Verify your chosen [InstanceType](/reference/rack-parameters/InstanceType) supports dedicated tenancy before enabling this parameter.

```bash
$ convox rack params set Tenancy=dedicated
```

## See Also

- [InstanceType](/reference/rack-parameters/InstanceType)
- [Tags](/reference/rack-parameters/Tags)
- [Private](/reference/rack-parameters/Private)
