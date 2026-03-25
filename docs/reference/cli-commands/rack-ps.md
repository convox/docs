---
title: "rack ps"
description: "List Rack system Processes."
---

# rack ps

List Rack system Processes. Shows internal services running on the Rack infrastructure, including their current status and start time.

## Syntax

```bash
$ convox rack ps
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--all` | `-a` | Show all Processes |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox rack ps
ID                   APP     SERVICE     STATUS   RELEASE   STARTED          COMMAND
rack-api-abc1234     system  api         running  RABCDEF   2 days ago
rack-router-def5678  system  router      running  RABCDEF   2 days ago
rack-monitor-ghi901  system  monitor     running  RABCDEF   2 days ago
```

## See Also

- [rack](/reference/cli-commands/rack)
- [ps](/reference/cli-commands/ps)
- [rack scale](/reference/cli-commands/rack-scale)
