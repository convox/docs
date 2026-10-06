---
title: "NLBPreserveClientIP"
description: "Forward the real client IP to targets behind the public Network Load Balancer."
---

# NLBPreserveClientIP

Forward the real client source IP to targets behind the public [NLB](/reference/rack-parameters/NLB). When `No` (default), the NLB source-NATs incoming connections to its VPC-internal IP, so applications see the NLB's address rather than the client's. When `Yes`, target tasks observe the real client IP.

| Setting | Value |
|:--|:--|
| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Prerequisites

Racks that use the default instance security group need no setup. On a Rack that sets a custom [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup), add an ingress rule to that group allowing all traffic from the public NLB security group (exported as `${Rack}:NLBSecurityGroup`) before enabling this parameter. The Rack reads the group and accepts the change once the rule exists. Per-port `preserve_client_ip: true` on a Service needs the same rule. See [Custom InstanceSecurityGroup](#custom-instancesecuritygroup) below.

## Use Cases

- Compliance requirements for real client IPs in application logs (HIPAA §164.312(b), PCI-DSS 10.2.1)
- Application-layer rate limiting or abuse mitigation keyed on client IP
- GeoIP-based routing or analytics
- Any workload where the NLB's internal IP in logs is operationally unhelpful

## Additional Information

```bash
$ convox rack params set NLBPreserveClientIP=Yes
```

Targets accept the forwarded traffic through ingress rules sourced from the NLB security group (`NLBSecurity`). The Rack adds one to the ECS instances' security group (`InstancesSecurity`) while [NLB](/reference/rack-parameters/NLB) is `Yes`. Fargate and [Isolate](/reference/app-parameters/Isolate) Services run `awsvpc`-mode tasks behind `ip`-type target groups, so each of those Services' `Security` security group has an equivalent rule for every `nlb:` port.

The setting applies to every listener on the public NLB. An App's listeners pick up a change on the App's next release promote (`convox deploy` or `convox releases promote`). Per-port [preserve_client_ip:](/application/services#nlb) overrides the Rack default for a single listener.

### Custom InstanceSecurityGroup

Requires rack version 20261005214736 or newer.

A custom [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup) replaces `InstancesSecurity` on the ECS instances, so the Rack's NLB ingress rule does not reach them, and Convox does not change a security group it does not own. Add an ingress rule to your group that allows all traffic from the Rack's exported `${Rack}:NLBSecurityGroup`, then enable the parameter:

```bash
$ aws ec2 authorize-security-group-ingress \
    --group-id sg-0123456789abcdef0 \
    --source-group $(aws cloudformation describe-stacks \
      --stack-name <rack> \
      --query 'Stacks[0].Outputs[?OutputKey==`NLBSecurityGroup`].OutputValue' \
      --output text) \
    --protocol all
$ convox rack params set NLBPreserveClientIP=Yes
Updating parameters... OK
```

Only a rule that allows all traffic (`--protocol all`) counts; the Rack does not accept a rule for a port or a port range. Tasks on the ECS instances listen on dynamic host ports, so a rule for the listener port alone does not reach them. Without the rule the Rack refuses the change and names both groups:

```text
preserve client IP on the public NLB needs an ingress rule on InstanceSecurityGroup sg-0123456789abcdef0 allowing all traffic from the NLB security group sg-0fedcba9876543210 (production:NLBSecurityGroup); add that rule and retry
```

While [NLB](/reference/rack-parameters/NLB) is `Yes`, the Rack checks your group every time this parameter is set to `Yes`, and on these commands:

| Command | Result |
|:--|:--|
| `convox rack params set InstanceSecurityGroup=<sg>` while this parameter is `Yes`, or while a deployed App sets `preserve_client_ip: true` on a public `nlb:` port | Accepted only when `<sg>` has the rule, even when the same command sets `NLBPreserveClientIP=No` |
| `convox deploy` or `convox releases promote` with `preserve_client_ip: true` on a public `nlb:` port | Accepted only when the group has the rule. See [Per-port enforcement](#per-port-enforcement) |

Deployed listeners keep client IP preservation until their App's next release promote. After setting `NLBPreserveClientIP=No`, redeploy every App with public `nlb:` ports before changing `InstanceSecurityGroup`. Keep the rule in place while client IP preservation is in use; without it the NLB cannot reach tasks on the ECS instances.

`NLB=Yes` is refused on a Rack with a custom InstanceSecurityGroup while this parameter is `Yes`, because the NLB security group does not exist yet:

```text
cannot enable NLB while NLBPreserveClientIP=Yes on a rack with a custom InstanceSecurityGroup; set NLBPreserveClientIP=No in this command, add an ingress rule to sg-0123456789abcdef0 allowing all traffic from the new production:NLBSecurityGroup, then set NLBPreserveClientIP=Yes
```

Run `convox rack params set NLB=Yes NLBPreserveClientIP=No`, wait for the update to complete, add the rule from the new `${Rack}:NLBSecurityGroup` export, then set `NLBPreserveClientIP=Yes`.

Before setting `NLB=No` or running [rack uninstall](/reference/cli-commands/rack-uninstall), remove the rule from your group:

```bash
$ aws ec2 revoke-security-group-ingress \
    --group-id sg-0123456789abcdef0 \
    --source-group sg-0fedcba9876543210 \
    --protocol all
```

EC2 does not delete a security group that another group's rule references. With the rule in place, `NLB=No` leaves the NLB security group behind, and an uninstall ends with the Rack stack in `DELETE_FAILED`.

### Per-port enforcement

On a Rack with a custom InstanceSecurityGroup, release promote (`convox deploy` or `convox releases promote`) checks every `nlb:` port that sets `preserve_client_ip: true`, public or internal, whatever the value of this parameter. The Release is accepted once the group has the rule for that port's NLB (`${Rack}:NLBSecurityGroup` for `scheme: public`, `${Rack}:NLBInternalSecurityGroup` for `scheme: internal`), and refused without it:

```text
service web nlb port 443: preserve client IP on the public NLB needs an ingress rule on InstanceSecurityGroup sg-0123456789abcdef0 allowing all traffic from the NLB security group sg-0fedcba9876543210 (production:NLBSecurityGroup); add that rule and retry
```

## See Also

- [NLBInternalPreserveClientIP](/reference/rack-parameters/NLBInternalPreserveClientIP)
- [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup)
- [services.nlb field](/application/services#nlb)
- [Network Load Balancing](/networking/nlb#preserve-client-ip)
