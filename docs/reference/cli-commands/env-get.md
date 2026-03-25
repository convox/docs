---
title: "env get"
description: "Get the value of a single environment variable."
---

# env get

Get the value of a single environment variable. Returns just the value, making it suitable for use in scripts and pipelines.

## Syntax

```bash
$ convox env get <var>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox env get DATABASE_URL -a myapp
postgres://dbhost:5432/myapp_production
```

## See Also

- [env](/reference/cli-commands/env)
- [env set](/reference/cli-commands/env-set)
- [env edit](/reference/cli-commands/env-edit)
- [env unset](/reference/cli-commands/env-unset)
- [Environment](/application/environment)
