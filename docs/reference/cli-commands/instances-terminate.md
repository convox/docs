---
title: "instances terminate"
description: "Terminate an instance."
---

# instances terminate

Terminate an instance. The Rack's Auto Scaling group launches a replacement instance. This is useful for recycling unhealthy instances or forcing fresh instance provisioning.

## Syntax

```bash
$ convox instances terminate <instance>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox instances terminate i-0a1b2c3d
Terminating instance... OK
```

## See Also

- [instances](/reference/cli-commands/instances)
- [rack scale](/reference/cli-commands/rack-scale)
- [Scaling](/scaling/scaling)
