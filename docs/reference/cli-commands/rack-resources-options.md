---
title: "rack resources options"
description: "List available options for a Resource type."
---

# rack resources options

List the configurable parameters available for a given Resource type. Use this to discover which options can be set when creating or updating a Resource.

## Syntax

```bash
$ convox rack resources options <type>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox rack resources options postgres
NAME                        DEFAULT      DESCRIPTION
AllocatedStorage            20           Allocated storage size (GB)
AllowMajorVersionUpgrade    false
AutoMinorVersionUpgrade     true
BackupRetentionPeriod       1            The automatic RDS backup retention period, (default 1 day)
Database                    app          Default database name
DatabaseSnapshotIdentifier               ARN of database snapshot to restore
Encrypted                   false        Encrypt database with KMS
EngineVersion               17           Version of Postgres
Family                      postgres17   Postgres version family
InstanceType                db.t3.micro  Instance class for database nodes
MaxConnections                           ParameterGroup max_connections value, i.e. '{DBInstanceClassMemory/15000000}'
MultiAZ                     false        Multiple availability zone
Password                    (generated)  Server password
Username                    postgres     Server username
```

## See Also

- [rack resources create](/reference/cli-commands/rack-resources-create)
- [rack resources update](/reference/cli-commands/rack-resources-update)
- [rack resources types](/reference/cli-commands/rack-resources-types)
- [Rack and External Resources](/management/rack-resources)
