---
title: "services restart"
description: "Restart all Processes for a specific Service."
---

# services restart

Restart all Processes for a specific Service. Unlike `restart`, which restarts every Service in the App, this command targets a single Service. This is useful when you need to refresh one component without disrupting the rest of the App.

## Syntax

```bash
$ convox services restart <service>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox services restart web -a myapp
Restarting web... OK
```

## See Also

- [restart](/reference/cli-commands/restart)
- [services](/reference/cli-commands/services)
- [ps stop](/reference/cli-commands/ps-stop)
- [Services](/application/services)
