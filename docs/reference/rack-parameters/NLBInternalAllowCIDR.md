---
title: "NLBInternalAllowCIDR"
description: "Comma-delimited IPv4 CIDR allowlist for the internal Network Load Balancer (max 5 entries)."
---

# NLBInternalAllowCIDR

Comma-delimited list of up to 5 IPv4 CIDRs permitted to reach listeners on the internal [NLBInternal](/reference/rack-parameters/NLBInternal). Each non-empty entry becomes an ingress rule on the `NLBInternalSecurity` security group. Per-port `allow_cidr:` entries declared in a Service's [nlb:](/application/services#nlb) field stack additively on top of this list. **IPv4 only**; IPv6 CIDRs are rejected at `rack params set`.

| Default value  | `""` (empty: falls back to the Rack's VPC CIDR)      |
| Allowed values | Comma-delimited list of up to 5 canonical IPv4 CIDRs |

## Use Cases

- Limiting internal NLB reachability to a subset of the VPC (specific subnets, peered-VPC ranges, VPN CIDRs)
- Opening the internal NLB to a peered VPC without opening the entire VPC
- Exposing the internal NLB to VPN clients whose CIDR is outside the Rack's VPC CIDR

## Additional Information

The **default is empty**, which is intentionally different from the public [NLBAllowCIDR](/reference/rack-parameters/NLBAllowCIDR). An empty internal-allowlist triggers a CloudFormation fallback that permits the Rack's VPC CIDR — preserving the "reachable from inside the VPC" posture an internal NLB is installed for.

```bash
$ convox rack params set NLBInternalAllowCIDR=10.0.0.0/16,10.1.0.0/16
```

**Setting an explicit value replaces the VPC CIDR fallback.** If you set `NLBInternalAllowCIDR=10.1.0.0/16` on a Rack whose VPC is `10.0.0.0/16`, the Rack's own instances lose in-VPC reachability to the internal NLB. Include the Rack's VPC CIDR explicitly if you still need intra-rack reachability:

```bash
$ convox rack params set NLBInternalAllowCIDR=10.0.0.0/16,10.1.0.0/16
```

Validation is identical to [NLBAllowCIDR](/reference/rack-parameters/NLBAllowCIDR#additional-information): IPv4 only, canonical form, up to 5 entries, no duplicates, no surrounding whitespace.

### Relationship to per-port `allow_cidr`

Same model as the public allowlist: per-port entries add ingress rules to the `NLBInternalSecurity` security group scoped to the listener port — they do not replace the rack-level entries. See [services.nlb](/application/services#nlb).

## See Also

- [NLBInternal](/reference/rack-parameters/NLBInternal)
- [NLBAllowCIDR](/reference/rack-parameters/NLBAllowCIDR)
- [VPCCIDR](/reference/rack-parameters/VPCCIDR)
- [Network Load Balancing](/networking/nlb#ingress-allowlist)
