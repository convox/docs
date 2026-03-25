---
title: "rack resources proxy"
description: "Proxy a local port to a Rack Resource."
---

# rack resources proxy

Proxy a local port to a Rack Resource. This is useful for accessing Rack-level Resources from your local machine using database clients or other tools.

## Syntax

```bash
$ convox rack resources proxy <resource>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--port` | `-p` | Local port |
| `--rack` | `-r` | Rack name |
| `--tls` | `-t` | Wrap connection in TLS |

## Example Usage

```bash
$ convox rack resources proxy shared-postgres -p 5432
proxying localhost:5432 to host.example.com:5432
```

## See Also

- [rack resources](/reference/cli-commands/rack-resources)
- [rack resources url](/reference/cli-commands/rack-resources-url)
- [resources proxy](/reference/cli-commands/resources-proxy)
- [proxy](/reference/cli-commands/proxy)
- [Rack and External Resources](/management/rack-resources)
