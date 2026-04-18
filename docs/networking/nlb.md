---
title: "Network Load Balancing"
description: "Expose TCP services through a Convox Rack via a Network Load Balancer."
---

# Network Load Balancing

Convox v2 Racks can provision Network Load Balancers (NLBs) alongside the default Application Load Balancer. NLBs operate at Layer 4 (TCP), allowing services to accept traffic on arbitrary ports and protocols that the ALB cannot handle — MQTT, Redis, raw TCP gRPC, game server protocols, and similar.

NLBs are opt-in at the Rack level and again per-service. The existing ALB routing is unaffected.

## Architecture

A Rack can host up to two shared NLBs:

- **Public NLB** — internet-facing, allocated one AWS Elastic IP per availability zone. Enabled via the [NLB](/reference/rack-parameters/NLB) rack parameter.
- **Internal NLB** — VPC-internal (no EIPs), scheme `internal`. Enabled via the [NLBInternal](/reference/rack-parameters/NLBInternal) rack parameter, which requires [Internal](/reference/rack-parameters/Internal)=`Yes`.

Services opt in per-port via the [nlb:](/application/services#nlb) field in `convox.yml`. Each declared port becomes a dedicated Listener and TargetGroup on the appropriate Rack NLB. Port values are unique per scheme per Rack — two services cannot share a public NLB port, though the same port can exist on both the public and internal NLB simultaneously.

## Enabling

```bash
$ convox rack params set NLB=Yes
```

Before running this, verify your AWS account has Elastic IP quota headroom. A public NLB consumes 2 EIPs on a 2-AZ Rack or 3 EIPs on an HA 3-AZ Rack. If the Rack already has `Private=Yes` (which consumes 2-3 EIPs for NAT gateways), a `Private=Yes` HA 3-AZ Rack will need 6 EIPs total — above the AWS default of 5 per region.

```bash
$ aws service-quotas get-service-quota \
    --service-code ec2 --quota-code L-0263D0A3
```

For internal NLB support:

```bash
$ convox rack params set Internal=Yes NLBInternal=Yes
```

After the CloudFormation update completes (typically 5-10 minutes), `convox rack` will show the NLB DNS name(s) and the allocated EIPs. Clients connect directly to the NLB DNS hostname or to an EIP.

## Service Configuration

```yaml
services:
  mqtt-broker:
    nlb:
      - port: 1883
        protocol: tcp
        scheme: public
```

See [services.nlb](/application/services#nlb) for the full field reference.

Services can declare multiple `nlb:` entries combining public and internal schemes. A service that needs both HTTP (via the ALB) and raw TCP (via the NLB) can use `port:` and `nlb:` together on the same container port.

If a service declares an `nlb:` port whose `scheme` does not match an enabled Rack NLB, the deploy is rejected at release promote with a clear error:

```
service X declares public nlb port 8443 but rack does not have NLB enabled; run 'convox rack params set NLB=Yes' first
```

Set the corresponding Rack parameter first, then redeploy.

## Known Limitations

### TCP only

Only TCP listeners are supported in this release. UDP and TLS listeners are planned for a future release. Services that need TLS termination must terminate TLS in the application.

### No real client IP

NLB target groups are configured with `preserve_client_ip.enabled=false` so the existing per-service security group rule (VPC CIDR ingress) covers both public and internal traffic uniformly. Application logs will record the NLB's VPC-internal IP rather than the real client address. Compliance frameworks that require real client IPs (HIPAA §164.312(b), PCI-DSS 10.2.1) are not satisfied by the v1 implementation — defer adoption until a future release enables per-service client-IP preservation.

### Plaintext traffic on public NLB

Public NLB listeners carry plaintext TCP. Customers must terminate TLS in the application until NLB-level TLS support ships.

### No operator-level CIDR allowlist

Ports declared with `scheme: public` are reachable from `0.0.0.0/0`. There is no Rack-level allowlist in v1; restrict access at the application layer.

### Cross-zone load balancing is off

AWS's default for NLB is off. With per-AZ EIPs, clients hitting one EIP only reach targets in that AZ. Services with uneven target distribution across AZs may hotspot. A future release will add a toggle.

### 50 listeners per NLB

The AWS default quota is 50 listeners per load balancer. Because Racks share one public NLB and one internal NLB across all services, the combined total of `scheme: public` NLB ports (and similarly `scheme: internal`) across all apps on the Rack cannot exceed 50 without a quota increase.

### Disable procedure

To disable NLB on a Rack, first remove the `nlb:` block from every service in every app and redeploy each. Only then will `convox rack params set NLB=No` succeed. The disable step releases the EIPs — if you re-enable later, the Rack will be assigned new EIPs and a new NLB DNS name. Re-validate any customer DNS pointing at the Rack after a disable/re-enable cycle.

### Two concurrent deploys claiming the same port

Two apps concurrently deploying with the same NLB listener port will both pass validation. One CloudFormation stack update will succeed; the other will fail with an ELBv2 duplicate-listener error via stack events. A future release adds proactive rack-wide conflict detection.

### Downgrade

Before downgrading a Rack to a version that predates NLB support, set both `NLB=No` and `NLBInternal=No` and wait for the CloudFormation update to complete. Otherwise the downgrade will fail with `Parameters: [NLB, NLBInternal] do not exist in the template`.

### NLB-only services require `Fargate` or `Isolate`

A service that declares `nlb:` ports without a `port:` field and runs on a default EC2 launch rack (no Fargate, no Isolate) inherits the Rack's `InstancesSecurity` default network mode. ECS target registration for the NLB target group may require the service-linked role `AWSServiceRoleForECS`, which is auto-created by AWS on first ECS usage. If you see target registration failures on NLB-only services, confirm the role exists in your AWS account.

## See Also

- [NLB rack parameter](/reference/rack-parameters/NLB)
- [NLBInternal rack parameter](/reference/rack-parameters/NLBInternal)
- [Load Balancing (ALB)](/networking/load-balancing)
- [Internal Services](/networking/internal-services)
- [services.nlb field](/application/services#nlb)
