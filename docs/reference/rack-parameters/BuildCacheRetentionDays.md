---
title: "BuildCacheRetentionDays"
description: "Set how many days build cache images are retained in the Rack build cache repository before the lifecycle policy expires them."
---

# BuildCacheRetentionDays

Number of days to retain cache images in the `{rack}-build-cache` ECR repository before the lifecycle policy expires them. This value applies only when [BuildCacheCleanup](/reference/rack-parameters/BuildCacheCleanup) is `Yes`. With cleanup off there is no expiry policy and this value has no effect.

| Default value  | `30`         |
| Allowed values | `1` to `365` |

## Use Cases

- Capping storage cost in the build cache repository by expiring stale layers on a fixed schedule
- Keeping recent cache layers long enough to stay warm across frequent Builds while still trimming old ones
- Aligning cache retention with an organization's image retention or compliance window

## Additional Information

When [BuildCache](/reference/rack-parameters/BuildCache) is `Yes`, the Rack provisions a dedicated ECR repository named `{rack}-build-cache` that stores intermediate cache images shared across Builds. By default these images are kept indefinitely. Setting [BuildCacheCleanup](/reference/rack-parameters/BuildCacheCleanup) to `Yes` attaches an ECR lifecycle policy to that repository, and `BuildCacheRetentionDays` controls the age threshold the policy uses.

The lifecycle policy expires any cache image more than `BuildCacheRetentionDays` days past its push time, measured from when the image was last pushed. Lower the value to reclaim ECR storage sooner, or raise it to keep cache layers warm for longer between Builds.

```bash
$ convox rack params set BuildCacheRetentionDays=14
```

Setting this parameter while cleanup is off changes the stored value but has no operational effect until cleanup is enabled. To turn on expiry and set the window in one update, batch both parameters into a single command.

```bash
$ convox rack params set BuildCacheCleanup=Yes BuildCacheRetentionDays=14
```

### Prerequisites

The build cache repository only exists when [BuildCache](/reference/rack-parameters/BuildCache) is `Yes`. With the build cache disabled there is no repository, no lifecycle policy, and this parameter has no effect. Gen1 Apps do not use the build cache, so this parameter does not apply to them.

### Related Rack parameters

- [BuildCache](/reference/rack-parameters/BuildCache) creates the dedicated build cache repository
- [BuildCacheCleanup](/reference/rack-parameters/BuildCacheCleanup) attaches the lifecycle policy that expires cache images (this parameter has no effect unless that one is `Yes`)

## See Also

- [BuildCache](/reference/rack-parameters/BuildCache)
- [BuildCacheCleanup](/reference/rack-parameters/BuildCacheCleanup)
- [Builds](/deployment/builds)
- [Rack Parameters](/reference/rack-parameters)
