---
title: "PrivateApiSecurityGroup"
description: "Restrict access to the private Rack API by specifying an allowed security group."
---

# PrivateApiSecurityGroup

Custom security group for restricting access to the private Rack API endpoint. This allows you to lock down the private API to only accept connections from instances or resources that belong to the specified security group when [PrivateApi](/reference/rack-parameters/PrivateApi) is set to `Yes`.

| Default value  | "" |

## Use Cases

- Restricting Rack API access to only a bastion host or jump box security group
- Limiting API access to a specific set of CI/CD runners that manage deployments
- Adding an extra layer of network-level access control on top of the API password authentication

## Additional Information

This parameter only takes effect when [PrivateApi](/reference/rack-parameters/PrivateApi) is set to `Yes`. When `PrivateApi` is `No`, the API is publicly accessible and this security group restriction does not apply.

The value should be a valid security group ID (e.g., `sg-0123456789abcdef0`) in the same VPC as the Rack. Resources in that security group will be allowed to connect to the Rack API's internal load balancer.

```bash
$ convox rack params set PrivateApi=Yes PrivateApiSecurityGroup=sg-0123456789abcdef0
```

## See Also

- [PrivateApi](/reference/rack-parameters/PrivateApi)
- [Private](/reference/rack-parameters/Private)
- [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup)
- [Password](/reference/rack-parameters/Password)
