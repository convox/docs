---
title: "env set"
description: "Set one or more environment variables for an App."
---

# env set

Set one or more environment variables for an App. Values can also be piped via stdin. Setting variables creates a new Release with the updated environment configuration.

## Syntax

```bash
$ convox env set <key=value> [key=value]...
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--id` | | Send logs to stderr, Release ID to stdout |
| `--promote` | `-p` | Promote the resulting Release |
| `--rack` | `-r` | Rack name |
| `--replace` | | Replace all environment variables with the given ones |
| `--wait` | `-w` | Wait for completion |

## Example Usage

```bash
$ convox env set DATABASE_URL=postgres://dbhost:5432/myapp_production SECRET_KEY_BASE=abc123 -a myapp
Setting DATABASE_URL, SECRET_KEY_BASE... OK
Release: RABCDEF
```

```bash
$ cat .env | convox env set -a myapp
Setting DATABASE_URL, SECRET_KEY_BASE... OK
Release: RABCDEG
```

## See Also

- [env](/reference/cli-commands/env)
- [env unset](/reference/cli-commands/env-unset)
- [env get](/reference/cli-commands/env-get)
- [env edit](/reference/cli-commands/env-edit)
- [Environment](/application/environment)
- [Releases](/deployment/releases)
