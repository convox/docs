---
title: "ps info"
description: "Display information about a running Process."
---

# ps info

Display detailed information about a Process, including its ID, App, command, instance, associated Release, Service name, start time, and current status.

## Syntax

```bash
$ convox ps info <pid>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox ps info 5e3c8576b942 -a myapp
Id        5e3c8576b942
App       myapp
Command   bundle exec rails server
Instance  i-0a1b2c3d4e5f67890
Release   RABCDEF
Service   web
Started   2 hours ago
Status    running
```

## See Also

- [ps](/reference/cli-commands/ps)
- [ps stop](/reference/cli-commands/ps-stop)
- [exec](/reference/cli-commands/exec)
