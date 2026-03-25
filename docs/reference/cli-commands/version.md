---
title: "version"
description: "Display client and server version information."
---

# version

Display the version of the local Convox CLI (client) and the connected Rack (server). This is useful for verifying compatibility between your CLI and Rack.

## Syntax

```bash
$ convox version
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox version
client: 3.18.5
server: 3.18.4 (myorg/production)
```

## See Also

- [update](/reference/cli-commands/update)
- [rack](/reference/cli-commands/rack)
- [rack update](/reference/cli-commands/rack-update)
