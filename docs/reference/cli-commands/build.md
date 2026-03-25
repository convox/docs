---
title: "build"
description: "Create a Build from the specified directory without promoting it."
---

# build

Create a Build from the specified directory (default: `.`) without promoting the resulting Release. This is useful when you want to build and test before promoting. Use `convox deploy` to build and promote in one step.

## Syntax

```bash
$ convox build [dir]
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
| `--wildcard-domain` | | Enable wildcard domain |

## Example Usage

```bash
$ convox build -a myapp -d "feature branch build"
Packaging source... OK
Uploading source... OK
Starting build... OK
Running: docker build -t 9836064b .
Sending build context to Docker daemon  4.096kB
Step 1/4 : FROM node:18-alpine
 ---> 2a5b3c4d5e6f
Step 2/4 : WORKDIR /app
 ---> Using cache
Step 3/4 : COPY . .
 ---> 7f8a9b0c1d2e
Step 4/4 : CMD ["node", "server.js"]
 ---> 3e4f5a6b7c8d
Successfully built 3e4f5a6b7c8d
Build:   BABCDEF
Release: RABCDEF
```

## See Also

- [deploy](/reference/cli-commands/deploy)
- [builds](/reference/cli-commands/builds)
- [builds info](/reference/cli-commands/builds-info)
- [Builds](/deployment/builds)
