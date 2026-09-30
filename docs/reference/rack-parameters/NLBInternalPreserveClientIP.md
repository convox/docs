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

Same constraint as [NLBPreserveClientIP](/reference/rack-parameters/NLBPreserveClientIP#prerequisites): this parameter cannot be enabled on a Rack that sets a user-supplied [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup). On such Racks, add an ingress rule to your custom security group (all protocols, sourced from `${Rack}:NLBInternalSecurityGroup`) before attempting to enable this parameter. The example on the [NLBPreserveClientIP](/reference/rack-parameters/NLBPreserveClientIP#incompatibility-with-a-custom-instancesecuritygroup) page applies with `` OutputKey==`NLBInternalSecurityGroup` `` in the query.

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

### Incompatibility with a custom InstanceSecurityGroup

Racks that set [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup) cannot enable this parameter. The Rack rejects the change with an error that references `${Rack}:NLBInternalSecurityGroup`:

```text
cannot enable NLBInternalPreserveClientIP on a rack with a user-supplied InstanceSecurityGroup; your instance SG must add an ingress rule from the NLB security group (exported as ${Rack}:NLBInternalSecurityGroup) for the NLB listener ports before this feature can be enabled safely
```

The inverse direction (setting `InstanceSecurityGroup` while this is already `Yes`) is also blocked unless the same call disables preserve-client-IP. See [NLBPreserveClientIP](/reference/rack-parameters/NLBPreserveClientIP#incompatibility-with-a-custom-instancesecuritygroup) for the full rationale.

## See Also

- [NLBPreserveClientIP](/reference/rack-parameters/NLBPreserveClientIP)
- [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup)
- [services.nlb field](/application/services#nlb)
- [Network Load Balancing](/networking/nlb#preserve-client-ip)
