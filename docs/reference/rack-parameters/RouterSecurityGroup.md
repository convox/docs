---
title: "RouterSecurityGroup"
description: "Assign custom security groups to the public-facing Convox Rack router."
---

# RouterSecurityGroup

Comma-delimited list of security group IDs to assign to the Rack's public router. If blank, the Rack creates its own router security group that allows inbound TCP traffic on ports 80 and 443 from all IP addresses (`0.0.0.0/0`).

| Default value  | "" |

## Use Cases

- Restricting public access to the Rack router to specific IP ranges or corporate networks
- Applying organization-mandated security groups that enforce logging or compliance rules
- Using a shared security group across multiple Racks for uniform ingress policies

## Additional Information

When this parameter is blank, the Rack creates a default security group permitting inbound TCP on ports 80 and 443 from `0.0.0.0/0` (all internet traffic). If you supply custom security group IDs, you are fully responsible for configuring the correct ingress rules to allow traffic to reach your applications.

This parameter controls the security group for the public-facing router only. To customize the security group for the internal router, use [RouterInternalSecurityGroup](/reference/rack-parameters/RouterInternalSecurityGroup).

```bash
$ convox rack params set RouterSecurityGroup=sg-0abc123def456
```

## See Also

- [RouterInternalSecurityGroup](/reference/rack-parameters/RouterInternalSecurityGroup)
- [RouterMitigationMode](/reference/rack-parameters/RouterMitigationMode)
- [SslPolicy](/reference/rack-parameters/SslPolicy)
- [WhiteList](/reference/rack-parameters/WhiteList)
