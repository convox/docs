---
title: "instances"
description: "List Rack instances."
---

# instances

List Rack instances. Shows each instance's ID, status, type, startup time, and resource usage. Use this to monitor the health and capacity of the underlying infrastructure.

## Syntax

```bash
$ convox instances
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox instances
ID          STATUS   STARTED     PS  CPU    MEM    PUBLIC         PRIVATE
i-0a1b2c3d  running  2 days ago  12  31.2%  48.5%  54.1.2.3      10.0.1.100
i-4e5f6a7b  running  5 days ago  8   22.7%  35.1%  54.4.5.6      10.0.2.101
i-8c9d0e1f  running  5 days ago  10  27.4%  42.3%  54.7.8.9      10.0.3.102
```

## See Also

- [instances ssh](/reference/cli-commands/instances-ssh)
- [instances terminate](/reference/cli-commands/instances-terminate)
- [instances keyroll](/reference/cli-commands/instances-keyroll)
- [rack scale](/reference/cli-commands/rack-scale)
- [Scaling](/scaling/scaling)
