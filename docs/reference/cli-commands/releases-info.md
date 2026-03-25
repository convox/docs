---
title: "releases info"
description: "Display information about a Release."
---

# releases info

Display detailed information about a Release, including its ID, associated Build, creation timestamp, description, and environment variables.

## Syntax

```bash
$ convox releases info <release>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox releases info RABCDEF -a myapp
Id           RABCDEF
Build        BABCDEF
Created      2025-01-15T12:00:00Z
Description  deployed via cli
Env          DATABASE_URL=postgres://host/myapp
             SECRET_KEY_BASE=abc123def456
```

## See Also

- [releases](/reference/cli-commands/releases)
- [releases manifest](/reference/cli-commands/releases-manifest)
- [releases promote](/reference/cli-commands/releases-promote)
- [Releases](/deployment/releases)
