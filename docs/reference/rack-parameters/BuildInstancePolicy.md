---
title: "BuildInstancePolicy"
description: "Attach an additional IAM policy to the Convox Rack build instances."
---

# BuildInstancePolicy

ARN of an additional IAM policy to add to the build cluster instances. This is similar to [InstancePolicy](/reference/rack-parameters/InstancePolicy) but applies only to build instances.

| Default value  | "" |

## Use Cases

- Granting build instances access to private ECR registries in other AWS accounts for pulling base images
- Allowing build instances to read secrets from AWS Secrets Manager or SSM Parameter Store during the build process
- Providing access to S3 buckets that contain build artifacts or dependencies

## Additional Information

The policy ARN must point to a valid IAM policy in the same AWS account. The policy is attached in addition to the default permissions that Convox assigns to build instances.

```bash
$ convox rack params set BuildInstancePolicy=arn:aws:iam::123456789012:policy/custom-build-policy
```

## See Also

- [InstancePolicy](/reference/rack-parameters/InstancePolicy)
- [BuildInstance](/reference/rack-parameters/BuildInstance)
- [BuildInstanceSecurityGroup](/reference/rack-parameters/BuildInstanceSecurityGroup)
