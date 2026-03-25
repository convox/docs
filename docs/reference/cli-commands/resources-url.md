---
title: "resources url"
description: "Get the URL for an App Resource."
---

# resources url

Get the connection string URL used by the App to access a specific Resource. This is the same URL injected into the App's environment.

## Syntax

```bash
$ convox resources url <resource>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox resources url database -a myapp
postgres://user:pass@host.example.com:5432/app
```

## See Also

- [resources](/reference/cli-commands/resources)
- [resources info](/reference/cli-commands/resources-info)
- [resources proxy](/reference/cli-commands/resources-proxy)
- [rack resources url](/reference/cli-commands/rack-resources-url)
- [Accessing Resources](/management/resources)
