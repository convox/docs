---
title: "apps import"
description: "Import an App from a previously exported tarball."
---

# apps import

Import an App from a previously exported tarball. The tarball should have been created with `convox apps export`. This restores the App's Build and environment onto the target Rack.

## Syntax

```bash
$ convox apps import [app]
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--file` | `-f` | Import from file |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox apps import myapp -f myapp-export.tgz
Creating app myapp... OK
Importing build... OK, RABCDEF
Importing env... OK, RABCDEG
Promoting RABCDEG... OK
```

## See Also

- [apps export](/reference/cli-commands/apps-export)
- [builds import](/reference/cli-commands/builds-import)
