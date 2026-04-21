---
title: "NLBInternal"
description: "Enable an internal Network Load Balancer on a Convox Rack for TCP services that opt in."
---

# NLBInternal

Internal Network Load Balancer (NLB) for the Rack. When enabled, creates a VPC-internal Layer 4 load balancer that services can opt into via the [nlb:](/application/services#nlb) field (`scheme: internal`) in `convox.yml`.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Use Cases

- Cross-service TCP traffic between apps inside the same VPC (peered VPCs, VPN clients)
- Dedicated internal endpoints for databases, caches, or back-end services that should never be reachable from the public internet
- Internal-only equivalents of public TCP services

## Prerequisites

`NLBInternal=Yes` requires `Internal=Yes`. Set both in a single call:

```bash
$ convox rack params set Internal=Yes NLBInternal=Yes
```

Unlike the public [NLB](/reference/rack-parameters/NLB), the internal NLB does not consume any Elastic IPs — AWS assigns private IPs from the selected subnets.

## Additional Information

When set to `Yes`, Convox provisions an `AWS::ElasticLoadBalancingV2::LoadBalancer` of type `network` with `Scheme: internal`. Subnet placement depends on the Rack's `Private` parameter:

- `Private=Yes` → internal NLB lives in the Rack's private subnets (`SubnetPrivate0/1/2`)
- `Private=No` → internal NLB lives in the Rack's public subnets with `Scheme=internal` (matching the existing `RouterInternal` ALB placement pattern)

Services opt in per-port using `scheme: internal` under the `nlb:` field in `convox.yml`. Each declared port becomes a dedicated Listener + TargetGroup on the shared internal NLB.

The internal NLB's DNS name is visible via `convox rack` once CloudFormation completes.

### Disable

Before setting `NLBInternal=No`, remove any `nlb:` entries with `scheme: internal` from all deployed apps and redeploy. The disable will be refused with a list of blocking apps otherwise.

### Known Limitations (v1)

Internal NLBs do not add exposure beyond what the Rack's `InstancesSecurity` security group already permits across the VPC CIDR. They do publish stable DNS names, which slightly lowers discovery cost for an in-VPC attacker pivoting between apps — rely on application-layer authentication, not network obscurity.

See [Network Load Balancing](/networking/nlb#known-limitations) for the shared limitations (no real client IP preservation, cross-zone load balancing off by default, 50-listener quota).

## See Also

- [NLB](/reference/rack-parameters/NLB)
- [Internal](/reference/rack-parameters/Internal)
- [InternalOnly](/reference/rack-parameters/InternalOnly)
- [services.nlb field](/application/services#nlb)
