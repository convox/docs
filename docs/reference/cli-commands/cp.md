---
title: "cp"
description: "Copy files to or from a running Process."
---

# cp

Copy files to or from a Process. Prefix a path with `pid:` to reference a running Process. Process paths must be absolute. This is useful for extracting logs, uploading configuration files, or retrieving generated artifacts.

## Syntax

```bash
$ convox cp <[pid:]src> <[pid:]dst>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox cp web-abc1234-def5:/app/log/production.log ./production.log -a myapp
```

```bash
$ convox cp ./config.yml web-abc1234-def5:/app/config/config.yml -a myapp
```

## See Also

- [exec](/reference/cli-commands/exec)
- [run](/reference/cli-commands/run)
- [ps](/reference/cli-commands/ps)
- [One-off Commands](/management/one-off-commands)
