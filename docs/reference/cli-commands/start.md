---
title: "start"
description: "Start an App for local development."
---

# start

Start an App for local development. Builds and runs your App locally using Docker, with optional code sync for live reloading. This lets you develop and test against a local environment that mirrors your production Rack setup.

## Syntax

```bash
$ convox start [service] [service...]
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--generation` | `-g` | Generation |
| `--manifest` | `-m` | Manifest file |
| `--no-build` | | Skip build |
| `--no-cache` | | Build without layer cache |
| `--no-sync` | | Do not sync local changes into running containers |
| `--rack` | `-r` | Rack name |
| `--shift` | `-s` | Shift local port numbers (generation 1 only) |

## Example Usage

```bash
$ convox start
build   | building: web
build   | Step 1/4 : FROM ruby:3.2
build   | Step 2/4 : COPY . /app
build   | Step 3/4 : WORKDIR /app
build   | Step 4/4 : CMD ["bundle", "exec", "rails", "server"]
build   | Successfully built a1b2c3d4e5f6
build   | running: docker tag abc123 myapp/web
convox  | starting sync for web at .
web     | [2025-01-15 12:00:00] INFO  WEBrick::HTTPServer#start: pid=1 port=3000
web     | listening on 0.0.0.0:3000
```

## See Also

- [deploy](/reference/cli-commands/deploy)
- [build](/reference/cli-commands/build)
- [test](/reference/cli-commands/test)
- [Running Locally](/development/running-locally)
