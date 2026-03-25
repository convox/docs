---
title: "rack params set"
description: "Set one or more Rack parameters."
---

# rack params set

Set one or more Rack parameters. Changes may trigger a Rack update that modifies infrastructure. Multiple parameters can be set in a single command.

## Syntax

```bash
$ convox rack params set <Key=Value> [Key=Value]...
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

## Example Usage

```bash
$ convox rack params set InstanceType=t3.medium InstanceCount=5
Updating parameters... OK
```

## See Also

- [rack params](/reference/cli-commands/rack-params)
- [apps params set](/reference/cli-commands/apps-params-set)
- [Rack Parameters](/reference/rack-parameters)
- [rack wait](/reference/cli-commands/rack-wait)
