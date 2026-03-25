---
title: "InstancePolicy"
description: "Attach an additional IAM policy to Convox Rack cluster instances for custom AWS permissions."
---

# InstancePolicy

ARN of an additional IAM policy to add to the instance-level role. This allows your containers running on the Rack instances to access additional AWS services beyond the default permissions.

| Default value  | "" |

## Use Cases

- Granting containers access to specific S3 buckets, DynamoDB tables, or other AWS resources
- Attaching a policy that allows instances to publish to SNS topics or SQS queues
- Adding read access to AWS Secrets Manager or Systems Manager Parameter Store for application secrets

## Additional Information

The value should be a full IAM policy ARN, for example:

```bash
$ convox rack params set InstancePolicy=arn:aws:iam::123456789012:policy/my-custom-policy
```

This policy is attached to the IAM role used by all runtime instances in the cluster. If you need to apply a separate policy only to build instances, use [BuildInstancePolicy](/reference/rack-parameters/BuildInstancePolicy) instead.

## See Also

- [BuildInstancePolicy](/reference/rack-parameters/BuildInstancePolicy)
- [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup)
- [InstanceType](/reference/rack-parameters/InstanceType)
