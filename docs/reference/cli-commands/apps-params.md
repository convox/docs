---
title: "apps params"
description: "Display App parameters and their current values."
---

# apps params

Display App parameters and their current values. Parameters control infrastructure-level settings for the App. Use `convox apps params set` to modify parameters.

## Syntax

```bash
$ convox apps params [app]
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox apps params myapp
FargateServices  No
Internal         No
RedirectHttps    Yes
TaskTags         No
```

## See Also

- [apps params set](/reference/cli-commands/apps-params-set)
- [rack params](/reference/cli-commands/rack-params)
- [App Parameters](/reference/app-parameters)
