---
title: "restart"
description: "Restart all Processes for all Services of an App."
---

# restart

Restart all Processes for all Services of an App. Every running Process is stopped and replaced with a new one. For restarting a single Service, use `services restart` instead.

## Syntax

```bash
$ convox restart
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox restart -a myapp
Restarting api... OK
Restarting web... OK
Restarting worker... OK
```

## See Also

- [services restart](/reference/cli-commands/services-restart)
- [services](/reference/cli-commands/services)
- [ps](/reference/cli-commands/ps)
- [Services](/application/services)
