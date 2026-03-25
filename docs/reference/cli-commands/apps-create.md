---
title: "apps create"
description: "Create a new App on the current Rack."
---

# apps create

Create a new App on the current Rack. By default, Apps are created as generation 2. After creation, you can deploy code to the App using `convox deploy`.

## Syntax

```bash
$ convox apps create [name]
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--generation` | `-g` | App generation (default: `2`) |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

## Example Usage

```bash
$ convox apps create myapp
Creating myapp... OK
```

## See Also

- [apps delete](/reference/cli-commands/apps-delete)
- [apps info](/reference/cli-commands/apps-info)
- [deploy](/reference/cli-commands/deploy)
- [Creating an Application](/deployment/creating-an-application)
