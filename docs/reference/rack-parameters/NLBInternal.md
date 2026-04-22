---
title: "NLBInternal"
description: "Enable an internal Network Load Balancer on a Convox Rack for TCP services that opt in."
---

# NLBInternal

Internal Network Load Balancer (NLB) for the Rack. When enabled, creates a VPC-internal Layer 4 load balancer that Services can opt into via the [nlb:](/application/services#nlb) field (`scheme: internal`) in `convox.yml`.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Use Cases

- Cross-service TCP traffic between Apps inside the same VPC (peered VPCs, VPN clients)
- Dedicated internal endpoints for databases, caches, or back-end Services that should never be reachable from the public internet
- Stable DNS for an internal listener consumed by other AWS accounts or peered VPCs

## Prerequisites

`NLBInternal=Yes` requires `Internal=Yes`. Set both in a single call:

```bash
$ convox rack params set Internal=Yes NLBInternal=Yes
```

Unlike the public [NLB](/reference/rack-parameters/NLB), the internal NLB does not consume any Elastic IPs — AWS assigns private IPs from the selected subnets.

## Additional Information

When set to `Yes`, Convox provisions:

- An `AWS::ElasticLoadBalancingV2::LoadBalancer` of type `network` with `Scheme: internal`
- A dedicated `NLBInternalSecurity` security group (exported as `${Rack}:NLBInternalSecurityGroup`)
- Ingress rules derived from [NLBInternalAllowCIDR](/reference/rack-parameters/NLBInternalAllowCIDR) (defaults to the Rack's [VPCCIDR](/reference/rack-parameters/VPCCIDR) when empty) plus any per-port [allow_cidr:](/application/services#nlb) entries declared by Services

Subnet placement depends on the Rack's `Private` parameter:

- `Private=Yes` → internal NLB lives in the Rack's private subnets (`SubnetPrivate0/1/2`)
- `Private=No` → internal NLB lives in the Rack's public subnets with `Scheme=internal` (matching the existing `RouterInternal` ALB placement pattern)

Services opt in per-port using `scheme: internal` under the `nlb:` field in `convox.yml`. Each declared port becomes a dedicated Listener + TargetGroup on the shared internal NLB.

The internal NLB's DNS name is visible via `convox rack` once CloudFormation completes.

### Related Rack parameters

- [NLBInternalAllowCIDR](/reference/rack-parameters/NLBInternalAllowCIDR) — CIDR allowlist (defaults to VPC CIDR when empty)
- [NLBInternalCrossZone](/reference/rack-parameters/NLBInternalCrossZone) — enable cross-zone load balancing on internal listeners
- [NLBInternalPreserveClientIP](/reference/rack-parameters/NLBInternalPreserveClientIP) — forward real client IP to target tasks
- [NLBInternalDeletionProtection](/reference/rack-parameters/NLBInternalDeletionProtection) — block accidental internal NLB deletion

### Disable

Before setting `NLBInternal=No`, remove any `nlb:` entries with `scheme: internal` from all deployed Apps and redeploy. The disable is refused with a list of blocking Apps otherwise.

If [NLBInternalDeletionProtection](/reference/rack-parameters/NLBInternalDeletionProtection)=`Yes`, `NLBInternal=No` is also rejected pre-flight; disable deletion protection first, wait for the update to complete, then toggle `NLBInternal=No` in a follow-up call.

### Known Limitations

Internal NLBs publish stable DNS names, which slightly lowers discovery cost for an in-VPC attacker pivoting between Apps — rely on application-layer authentication, not network obscurity. Tighten reachability with [NLBInternalAllowCIDR](/reference/rack-parameters/NLBInternalAllowCIDR) or per-port [allow_cidr:](/application/services#nlb) if the default VPC-wide posture is broader than needed.

See [Network Load Balancing](/networking/nlb#known-limitations) for the shared limitations list (50-listener quota, concurrent-deploy caveats, downgrade procedure).

## See Also

- [NLB](/reference/rack-parameters/NLB)
- [Internal](/reference/rack-parameters/Internal)
- [NLBInternalAllowCIDR](/reference/rack-parameters/NLBInternalAllowCIDR)
- [services.nlb field](/application/services#nlb)
