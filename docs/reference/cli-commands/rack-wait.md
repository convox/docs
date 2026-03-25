---
title: "rack wait"
description: "Wait for the Rack to finish updating."
---

# rack wait

Wait for the Rack to finish updating. Blocks until the Rack reaches a stable state, making it useful in CI/CD pipelines or scripts that depend on the Rack being ready.

## Syntax

```bash
$ convox rack wait
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox rack wait
Waiting for rack... OK
```

## See Also

- [rack update](/reference/cli-commands/rack-update)
- [rack](/reference/cli-commands/rack)
- [apps wait](/reference/cli-commands/apps-wait)
- [Rack Statuses](/reference/rack-statuses)
