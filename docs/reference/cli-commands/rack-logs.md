---
title: "rack logs"
description: "Stream logs for the Rack."
---

# rack logs

Stream logs for the Rack. Shows the output of the Rack API (`service/web`) and the Rack monitor (`service/monitor`), the output of every Build (`build/build`), and CloudFormation events for the Rack stack, its Rack Resources, and Apps that do not log to CloudWatch (`system/cloudformation`). This is useful for diagnosing Rack-level issues that are not visible in App logs.

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
2025-01-15T12:00:00Z service/web/7c2e9f4a1b3d4e5f8a6b0c1d2e3f4a5b id=a1b2c3d4e5f6 ns=api at=SystemGet method="GET" path="/system" response=200 elapsed=70.493
2025-01-15T12:00:05Z service/monitor/3f8a1c6e2d4b4a7f9e0c5b1d8a2f6e3c ns=workers.monitor tick
2025-01-15T12:01:10Z system/cloudformation aws/cfm production UPDATE_IN_PROGRESS Instances Received SUCCESS signal with UniqueId i-0a1b2c3d4e5f67890
2025-01-15T12:04:25Z system/cloudformation aws/cfm production UPDATE_COMPLETE production
```

## See Also

- [logs](/reference/cli-commands/logs)
- [Logs](/management/logs)
- [Debugging](/management/debugging)
