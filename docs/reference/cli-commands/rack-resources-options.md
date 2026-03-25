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
NAME       DEFAULT      DESCRIPTION
MultiAZ    false        Run across multiple availability zones
Storage    20           Storage size in GB
Class      db.t3.micro  Instance class
Encrypted  true         Encrypt data at rest
Version    13           PostgreSQL engine version
```

## See Also

- [rack resources create](/reference/cli-commands/rack-resources-create)
- [rack resources update](/reference/cli-commands/rack-resources-update)
- [rack resources types](/reference/cli-commands/rack-resources-types)
- [Rack and External Resources](/management/rack-resources)
