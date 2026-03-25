---
title: "resources proxy"
description: "Proxy a local port to an App Resource."
---

# resources proxy

Proxy a local port to an App Resource. This is useful for connecting local tools such as database clients, GUI managers, or scripts to a remote Resource running inside the Rack.

## Syntax

```bash
$ convox resources proxy <resource>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--port` | `-p` | Local port |
| `--rack` | `-r` | Rack name |
| `--tls` | `-t` | Wrap connection in TLS |

## Example Usage

```bash
$ convox resources proxy database -a myapp -p 5432
proxying localhost:5432 to host.example.com:5432
```

## See Also

- [resources](/reference/cli-commands/resources)
- [resources url](/reference/cli-commands/resources-url)
- [rack resources proxy](/reference/cli-commands/rack-resources-proxy)
- [proxy](/reference/cli-commands/proxy)
- [Accessing Resources](/management/resources)
