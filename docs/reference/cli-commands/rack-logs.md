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
2025-01-15T12:00:00Z system/cloudformation aws/cfm myapp UPDATE_IN_PROGRESS AWS::CloudFormation::Stack User Initiated
2025-01-15T12:00:05Z system/ecs aws/ecs (service myapp-web) has started 1 tasks: (task abc1234).
2025-01-15T12:00:20Z system/ecs aws/ecs (service myapp-web) has reached a steady state.
2025-01-15T12:00:25Z system/cloudformation aws/cfm myapp UPDATE_COMPLETE AWS::CloudFormation::Stack
```

## See Also

- [logs](/reference/cli-commands/logs)
- [Logs](/management/logs)
- [Debugging](/management/debugging)
