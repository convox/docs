---
title: "apps info"
description: "Display detailed information about an App."
---

# apps info

Display detailed information about an App, including its name, status, generation, lock state, current Release, and router endpoint.

## Syntax

```bash
$ convox apps info [app]
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox apps info myapp
Name        myapp
Status      running
Generation  2
Locked      false
Release     RABCDEF
Router      router.0a1b2c3d4e5f.convox.cloud
```

## See Also

- [apps](/reference/cli-commands/apps)
- [apps params](/reference/cli-commands/apps-params)
- [App Statuses](/reference/app-statuses)
