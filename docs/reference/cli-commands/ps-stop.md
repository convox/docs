---
title: "ps stop"
description: "Stop a running Process."
---

# ps stop

Stop a running Process. The Service scheduler will automatically start a replacement Process to maintain the desired count. This is useful for recycling a specific Process without affecting overall availability.

## Syntax

```bash
$ convox ps stop <pid>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox ps stop web-abc1234-def5 -a myapp
Stopping web-abc1234-def5... OK
```

## See Also

- [ps](/reference/cli-commands/ps)
- [ps info](/reference/cli-commands/ps-info)
- [services restart](/reference/cli-commands/services-restart)
