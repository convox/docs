---
title: "rack access key rotate"
description: "Rotate the Rack access key."
---

# rack access key rotate

Rotate the Rack access key. Generates a new API key and invalidates the previous one. This is useful for maintaining security by regularly cycling credentials.

## Syntax

```bash
$ convox rack access key rotate
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox rack access key rotate
OK
```

## See Also

- [rack access](/reference/cli-commands/rack-access)
- [rack](/reference/cli-commands/rack)
- [instances keyroll](/reference/cli-commands/instances-keyroll)
