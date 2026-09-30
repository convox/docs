---
title: "PermissionsBoundary"
description: "Set an IAM permissions boundary on every IAM role and user a Convox Rack creates."
---

# PermissionsBoundary

ARN of an IAM managed policy that the Rack sets as the [permissions boundary](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) on every IAM role and user it creates, except the Rack's own `ApiRole`. With the same boundary also set on `ApiRole` by an account administrator, no role or user the Rack creates or changes can hold more than the boundary policy allows.

| Setting | Value |
|:--|:--|
| Default value | "" |
| Allowed values | blank, or an IAM managed policy ARN (`arn:aws:iam::<account>:policy/<path><name>`) |

## Use Cases

- Capping the permissions of every role and user a Rack creates, including App roles that get grants from [IamPolicy](/reference/app-parameters/IamPolicy) or a Service's [`policies`](/application/services#policies)
- Meeting an account policy that requires a permissions boundary on every IAM role
- Limiting the Rack's `ApiRole` together with [ApiRoleScoped](/reference/rack-parameters/ApiRoleScoped)

## Additional Information

The Rack's IAM permissions for managing roles, policies, instance profiles and users apply only to entities under the `/convox/` path in the Rack's own AWS account, which is where the Rack creates all of them. `iam:GetRole` and `iam:PassRole` cover every role in the account, so CloudFormation can create new roles and a user-supplied role such as a Generation 1 `TaskRole` keeps working. This parameter adds a boundary on top of that scope.

### Where the boundary applies

| Roles and users | Receive the boundary |
|:----------------|:---------------------|
| Rack roles: `InstancesRole`, `BuildInstancesRole`, `InstancesLifecycleRole`, `InstancesLifecycleHandlerRole`, `SpotFleetRole`, `ServiceRole`, `CustomTopicRole`, `SecureEnvironmentRole` | when the parameter changes |
| App roles, Generation 1 and Generation 2, including Service and Timer roles | on the App's next deploy or promote. A new App gets it when it is created |
| Rack resource roles and users | on the next `convox rack resources update <name>` |
| `ApiRole` | never. An account administrator sets it, see step 6 below |

### Accepted values

The value must be the ARN of an IAM managed policy. Policy paths and names can use letters, digits and `_.,=@-`. Any other value is refused by CloudFormation before the update starts:

```bash
$ convox rack params set PermissionsBoundary=not-an-arn
Updating parameters... ERROR: ValidationError: Parameter 'PermissionsBoundary' must match pattern ^$|^arn:aws[a-zA-Z-]*:iam::[0-9]{12}:policy/[A-Za-z0-9_.,=@/-]+$
```

Create the boundary policy at a path outside `/convox/`, so the Rack cannot change it. A starting point is the [reference boundary policy](#reference-boundary-policy) below.

### Enabling the boundary

1. Update the Rack to a version that has this parameter (`convox rack update`).
2. An account administrator creates the boundary policy in IAM, at a path outside `/convox/`.
3. Set the parameter:

   ```bash
   $ convox rack params set PermissionsBoundary=arn:aws:iam::123456789012:policy/convox-boundary/rack-boundary --wait
   ```

4. Deploy every App, and update every Rack resource, so their roles and users receive the boundary:

   ```bash
   $ convox deploy -a myapp
   $ convox rack resources update my-syslog
   ```

5. Check that every role and user under `/convox/` in the account carries the boundary. In an account with more than one Rack, complete steps 3 and 4 on every Rack first. This lists the ones that do not. At this point it should return only the `ApiRole` of each Rack, plus any roles a v3 Rack or the Console's AWS integration created under `/convox/`, which no V2 Rack parameter can bound:

   ```bash
   $ aws iam get-account-authorization-details --filter Role User --query "[RoleDetailList[?Path=='/convox/' && PermissionsBoundary==null].RoleName, UserDetailList[?Path=='/convox/' && PermissionsBoundary==null].UserName][]"
   ```

6. After step 4 has finished for every App and Rack resource, the administrator sets the same boundary on the Rack's `ApiRole`. While `ApiRole` carries the boundary, a failed deploy of an App whose roles do not carry it yet cannot roll back, and the App stack stops in `UPDATE_ROLLBACK_FAILED` (see [Recovering a stuck rollback](#recovering-a-stuck-rollback)).

   ```bash
   $ aws cloudformation describe-stack-resource --stack-name <rack> --logical-resource-id ApiRole --query StackResourceDetail.PhysicalResourceId --output text
   <rack>-ApiRole-<suffix>
   $ aws iam put-role-permissions-boundary --role-name <rack>-ApiRole-<suffix> --permissions-boundary arn:aws:iam::123456789012:policy/convox-boundary/rack-boundary
   ```

### Disabling the boundary

1. The administrator removes the boundary from `ApiRole`:

   ```bash
   $ aws iam delete-role-permissions-boundary --role-name <rack>-ApiRole-<suffix>
   ```

2. Clear the parameter:

   ```bash
   $ convox rack params set PermissionsBoundary= --wait
   ```

3. Deploy every App and update every Rack resource, which removes the boundary from their roles and users. After this, `convox releases promote` of a Generation 1 Release that was first promoted while the parameter was set puts the boundary back on that App's roles. Use `convox releases rollback` or `convox deploy` instead.

### Refused updates

The Rack API refuses these updates before any change starts:

| Update | Refused when | Error |
|:-------|:-------------|:------|
| Any change to `PermissionsBoundary` | `ApiRole` has a permissions boundary | `remove the permissions boundary from ApiRole first` |
| A version change and a `PermissionsBoundary` change in one update | both change | `set PermissionsBoundary in a separate update` |
| A version change to a release without this parameter | `PermissionsBoundary` is set | `clear PermissionsBoundary before moving to a version without it` |

```bash
$ convox rack params set PermissionsBoundary= --wait
Updating parameters... ERROR: remove the permissions boundary from ApiRole first
```

### Limits while enabled

- App roles are capped by the boundary. Grants from [IamPolicy](/reference/app-parameters/IamPolicy) or a Service's [`policies`](/application/services#policies) beyond what the boundary allows stop applying after the App's next deploy.
- `convox releases promote` of a Generation 1 Release that was first promoted before the parameter was set removes the boundary from that App's roles, and fails once `ApiRole` carries the boundary. `convox releases rollback` and `convox deploy` create a new Release, which gets the boundary.
- A Generation 1 [`TaskRole`](/gen1/app-parameters#taskrole) outside `/convox/` must be allowed by an `iam:PassRole` statement in the boundary policy, or the App's next deploy fails. The reference policy below allows passing `/convox/` roles only.
- CloudTrail records an `AccessDenied` for `iam:GetPolicy` on the boundary policy each time CloudFormation creates or updates a Rack or App role. The update still completes.

### Recovering a stuck rollback

An IAM administrator runs `continue-update-rollback` on the App stack with their own credentials:

```bash
$ aws cloudformation continue-update-rollback --stack-name <rack>-<app>
```

The Rack's stacks have no CloudFormation service role, so CloudFormation finishes the rollback with the administrator's permissions and the boundary can stay on `ApiRole`. Run it on the App stack, not on a nested Service stack. `convox apps` shows the App as `failed` until then.

### Downgrading

Rack versions older than this parameter restore the previous IAM policy and cannot remove a boundary left on a role. The Rack refuses the version change while the parameter is set:

```bash
$ convox rack update 20260826164715
Updating to 20260826164715... ERROR: clear PermissionsBoundary before moving to a version without it
```

Before moving to an older version, follow [Disabling the boundary](#disabling-the-boundary) in full, including the deploy of every App and the update of every Rack resource. The Rack checks only the parameter, not the roles.

## Reference Boundary Policy

This policy allows what a Rack and the roles it creates need. Outside `/convox/`, its only IAM writes are service-linked role creation and IAM server certificates. It allows no `sts:AssumeRole*`, AWS Organizations or account management action, so App grants of those actions from IamPolicy or `policies` stop applying under it. Replace `123456789012` with the account ID and `convox-boundary/rack-boundary` with the policy's own path and name; the `iam:PermissionsBoundary` condition must name this policy's ARN. In AWS GovCloud (US) the ARNs start with `arn:aws-us-gov:`.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "Services",
      "Effect": "Allow",
      "NotAction": ["iam:*", "organizations:*", "account:*", "sts:AssumeRole*"],
      "Resource": "*"
    },
    {
      "Sid": "ConvoxIamWithBoundary",
      "Effect": "Allow",
      "Action": [
        "iam:AttachRolePolicy",
        "iam:AttachUserPolicy",
        "iam:CreateRole",
        "iam:CreateUser",
        "iam:PutRolePermissionsBoundary",
        "iam:PutRolePolicy",
        "iam:PutUserPermissionsBoundary",
        "iam:PutUserPolicy",
        "iam:UpdateAssumeRolePolicy"
      ],
      "Resource": [
        "arn:aws:iam::123456789012:role/convox/*",
        "arn:aws:iam::123456789012:user/convox/*"
      ],
      "Condition": {
        "StringEquals": { "iam:PermissionsBoundary": "arn:aws:iam::123456789012:policy/convox-boundary/rack-boundary" }
      }
    },
    {
      "Sid": "ConvoxIam",
      "Effect": "Allow",
      "Action": [
        "iam:AddRoleToInstanceProfile",
        "iam:CreateAccessKey",
        "iam:CreateInstanceProfile",
        "iam:CreatePolicy",
        "iam:CreatePolicyVersion",
        "iam:DeleteAccessKey",
        "iam:DeleteInstanceProfile",
        "iam:DeletePolicy",
        "iam:DeletePolicyVersion",
        "iam:DeleteRole",
        "iam:DeleteRolePolicy",
        "iam:DeleteUser",
        "iam:DeleteUserPolicy",
        "iam:DetachRolePolicy",
        "iam:DetachUserPolicy",
        "iam:RemoveRoleFromInstanceProfile",
        "iam:SetDefaultPolicyVersion",
        "iam:TagInstanceProfile",
        "iam:TagPolicy",
        "iam:TagRole",
        "iam:TagUser",
        "iam:UntagInstanceProfile",
        "iam:UntagPolicy",
        "iam:UntagRole",
        "iam:UntagUser"
      ],
      "Resource": [
        "arn:aws:iam::123456789012:instance-profile/convox/*",
        "arn:aws:iam::123456789012:policy/convox/*",
        "arn:aws:iam::123456789012:role/convox/*",
        "arn:aws:iam::123456789012:user/convox/*"
      ]
    },
    {
      "Sid": "IamRead",
      "Effect": "Allow",
      "Action": ["iam:Get*", "iam:List*"],
      "Resource": "*"
    },
    {
      "Sid": "PassConvoxRoles",
      "Effect": "Allow",
      "Action": "iam:PassRole",
      "Resource": "arn:aws:iam::123456789012:role/convox/*",
      "Condition": {
        "StringEquals": {
          "iam:PassedToService": [
            "application-autoscaling.amazonaws.com",
            "autoscaling.amazonaws.com",
            "ec2.amazonaws.com",
            "ecs.amazonaws.com",
            "ecs-tasks.amazonaws.com",
            "events.amazonaws.com",
            "lambda.amazonaws.com",
            "spotfleet.amazonaws.com"
          ]
        }
      }
    },
    {
      "Sid": "ServiceLinkedRolesAndCertificates",
      "Effect": "Allow",
      "Action": [
        "iam:CreateServiceLinkedRole",
        "iam:DeleteServerCertificate",
        "iam:UploadServerCertificate"
      ],
      "Resource": "*"
    }
  ]
}
```

## See Also

- [ApiRoleScoped](/reference/rack-parameters/ApiRoleScoped)
- [Security](/reference/security)
- [IamPolicy](/reference/app-parameters/IamPolicy)
- [Rack Parameters](/reference/rack-parameters) for a full list of available parameters
