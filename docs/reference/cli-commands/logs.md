---
title: "logs"
description: "Stream logs for an App."
---

# logs

Stream logs for an App. Aggregates output from all running Processes across every Service. Use `--filter` to narrow results to specific patterns and `--no-follow` for a one-time snapshot instead of a continuous stream.

## Syntax

```bash
$ convox logs
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--filter` | | Filter logs by a pattern |
| `--no-follow` | | Do not follow log output |
| `--rack` | `-r` | Rack name |
| `--since` | | Show logs since a duration (default: `2m`) |

## Example Usage

```bash
$ convox logs -a myapp --since 5m
2025-01-15T12:00:00Z service/web/0a1b2c3d Started worker process
2025-01-15T12:00:01Z service/web/0a1b2c3d Listening on 0.0.0.0:3000
2025-01-15T12:00:05Z service/web/0a1b2c3d GET /health 200 1ms
2025-01-15T12:00:10Z service/worker/4e5f6a7b Processing job id=8842
2025-01-15T12:00:12Z service/worker/4e5f6a7b Job complete id=8842 duration=2.1s
```

## See Also

- [rack logs](/reference/cli-commands/rack-logs)
- [builds logs](/reference/cli-commands/builds-logs)
- [Logs](/management/logs)
