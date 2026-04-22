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

## NLB deletion protection interlock

Uninstall is rejected pre-flight if either [NLBDeletionProtection](/reference/rack-parameters/NLBDeletionProtection) or [NLBInternalDeletionProtection](/reference/rack-parameters/NLBInternalDeletionProtection) is `Yes`:

```text
cannot uninstall rack while NLB deletion protection is enabled;
run 'convox rack params set NLBDeletionProtection=No NLBInternalDeletionProtection=No' first
(current: NLBDeletionProtection=Yes)
```

The interlock fires as soon as either flag is set, even if the rack has only a public NLB or only an internal NLB. `convox rack params set NLBDeletionProtection=No NLBInternalDeletionProtection=No` is safe on Racks where one of the two flags was never enabled — setting a parameter to its existing default is a no-op.

The interlock runs before any destructive output. Clear the protection flags, wait for the rack update to complete, then rerun `convox rack uninstall`.

## See Also

- [rack install](/reference/cli-commands/rack-install)
- [rack](/reference/cli-commands/rack)
- [NLBDeletionProtection](/reference/rack-parameters/NLBDeletionProtection)
