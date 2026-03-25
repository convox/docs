---
title: "DynamoDbTablePointInTimeRecoveryEnabled"
description: "Enables point-in-time recovery (PITR) on DynamoDB tables created by a Convox Rack."
---

# DynamoDbTablePointInTimeRecoveryEnabled

Point-in-time recovery for the Rack's DynamoDB table.

| Default value  | `false`         |
| Allowed values | `true`, `false` |

## Use Cases

- Enable on production racks to allow restoring DynamoDB table data to any point within the last 35 days.
- Required by compliance standards (e.g., SOC 2, HIPAA) that mandate continuous backup capabilities for data stores.
- Provides an additional safety net against accidental data corruption or deletion at the item level.

## Additional Information

When enabled, AWS DynamoDB continuously backs up your table data, allowing you to restore the table to any second within the preceding 35 days. This is in addition to any on-demand backups you may create.

Point-in-time recovery incurs additional AWS charges based on the size of the table. See the [AWS DynamoDB pricing page](https://aws.amazon.com/dynamodb/pricing/) for details.

```bash
$ convox rack params set DynamoDbTablePointInTimeRecoveryEnabled=true
```

## See Also

- [DynamoDbTableDeletionProtectionEnabled](/reference/rack-parameters/DynamoDbTableDeletionProtectionEnabled)
- [Encryption](/reference/rack-parameters/Encryption)
- [Rack Parameters](/reference/rack-parameters) for a full list of available parameters
