---
title: "NLBDeletionProtection"
description: "Block accidental deletion of the public Network Load Balancer."
---

# NLBDeletionProtection

Enable AWS deletion protection on the public [NLB](/reference/rack-parameters/NLB). When `Yes`, any operation that would delete the load balancer is rejected pre-flight — including `convox rack params set NLB=No` and `convox rack uninstall`. Unset this before intentionally tearing down the NLB.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Use Cases

- Production Racks where accidental disable of the public NLB would take down customer-facing TCP listeners
- Racks with `rack params set` automation where an errant flag could flip `NLB=No` unintentionally
- Shared Racks where multiple operators have `params set` permission

## Additional Information

```bash
$ convox rack params set NLBDeletionProtection=Yes
```

The setting is applied to the underlying AWS load balancer via the `deletion_protection.enabled` attribute. No Service downtime occurs when toggling the protection state.

### Interlocks

Convox refuses to accept the destructive operations while protection is on, before CloudFormation touches any resource:

- `convox rack params set NLB=No` while `NLBDeletionProtection=Yes` — rejected with:

  ```
  cannot disable NLB while NLBDeletionProtection=Yes; unset protection
  first, wait for the update to complete, then toggle NLB off
  ```

- `convox rack uninstall` while either `NLBDeletionProtection=Yes` or `NLBInternalDeletionProtection=Yes` — rejected pre-flight with a parallel error.

The pre-flight block exists to avoid a known CloudFormation failure mode: the stack update starts, CloudFormation reaches the delete step, AWS refuses because of deletion protection, and the whole stack lands in `UPDATE_ROLLBACK_FAILED`. Unset protection first, wait for the update to complete, then run the disable command in a follow-up call.

### Scope

This parameter protects only the public NLB. Protect the internal NLB with [NLBInternalDeletionProtection](/reference/rack-parameters/NLBInternalDeletionProtection).

## See Also

- [NLB](/reference/rack-parameters/NLB)
- [NLBInternalDeletionProtection](/reference/rack-parameters/NLBInternalDeletionProtection)
- [rack uninstall](/reference/cli-commands/rack-uninstall)
- [Network Load Balancing](/networking/nlb#deletion-protection)
