---
title: "rack sync"
description: "Sync v2 Rack API URL."
---

# rack sync

Sync v2 Rack API URL. Updates the locally stored API endpoint for the Rack to ensure it matches the current infrastructure. This command is used with v2 Racks to ensure the locally stored API endpoint is current. Run this if you experience connectivity issues after a Rack update or IP change.

## Syntax

```bash
$ convox rack sync
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--name` | `-n` | Rack name (displays the Rack API URL) |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox rack sync
Synchronizing rack API URL...
OK

$ convox rack sync --name production
Fetching rack API URL...
url=https://convox:abc123def456@rack.0a1b2c3d4e5f.convox.cloud
OK
```

## See Also

- [rack](/reference/cli-commands/rack)
- [login](/reference/cli-commands/login)
- [switch](/reference/cli-commands/switch)
