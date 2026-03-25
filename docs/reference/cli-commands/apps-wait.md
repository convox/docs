---
title: "apps wait"
description: "Wait for an App to finish updating."
---

# apps wait

Wait for an App to finish updating. This command blocks until the App reaches a stable state, making it useful in CI/CD pipelines and scripts that need to proceed only after an update is complete.

## Syntax

```bash
$ convox apps wait [app]
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox apps wait myapp
Waiting for app... OK
```

## See Also

- [apps info](/reference/cli-commands/apps-info)
- [apps cancel](/reference/cli-commands/apps-cancel)
- [rack wait](/reference/cli-commands/rack-wait)
- [App Statuses](/reference/app-statuses)
