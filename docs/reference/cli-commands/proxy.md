---
title: "proxy"
description: "Proxy one or more local ports to hosts inside the Rack."
---

# proxy

Proxy one or more local ports to hosts inside the Rack. If no local port is specified, the host port is used. This is useful for accessing internal Services or infrastructure that are not publicly exposed.

## Syntax

```bash
$ convox proxy <[port:]host:hostport> [[port:]host:hostport]...
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |
| `--tls` | `-t` | Wrap connection in TLS |

## Example Usage

```bash
$ convox proxy 5432:internal-db.example.com:5432
proxying localhost:5432 to internal-db.example.com:5432
```

## See Also

- [resources proxy](/reference/cli-commands/resources-proxy)
- [rack resources proxy](/reference/cli-commands/rack-resources-proxy)
- [Load Balancing](/networking/load-balancing)
