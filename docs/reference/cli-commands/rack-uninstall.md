---
title: "rack uninstall"
description: "Uninstall a Rack and remove all its infrastructure."
---

# rack uninstall

Uninstall a Rack. The CLI lists the CloudFormation stacks of the Rack and of every App and Rack Resource on it, asks for confirmation, and deletes them. The Rack's and each App's S3 buckets and the ECR repositories that hold App images stay in your AWS account, and PostgreSQL, MySQL and MariaDB databases, whether App resources or Rack Resources, leave a final RDS snapshot. The bucket of an `s3` Rack Resource is deleted with its stack, and that delete fails while the bucket still holds objects. Use `--force` to skip the confirmation prompt; it is required when stdin is not a terminal.

For a Rack installed with [rack install](/reference/cli-commands/rack-install), the CLI also removes it from [racks](/reference/cli-commands/racks) and clears it as the selected Rack. Select another Rack with [switch](/reference/cli-commands/switch) afterwards. The CLI does this before the confirmation prompt, so it also happens when you answer `N` or the uninstall is rejected, for example by the NLB deletion protection interlock below. `-r <name>` then goes to the host you logged in to with [login](/reference/cli-commands/login), where it can select a Console Rack whose name contains `<name>`. Run commands for the uninstalled Rack with `CONVOX_HOST` set to the API host that `rack install` printed after `Starting... OK`.

## Syntax

```bash
$ convox rack uninstall <type> <name>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--force` | `-f` | Force uninstall (required for non-interactive use) |

## Example Usage

```bash
$ convox rack uninstall aws production
The following stacks will be deleted:
  production-myapp
  production
Delete everything? [y/N]: y
disabled autoscaler rule: production-InstancesAutoscalerEvent-1A2B3C4D5E6F
production-BuildInstances-7G8H9I0J1K2L
production-Instances-3M4N5O6P7Q8R
Deleting stack: production-myapp
DELETE_IN_PROGRESS    production-myapp      AWS::CloudFormation::Stack
DELETE_IN_PROGRESS    ServiceWeb            AWS::CloudFormation::Stack
DELETE_SKIPPED        Registry              AWS::ECR::Repository
DELETE_SKIPPED        Settings              AWS::S3::Bucket
DELETE_COMPLETE       ServiceWeb            AWS::CloudFormation::Stack
Deleting stack: production
DELETE_IN_PROGRESS    production                           AWS::CloudFormation::Stack
DELETE_SKIPPED        Settings                             AWS::S3::Bucket
DELETE_COMPLETE       CustomerManagedKey                   AWS::KMS::Key
```

## NLB deletion protection interlock

Uninstall is rejected pre-flight if either [NLBDeletionProtection](/reference/rack-parameters/NLBDeletionProtection) or [NLBInternalDeletionProtection](/reference/rack-parameters/NLBInternalDeletionProtection) is `Yes`:

```text
cannot uninstall rack while NLB deletion protection is enabled; run 'convox rack params set NLBDeletionProtection=No NLBInternalDeletionProtection=No' first (current: NLBDeletionProtection=Yes)
```

The interlock fires as soon as either flag is set, even if the rack has only a public NLB or only an internal NLB. `convox rack params set NLBDeletionProtection=No NLBInternalDeletionProtection=No` is safe on Racks where one of the two flags was never enabled, because setting a parameter to its existing default is a no-op.

The interlock runs before any destructive output. Clear the protection flags, wait for the rack update to complete, then rerun `convox rack uninstall`.

## Custom InstanceSecurityGroup with an NLB

On a Rack with a custom [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup), remove any ingress rule on that group that references the Rack's NLB security group (`${Rack}:NLBSecurityGroup` or `${Rack}:NLBInternalSecurityGroup`) before uninstalling. EC2 does not delete a security group that another group's rule references, so the Rack stack delete ends in `DELETE_FAILED` while the rule exists.

## See Also

- [rack install](/reference/cli-commands/rack-install)
- [rack](/reference/cli-commands/rack)
- [NLBDeletionProtection](/reference/rack-parameters/NLBDeletionProtection)
