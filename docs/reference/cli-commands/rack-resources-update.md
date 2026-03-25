---
title: "rack resources update"
description: "Update options for a Rack Resource."
---

# rack resources update

Update the configurable options for an existing Rack Resource. Depending on the option changed, the Resource may be recreated, which can cause brief downtime and a new connection URL.

## Syntax

```bash
$ convox rack resources update <name> [Option=Value]...
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

## Example Usage

```bash
$ convox rack resources update shared-postgres Storage=50 --wait
Updating resource... OK
```

## See Also

- [rack resources info](/reference/cli-commands/rack-resources-info)
- [rack resources options](/reference/cli-commands/rack-resources-options)
- [rack resources](/reference/cli-commands/rack-resources)
- [Rack and External Resources](/management/rack-resources)
