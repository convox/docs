---
title: "Network Load Balancing"
description: "Expose TCP and TLS services through a Convox Rack via a shared Network Load Balancer."
---

# Network Load Balancing

Convox v2 Racks can provision Network Load Balancers (NLBs) alongside the default Application Load Balancer. NLBs operate at Layer 4, allowing Services to accept traffic on arbitrary ports — raw TCP for protocols the ALB cannot handle (MQTT, Redis, raw TCP gRPC, game-server protocols) and TLS-terminating listeners when you need ACM-issued certificates on a non-HTTP port.

NLBs are opt-in at the Rack level and again per-Service. The existing ALB routing is unaffected.

## Architecture

A Rack can host up to two shared NLBs:

- **Public NLB** — internet-facing, allocated one AWS Elastic IP per Availability Zone. Enabled via the [NLB](/reference/rack-parameters/NLB) rack parameter. Mutually exclusive with [InternalOnly](/reference/rack-parameters/InternalOnly)=`Yes`; Racks configured as InternalOnly must use the internal NLB.
- **Internal NLB** — VPC-internal (no EIPs), scheme `internal`. Enabled via the [NLBInternal](/reference/rack-parameters/NLBInternal) rack parameter, which requires [Internal](/reference/rack-parameters/Internal)=`Yes`.

Each NLB is fronted by a dedicated security group (`NLBSecurity` for the public NLB, `NLBInternalSecurity` for the internal NLB). The security groups are exported from the Rack stack as `${Rack}:NLBSecurityGroup` and `${Rack}:NLBInternalSecurityGroup` for operators who need to reference them from custom infrastructure. Rack-level allowlists ([NLBAllowCIDR](/reference/rack-parameters/NLBAllowCIDR) / [NLBInternalAllowCIDR](/reference/rack-parameters/NLBInternalAllowCIDR)) and per-service `allow_cidr:` entries attach ingress rules to these groups.

Services opt in per-port via the [nlb:](/application/services#nlb) field in `convox.yml`. Each declared port becomes a dedicated Listener and TargetGroup on the appropriate Rack NLB.

Within a single App's `convox.yml`, every NLB port must be unique across all Services regardless of scheme — the manifest validator rejects a second Service declaring the same port number, even if one is `public` and the other `internal`. Across Apps, the public NLB and internal NLB are separate AWS load balancers, so port 443 on App A's public listener and port 443 on App B's internal listener coexist without conflict. Two Apps claiming the same port on the *same* scheme is not caught at manifest time; the second deploy fails later at CloudFormation stack update (see "Two concurrent deploys" under Known Limitations).

## Enabling the Rack NLBs

```bash
$ convox rack params set NLB=Yes
Updating parameters... NLB provisioning typically takes 5-10 minutes; check status with 'convox rack'.
OK
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
Version       20260421192651
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

```text
service api declares public nlb port 443 but rack does not have NLB enabled;
run 'convox rack params set NLB=Yes' first
```

TLS listeners are validated at release promote too — the referenced certificate ARN must exist in the Rack's region and account, and ACM certificates must be in `ISSUED` state. Typical failure messages:

```text
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

`convox services` adds an `NLB PORTS` column when any Service on the App declares `nlb:` ports. Each cell is `PORT:CONTAINERPORT`, suffixed with `/tls` when the protocol is TLS, `(internal)` when the scheme is internal, and a bracket `[cz=... allow=N pcip=...]` when the port has per-port overrides for cross-zone, allow-CIDR, or preserve-client-IP (see [Per-port NLB attributes](#per-port-nlb-attributes) below).

```bash
$ convox services -a broker
SERVICE  DOMAIN                              PORTS      NLB PORTS
api      api.broker.0a1b2c.convox.cloud      443:3000   443:3000/tls 50051:50051(internal)
worker
```

## Changing a listener's protocol

Switching an `nlb:` entry from `protocol: tcp` to `protocol: tls` (or vice versa) on an existing port modifies the listener in place via AWS `ModifyListener`. AWS documents this as a no-interruption update, though clients holding an open connection at the exact protocol-boundary moment may observe a brief disruption — switch during a low-traffic window.

## Rack-level NLB configuration

Four behaviors are configurable at the Rack level and, for three of them, also at the per-port level. Rack params apply to every NLB listener on that scheme unless a `convox.yml` per-port field overrides them; per-port values layer on top of the Rack defaults without removing them.

All NLB rack parameters can be listed with:

```bash
$ convox rack params -g nlb
```

See [rack params](/reference/cli-commands/rack-params) for other group filters and the `--reveal` flag.

### Ingress allowlist

Traffic reaching an NLB listener is gated by the listener's security group.

- [NLBAllowCIDR](/reference/rack-parameters/NLBAllowCIDR) — comma-delimited list of CIDRs permitted on the public NLB. Default `0.0.0.0/0`. Maximum 5 entries.
- [NLBInternalAllowCIDR](/reference/rack-parameters/NLBInternalAllowCIDR) — same for the internal NLB. Empty (default) means the Rack falls back to the VPC CIDR block. Setting an explicit value **replaces** that fallback, so include your VPC CIDR if you still need in-rack reachability.

Both accept IPv4 CIDRs only. Non-canonical forms (host bits set, e.g. `10.0.0.1/24`) and IPv6 are rejected at `convox rack params set` with a concrete error. Duplicates and leading/trailing whitespace are rejected for the same reason.

```bash
$ convox rack params set NLBAllowCIDR=10.0.0.0/16,192.168.0.0/24
```

Per-port overrides are additive: a Service declaring `allow_cidr:` adds ingress rules for those CIDRs in addition to the rack-level list — the per-port entries never replace the rack-level ones. Omit the field entirely to keep rack-level defaults only (an empty `allow_cidr: []` has the same effect as omitting the field).

To scope a specific listener **narrower** than the rack default, you have to reduce `NLBAllowCIDR` itself, since per-port entries only add to the accepted set and cannot remove from it. A common pattern is a tight rack-level list plus selective per-port broadenings — e.g. `NLBAllowCIDR=10.0.0.0/16` rack-wide, with a single internet-facing listener declaring `allow_cidr: ["0.0.0.0/0"]`.

### Cross-zone load balancing

By default, AWS NLBs only route traffic to targets in the same Availability Zone as the listener. This minimizes cross-AZ data transfer charges but can hotspot a zone if targets are unevenly distributed.

- [NLBCrossZone](/reference/rack-parameters/NLBCrossZone) — set to `Yes` to enable cross-zone load balancing on every public NLB listener. Off by default.
- [NLBInternalCrossZone](/reference/rack-parameters/NLBInternalCrossZone) — same for the internal NLB.

Enabling cross-zone makes AWS bill you for inter-AZ traffic between the listener and targets. Per-port `cross_zone: true` overrides the rack default for a single listener only, letting you enable it surgically for services that need uniform distribution without paying the cost fleet-wide.

### Preserve client IP

By default, NLB target-group traffic is source-NAT'd to the NLB's VPC-internal IP, so application logs record the load balancer's address rather than the real client IP. Enabling preserve-client-IP forwards the real source address.

- [NLBPreserveClientIP](/reference/rack-parameters/NLBPreserveClientIP) — set to `Yes` on the public NLB. Off by default.
- [NLBInternalPreserveClientIP](/reference/rack-parameters/NLBInternalPreserveClientIP) — same for the internal NLB.

When enabled, Convox adds an ingress rule on the ECS instance security group sourced from the NLB security group, allowing traffic from arbitrary client IPs to reach targets. Compliance frameworks that require real client IPs (HIPAA §164.312(b), PCI-DSS 10.2.1) are satisfied by this configuration.

This feature is **incompatible with a customer-supplied [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup)**. Convox cannot modify a security group it does not own — on Racks where `InstanceSecurityGroup` is set, enabling `NLBPreserveClientIP=Yes` is rejected at `rack params set`:

```text
cannot enable NLBPreserveClientIP on a rack with a customer-supplied
InstanceSecurityGroup; your instance SG must add an ingress rule from the
NLB security group (exported as ${Rack}:NLBSecurityGroup) for the NLB
listener ports before this feature can be enabled safely
```

Operators on custom-SG racks must add an ingress rule on their SG sourced from the Rack's NLB security group (exported as `${Rack}:NLBSecurityGroup` or `${Rack}:NLBInternalSecurityGroup`) before enabling preserve-client-IP. The inverse direction is also blocked: setting `InstanceSecurityGroup` while `NLBPreserveClientIP=Yes` is already in force is rejected unless the same call also disables preserve-client-IP.

Per-port `preserve_client_ip: true` is also rejected at release-promote on custom-SG racks.

### Deletion protection

- [NLBDeletionProtection](/reference/rack-parameters/NLBDeletionProtection) — set to `Yes` to block accidental deletion of the public NLB.
- [NLBInternalDeletionProtection](/reference/rack-parameters/NLBInternalDeletionProtection) — same for the internal NLB.

When deletion protection is on, `NLB=No` (or `NLBInternal=No`) and `convox rack uninstall` are rejected pre-flight:

```text
cannot disable NLB while NLBDeletionProtection=Yes; unset protection first,
wait for the update to complete, then toggle NLB off
```

The interlock catches the common pitfall where a Rack update to `NLB=No` succeeds but CloudFormation then fails to delete the protected load balancer, leaving the stack in `UPDATE_ROLLBACK_FAILED`. Disable protection first, wait for the update to complete, then run the disable command in a follow-up call.

## Per-port NLB attributes

Three per-port fields on a `services.nlb:` entry override the Rack-level defaults for that single listener without changing the rack-wide value:

- `cross_zone:` — `true` or `false`. Enables or disables cross-zone load balancing on this listener only.
- `allow_cidr:` — list of IPv4 CIDRs. Adds ingress rules to the scheme's NLB security group scoped to this listener's port. Stacks additively on top of the rack-level [NLBAllowCIDR](/reference/rack-parameters/NLBAllowCIDR) or [NLBInternalAllowCIDR](/reference/rack-parameters/NLBInternalAllowCIDR).
- `preserve_client_ip:` — `true` or `false`. Forwards real client IP for this listener's target group.

Example combining all three:

```yaml
services:
  web:
    image: example/web
    port: 3000/http
    nlb:
      - port: 8443
        containerPort: 3000
        protocol: tls
        scheme: public
        certificate: arn:aws:acm:us-east-1:123456789012:certificate/abcd1234-...
        cross_zone: true
        allow_cidr:
          - 10.0.0.0/24
          - 10.1.0.0/24
        preserve_client_ip: false
```

`convox services` surfaces the per-port overrides in a bracketed suffix:

```bash
$ convox services -a broker
SERVICE  DOMAIN                              PORTS      NLB PORTS
web      web.broker.0a1b2c.convox.cloud      443:3000   8443:3000/tls[cz=true allow=2 pcip=false]
api                                                     9443:8080(internal)[allow=1]
worker
```

Short forms:

- `cz=` — cross-zone override (`true` or `false`)
- `allow=N` — count of per-port `allow_cidr:` entries (not the rack-level total)
- `pcip=` — preserve-client-IP override (`true` or `false`)

A key appears only when the listener explicitly overrides the corresponding rack-level default. Reading `pcip=false`, for example, means this listener has `preserve_client_ip: false` set explicitly; it does not imply the rack-level default is `Yes`. Ports that inherit every rack-level default carry no bracket at all.

See [services.nlb](/application/services#nlb) for the full field reference.

## Known Limitations

### 50 listeners per NLB

The AWS default quota is 50 listeners per load balancer. Because Racks share one public NLB and one internal NLB across all Services, the combined total of `scheme: public` NLB ports (and similarly `scheme: internal`) across all Apps on the Rack cannot exceed 50 without a quota increase.

### Disable procedure

To fully disable NLB on a production Rack, run the following in order. Each step must complete before the next one begins:

1. **Clear deletion protection** if it is on:

   ```bash
   $ convox rack params set NLBDeletionProtection=No NLBInternalDeletionProtection=No
   ```

   Wait for the Rack update to complete (`convox rack` returns to `running`).

2. **Remove every `nlb:` block** from `convox.yml` for every App that declares one and `convox deploy` each. A Rack with at least one App still referencing the NLB rejects the disable:

   ```text
   cannot disable NLB: apps myapp/web still declare public nlb ports;
   remove nlb: from their manifests and redeploy first
   ```

3. **Flip the Rack parameters off**:

   ```bash
   $ convox rack params set NLB=No NLBInternal=No
   ```

The disable releases the EIPs — if you re-enable later, the Rack is assigned new EIPs and a new NLB DNS name. Re-validate any customer DNS pointing at the Rack after a disable/re-enable cycle.

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
