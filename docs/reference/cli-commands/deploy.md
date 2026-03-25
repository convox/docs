---
title: "deploy"
description: "Create a Build from a directory and promote the resulting Release."
---

# deploy

Create a Build from the specified directory (default: `.`) and promote the resulting Release. This is the most common way to deploy code changes, combining `build` and `releases promote` into a single command. The deploy process packages your source, uploads it to the Rack, builds a Docker image, and promotes the new Release.

## Syntax

```bash
$ convox deploy [dir]
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--build-args` | | Docker build arguments (can be specified multiple times) |
| `--description` | `-d` | Build description |
| `--development` | | Create a development Build |
| `--id` | | Send logs to stderr, Release ID to stdout |
| `--manifest` | `-m` | Path to manifest file |
| `--no-cache` | | Build without Docker cache |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |
| `--wildcard-domain` | | Enable wildcard domain |

## Example Usage

```bash
$ convox deploy -a myapp
Packaging source... OK
Uploading source... OK
Starting build... OK
Authenticating 123456789012.dkr.ecr.us-east-1.amazonaws.com: Login Succeeded
Running: docker build -t 9836064b .
Sending build context to Docker daemon  4.096kB
Step 1/4 : FROM ruby:3.2
Step 2/4 : COPY . /app
Step 3/4 : WORKDIR /app
Step 4/4 : CMD ["bundle", "exec", "rails", "server"]
Successfully built a1b2c3d4e5f6
Running: docker tag 9836064b convox/myapp:web.BABCDEF
Running: docker push convox/myapp:web.BABCDEF
Build:   BABCDEF
Release: RABCDEF
Promoting RABCDEF...
2025-01-15T12:00:00Z system/k8s/web Scaled up replica set web-abc1234 to 1
2025-01-15T12:00:01Z system/k8s/web Started container main
OK
```

## See Also

- [build](/reference/cli-commands/build)
- [apps create](/reference/cli-commands/apps-create)
- [releases promote](/reference/cli-commands/releases-promote)
- [releases rollback](/reference/cli-commands/releases-rollback)
- [Builds](/deployment/builds)
- [Releases](/deployment/releases)
- [Rolling Updates](/deployment/rolling-updates)
