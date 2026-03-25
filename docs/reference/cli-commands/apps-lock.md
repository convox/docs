---
title: "apps lock"
description: "Enable termination protection for an App."
---

# apps lock

Enable termination protection for an App. A locked App cannot be deleted until it is explicitly unlocked. This is useful for protecting production Apps from accidental deletion.

## Syntax

```bash
$ convox apps lock [app]
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox apps lock myapp
Locking myapp... OK
```

## See Also

- [apps unlock](/reference/cli-commands/apps-unlock)
- [apps delete](/reference/cli-commands/apps-delete)
- [apps info](/reference/cli-commands/apps-info)
