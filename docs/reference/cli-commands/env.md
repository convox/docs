---
title: "env"
description: "List environment variables for an App."
---

# env

List environment variables for an App. Displays all currently set variables and their values. Environment variables are tied to Releases; changing them creates a new Release.

## Syntax

```bash
$ convox env
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox env -a myapp
DATABASE_URL=postgres://dbhost:5432/myapp_production
RAILS_ENV=production
REDIS_URL=redis://redis.example.com:6379
SECRET_KEY_BASE=abc123def456ghi789jkl012mno345
```

## See Also

- [env set](/reference/cli-commands/env-set)
- [env get](/reference/cli-commands/env-get)
- [env edit](/reference/cli-commands/env-edit)
- [env unset](/reference/cli-commands/env-unset)
- [Environment](/application/environment)
