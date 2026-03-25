---
title: "rack resources"
description: "List Rack-level Resources."
---

# rack resources

List the Resources managed at the Rack level. Unlike App Resources defined in `convox.yml`, Rack Resources exist independently and can be linked to one or more Apps.

## Syntax

```bash
$ convox rack resources
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox rack resources
NAME             TYPE       STATUS
shared-postgres  postgres   running
shared-redis     redis      running
```

## See Also

- [rack resources create](/reference/cli-commands/rack-resources-create)
- [rack resources info](/reference/cli-commands/rack-resources-info)
- [resources](/reference/cli-commands/resources)
- [Rack and External Resources](/management/rack-resources)
