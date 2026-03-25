---
title: "releases manifest"
description: "Display the convox.yml manifest for a Release."
---

# releases manifest

Display the convox.yml manifest for a Release. Useful for verifying what configuration was deployed with a specific Release or comparing manifest changes between Releases.

## Syntax

```bash
$ convox releases manifest <release>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox releases manifest RABCDEF -a myapp
environment:
  - PORT=3000
services:
  web:
    build: .
    port: 3000
    health: /health
    scale:
      count: 2
      memory: 512
  worker:
    build: ./worker
    command: bundle exec sidekiq
```

## See Also

- [releases info](/reference/cli-commands/releases-info)
- [releases](/reference/cli-commands/releases)
- [convox.yml Reference](/application/convox-yml)
- [Releases](/deployment/releases)
