---
title: "rack resources delete"
description: "Delete a Rack-level Resource."
---

# rack resources delete

Delete a Rack-level Resource. The Resource must be unlinked from all Apps before it can be deleted.

## Syntax

```bash
$ convox rack resources delete <name>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

## Example Usage

```bash
$ convox rack resources delete shared-postgres --wait
Deleting resource... OK
```

## See Also

- [rack resources create](/reference/cli-commands/rack-resources-create)
- [rack resources unlink](/reference/cli-commands/rack-resources-unlink)
- [rack resources](/reference/cli-commands/rack-resources)
- [Rack and External Resources](/management/rack-resources)
