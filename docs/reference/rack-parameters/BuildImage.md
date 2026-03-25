---
title: "BuildImage"
description: "Override the default builder image used during Convox builds."
---

# BuildImage

Custom builder image for `convox build`. When blank, the default Convox builder image is used.

| Default value  | "" |

## Use Cases

- Using a custom builder image that includes additional build tools, language runtimes, or certificates
- Pinning to a specific builder image version for reproducible builds
- Testing a new builder image before rolling it out to all Racks

## Additional Information

The builder image must be accessible from the Rack's build instance. If you are using a private ECR repository, ensure the build instance IAM role has the necessary pull permissions.

```bash
$ convox rack params set BuildImage=myregistry/custom-builder:latest
```

## See Also

- [BuildInstance](/reference/rack-parameters/BuildInstance)
- [BuildMethod](/reference/rack-parameters/BuildMethod)
- [BuildCpu](/reference/rack-parameters/BuildCpu)
- [BuildMemory](/reference/rack-parameters/BuildMemory)
