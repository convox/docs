---
title: "rack uninstall"
description: "Uninstall a Rack and remove all its infrastructure."
---

# rack uninstall

Uninstall a Rack. This permanently removes all Rack infrastructure including clusters, networking, and storage. Use `--force` for non-interactive use to skip the confirmation prompt.

## Syntax

```bash
$ convox rack uninstall <type> <name>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--force` | `-f` | Force uninstall (required for non-interactive use) |

## Example Usage

```bash
$ convox rack uninstall aws production
Uninstalling Rack production...
This will permanently destroy all infrastructure. Are you sure? (y/n): y
Deleting EKS cluster...
Deleting VPC...
Deleting IAM resources...
Rack production uninstalled successfully
```

## See Also

- [rack install](/reference/cli-commands/rack-install)
- [rack](/reference/cli-commands/rack)
