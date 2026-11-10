---
title: "BuildImage"
description: "Override the default builder image used during Convox builds."
---

# BuildImage

Custom builder image for `convox build`. When blank, the default Convox builder image for the Rack's [BuildMethod](/reference/rack-parameters/BuildMethod) is used:

| BuildMethod | Default builder image |
|:--|:--|
| `ec2` | `convox/rack:<version>`, or `convox/rack:<version>-arm64` on Racks with ARM instances |
| `fargate` | `convox/build:<version>`, or `convox/build:<version>-arm64` on Racks with ARM instances |

| Setting | Value |
|:--|:--|
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

### Custom Builder Images for EC2 Builds

The `convox/rack` image has no shell or package manager, so a custom builder image cannot add tools to it with `RUN`. Start from your own base and copy the Convox build binaries and the Docker CLI in:

```dockerfile
FROM ubuntu:24.04
RUN apt-get update && apt-get install -y ca-certificates <your tools>
COPY --from=convox/rack:<version> /go/bin/ /go/bin/
COPY --from=convox/rack:<version> /usr/bin/docker /usr/bin/docker
COPY --from=convox/rack:<version> /usr/local/lib/docker/cli-plugins/ /usr/local/lib/docker/cli-plugins/
ENV PATH=$PATH:/go/bin
```

Use the Rack's version for `<version>`, and the `-arm64` tag on Racks with ARM instances.

## See Also

- [BuildInstance](/reference/rack-parameters/BuildInstance)
- [BuildMethod](/reference/rack-parameters/BuildMethod)
- [BuildCpu](/reference/rack-parameters/BuildCpu)
- [BuildMemory](/reference/rack-parameters/BuildMemory)
