---
title: "rack resources info"
description: "Get information about a Rack Resource."
---

# rack resources info

Display detailed information about a Rack Resource, including its name, type, status, configurable options, connection URL, and any linked Apps.

## Syntax

```bash
$ convox rack resources info <resource>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox rack resources info shared-postgres
Name     shared-postgres
Type     postgres
Status   running
Options  AllocatedStorage=20
         AllowMajorVersionUpgrade=false
         AutoMinorVersionUpgrade=true
         BackupRetentionPeriod=1
         Database=app
         DatabaseSnapshotIdentifier=
         Encrypted=false
         EngineVersion=17
         Family=postgres17
         InstanceType=db.t3.micro
         MaxConnections=
         MultiAZ=true
         Username=postgres
URL      postgres://postgres:password@myrack-shared-postgres.abcdefghijkl.us-east-1.rds.amazonaws.com:5432/app
```

## See Also

- [rack resources](/reference/cli-commands/rack-resources)
- [rack resources url](/reference/cli-commands/rack-resources-url)
- [rack resources options](/reference/cli-commands/rack-resources-options)
- [rack resources update](/reference/cli-commands/rack-resources-update)
- [Rack and External Resources](/management/rack-resources)
