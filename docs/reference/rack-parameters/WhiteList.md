---
title: "WhiteList"
description: "Restrict access to the Convox Rack API by specifying allowed CIDR ranges."
---

# WhiteList

Comma-delimited list of CIDRs to allow access to the Rack API. For example: `10.0.0.0/24,172.10.0.1/32`. A maximum of 4 CIDRs can be specified.

| Default value  | "" |

## Use Cases

- Restricting Rack API access to your office or VPN IP ranges for security compliance
- Limiting API access to specific CI/CD system IPs that need to deploy to the Rack
- Locking down a production Rack to prevent unauthorized access from unknown networks

## Additional Information

**Attention!** Please be careful to consider all required connections to the Rack API before enabling WhiteList. You can block your own access and ability to edit this parameter from the CLI if misconfigured.

Before setting the WhiteList, ensure you include:

- Your current IP address or network range
- Any CI/CD systems that deploy to this Rack
- Any monitoring or management tools that connect to the Rack API

```bash
$ convox rack params set WhiteList="203.0.113.0/24,198.51.100.10/32"
```

If you accidentally lock yourself out, you will need to modify the parameter directly through the AWS CloudFormation console.

The [InstancesIpToIncludInWhiteListing](/reference/rack-parameters/InstancesIpToIncludInWhiteListing) parameter controls whether build and workload instance IPs are automatically included in the whitelist when the Rack is public.

## See Also

- [InstancesIpToIncludInWhiteListing](/reference/rack-parameters/InstancesIpToIncludInWhiteListing)
- [PrivateApi](/reference/rack-parameters/PrivateApi)
- [PrivateApiSecurityGroup](/reference/rack-parameters/PrivateApiSecurityGroup)
- [RouterSecurityGroup](/reference/rack-parameters/RouterSecurityGroup)
