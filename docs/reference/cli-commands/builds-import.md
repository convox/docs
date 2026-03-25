---
title: "builds import"
description: "Import a Build from a previously exported tarball."
---

# builds import

Import a Build from a previously exported tarball. The tarball should have been created with `convox builds export`. Reads from stdin by default, or from a file specified with the `--file` flag.

## Syntax

```bash
$ convox builds import
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--file` | `-f` | Import from file |
| `--id` | | Send logs to stderr, Release ID to stdout |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox builds import -a myapp -f myapp-build.tgz
Importing build... OK, RABCDE2
```

## See Also

- [builds export](/reference/cli-commands/builds-export)
- [apps import](/reference/cli-commands/apps-import)
- [builds info](/reference/cli-commands/builds-info)
- [Builds](/deployment/builds)
