---
title: "rack scale"
description: "Display or update Rack scaling parameters."
---

# rack scale

Display or update Rack scaling parameters. When called without flags, displays the current scaling configuration including autoscale status, instance count, and instance type.

## Syntax

```bash
$ convox rack scale
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--count` | `-c` | Instance count |
| `--rack` | `-r` | Rack name |
| `--type` | `-t` | Instance type |

## Example Usage

```bash
$ convox rack scale
Autoscale  true
Count      3
Status     running
Type       t3.medium

$ convox rack scale --count 5 --type t3.large
Scaling rack... OK
```

## See Also

- [rack params](/reference/cli-commands/rack-params)
- [rack](/reference/cli-commands/rack)
- [instances](/reference/cli-commands/instances)
- [scale](/reference/cli-commands/scale)
- [Scaling](/scaling/scaling)
