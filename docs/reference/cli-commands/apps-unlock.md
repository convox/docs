---
title: "apps unlock"
description: "Disable termination protection for an App."
---

# apps unlock

Disable termination protection for an App. Once unlocked, the App can be deleted with `convox apps delete`.

## Syntax

```bash
$ convox apps unlock [app]
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox apps unlock myapp
Unlocking myapp... OK
```

## See Also

- [apps lock](/reference/cli-commands/apps-lock)
- [apps delete](/reference/cli-commands/apps-delete)
- [apps info](/reference/cli-commands/apps-info)
