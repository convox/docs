---
title: "NLB"
description: "Enable a public Network Load Balancer on a Convox Rack for TCP services that opt in."
---

# NLB

Public Network Load Balancer (NLB) for the Rack. When enabled, creates an internet-facing Layer 4 load balancer that Services can opt into via the [nlb:](/application/services#nlb) field in `convox.yml`. The existing Application Load Balancer (ALB) is unaffected — the NLB is a sidecar for TCP workloads that the ALB cannot handle.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Use Cases

- Exposing non-HTTP TCP protocols (MQTT, Redis, raw TCP gRPC, game servers)
- Portable static IPs via AWS Elastic IPs allocated per availability zone
- Low-overhead Layer 4 routing with source IP preservation at the VPC boundary

## Prerequisites

Enabling `NLB=Yes` consumes 2 Elastic IPs on a 2-AZ Rack or 3 EIPs on an HA 3-AZ Rack. `Private=Yes` Racks already consume 2-3 EIPs for NAT gateways. Before enabling NLB on a `Private=Yes` HA 3-AZ Rack (total 6 EIPs), verify your account has quota headroom:

```bash
$ aws service-quotas get-service-quota \
    --service-code ec2 --quota-code L-0263D0A3
```

The default regional quota is 5 EIPs. Request an increase through AWS Service Quotas if needed before enabling.

## Additional Information

When set to `Yes`, Convox provisions:

- An `AWS::ElasticLoadBalancingV2::LoadBalancer` of type `network` with `SubnetMappings` that attach the EIPs to the Rack's public subnets
- A dedicated `NLBSecurity` security group (exported as `${Rack}:NLBSecurityGroup`) that gates all inbound traffic to public NLB listeners
- Ingress rules derived from [NLBAllowCIDR](/reference/rack-parameters/NLBAllowCIDR) (defaults to `0.0.0.0/0`) plus any per-port [allow_cidr:](/application/services#nlb) entries declared by Services

Services opt in per-port using the [nlb:](/application/services#nlb) field in `convox.yml`. Each declared port becomes a dedicated `AWS::ElasticLoadBalancingV2::Listener` + `TargetGroup` on the shared Rack NLB.

```bash
$ convox rack params set NLB=Yes
```

NLB provisioning typically takes 5-10 minutes. Check status with `convox rack`; the output will show the NLB DNS name and allocated EIPs once CloudFormation completes.

This parameter cannot be enabled on an [InternalOnly](/reference/rack-parameters/InternalOnly) Rack. Use [NLBInternal](/reference/rack-parameters/NLBInternal) instead.

### Related Rack parameters

- [NLBAllowCIDR](/reference/rack-parameters/NLBAllowCIDR) — CIDR allowlist for ingress to public NLB listeners (default `0.0.0.0/0`, max 5 entries)
- [NLBCrossZone](/reference/rack-parameters/NLBCrossZone) — enable cross-zone load balancing on public listeners
- [NLBPreserveClientIP](/reference/rack-parameters/NLBPreserveClientIP) — forward real client IP to target tasks
- [NLBDeletionProtection](/reference/rack-parameters/NLBDeletionProtection) — block accidental NLB deletion

List all ten NLB-related parameters at once:

```bash
$ convox rack params -g nlb
```

### Disable

Before setting `NLB=No`, remove the `nlb:` field from every Service in every App deployed on the Rack and redeploy each. The disable is refused with a list of blocking Apps otherwise. This is intentional — disabling the NLB while Apps still reference it would break those Apps' next deploy via `Fn::ImportValue` resolution failure.

If [NLBDeletionProtection](/reference/rack-parameters/NLBDeletionProtection)=`Yes`, `NLB=No` is also rejected pre-flight; disable deletion protection in an earlier `rack params set` call, wait for the update to complete, then toggle `NLB=No` in a follow-up call.

### Known Limitations

See [Network Load Balancing](/networking/nlb#known-limitations) for limitations that apply to every NLB listener.

## See Also

- [NLBInternal](/reference/rack-parameters/NLBInternal)
- [NLBAllowCIDR](/reference/rack-parameters/NLBAllowCIDR)
- [services.nlb field](/application/services#nlb)
