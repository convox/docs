---
title: "NLBCrossZone"
description: "Enable cross-zone load balancing on the public Network Load Balancer."
---

# NLBCrossZone

Enable cross-zone load balancing on the public [NLB](/reference/rack-parameters/NLB). When `No`, each Availability Zone's listener routes only to targets in that AZ — the AWS default. When `Yes`, every listener routes to targets in every AZ, spreading traffic evenly at the cost of cross-AZ data transfer.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Use Cases

- Services with uneven target distribution across AZs where a single AZ would otherwise hotspot
- Workloads sensitive to latency variance between AZs where uniform spread improves p99
- Clients that resolve one of the NLB's Elastic IPs via DNS but expect traffic to fan out across the full target fleet

## Additional Information

```bash
$ convox rack params set NLBCrossZone=Yes
```

The update applies to every existing and future listener on the public NLB. No Service redeploy is required.

### Cost implications

Enabling cross-zone load balancing causes AWS to bill you for traffic crossing Availability Zone boundaries between the NLB and target tasks. Volume-heavy Services (video streaming, bulk file transfer) can see a meaningful cost increase. For targeted cross-zone distribution on a single Service instead of the whole fleet, use the per-port [cross_zone:](/application/services#nlb) override.

### Relationship to per-port `cross_zone`

A Service's `nlb:` entry with `cross_zone: true` forces cross-zone on that listener even when `NLBCrossZone=No`. `cross_zone: false` disables it on that listener even when `NLBCrossZone=Yes`. Omitting the field inherits the rack default.

## See Also

- [NLB](/reference/rack-parameters/NLB)
- [NLBInternalCrossZone](/reference/rack-parameters/NLBInternalCrossZone)
- [services.nlb field](/application/services#nlb)
- [Network Load Balancing](/networking/nlb#cross-zone-load-balancing)
