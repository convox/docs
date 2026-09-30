---
title: "rack update"
description: "Update the Rack to the next version or a specific version."
---

# rack update

Update the Rack to the next version, or to a specific version if specified. Updates are applied as rolling updates to minimize downtime. If no version is provided, the Rack updates to the next required version after its current one, or to the latest published version when no later version is required.

## Syntax

```bash
$ convox rack update [version]
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

## Example Usage

```bash
$ convox rack update
Updating to 20260929232050... OK

$ convox rack update 20260826164715
Updating to 20260826164715... OK
```

With `--wait`, the command streams the Rack's update logs after `Updating to <version>...` and prints `OK` once the Rack returns to `running`.

A version change to a release that does not have the [PermissionsBoundary](/reference/rack-parameters/PermissionsBoundary) parameter is refused while that parameter is set:

```bash
$ convox rack update 20260826164715
Updating to 20260826164715... ERROR: clear PermissionsBoundary before moving to a version without it
```

See [Downgrading](/reference/rack-parameters/PermissionsBoundary#downgrading) for the steps before moving to an older version.

## See Also

- [rack releases](/reference/cli-commands/rack-releases)
- [rack wait](/reference/cli-commands/rack-wait)
- [rack](/reference/cli-commands/rack)
- [Rack Updates](/management/rack-updates)
