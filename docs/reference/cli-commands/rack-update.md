---
title: "rack update"
description: "Update the Rack to the next version or a specific version."
---

# rack update

Update the Rack to the next version, or to a specific version if specified. Updates are applied as rolling updates to minimize downtime. If no version is provided, the Rack updates to the next sequential version.

## Syntax

```bash
$ convox rack update [version]
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

## Example Usage

```bash
$ convox rack update
Updating to 3.18.6... OK

$ convox rack update 3.19.0 --wait
Updating to 3.19.0... OK
```

## See Also

- [rack releases](/reference/cli-commands/rack-releases)
- [rack wait](/reference/cli-commands/rack-wait)
- [rack](/reference/cli-commands/rack)
- [Rack Updates](/management/rack-updates)
