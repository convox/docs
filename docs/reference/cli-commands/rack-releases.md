---
title: "rack releases"
description: "List Rack version history."
---

# rack releases

List Rack version history. Shows the sequence of Rack updates, including the version installed and when it was updated.

## Syntax

```bash
$ convox rack releases
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox rack releases
VERSION  UPDATED
3.18.5   2 days ago
3.18.4   1 week ago
3.18.3   2 weeks ago
3.18.2   3 weeks ago
```

## See Also

- [rack update](/reference/cli-commands/rack-update)
- [rack](/reference/cli-commands/rack)
- [releases](/reference/cli-commands/releases)
- [Rack Updates](/management/rack-updates)
