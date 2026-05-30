---
title: "BuildCache"
description: "Enable persistent Docker layer caching for Convox Rack builds using a dedicated ECR repository."
---

# BuildCache

Persistent build layer caching for the Rack. When enabled, App builds push their Docker layers to a dedicated `{rack}-build-cache` ECR repository and reuse those layers on later builds, which speeds up rebuilds when the relevant Dockerfile stages have not changed. The cache repository is separate from the per-App image registries that hold built Releases.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Use Cases

- Cutting rebuild times for Apps with expensive dependency-install or compile stages
- Reusing unchanged base and dependency layers across consecutive Builds and across Apps
- Keeping cache artifacts isolated in their own ECR repository, away from App Release images

## Additional Information

When set to `Yes`, Convox creates an `AWS::ECR::Repository` named `{rack}-build-cache` and points each Build at it. The build method determines how layers are cached:

- The `ec2` build method uses `docker buildx` with registry-backed caching (`--cache-from` and `--cache-to` of `type=registry`).
- The `fargate` build method uses kaniko with `--cache=true` and `--cache-repo` pointed at the same repository.

Enable the cache with a single parameter change:

```bash
$ convox rack params set BuildCache=Yes
```

The cache repository is provisioned only while `BuildCache=Yes`. Setting it back to `No` removes the repository and deletes every cache image it holds. Re-enabling later starts with a cold cache.

Old cache images are not pruned by this parameter. Control expiry of stale cache images separately with [BuildCacheCleanup](/reference/rack-parameters/BuildCacheCleanup) and [BuildCacheRetentionDays](/reference/rack-parameters/BuildCacheRetentionDays). With cleanup off, cache images are retained indefinitely.

### Gen1 limitation

Gen1 Apps do not use the build cache. The dedicated cache repository is wired only into App builds, so enabling `BuildCache=Yes` has no effect on gen1 Apps.

### Build method differences

`fargate` (kaniko) caching is less aggressive than `ec2` (buildkit) caching. Kaniko does not resume after the first cache miss in a build, does not cache `COPY` layers by default, and has limited multi-stage support. Builds that depend heavily on layer reuse cache more effectively on the `ec2` build method. See [BuildMethod](/reference/rack-parameters/BuildMethod) to choose between them.

### Buildx availability on EC2

On the `ec2` build method, caching requires the Docker buildx plugin. If buildx is not available on the build host, the Build proceeds without caching and prints a warning rather than failing.

## See Also

- [BuildCacheCleanup](/reference/rack-parameters/BuildCacheCleanup)
- [BuildCacheRetentionDays](/reference/rack-parameters/BuildCacheRetentionDays)
- [BuildMethod](/reference/rack-parameters/BuildMethod)
- [Builds](/deployment/builds)
- [Rack Parameters](/reference/rack-parameters)
