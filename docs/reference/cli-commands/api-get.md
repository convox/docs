---
title: "api get"
description: "Query the Rack API directly and return JSON output."
---

# api get

Query the Rack API at the specified path and return raw JSON output. This is useful for automation scripts, debugging, and accessing API endpoints not covered by other CLI commands.

## Syntax

```bash
$ convox api get <path>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox api get /apps
[
  {
    "name": "myapp",
    "generation": "2",
    "locked": false,
    "release": "RABCDEFGHIJ",
    "status": "running"
  }
]
```

## See Also

- [rack](/reference/cli-commands/rack)
- [version](/reference/cli-commands/version)
