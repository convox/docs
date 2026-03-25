---
title: "builds"
description: "List Builds for an App."
---

# builds

List Builds for an App, showing each Build's ID, status, associated Release, description, and start time. The most recent Builds are shown first.

## Syntax

```bash
$ convox builds
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--limit` | `-l` | Number of Builds to list |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox builds -a myapp
ID       STATUS    RELEASE  STARTED      ELAPSED  DESCRIPTION
BABCDEF  complete  RABCDEF  2 hours ago  47s      feature branch build
BABCDE0  complete  RABCDE0  1 day ago    1m32s    initial deploy
BABCDE1  failed    RABCDE1  2 days ago   12s      dependency update
```

## See Also

- [build](/reference/cli-commands/build)
- [builds info](/reference/cli-commands/builds-info)
- [builds logs](/reference/cli-commands/builds-logs)
- [Builds](/deployment/builds)
