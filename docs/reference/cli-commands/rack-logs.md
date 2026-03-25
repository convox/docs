---
title: "rack logs"
description: "Stream logs for the Rack."
---

# rack logs

Stream logs for the Rack. Shows system-level logs including scheduler events, load balancer activity, and infrastructure changes. This is useful for diagnosing Rack-level issues that are not visible in App logs.

## Syntax

```bash
$ convox rack logs
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--filter` | | Filter logs by a pattern |
| `--no-follow` | | Do not follow log output |
| `--rack` | `-r` | Rack name |
| `--since` | | Show logs since a duration (default: `2m`) |

## Example Usage

```bash
$ convox rack logs --since 10m
2025-01-15T12:00:00Z system/k8s/scheduler  Successfully assigned myapp/web-abc1234 to node ip-10-0-1-50
2025-01-15T12:00:02Z system/k8s/web        Pulling image "123456789012.dkr.ecr.us-east-1.amazonaws.com/convox/myapp:web.BABCDEF"
2025-01-15T12:00:08Z system/k8s/web        Started container main
2025-01-15T12:00:15Z system/elb             Registered instance i-0a1b2c3d4e on target group myapp-web
2025-01-15T12:00:20Z system/elb             Health check passed for target i-0a1b2c3d4e
```

## See Also

- [logs](/reference/cli-commands/logs)
- [Logs](/management/logs)
- [Debugging](/management/debugging)
