---
title: "releases"
description: "List Releases for an App."
---

# releases

List Releases for an App. Each Release represents a specific Build plus environment configuration. The output includes the Release ID, status, associated Build, creation timestamp, and description.

## Syntax

```bash
$ convox releases
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--limit` | `-l` | Number of Releases to list |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox releases -a myapp
ID       STATUS  BUILD    CREATED      DESCRIPTION
RABCDEF  active  BABCDEF  2 hours ago  deployed via cli
RABCDEG          BABCDEG  1 day ago    env update
RABCDEH          BABCDEH  2 days ago   initial deploy
```

## See Also

- [releases info](/reference/cli-commands/releases-info)
- [releases promote](/reference/cli-commands/releases-promote)
- [releases rollback](/reference/cli-commands/releases-rollback)
- [deploy](/reference/cli-commands/deploy)
- [rack releases](/reference/cli-commands/rack-releases)
- [Releases](/deployment/releases)
