---
title: "env unset"
description: "Remove one or more environment variables from an App."
---

# env unset

Remove one or more environment variables from an App. Unsetting variables creates a new Release with the updated environment configuration.

## Syntax

```bash
$ convox env unset <key> [key]...
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--id` | | Send logs to stderr, Release ID to stdout |
| `--promote` | `-p` | Promote the resulting Release |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

## Example Usage

```bash
$ convox env unset DEBUG_MODE LEGACY_FLAG -a myapp
Unsetting DEBUG_MODE, LEGACY_FLAG... OK
Release: RABCDEF
```

## See Also

- [env edit](/reference/cli-commands/env-edit)
- [env set](/reference/cli-commands/env-set)
- [env get](/reference/cli-commands/env-get)
- [env](/reference/cli-commands/env)
- [Environment](/application/environment)
