---
title: "builds info"
description: "Display detailed information about a Build."
---

# builds info

Display detailed information about a Build, including its ID, status, associated Release, description, start time, and elapsed duration.

## Syntax

```bash
$ convox builds info <build>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox builds info BABCDEF -a myapp
Id           BABCDEF
Status       complete
Release      RABCDEF
Description  feature branch build
Started      2 hours ago
Elapsed      47s
```

## See Also

- [builds](/reference/cli-commands/builds)
- [builds logs](/reference/cli-commands/builds-logs)
- [build](/reference/cli-commands/build)
- [Builds](/deployment/builds)
