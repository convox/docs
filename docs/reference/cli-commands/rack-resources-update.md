---
title: "rack resources update"
description: "Update options for a Rack Resource."
---

# rack resources update

Update the configurable options for an existing Rack Resource. Depending on the option changed, the Resource may be recreated, which can cause brief downtime and a new connection URL.

Options you leave out keep their current values. Option names must match the names that [rack resources options](/reference/cli-commands/rack-resources-options) lists, including case: an option the Resource type does not declare is ignored, so `url=` does not change a webhook's `Url`.

## Syntax

```bash
$ convox rack resources update <name> [Option=Value]...
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

## Example Usage

```bash
$ convox rack resources update shared-postgres AllocatedStorage=50 --wait
Updating resource... OK
```

To point a webhook at a new URL (requires rack version 20261005214736 or newer):

```bash
$ convox rack resources update my-hook Url=https://hooks.example.com/new --wait
Updating resource... OK
```

An update that changes nothing prints `ERROR: ValidationError: No updates are to be performed.` and leaves the Resource as it is. See [webhook](/management/rack-resources#webhook) for the `Url` rules on webhooks.

## See Also

- [rack resources info](/reference/cli-commands/rack-resources-info)
- [rack resources options](/reference/cli-commands/rack-resources-options)
- [rack resources](/reference/cli-commands/rack-resources)
- [Rack and External Resources](/management/rack-resources)
