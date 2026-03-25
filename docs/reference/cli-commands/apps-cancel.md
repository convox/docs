---
title: "apps cancel"
description: "Cancel an in-progress App update and roll back to the last active Release."
---

# apps cancel

Cancel an in-progress App update and roll back to the last active Release. This is useful when an update is taking too long or you need to revert a problematic change quickly.

## Syntax

```bash
$ convox apps cancel [app]
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

## Example Usage

```bash
$ convox apps cancel myapp
Cancelling deployment of myapp...
Rewriting last active release...
OK
```

## See Also

- [apps info](/reference/cli-commands/apps-info)
- [apps wait](/reference/cli-commands/apps-wait)
- [releases rollback](/reference/cli-commands/releases-rollback)
- [Rolling Back](/deployment/rolling-back)
