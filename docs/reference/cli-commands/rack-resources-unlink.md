---
title: "rack resources unlink"
description: "Unlink a Rack Resource from an App."
---

# rack resources unlink

Unlink a Rack Resource from an App, removing the Resource's connection URL from the App's environment variables.

## Syntax

```bash
$ convox rack resources unlink <resource>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

## Example Usage

```bash
$ convox rack resources unlink shared-postgres -a myapp --wait
Unlinking from myapp... OK
```

## See Also

- [rack resources link](/reference/cli-commands/rack-resources-link)
- [rack resources delete](/reference/cli-commands/rack-resources-delete)
- [rack resources](/reference/cli-commands/rack-resources)
- [Rack and External Resources](/management/rack-resources)
