---
title: "NLBPreserveClientIP"
description: "Forward the real client IP to targets behind the public Network Load Balancer."
---

# NLBPreserveClientIP

Forward the real client source IP to targets behind the public [NLB](/reference/rack-parameters/NLB). When `No` (default), the NLB source-NATs incoming connections to its VPC-internal IP — applications see the NLB's address rather than the client's. When `Yes`, target tasks observe the real client IP.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Prerequisites

This parameter is **incompatible with a Rack that sets a customer-supplied [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup)**. Convox cannot add the required ingress rule to an SG it does not own. If your Rack uses a custom instance SG, add the ingress rule yourself (sourced from `${Rack}:NLBSecurityGroup`) before enabling this parameter — see [Incompatibility with customer InstanceSecurityGroup](#incompatibility-with-customer-instancesecuritygroup) below. Per-port `preserve_client_ip: true` on a Service is blocked on the same Racks.

## Use Cases

- Compliance requirements for real client IPs in application logs (HIPAA §164.312(b), PCI-DSS 10.2.1)
- Application-layer rate limiting or abuse mitigation keyed on client IP
- GeoIP-based routing or analytics
- Any workload where the NLB's internal IP in logs is operationally unhelpful

## Additional Information

```bash
$ convox rack params set NLBPreserveClientIP=Yes
```

When enabled, Convox attaches an ingress rule on the ECS instances' security group (`InstancesSecurity`) sourced from the NLB security group (`NLBSecurity`), allowing the real-IP forwarded traffic to reach targets. For Fargate and [Isolate](/reference/app-parameters/Isolate) Services, the equivalent rule is added to each Service's `Security` security group so `ip`-type target groups (used by `awsvpc`-mode tasks) work correctly.

The setting applies to every existing and future listener on the public NLB. Per-port [preserve_client_ip:](/application/services#nlb) overrides the rack default for a single listener.

### Incompatibility with customer InstanceSecurityGroup

Racks that set [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup) to a customer-managed security group cannot enable `NLBPreserveClientIP=Yes` through Convox — the CloudFormation stack cannot attach ingress rules to an SG Convox does not own. The validator rejects the change:

```text
cannot enable NLBPreserveClientIP on a rack with a customer-supplied
InstanceSecurityGroup; your instance SG must add an ingress rule from the
NLB security group (exported as ${Rack}:NLBSecurityGroup) for the NLB
listener ports before this feature can be enabled safely
```

The fix on customer-SG Racks is to add the ingress rule to your SG manually — sourced from the Rack's exported `${Rack}:NLBSecurityGroup`, one rule per NLB listener port you expose — before attempting to enable preserve-client-IP. Example:

```bash
$ aws ec2 authorize-security-group-ingress \
    --group-id sg-customer-instance \
    --source-group $(aws cloudformation describe-stacks \
      --stack-name <rack> \
      --query 'Stacks[0].Outputs[?OutputKey==`NLBSecurityGroup`].OutputValue' \
      --output text) \
    --protocol tcp --port <listener-port>
```

Repeat the command for each NLB listener port declared in your Apps' `convox.yml`.

The inverse direction is also blocked. Setting `InstanceSecurityGroup` to a non-empty value while `NLBPreserveClientIP=Yes` is already in force is rejected unless the same `rack params set` call also disables preserve-client-IP — preventing a silent-breakage transition where the SG swap leaves the old ingress rule behind but the new SG does not allow the real-IP traffic.

### Per-port enforcement

If the Rack parameter is `No` but a Service declares `preserve_client_ip: true` on an `nlb:` entry, release promote runs the same interlock on that scheme and rejects the Release on customer-SG Racks with a parallel error.

## See Also

- [NLBInternalPreserveClientIP](/reference/rack-parameters/NLBInternalPreserveClientIP)
- [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup)
- [services.nlb field](/application/services#nlb)
- [Network Load Balancing](/networking/nlb#preserve-client-ip)
