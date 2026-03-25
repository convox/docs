---
title: "DynamoDbTableDeletionProtectionEnabled"
description: "Enables deletion protection on DynamoDB tables created by a Convox Rack to prevent accidental data loss."
---

# DynamoDbTableDeletionProtectionEnabled

Deletion protection for the Rack's DynamoDB table. When enabled, the table cannot be deleted by any user or process. This setting is disabled by default.

| Default value  | `false`         |
| Allowed values | `true`, `false` |

## Use Cases

- Enable on production racks to prevent accidental deletion of DynamoDB tables that store critical Rack state.
- Required by compliance frameworks that mandate deletion protection on persistent data stores.
- Protect against infrastructure-as-code mistakes or manual CloudFormation stack operations that could remove tables.

## Additional Information

When set to `true`, any attempt to delete the DynamoDB table (whether through the AWS Console, CLI, or CloudFormation) will fail until deletion protection is explicitly disabled. This adds a safety layer against accidental data loss.

Note that enabling this parameter does not affect read/write access to the table -- it only prevents the table itself from being deleted.

```bash
$ convox rack params set DynamoDbTableDeletionProtectionEnabled=true
```

## See Also

- [DynamoDbTablePointInTimeRecoveryEnabled](/reference/rack-parameters/DynamoDbTablePointInTimeRecoveryEnabled)
- [Encryption](/reference/rack-parameters/Encryption)
- [Rack Parameters](/reference/rack-parameters) for a full list of available parameters
