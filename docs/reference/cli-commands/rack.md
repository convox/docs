---
title: "rack"
description: "Display information about the current Rack."
---

# rack

Display information about the current Rack, including its name, provider, region, router endpoint, status, and version.

## Syntax

```bash
$ convox rack
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox rack
Name      production
Provider  aws
Region    us-east-1
Router    router.0a1b2c3d4e5f.convox.cloud
Status    running
Version   3.18.5
```

## See Also

- [racks](/reference/cli-commands/racks)
- [rack update](/reference/cli-commands/rack-update)
- [rack install](/reference/cli-commands/rack-install)
- [rack uninstall](/reference/cli-commands/rack-uninstall)
- [rack params](/reference/cli-commands/rack-params)
- [Rack Statuses](/reference/rack-statuses)
