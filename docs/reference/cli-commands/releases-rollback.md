---
title: "releases rollback"
description: "Copy an old Release forward and promote it to revert a deployment."
---

# releases rollback

Copy an old Release forward (preserving its Build and environment) and promote the new Release. This is the safest way to revert a deployment because it creates a new Release rather than modifying history. The original Release remains in the history for auditing.

## Syntax

```bash
$ convox releases rollback <release>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--id` | | Send logs to stderr, Release ID to stdout |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

## Example Usage

```bash
$ convox releases rollback RABCDEH -a myapp
Rolling back to RABCDEH... OK, RABCDEJ
Promoting RABCDEJ... OK
```

## See Also

- [releases promote](/reference/cli-commands/releases-promote)
- [releases info](/reference/cli-commands/releases-info)
- [apps cancel](/reference/cli-commands/apps-cancel)
- [deploy](/reference/cli-commands/deploy)
- [Releases](/deployment/releases)
- [Rolling Back](/deployment/rolling-back)
