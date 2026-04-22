---
title: "NLBAllowCIDR"
description: "Comma-delimited IPv4 CIDR allowlist for the public Network Load Balancer (max 5 entries)."
---

# NLBAllowCIDR

Comma-delimited list of up to 5 IPv4 CIDRs permitted to reach listeners on the public [NLB](/reference/rack-parameters/NLB). Each non-empty entry becomes an ingress rule on the `NLBSecurity` security group. Per-port `allow_cidr:` entries declared in a Service's [nlb:](/application/services#nlb) field stack additively on top of this list — they never remove rack-level entries. **IPv4 only**; IPv6 CIDRs are rejected at `rack params set`.

| Default value  | `0.0.0.0/0` (all IPv4 addresses)                     |
| Allowed values | Comma-delimited list of up to 5 canonical IPv4 CIDRs |

## Use Cases

- Restricting a public NLB to known corporate egress IPs or partner CIDR ranges
- Complementing application-layer auth with a network-layer narrow list
- Staging a launch behind a small test-user allowlist before opening to `0.0.0.0/0`

## Additional Information

```bash
$ convox rack params set NLBAllowCIDR=203.0.113.0/24,198.51.100.0/24
```

Validation runs at `convox rack params set`. The Rack rejects:

- More than 5 non-empty entries
- IPv6 CIDRs (`2001:db8::/32`)
- Non-canonical forms where host bits are set (`10.0.0.1/24` → use `10.0.0.0/24`)
- Duplicate entries
- Leading or trailing whitespace around any entry

Each failure is reported with the offending entry and the fix. `convox rack params set` does not apply a partial update — if any entry is invalid, none are applied:

```text
$ convox rack params set NLBAllowCIDR=10.0.0.1/24,192.168.1.0/24
NLBAllowCIDR entry "10.0.0.1/24" is not canonical; use "10.0.0.0/24" instead

$ convox rack params set NLBAllowCIDR=2001:db8::/32
NLBAllowCIDR entry "2001:db8::/32" is not a valid IPv4 CIDR (IPv6 not supported)
```

The 5-entry limit is a CloudFormation constraint (one `AWS::EC2::SecurityGroupIngress` resource per slot). Larger allowlists must be applied per-port through [services.nlb allow_cidr](/application/services#nlb), which uses a separate per-port ingress rule and is not subject to the rack-level slot cap.

### Relationship to per-port `allow_cidr`

A Service's `allow_cidr:` is additive, not replacive. If the rack-level list is `203.0.113.0/24` and a Service declares `allow_cidr: [10.0.0.0/16]` on a listener port, that port accepts traffic from **both** `203.0.113.0/24` and `10.0.0.0/16`.

To scope a specific listener more narrowly than the rack default, declare an `allow_cidr:` that omits the broad ranges and reduce the rack default separately.

### Setting to empty

Setting `NLBAllowCIDR=""` removes all ingress rules on the public NLB's security group, which blocks every public NLB listener from receiving traffic. This is a kill-switch posture; use it only to freeze traffic during incident response. Restore an explicit CIDR list (or the `0.0.0.0/0` default) to re-open the NLB.

## See Also

- [NLB](/reference/rack-parameters/NLB)
- [NLBInternalAllowCIDR](/reference/rack-parameters/NLBInternalAllowCIDR)
- [services.nlb field](/application/services#nlb)
- [Network Load Balancing](/networking/nlb#ingress-allowlist)
