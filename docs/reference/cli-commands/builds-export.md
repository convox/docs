---
title: "builds export"
description: "Export a Build to a tarball for portability between Racks."
---

# builds export

Export a Build to a tarball. The exported tarball can be imported on another Rack using `convox builds import`, making it easy to promote tested Builds across environments.

## Syntax

```bash
$ convox builds export <build>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--file` | `-f` | Export to file |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox builds export BABCDEF -a myapp -f myapp-build.tgz
Exporting build... OK
```

## See Also

- [builds import](/reference/cli-commands/builds-import)
- [apps export](/reference/cli-commands/apps-export)
- [builds info](/reference/cli-commands/builds-info)
- [Builds](/deployment/builds)
