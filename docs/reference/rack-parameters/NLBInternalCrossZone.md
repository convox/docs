---
title: "NLBInternalCrossZone"
description: "Enable cross-zone load balancing on the internal Network Load Balancer."
---

# NLBInternalCrossZone

Enable cross-zone load balancing on the internal [NLBInternal](/reference/rack-parameters/NLBInternal). Same semantics as [NLBCrossZone](/reference/rack-parameters/NLBCrossZone) but scoped to the internal NLB.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Use Cases

- Internal Services with uneven target distribution across AZs
- In-VPC RPC or database-like workloads that benefit from uniform spread
- Latency-sensitive internal traffic where a hot AZ degrades p99

## Additional Information

```bash
$ convox rack params set NLBInternalCrossZone=Yes
```

Applies to every existing and future listener on the internal NLB. No Service redeploy is required.

Cost implications mirror [NLBCrossZone](/reference/rack-parameters/NLBCrossZone#cost-implications) — cross-AZ data transfer charges apply.

Per-port [cross_zone:](/application/services#nlb) on a Service with `scheme: internal` overrides this rack default for that listener only.

## See Also

- [NLBInternal](/reference/rack-parameters/NLBInternal)
- [NLBCrossZone](/reference/rack-parameters/NLBCrossZone)
- [services.nlb field](/application/services#nlb)
- [Network Load Balancing](/networking/nlb#cross-zone-load-balancing)
