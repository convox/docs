---
title: "apps"
description: "List all Apps on the current Rack."
---

# apps

List all Apps on the current Rack. The output includes each App's name and its current status.

## Syntax

```bash
$ convox apps
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox apps
APP       STATUS    RELEASE
api       running   RABCDEF
web       running   RABCDEG
worker    updating  RABCDEH
```

## See Also

- [apps create](/reference/cli-commands/apps-create)
- [apps info](/reference/cli-commands/apps-info)
- [apps delete](/reference/cli-commands/apps-delete)
- [App Statuses](/reference/app-statuses)
