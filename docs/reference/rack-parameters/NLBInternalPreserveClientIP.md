---
title: "NLBInternalPreserveClientIP"
description: "Forward the real client IP to targets behind the internal Network Load Balancer."
---

# NLBInternalPreserveClientIP

Forward the real client source IP to targets behind the internal [NLBInternal](/reference/rack-parameters/NLBInternal). Same semantics as [NLBPreserveClientIP](/reference/rack-parameters/NLBPreserveClientIP) but scoped to the internal NLB.

| Setting | Value |
|:--|:--|
| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Prerequisites

Same requirement as [NLBPreserveClientIP](/reference/rack-parameters/NLBPreserveClientIP#prerequisites): on a Rack that sets a custom [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup), add an ingress rule to that group allowing all traffic from the internal NLB security group (exported as `${Rack}:NLBInternalSecurityGroup`) before enabling this parameter. The example on the [NLBPreserveClientIP](/reference/rack-parameters/NLBPreserveClientIP#custom-instancesecuritygroup) page applies with `` OutputKey==`NLBInternalSecurityGroup` `` in the query.

## Use Cases

- In-VPC audit logs that need the calling Service's real IP rather than the internal NLB's address
- Per-client rate limiting in a microservice topology sitting behind an internal NLB
- Compliance logging on internal-only workloads
- Analytics on traffic arriving from peered VPCs or VPN clients

## Additional Information

```bash
$ convox rack params set NLBInternalPreserveClientIP=Yes
```

Applies to every listener on the internal NLB. An App's listeners pick up a change on the App's next release promote (`convox deploy` or `convox releases promote`). Per-port [preserve_client_ip:](/application/services#nlb) on a Service with `scheme: internal` overrides this Rack default for a single listener.

### Custom InstanceSecurityGroup

Requires rack version 20261005214736 or newer.

On a Rack with a custom [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup), the Rack accepts this parameter once that group has an ingress rule allowing all traffic from `${Rack}:NLBInternalSecurityGroup`, and refuses it otherwise:

```text
preserve client IP on the internal NLB needs an ingress rule on InstanceSecurityGroup sg-0123456789abcdef0 allowing all traffic from the NLB security group sg-0fedcba9876543210 (production:NLBInternalSecurityGroup); add that rule and retry
```

The rest of [Custom InstanceSecurityGroup](/reference/rack-parameters/NLBPreserveClientIP#custom-instancesecuritygroup) on the NLBPreserveClientIP page applies to the internal NLB, with `NLBInternal`, `NLBInternalPreserveClientIP`, `${Rack}:NLBInternalSecurityGroup`, and `scheme: internal` ports in place of the public ones: setting `InstanceSecurityGroup` to a new group needs the rule, `NLBInternal=Yes` is refused while this parameter is `Yes`, and the rule must be removed before `NLBInternal=No` or an uninstall.

## See Also

- [NLBPreserveClientIP](/reference/rack-parameters/NLBPreserveClientIP)
- [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup)
- [services.nlb field](/application/services#nlb)
- [Network Load Balancing](/networking/nlb#preserve-client-ip)
