---
title: "Service Tags"
description: "Enable ECS task tagging to track AWS costs by Rack, App, and Service."
---

# Service Tags

Service tags enable AWS resource tagging on your ECS tasks, allowing you to break down AWS billing data by Rack, App, and Service. When enabled, Convox propagates tags from ECS Services to all launched tasks, making them available in AWS Cost Explorer and billing reports.

## Prerequisites

Before enabling Service tags:

1. **Opt in to the new ARN and resource ID format** in your AWS account. Amazon requires this for ECS tag propagation. See the [AWS blog post on the new ARN format](https://aws.amazon.com/blogs/compute/migrating-your-amazon-ecs-deployment-to-the-new-arn-and-resource-id-format-2/) for instructions.

2. **Activate user-defined cost allocation tags** in the AWS Billing console. AWS does not include user-defined tags in billing reports until you explicitly activate them. See [AWS documentation on activating tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/activating-tags.html).

## Enabling Task Tags

Enable tag propagation on a per-App basis using the `TaskTags` App parameter:

```bash
$ convox apps params set TaskTags=Yes -a myapp
```

You can also set this in your `convox.yml` manifest:

```yaml
params:
  TaskTags: "Yes"
```

## What Gets Tagged

When `TaskTags` is enabled, Convox configures two ECS properties on each Service:

| Property | Value | Effect |
|----------|-------|--------|
| `EnableECSManagedTags` | `true` | AWS applies its own managed tags (cluster, service name) to tasks |
| `PropagateTags` | `SERVICE` | Tags on the ECS Service resource propagate to all tasks it launches |

This means every ECS task inherits the tags from its parent ECS Service, which include the App name, Service name, and other CloudFormation-generated tags.

## Viewing Tags in AWS

After enabling `TaskTags` and deploying, tags appear on ECS tasks within a few minutes. You can verify them in the AWS Console:

1. Navigate to **ECS** > **Clusters** > your cluster > **Tasks**
2. Select a running task
3. View the **Tags** tab

## Using Tags for Cost Allocation

Once tags are propagating and activated in AWS Billing:

1. Navigate to **AWS Cost Explorer**
2. Group costs by the relevant tag key (e.g., `aws:ecs:serviceName`, `App`)
3. Filter and break down costs by Rack, App, or Service

This allows you to attribute infrastructure costs to specific applications and services for chargeback or budgeting purposes.

## Limitations

- Task tags apply to Generation 2 Apps only. Generation 1 Apps do not support this feature.
- [Timers](/application/timers) do not support tag propagation. Tasks launched by timers will not inherit Service tags.
- Tag propagation applies to ECS tasks only. Other AWS resources (load balancers, S3 buckets) are tagged separately by CloudFormation.

## See Also

- [App Parameters](/reference/app-parameters)
- [Scaling](/scaling/scaling)
- [Services](/application/services)
