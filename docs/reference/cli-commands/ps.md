---
title: "ps"
description: "List running Processes for an App."
---

# ps

List running Processes for an App. Shows each Process's ID, Service name, associated Release, start time, and command. Use the `--service` flag to filter by a specific Service.

## Syntax

```bash
$ convox ps
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |
| `--release` | | Filter by Release |
| `--service` | `-s` | Filter by Service |

## Example Usage

```bash
$ convox ps -a myapp
ID                SERVICE  STATUS   RELEASE  STARTED      COMMAND
web-abc1234-def5  web      running  RABCDEF  2 hours ago  bundle exec rails server
web-ghi6789-jkl0  web      running  RABCDEF  2 hours ago  bundle exec rails server
worker-mno1234    worker   running  RABCDEF  2 hours ago  bundle exec sidekiq
```

## See Also

- [ps info](/reference/cli-commands/ps-info)
- [ps stop](/reference/cli-commands/ps-stop)
- [exec](/reference/cli-commands/exec)
- [run](/reference/cli-commands/run)
- [services](/reference/cli-commands/services)
- [rack ps](/reference/cli-commands/rack-ps)
- [One-off Commands](/management/one-off-commands)
