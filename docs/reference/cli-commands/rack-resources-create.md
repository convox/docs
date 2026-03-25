---
title: "rack resources create"
description: "Create a Rack-level Resource."
---

# rack resources create

Create a new Resource at the Rack level. The type specifies the kind of Resource to create (e.g., `postgres`, `redis`, `memcached`). Optional key-value pairs configure the Resource.

## Syntax

```bash
$ convox rack resources create <type> [Option=Value]...
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--name` | `-n` | Resource name |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

## Example Usage

```bash
$ convox rack resources create postgres MultiAZ=true -n shared-postgres --wait
Creating resource... OK, shared-postgres
```

## See Also

- [rack resources delete](/reference/cli-commands/rack-resources-delete)
- [rack resources info](/reference/cli-commands/rack-resources-info)
- [rack resources types](/reference/cli-commands/rack-resources-types)
- [rack resources link](/reference/cli-commands/rack-resources-link)
- [Rack and External Resources](/management/rack-resources)
