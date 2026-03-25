---
title: "scale"
description: "Display or update scaling parameters for Services."
---

# scale

Display or update scaling parameters for Services. When called without `--count`, `--cpu`, or `--memory`, displays the current scaling configuration for all Services. When flags are provided, updates the specified Service's scaling parameters.

## Syntax

```bash
$ convox scale [service]
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--count` | | Desired Process count |
| `--cpu` | | CPU units (1024 = 1 vCPU) |
| `--memory` | | Memory in MB |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

## Example Usage

```bash
$ convox scale -a myapp
SERVICE  DESIRED  RUNNING  CPU  MEMORY
web      2        2        256  512
worker   1        1        512  1024

$ convox scale web --count 3 -a myapp
```

## See Also

- [services](/reference/cli-commands/services)
- [ps](/reference/cli-commands/ps)
- [rack scale](/reference/cli-commands/rack-scale)
- [Scaling](/scaling/scaling)
- [Services](/application/services)
