---
title: "ResourcePassword"
description: "Override the auto-generated password for embedded database resources defined in convox.yml."
---

# ResourcePassword

Password override for embedded database resources defined in `convox.yml`. When blank, a password is generated automatically from the CloudFormation stack ID.

| Default value  | "" |

## Use Cases

- Set a specific password when migrating an existing database into Convox and you need to preserve the original credentials
- Set when external systems need to connect to the database with a known password
- Leave blank (default) to let Convox auto-generate a password derived from the CloudFormation stack ID

## Additional Information

When this parameter is blank, Convox generates a password from the CloudFormation stack ID. This password is stable across deployments (it does not change unless you explicitly set a new one).

This parameter applies to all embedded resources in the application. You cannot set different passwords for different resources using this parameter. The password value is marked as `NoEcho` in CloudFormation, meaning it will not appear in stack outputs or console displays.

```bash
$ convox apps params set ResourcePassword=my-secure-password
```

> Changing this parameter on an existing application will update the password for the resource, which may cause connection failures if running services still have the old password cached. Plan a deployment immediately after changing this value.

## See Also

- [AutoMinorVersionUpgrade](/reference/app-parameters/AutoMinorVersionUpgrade)
- [Rack Parameters](/reference/rack-parameters)
