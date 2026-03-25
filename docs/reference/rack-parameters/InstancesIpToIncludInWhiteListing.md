---
title: "InstancesIpToIncludInWhiteListing"
description: "Control which instance IPs are automatically added to the whitelist on public Convox Racks."
---

# InstancesIpToIncludInWhiteListing

Automatic IP whitelist inclusion for Rack instances. Controls which instance types (build, workload, or both) have their IPs added to the whitelist.

| Default value  | `Both`                                |
| Allowed values | `Both`, `Build`, `Workload`, `None`   |

## Use Cases

- Setting to `Workload` when build instances do not need direct access to the Rack's load balancer
- Setting to `None` when you manage IP whitelisting entirely through an external firewall or WAF
- Keeping as `Both` to ensure both build and workload instances can reach the Rack endpoints without manual IP management

## Additional Information

- `Both` -- Automatically adds the public IPs of both build and workload instances to the whitelist.
- `Build` -- Only adds the build instance IPs to the whitelist.
- `Workload` -- Only adds the workload (runtime) instance IPs to the whitelist.
- `None` -- Does not automatically add any instance IPs to the whitelist.

This parameter only takes effect when the Rack is public and the [WhiteList](/reference/rack-parameters/WhiteList) parameter is configured. When whitelisting is not in use, this parameter has no effect.

```bash
$ convox rack params set InstancesIpToIncludInWhiteListing=Workload
```

## See Also

- [WhiteList](/reference/rack-parameters/WhiteList)
- [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup)
- [Private](/reference/rack-parameters/Private)
