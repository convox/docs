---
title: "BuildCacheCleanup"
description: "Automatically expire old build cache images in the dedicated ECR cache repository on a Convox Rack."
---

# BuildCacheCleanup

Controls whether the Rack attaches an ECR lifecycle policy that expires old images in the build cache repository. When enabled, cache images older than [BuildCacheRetentionDays](/reference/rack-parameters/BuildCacheRetentionDays) are expired automatically. This parameter has effect only when [BuildCache](/reference/rack-parameters/BuildCache) is set to `Yes`.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Use Cases

- Bounding the storage footprint (and ECR storage cost) of a long-lived build cache repository
- Pruning cache layers that are no longer referenced by recent Builds
- Keeping the cache repository small without manually deleting images

## Additional Information

When [BuildCache](/reference/rack-parameters/BuildCache) is set to `Yes`, the Rack provisions a dedicated ECR repository named `{rack}-build-cache` to hold persistent build cache images. `BuildCacheCleanup` governs the retention behavior of that repository.

When set to `Yes`, Convox attaches a CloudFormation-managed ECR lifecycle policy to the `{rack}-build-cache` repository. The policy expires any cache image older than [BuildCacheRetentionDays](/reference/rack-parameters/BuildCacheRetentionDays) (counted from the date each image was pushed).

When set to `No`, no expiry policy is attached and cache images are kept indefinitely. The cache repository grows over time. This is a storage-cost consideration, not data loss, because cache images are regenerable on the next Build.

```bash
$ convox rack params set BuildCacheCleanup=Yes
```

Pair this with [BuildCacheRetentionDays](/reference/rack-parameters/BuildCacheRetentionDays) to choose how long cache images are retained before they expire. The retention window defaults to 30 days and accepts values from 1 to 365.

### Prerequisites

`BuildCacheCleanup` has effect only when [BuildCache](/reference/rack-parameters/BuildCache) is `Yes`. With `BuildCache=No`, the `{rack}-build-cache` repository does not exist, so there is nothing for the cleanup policy to apply to. Enable the build cache first, then set the cleanup behavior:

```bash
$ convox rack params set BuildCache=Yes BuildCacheCleanup=Yes BuildCacheRetentionDays=30
```

Gen1 apps do not use the build cache, so this parameter has no effect on gen1 Builds.

### Disable

Set `BuildCacheCleanup=No` to stop expiring cache images. The lifecycle policy is removed from the `{rack}-build-cache` repository on the next Rack update, and existing cache images are retained from that point forward.

```bash
$ convox rack params set BuildCacheCleanup=No
```

## See Also

- [BuildCache](/reference/rack-parameters/BuildCache)
- [BuildCacheRetentionDays](/reference/rack-parameters/BuildCacheRetentionDays)
- [Builds](/deployment/builds)
- [Rack Parameters](/reference/rack-parameters)
