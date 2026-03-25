---
title: "apps params set"
description: "Set one or more App parameters."
---

# apps params set

Set one or more App parameters. Changes may trigger an App update that modifies the underlying infrastructure. You can set multiple parameters in a single command.

## Syntax

```bash
$ convox apps params set <Key=Value> [Key=Value]...
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

## Example Usage

```bash
$ convox apps params set FargateServices=Yes -a myapp
Updating parameters... OK
```

## See Also

- [apps params](/reference/cli-commands/apps-params)
- [rack params set](/reference/cli-commands/rack-params-set)
- [App Parameters](/reference/app-parameters)
