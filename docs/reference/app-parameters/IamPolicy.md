---
title: "IamPolicy"
description: "Attach a custom IAM policy to the application's ECS Task Role for fine-grained AWS permissions."
---

# IamPolicy

ARN of a custom IAM policy to attach to the application's ECS [Task Role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html).
If the service has the [Policies](/application/services) parameter set, this will not apply at the service level.

| Default value  | "" |

## Use Cases

- Attach a custom IAM policy when your application needs to access AWS services such as S3, SQS, SNS, or DynamoDB
- Use when you want a single IAM policy shared across all services in the application
- Use as a simpler alternative to defining per-service policies in `convox.yml` when all services need the same permissions

## Additional Information

The value must be a valid IAM policy ARN (e.g., `arn:aws:iam::123456789012:policy/my-custom-policy`). The policy is attached as a managed policy to the ECS Task Role that Convox creates for the application.

If a service defines its own `Policies` attribute in `convox.yml`, the service-level policies take precedence and this app-level `IamPolicy` will **not** be applied to that service. Services without explicit policies will use the app-level `IamPolicy`.

```bash
$ convox apps params set IamPolicy=arn:aws:iam::123456789012:policy/my-app-policy
```

## See Also

- [Services](/application/services)
- [Rack Parameters](/reference/rack-parameters)
