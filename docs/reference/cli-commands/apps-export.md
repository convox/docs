---
title: "apps export"
description: "Export an App, including its Build and environment, to a tarball."
---

# apps export

Export an App to a tarball that includes its Build and environment variables. This is useful for migrating an App between Racks or creating a portable backup.

## Syntax

```bash
$ convox apps export [app]
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--file` | `-f` | Export to file |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox apps export myapp -f myapp-export.tgz
Exporting app myapp... OK
Exporting env... OK
Exporting build BABCDEF... OK
Packaging export... OK
```

## See Also

- [apps import](/reference/cli-commands/apps-import)
- [builds export](/reference/cli-commands/builds-export)
