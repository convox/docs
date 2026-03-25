---
title: "resources info"
description: "Get information about an App Resource."
---

# resources info

Display detailed information about a specific App Resource, including its name, type, and connection URL.

## Syntax

```bash
$ convox resources info <resource>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox resources info database -a myapp
Name  database
Type  postgres
URL   postgres://user:pass@host.example.com:5432/app
```

## See Also

- [resources](/reference/cli-commands/resources)
- [resources url](/reference/cli-commands/resources-url)
- [resources proxy](/reference/cli-commands/resources-proxy)
- [rack resources info](/reference/cli-commands/rack-resources-info)
- [Accessing Resources](/management/resources)
