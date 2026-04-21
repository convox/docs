---
title: "Network Load Balancing"
description: "Expose TCP and TLS services through a Convox Rack via a shared Network Load Balancer."
---

# Network Load Balancing

Convox v2 Racks can provision Network Load Balancers (NLBs) alongside the default Application Load Balancer. NLBs operate at Layer 4, allowing Services to accept traffic on arbitrary ports — raw TCP for protocols the ALB cannot handle (MQTT, Redis, raw TCP gRPC, game-server protocols) and TLS-terminating listeners when you need ACM-issued certificates on a non-HTTP port.

NLBs are opt-in at the Rack level and again per-Service. The existing ALB routing is unaffected.

## Architecture

A Rack can host up to two shared NLBs:

- **Public NLB** — internet-facing, allocated one AWS Elastic IP per Availability Zone. Enabled via the [NLB](/reference/rack-parameters/NLB) rack parameter.
- **Internal NLB** — VPC-internal (no EIPs), scheme `internal`. Enabled via the [NLBInternal](/reference/rack-parameters/NLBInternal) rack parameter, which requires [Internal](/reference/rack-parameters/Internal)=`Yes`.

Services opt in per-port via the [nlb:](/application/services#nlb) field in `convox.yml`. Each declared port becomes a dedicated Listener and TargetGroup on the appropriate Rack NLB.

Within a single App's `convox.yml`, every NLB port must be unique across all Services regardless of scheme — the manifest validator rejects a second Service declaring the same port number, even if one is `public` and the other `internal`. Across Apps, the public NLB and internal NLB are separate AWS load balancers, so port 443 on App A's public listener and port 443 on App B's internal listener coexist without conflict. Two Apps claiming the same port on the *same* scheme is not caught at manifest time; the second deploy fails later at CloudFormation stack update (see "Two concurrent deploys" under Known Limitations).

## Enabling the Rack NLBs

```bash
$ convox rack params set NLB=Yes
NLB provisioning typically takes 5-10 minutes; check status with 'convox rack'.
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

After the CloudFormation update completes, `convox rack` will show the NLB DNS name(s) and allocated EIPs:

```bash
$ convox rack
Name          production
Provider      aws
Region        us-east-1
Router        router.0a1b2c3d4e5f.convox.cloud
NLB           production-nlb-abc123.elb.us-east-1.amazonaws.com (52.1.2.3, 52.4.5.6, 52.7.8.9)
NLB Internal  production-nlb-internal-xyz789.elb.us-east-1.amazonaws.com
Status        running
Version       20260418101514
```

Clients connect directly to the NLB DNS hostname or to an EIP.

## Service Configuration

### Plain TCP

```yaml
services:
  mqtt-broker:
    image: eclipse-mosquitto:2
    nlb:
      - port: 1883
        protocol: tcp
        scheme: public
```

### TLS termination at the NLB

Attach an ACM or IAM server certificate to a listener with `protocol: tls` and a `certificate:` ARN. The NLB terminates TLS and forwards plaintext TCP to the container — backend Services do not need to hold the certificate material.

```yaml
services:
  api:
    image: example/api
    nlb:
      - port: 443
        protocol: tls
        containerPort: 8080
        scheme: public
        certificate: arn:aws:acm:us-east-1:123456789012:certificate/abcd1234-5678-90ab-cdef-1234567890ab
```

TLS listeners use `ELBSecurityPolicy-TLS13-1-2-2021-06` (TLS 1.2 and 1.3 with ECDHE ciphers). The target group protocol remains TCP — backends never see TLS traffic.

The `certificate:` field requires a full AWS ARN. `convox certs` lists both ACM-issued and IAM-imported certificates on the Rack, but displays them by Convox-synthesized short ID (ACM: `acm-<hash>`) or IAM server-certificate name — neither is a full ARN. Retrieve the ARN from AWS directly:

- **ACM**: copy from the AWS Console Certificate Manager page, or run `aws acm list-certificates --region <rack-region>`.
- **IAM**: construct as `arn:aws:iam::<account-id>:server-certificate/<name>` using the name `convox certs` already shows; `aws sts get-caller-identity` returns the account ID.

Both ARN formats are accepted:

- ACM: `arn:aws:acm:<region>:<account>:certificate/<uuid>`
- IAM server-certificate: `arn:aws:iam::<account>:server-certificate/<name>`

See [services.nlb](/application/services#nlb) for the full field reference.

### Mixing schemes and protocols

A Service can declare multiple `nlb:` entries combining public + internal schemes and `tcp` + `tls` protocols. A Service that needs both HTTP (via the ALB) and raw TCP (via the NLB) can use `port:` and `nlb:` together on the same `containerPort`.

```yaml
services:
  api:
    image: example/api
    port: 3000/http
    nlb:
      - port: 443
        protocol: tls
        containerPort: 3000
        scheme: public
        certificate: arn:aws:acm:us-east-1:123456789012:certificate/abcd1234-5678-90ab-cdef-1234567890ab
      - port: 50051
        protocol: tcp
        containerPort: 50051
        scheme: internal
```

### Release-time validation

If a Service declares an `nlb:` port whose `scheme` does not match an enabled Rack NLB, the deploy is rejected at release promote with a clear error:

```
service api declares public nlb port 443 but rack does not have NLB enabled;
run 'convox rack params set NLB=Yes' first
```

TLS listeners are validated at release promote too — the referenced certificate ARN must exist in the Rack's region and account, and ACM certificates must be in `ISSUED` state. Typical failure messages:

```
certificate arn:aws:acm:us-east-1:123456789012:certificate/...: not found in
  this region (is this cert in another region?)
certificate arn:aws:acm:us-east-1:123456789012:certificate/...: not usable
  (status: PENDING_VALIDATION)
certificate arn:aws:acm:us-east-1:999999999999:certificate/...: access denied
  (cross-account certificates are not supported)
certificate arn:aws:iam::123456789012:server-certificate/legacy: IAM server
  certificate not found
```

These fail immediately on `convox releases promote`, not as an opaque CloudFormation error ten minutes later.

### Viewing NLB ports on a Service

`convox services` adds an `NLB PORTS` column when any Service on the App declares `nlb:` ports. Format is `PORT:CONTAINERPORT` with `/tls` when the protocol is TLS and `(internal)` when the scheme is internal.

```bash
$ convox services -a broker
SERVICE  DOMAIN                              PORTS      NLB PORTS
api      api.broker.0a1b2c.convox.cloud      443:3000   443:3000/tls 50051:50051(internal)
worker
```

## Changing a listener's protocol

Switching an `nlb:` entry from `protocol: tcp` to `protocol: tls` (or vice versa) on an existing port modifies the listener in place via AWS `ModifyListener`. AWS documents this as a no-interruption update, though clients holding an open connection at the exact protocol-boundary moment may observe a brief disruption — switch during a low-traffic window.

## Known Limitations

### No real client IP

NLB target groups are configured with `preserve_client_ip.enabled=false` by design, so the per-Service security group rule (VPC CIDR ingress) covers both public and internal traffic uniformly. Application logs record the NLB's VPC-internal IP rather than the real client address. Compliance frameworks that require real client IPs (HIPAA §164.312(b), PCI-DSS 10.2.1) are not satisfied by this configuration — workloads with those requirements should terminate the Layer 4 listener in front of a proxy that injects the client IP, or wait for a future Convox option to toggle this attribute.

### No operator-level CIDR allowlist

Ports declared with `scheme: public` are reachable from `0.0.0.0/0`. There is no Rack-level allowlist in this release; restrict access at the application layer or with a custom security group on the target service.

### Cross-zone load balancing is off

NLB inherits AWS's default — cross-zone load balancing is disabled. With per-AZ EIPs, clients hitting one EIP only reach targets in that AZ. Services with uneven target distribution across AZs may hotspot. This is not currently configurable through a Rack parameter.

### 50 listeners per NLB

The AWS default quota is 50 listeners per load balancer. Because Racks share one public NLB and one internal NLB across all Services, the combined total of `scheme: public` NLB ports (and similarly `scheme: internal`) across all Apps on the Rack cannot exceed 50 without a quota increase.

### Disable procedure

To disable NLB on a Rack, first remove the `nlb:` block from every Service in every App and redeploy each. Only then will `convox rack params set NLB=No` succeed. The disable step releases the EIPs — if you re-enable later, the Rack will be assigned new EIPs and a new NLB DNS name. Re-validate any customer DNS pointing at the Rack after a disable/re-enable cycle.

### Two concurrent deploys claiming the same port

Two Apps concurrently deploying with the same NLB listener port both pass release-promote validation (which only inspects the single App's manifest). One CloudFormation stack update succeeds; the other fails with an ELBv2 duplicate-listener error surfaced via stack events. Avoid concurrent deploys that claim the same NLB port across Apps, and confirm port allocation manually when multiple teams share a Rack.

### Downgrade

Before downgrading a Rack to a version that predates NLB support, set both `NLB=No` and `NLBInternal=No` and wait for the CloudFormation update to complete. Otherwise the downgrade fails with `Parameters: [NLB, NLBInternal] do not exist in the template`.

### NLB-only Services on EC2 launch type

A Service that declares only `nlb:` ports (no `port:` field) and runs on a default EC2-launch Rack (no Fargate, no [Isolate](/reference/app-parameters/Isolate)) registers targets via the ECS service-linked role `AWSServiceRoleForECS`. AWS creates this role automatically on first ECS usage — if target registration fails on an NLB-only Service, confirm the role exists in the account. Fargate and Isolate Services use `awsvpc` mode and register targets by IP, so the role requirement does not apply.

## See Also

- [NLB rack parameter](/reference/rack-parameters/NLB)
- [NLBInternal rack parameter](/reference/rack-parameters/NLBInternal)
- [services.nlb field](/application/services#nlb)
- [Load Balancing (ALB)](/networking/load-balancing)
