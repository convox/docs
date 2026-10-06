---
title: "rack params set"
description: "Set one or more Rack parameters."
---

# rack params set

Set one or more Rack parameters. Changes may trigger a Rack update that modifies infrastructure. Multiple parameters can be set in a single command.

## Syntax

```bash
$ convox rack params set <Key=Value> [Key=Value]...
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

## Example Usage

```bash
$ convox rack params set InstanceType=t3.medium InstanceCount=5
Updating parameters... OK
```

When a single call transitions [NLB](/reference/rack-parameters/NLB) or [NLBInternal](/reference/rack-parameters/NLBInternal) from a non-`Yes` value to `Yes`, the command prints a provisioning-time hint before reporting `OK`:

```bash
$ convox rack params set NLB=Yes
Updating parameters... NLB provisioning typically takes 5-10 minutes; check status with 'convox rack'.
OK
```

The CloudFormation update continues asynchronously; use `convox rack` (or `--wait`) to confirm the NLB is reachable.

## Validation and interlocks

Several parameter combinations are rejected pre-flight (before CloudFormation touches any resource) so a misconfiguration cannot strand a stack in `UPDATE_ROLLBACK_FAILED`.

- **NLB allow-CIDR lists**: [NLBAllowCIDR](/reference/rack-parameters/NLBAllowCIDR) and [NLBInternalAllowCIDR](/reference/rack-parameters/NLBInternalAllowCIDR) accept up to 5 canonical IPv4 CIDRs. IPv6, non-canonical host-bit forms (`10.0.0.1/24`), duplicates, and surrounding whitespace are rejected with a specific error that names the offending entry.
- **NLB deletion protection vs disable**: setting `NLB=No` (or `NLBInternal=No`) while the corresponding `NLB*DeletionProtection` is `Yes` is rejected. Disable protection first, wait for the update to complete, then toggle the NLB off.
- **NLB preserve-client-IP vs custom InstanceSecurityGroup** (rack version 20261005214736 or newer): on a Rack with a custom [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup), setting [NLBPreserveClientIP](/reference/rack-parameters/NLBPreserveClientIP) or [NLBInternalPreserveClientIP](/reference/rack-parameters/NLBInternalPreserveClientIP) to `Yes` while that NLB is enabled is rejected unless the group has an ingress rule allowing all traffic from the NLB security group. Setting `InstanceSecurityGroup` to a new group while either parameter is `Yes`, or while a deployed App sets `preserve_client_ip: true`, needs the same rule on the new group, even when the same call sets the parameter to `No`. Enabling `NLB` or `NLBInternal` is rejected while that NLB's preserve parameter is `Yes`. See [NLBPreserveClientIP](/reference/rack-parameters/NLBPreserveClientIP#custom-instancesecuritygroup) for details.
- **PermissionsBoundary while ApiRole is bounded**: changing [PermissionsBoundary](/reference/rack-parameters/PermissionsBoundary) while the Rack's `ApiRole` carries a permissions boundary is rejected with `remove the permissions boundary from ApiRole first`. An account administrator removes it from `ApiRole` first. See [Disabling the boundary](/reference/rack-parameters/PermissionsBoundary#disabling-the-boundary).

The validator applies atomically; if any parameter in the call is invalid, none are applied.

## See Also

- [rack params](/reference/cli-commands/rack-params)
- [rack uninstall](/reference/cli-commands/rack-uninstall)
- [Rack Parameters](/reference/rack-parameters)
- [rack wait](/reference/cli-commands/rack-wait)
