---
title: "env edit"
description: "Edit environment variables interactively using your default editor."
---

# env edit

Edit environment variables interactively using `$EDITOR`. Opens the current environment in your default editor. Saving and closing the editor creates a new Release with the updated variables.

## Syntax

```bash
$ convox env edit
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--promote` | `-p` | Promote the resulting Release |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

## Example Usage

```bash
$ convox env edit -a myapp
Setting VAR1, VAR2... OK
Release: RABCDEF
```

## See Also

- [env](/reference/cli-commands/env)
- [env set](/reference/cli-commands/env-set)
- [env get](/reference/cli-commands/env-get)
- [env unset](/reference/cli-commands/env-unset)
- [Environment](/application/environment)
