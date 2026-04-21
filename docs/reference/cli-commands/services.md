---
title: "services"
description: "List Services for an App, showing name, domain, and ports."
---

# services

List all Services for an App. The output includes each Service's name, its assigned domain, and the ports it exposes. This is useful for verifying which endpoints are available after a deployment.

When any Service in the App declares [nlb:](/application/services#nlb) entries, an additional `NLB PORTS` column is shown. Each cell is `PORT:CONTAINERPORT`, suffixed with `/tls` for TLS-terminating listeners and `(internal)` for listeners on the internal NLB.

## Syntax

```bash
$ convox services
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox services -a myapp
SERVICE  DOMAIN                               PORTS
web      web.myapp.0a1b2c3d4e.convox.cloud    443:3000
api      api.myapp.0a1b2c3d4e.convox.cloud    443:4000
worker
```

With an App that declares `nlb:` ports:

```bash
$ convox services -a broker
SERVICE  DOMAIN                               PORTS      NLB PORTS
api      api.broker.0a1b2c3d4e.convox.cloud   443:3000   443:3000/tls 50051:50051(internal)
worker
```

## See Also

- [services restart](/reference/cli-commands/services-restart)
- [scale](/reference/cli-commands/scale)
- [ps](/reference/cli-commands/ps)
- [Services](/application/services)
- [Network Load Balancing](/networking/nlb)
