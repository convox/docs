---
title: "AutoMinorVersionUpgrade"
description: "Controls whether automatic minor version upgrades are enabled for database resources in your Convox App."
---

# AutoMinorVersionUpgrade

Automatic minor version upgrade control for database resources. Set to `false` to disable. When enabled, AWS will automatically apply minor engine version upgrades to RDS instances, ElastiCache clusters, and other supported database resources during their maintenance windows.

| Default value  | `true` |
| Allowed values | `true`, `false` |

## Use Cases

- Set to `false` when you need to pin database engine versions for compliance or compatibility reasons
- Set to `false` in production environments where you want to control exactly when database upgrades occur
- Leave as `true` (default) to automatically receive security patches and bug fixes for your database engines

## Additional Information

This parameter applies to all embedded database resources defined in your `convox.yml` (such as `postgres`, `mysql`, `redis`, etc.). It maps directly to the AWS `AutoMinorVersionUpgrade` property on the underlying CloudFormation resources.

When disabled, you are responsible for manually upgrading database engine versions. Leaving minor version upgrades disabled for extended periods may result in running database versions that no longer receive security patches from AWS.

```bash
$ convox apps params set AutoMinorVersionUpgrade=false
```

## See Also

- [ResourcePassword](/reference/app-parameters/ResourcePassword)
- [Rack Parameters](/reference/rack-parameters)
