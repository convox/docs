---
title: "NLBInternalDeletionProtection"
description: "Block accidental deletion of the internal Network Load Balancer."
---

# NLBInternalDeletionProtection

Enable AWS deletion protection on the internal [NLBInternal](/reference/rack-parameters/NLBInternal). Same semantics as [NLBDeletionProtection](/reference/rack-parameters/NLBDeletionProtection) but scoped to the internal NLB.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Use Cases

- Production Racks where internal-service traffic would be disrupted if the internal NLB were deleted
- Racks where internal Services are consumed by other AWS accounts or peered VPCs and an accidental delete would require DNS-level coordination to recover
- Shared Racks with multiple operators who have `rack params set` permission

## Additional Information

```bash
$ convox rack params set NLBInternalDeletionProtection=Yes
```

### Interlocks

- `convox rack params set NLBInternal=No` while `NLBInternalDeletionProtection=Yes` — rejected pre-flight with:

  ```
  cannot disable NLBInternal while NLBInternalDeletionProtection=Yes;
  unset protection first, wait for the update to complete, then toggle
  NLBInternal off
  ```

- `convox rack uninstall` while either `NLBDeletionProtection=Yes` or `NLBInternalDeletionProtection=Yes` is enabled — rejected pre-flight. The uninstall interlock is global across both schemes.

See [NLBDeletionProtection](/reference/rack-parameters/NLBDeletionProtection#interlocks) for the rationale.

## See Also

- [NLBInternal](/reference/rack-parameters/NLBInternal)
- [NLBDeletionProtection](/reference/rack-parameters/NLBDeletionProtection)
- [rack uninstall](/reference/cli-commands/rack-uninstall)
- [Network Load Balancing](/networking/nlb#deletion-protection)
