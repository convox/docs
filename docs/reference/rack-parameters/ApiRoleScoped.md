---
title: "ApiRoleScoped"
description: "Replace the PowerUserAccess policy on the Convox Rack API role with a policy limited to the AWS services the Rack uses."
---

# ApiRoleScoped

Replace the AWS managed `PowerUserAccess` policy on the Rack's `ApiRole` with a Convox policy limited to the AWS services the Rack uses.

| Setting | Value |
|:--|:--|
| Default value | `No` |
| Allowed values | `Yes`, `No` |

## Use Cases

- Removing `PowerUserAccess` from the Rack's IAM role to meet an account policy on AWS managed policies
- Limiting the AWS services the Rack API, builds and the instance autoscaler can call
- Capping the Rack together with [PermissionsBoundary](/reference/rack-parameters/PermissionsBoundary)

## Additional Information

`ApiRole` is the IAM role the Rack API, the build tasks and the instance autoscaler run as, and the role CloudFormation uses for Rack, App and resource updates. By default it carries `PowerUserAccess`. With `ApiRoleScoped=Yes` the Rack attaches its own managed policy, `<rack>-ApiPolicyScoped-<suffix>`, instead. It allows every action of these services on all resources in the account:

| Service | IAM prefix |
|:--------|:-----------|
| ACM | `acm` |
| Application Auto Scaling | `application-autoscaling` |
| EC2 Auto Scaling | `autoscaling`, `managed-fleets` |
| CloudFormation | `cloudformation` |
| CloudWatch | `cloudwatch` |
| CloudWatch Logs | `logs` |
| DynamoDB | `dynamodb` |
| EC2 | `ec2` |
| ECR | `ecr` |
| ECS | `ecs` |
| EFS | `elasticfilesystem` |
| ElastiCache | `elasticache` |
| Elastic Load Balancing | `elasticloadbalancing` |
| EventBridge | `events` |
| KMS | `kms` |
| Lambda | `lambda` |
| RDS | `rds` |
| Route 53 | `route53` |
| S3 | `s3` |
| SNS | `sns` |
| SQS | `sqs` |
| Systems Manager | `ssm` |

It also allows `iam:CreateServiceLinkedRole`, `iam:ListRoles` and `sts:GetServiceBearerToken`. It does not allow `sts:AssumeRole`.

These stay the same with either value:

- `ApiPolicyV2`, the Rack's policy for the IAM roles, policies, instance profiles and users it manages. To cap it, set [PermissionsBoundary](/reference/rack-parameters/PermissionsBoundary) and have an account administrator set the same boundary on `ApiRole`.
- The inline policies on `ApiRole`, which also grant Secrets Manager access to the Rack's own secrets and ECR Public image pulls, outside the services above.

### Enabling

Update the Rack first, then set the parameter by itself, not together with other parameters or a version change:

```bash
$ convox rack update --wait
$ convox rack params set ApiRoleScoped=Yes --wait
```

Or set it at install:

```bash
$ convox rack install aws -n production ApiRoleScoped=Yes
```

To go back to `PowerUserAccess`:

```bash
$ convox rack params set ApiRoleScoped=No --wait
```

### Limits while enabled

- On a Rack with [BuildMethod](/reference/rack-parameters/BuildMethod)=`fargate`, Dockerfile steps that call AWS with the build task's credentials can use only the services above.
- A Rack or App change that needs an action outside the list fails on this Rack. Send the denied action from the stack events to Convox support so it can be added.

### If an update fails with AccessDenied

A stack update that needs an action outside the list fails, and its rollback can fail for the same reason, leaving the stack in `UPDATE_ROLLBACK_FAILED`. An account administrator restores it:

1. Attach `PowerUserAccess` to the Rack's `ApiRole`. In AWS GovCloud (US) the policy ARN is `arn:aws-us-gov:iam::aws:policy/PowerUserAccess`.

   ```bash
   $ aws cloudformation describe-stack-resource --stack-name <rack> --logical-resource-id ApiRole --query StackResourceDetail.PhysicalResourceId --output text
   <rack>-ApiRole-<suffix>
   $ aws iam attach-role-policy --role-name <rack>-ApiRole-<suffix> --policy-arn arn:aws:iam::aws:policy/PowerUserAccess
   ```

2. If a stack is in `UPDATE_ROLLBACK_FAILED`, continue its rollback, without `--role-arn`. Use the failed stack when it is top level: the Rack stack, an App stack (`<rack>-<app>`) or a Rack resource stack (`<rack>-<resource>`). For a Service, Timer or resource inside an App, use its App stack.

   ```bash
   $ aws cloudformation continue-update-rollback --stack-name <stack>
   ```

3. Turn the option off:

   ```bash
   $ convox rack params set ApiRoleScoped=No --wait
   ```

4. Send the denied action from the stack events to Convox support.

### Downgrading

Rack versions older than this parameter attach `PowerUserAccess` again when you update to them. After updating back to this version or newer, check `convox rack params` and set `ApiRoleScoped=Yes` again if it reads `No`.

## See Also

- [PermissionsBoundary](/reference/rack-parameters/PermissionsBoundary)
- [Security](/reference/security)
- [BuildMethod](/reference/rack-parameters/BuildMethod)
- [Rack Parameters](/reference/rack-parameters) for a full list of available parameters
