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

This parameter is **incompatible with a Rack that sets a user-supplied [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup)**. Convox cannot add the required ingress rule to a security group it does not own. If your Rack uses a custom instance security group, add the ingress rule yourself (all protocols, sourced from `${Rack}:NLBSecurityGroup`) before enabling this parameter. See [Incompatibility with a custom InstanceSecurityGroup](#incompatibility-with-a-custom-instancesecuritygroup) below. Per-port `preserve_client_ip: true` on a Service is blocked on the same Racks.

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

### Incompatibility with a custom InstanceSecurityGroup

Racks that set [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup) to a security group you manage cannot enable `NLBPreserveClientIP=Yes`. The custom security group replaces `InstancesSecurity` on the ECS instances, so the Rack's NLB ingress rule does not reach them, and CloudFormation cannot attach rules to a security group Convox does not own. The Rack rejects the change:

```text
cannot enable NLBPreserveClientIP on a rack with a user-supplied InstanceSecurityGroup; your instance SG must add an ingress rule from the NLB security group (exported as ${Rack}:NLBSecurityGroup) for the NLB listener ports before this feature can be enabled safely
```

The fix on Racks with a custom InstanceSecurityGroup is to add the ingress rule to your security group manually, sourced from the Rack's exported `${Rack}:NLBSecurityGroup`, before attempting to enable preserve-client-IP. Allow all protocols from that group: tasks on the ECS instances listen on dynamic host ports, so a rule for the listener port alone does not reach them. Example:

```bash
$ aws ec2 authorize-security-group-ingress \
    --group-id sg-0123456789abcdef0 \
    --source-group $(aws cloudformation describe-stacks \
      --stack-name <rack> \
      --query 'Stacks[0].Outputs[?OutputKey==`NLBSecurityGroup`].OutputValue' \
      --output text) \
    --protocol all
```

The inverse direction is also blocked. Setting `InstanceSecurityGroup` on a Rack that does not have one while `NLBPreserveClientIP` is `Yes` is rejected unless the same `rack params set` call also sets `NLBPreserveClientIP=No`, because the new security group would not admit the forwarded traffic.

### Per-port enforcement

Release promote rejects a Release that sets `preserve_client_ip: true` on any `nlb:` port, public or internal, on a Rack with a custom InstanceSecurityGroup, whatever the value of this parameter:

```text
service web nlb port 443: cannot set preserve_client_ip=true on a rack with a user-supplied InstanceSecurityGroup; your instance SG must add an ingress rule from the NLB security group (exported as ${Rack}:NLBSecurityGroup / ${Rack}:NLBInternalSecurityGroup) for the NLB listener ports before this feature can be enabled safely
```

## See Also

- [NLBInternalPreserveClientIP](/reference/rack-parameters/NLBInternalPreserveClientIP)
- [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup)
- [services.nlb field](/application/services#nlb)
- [Network Load Balancing](/networking/nlb#preserve-client-ip)
