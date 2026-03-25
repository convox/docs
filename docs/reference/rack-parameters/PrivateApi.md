---
title: "PrivateApi"
description: "Place the Rack API Load Balancer in a private network, making it unreachable from the internet."
---

# PrivateApi

Private network placement for the Rack API load balancer. When set to `Yes`, the Rack API endpoint is only accessible from within the VPC or through a VPN/peering connection.

When set to `Yes`, the Rack API endpoint is only accessible from within the VPC or through a VPN/peering connection. This prevents any direct access to the Rack API from the public internet.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Use Cases

- Securing the Rack API endpoint so it is only accessible from your corporate network via VPN or VPC peering
- Meeting compliance requirements that prohibit management APIs from being publicly accessible
- Restricting API access to a specific security group using [PrivateApiSecurityGroup](/reference/rack-parameters/PrivateApiSecurityGroup)

## Additional Information

When `PrivateApi` is set to `Yes`, you will need network connectivity to the Rack's VPC to manage the Rack (e.g., via VPN, AWS Direct Connect, or VPC peering). The Convox CLI and any CI/CD pipelines must be able to reach the internal API endpoint.

You can further restrict which resources can reach the private API by specifying a [PrivateApiSecurityGroup](/reference/rack-parameters/PrivateApiSecurityGroup).

This parameter is independent of [Private](/reference/rack-parameters/Private). You can have a private API with public instances, or vice versa, though combining both is common for fully locked-down environments.

```bash
$ convox rack params set PrivateApi=Yes
```

## See Also

- [Private](/reference/rack-parameters/Private)
- [PrivateApiSecurityGroup](/reference/rack-parameters/PrivateApiSecurityGroup)
- [PrivateBuild](/reference/rack-parameters/PrivateBuild)
- [Password](/reference/rack-parameters/Password)
