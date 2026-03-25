---
title: "apps delete"
description: "Delete an App and all of its associated resources from the current Rack."
---

# apps delete

Delete an App and all of its associated resources from the current Rack. The App must be unlocked before it can be deleted. Use `convox apps unlock` if the App is currently locked.

## Syntax

```bash
$ convox apps delete <app>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

## Example Usage

```bash
$ convox apps delete myapp
Deleting myapp... OK
```

## See Also

- [apps create](/reference/cli-commands/apps-create)
- [apps lock](/reference/cli-commands/apps-lock)
- [apps unlock](/reference/cli-commands/apps-unlock)
