---
title: "PrivateBuild"
description: "Place only the Rack's build instances into a private subnet."
---

# PrivateBuild

Private subnet placement for build instances only. This parameter is unused if [Private](/reference/rack-parameters/Private) is set to `Yes`, since all instances are already in private subnets in that case.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Use Cases

- Keeping build instances private while allowing runtime instances to remain in public subnets
- Preventing build instances from having public IP addresses when builds pull from private registries or internal artifact stores
- Reducing the attack surface for build infrastructure without requiring a fully private Rack

## Additional Information

When `PrivateBuild` is set to `Yes`, build instances are placed in private subnets and require a NAT Gateway for outbound internet access (e.g., to pull base images from Docker Hub or other public registries).

When this parameter is enabled, registry authentication is handled entirely on the build cluster rather than through the main Rack. This means private registry credentials are not passed through the Rack API path, and all registry access stays within the build environment.

If [Private](/reference/rack-parameters/Private) is already set to `Yes`, this parameter has no additional effect because all instances, including build instances, are already in private subnets.

```bash
$ convox rack params set PrivateBuild=Yes
```

## See Also

- [Private](/reference/rack-parameters/Private)
- [PrivateApi](/reference/rack-parameters/PrivateApi)
- [BuildInstance](/reference/rack-parameters/BuildInstance)
- [BuildInstanceSecurityGroup](/reference/rack-parameters/BuildInstanceSecurityGroup)
