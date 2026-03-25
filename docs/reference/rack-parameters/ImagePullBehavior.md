---
title: "ImagePullBehavior"
description: "Control the Docker image pull strategy on Convox Rack ECS container instances."
---

# ImagePullBehavior

The behavior used to customize the Docker image pull process on container instances. See [ECS_IMAGE_PULL_BEHAVIOR](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-config.html) in the AWS docs.

| Default value  | `default`                                    |
| Allowed values | `default`, `always`, `once`, `prefer-cached` |

## Use Cases

- Setting to `always` to guarantee that the latest image version is pulled on every task launch
- Setting to `prefer-cached` to speed up container startup times by reusing locally cached images
- Setting to `once` to pull an image only once per instance, reducing ECR bandwidth and API calls

## Additional Information

- `default` -- The ECS agent uses Docker's default image pull behavior, which typically pulls the image if it is not already cached locally.
- `always` -- The image is always pulled from the registry, even if a cached version exists. This is useful when you want to ensure the latest image tag is always used.
- `once` -- The image is only pulled once per instance. Subsequent tasks reuse the cached image, even if the tag has been updated in the registry.
- `prefer-cached` -- The cached image is used if available. If no cached image exists, the image is pulled from the registry. This is the fastest option for startup time.

```bash
$ convox rack params set ImagePullBehavior=always
```

## See Also

- [InstanceType](/reference/rack-parameters/InstanceType)
- [InstanceCount](/reference/rack-parameters/InstanceCount)
