---
title: "rack resources url"
description: "Get the URL for a Rack Resource."
---

# rack resources url

Get the connection string for direct access to a Rack Resource. This is the same URL that is injected into linked Apps.

## Syntax

```bash
$ convox rack resources url <resource>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox rack resources url shared-postgres
postgres://user:pass@host.example.com:5432/app
```

## See Also

- [rack resources](/reference/cli-commands/rack-resources)
- [rack resources info](/reference/cli-commands/rack-resources-info)
- [rack resources proxy](/reference/cli-commands/rack-resources-proxy)
- [Rack and External Resources](/management/rack-resources)
