---
title: "releases promote"
description: "Promote a Release to make it the active deployment."
---

# releases promote

Promote a Release to make it the active deployment. If no Release is specified, promotes the most recent Release. This triggers a rolling update to deploy the new Release, gradually replacing old Processes with new ones.

## Syntax

```bash
$ convox releases promote [release]
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

## Example Usage

```bash
$ convox releases promote RABCDEF -a myapp
Promoting RABCDEF... OK
```

## See Also

- [releases rollback](/reference/cli-commands/releases-rollback)
- [deploy](/reference/cli-commands/deploy)
- [releases info](/reference/cli-commands/releases-info)
- [Releases](/deployment/releases)
- [Rolling Updates](/deployment/rolling-updates)
