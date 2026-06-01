---
title: "Services"
description: "Configure application services in convox.yml, including build settings, scaling, environment, ports, health checks, and deployment options."
---

# Services

A Service is a single container process in your application. Each Service you define under `services:` in `convox.yml` maps to a container that Convox builds (or pulls), runs, scales, and routes traffic to. A typical App has a `web` service for HTTP traffic plus any number of worker, agent, or internal Services.

This page documents the keys you can set on a Service in gen2 (`convox.yml`), the current and only actively developed generation. gen1 (`docker-compose.yml`) is End-of-Life.

## Definition

```yaml
services:
  web:
    build: .
    command: bin/web
    domain: ${HOST}
    drain: 30
    environment:
      - FOO=bar
      - HOST
    health: /
    image: ubuntu:22.04
    init: true
    internalAndExternal: true
    links:
      - other
    policies:
      - arn:aws:iam::aws:policy/AdministratorAccess
    port: 3000
    resources:
      - database
    scale:
      count: 2
      memory: 1024
      cpu: 512
    tags:
      team: platform
    termination:
      grace: 120
    privileged: true
```

## Image and Build

### build

Configuration options that define the build context and Dockerfile used.

Can be defined as either a string containing a path to use to build this Service:

```yaml
services:
  web:
    build: ./dir
```

or as an object:

```yaml
services:
  web:
    build:
      path: .
      manifest: ./path/to/Dockerfile
```

If you don't specify a build path then `.` is used by default. The default manifest is `Dockerfile`.

### image

Use an external Docker image to back this Service.

### command

Override the default command for this Service.

### init

Use a Docker-provided pid1 for intracontainer process management.

## Networking

### port

Defines the port on which an HTTP Service is listening.

If you'd like to use end-to-end encryption, have your application listenin on HTTPS (self-signed certificates are fine) and prefix the port with `https:`

If you'd like to run the GRPC Service, then prefix the port with `grpc:` for insecure grpc and `secure-grpc:` for secure grpc

#### Examples

* `port: 3000`
* `port: https:3001`
* `port: grpc:50051`
* `port: secure-grpc:50051`

### domain

See [Custom Domains](/deployment/custom-domains)

### internal

Flag a Service as internal, preventing access to it from outside your VPC. Defaults to `false`.

Your Rack must have the `Internal` param set to `Yes` to deploy internal Services. You can set it with:

```bash
$ convox rack params set Internal=Yes
```

### internalAndExternal

When set to `true`, the Service will have both an internal and an external load balancer. This allows the Service to receive traffic from both inside and outside the VPC.

```yaml
services:
  web:
    port: 3000
    internalAndExternal: true
```

Setting `internalAndExternal: true` implicitly enables `internal` on the Service. The primary endpoint uses the internal router, and a second set of listener rules is created on the external router.

Your Rack must have the `Internal` param set to `Yes` for the internal load balancer to be created.

When using `convox services`, the endpoint is displayed as:

```bash
$ convox services
SERVICE  DOMAIN                                                              PORTS
web      host.internal(internal), host.external(external)                    443:3000
```

If the Service has a custom [domain](/deployment/custom-domains), listener rules for that domain are registered on both routers.

### nlb

`nlb` lets a Service expose one or more TCP or TLS ports through the Rack's shared Network Load Balancer(s). Requires the [NLB](/reference/rack-parameters/NLB) or [NLBInternal](/reference/rack-parameters/NLBInternal) rack parameter to be enabled first.

```yaml
services:
  api:
    port: 3000/http
    nlb:
      - port: 443
        protocol: tls
        containerPort: 3000
        scheme: public
        certificate: arn:aws:acm:us-east-1:123456789012:certificate/abcd1234-5678-90ab-cdef-1234567890ab
      - port: 8443
        protocol: tcp
        containerPort: 8443
        scheme: public
      - port: 50051
        protocol: tcp
        containerPort: 50051
        scheme: internal
```

Each list entry is an object with:

- `port` (required, integer 1..65535) sets the port the NLB listens on (what clients connect to). Within a single App, port values must be unique across all `nlb:` entries regardless of scheme; across Apps, the same port can be used on the public and internal NLB without conflict (they are separate AWS load balancers). See [Network Load Balancing](/networking/nlb#architecture) for cross-App implications.
- `protocol` accepts `tcp` (default) or `tls`. `tls` terminates TLS at the NLB and forwards plaintext TCP to the container.
- `containerPort` is optional. The container port that receives traffic from this listener. Defaults to `port` when omitted.
- `scheme` accepts `public` (routes to the public Rack NLB) or `internal` (routes to the internal Rack NLB). Defaults to `public`.
- `certificate` is required when `protocol: tls`. Accepts a full ACM ARN (`arn:aws:acm:<region>:<account>:certificate/<uuid>`) or an IAM server-certificate ARN (`arn:aws:iam::<account>:server-certificate/<name>`). Must be in the same region and account as the Rack. ACM certificates must be in `ISSUED` state, otherwise the release is rejected at promote time. `convox certs` lists certificates but displays Convox short IDs (`acm-...`) or IAM server-certificate names, not ARNs, so copy the full ARN from the AWS Console or `aws acm list-certificates --region <rack-region>`. See [Network Load Balancing](/networking/nlb#tls-termination-at-the-nlb) for the full ARN lookup procedure.
- `cross_zone` is an optional boolean. Override the Rack-level [NLBCrossZone](/reference/rack-parameters/NLBCrossZone) (public) or [NLBInternalCrossZone](/reference/rack-parameters/NLBInternalCrossZone) (internal) default for this listener only. Enabling cross-zone load balancing spreads traffic across targets in every AZ and incurs cross-AZ data transfer charges. See [Cross-zone load balancing](/networking/nlb#cross-zone-load-balancing).
- `allow_cidr` is an optional list of IPv4 CIDRs. Adds ingress rules on the NLB security group scoped to this listener's port. Additive on top of the rack-level [NLBAllowCIDR](/reference/rack-parameters/NLBAllowCIDR) or [NLBInternalAllowCIDR](/reference/rack-parameters/NLBInternalAllowCIDR). Non-canonical CIDRs (host bits set), IPv6, duplicates, and leading/trailing whitespace are rejected at manifest parse time. See [Ingress allowlist](/networking/nlb#ingress-allowlist).
- `preserve_client_ip` is an optional boolean. Override the Rack-level [NLBPreserveClientIP](/reference/rack-parameters/NLBPreserveClientIP) or [NLBInternalPreserveClientIP](/reference/rack-parameters/NLBInternalPreserveClientIP) default for this listener only. Forwards the real client source IP to the target task instead of the NLB's VPC-internal IP. Rejected at release promote on Racks that use a customer-supplied [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup) (the required NLB-source ingress rule cannot be added to an SG Convox does not own). See [Preserve client IP](/networking/nlb#preserve-client-ip).

TLS listeners use `ELBSecurityPolicy-TLS13-1-2-2021-06` (TLS 1.2 and 1.3, ECDHE only). The target group protocol stays TCP, so backends never see TLS traffic.

Services can combine `nlb:` with `port:` to expose both HTTP (via the ALB) and raw TCP or TLS (via the NLB) on the same `containerPort`. Services with `agent: enabled: true` cannot use `nlb:` because agent mode pins one task per instance, which is incompatible with horizontally scaled NLB target groups.

See [Network Load Balancing](/networking/nlb) for setup, TLS details, per-port attribute semantics, and limitations.

### sticky

Toggle load balancer stickiness (using a cookie to keep a user associated with a single container) which helps some applications maintain consistency during rolling deploys. Defaults to `true`.

## Scaling and Resources

### scale

Set the initial scale parameters for this Service. Default CPU is `256` (0.25 vCPU) and default memory is `512` (MiB).

### agent

The `agent` attribute may be used to define that this Service should start one container on every instance.

This is useful for Services that gather metrics or perform other instance-level behaviors.

You can use this attribute in one of two format:

```yaml
services:
  monitor:
    agent: true
```

or if your agent needs to open host-level ports then use this format:

```yaml
services:
  datadog:
    agent:
      enabled: true
      ports:
        - 8125/udp
        - 8126/tcp
```

### singleton

Controls deployment behavior. When set to true existing containers for this Service will be stopped before new containers are deployed.

## Environment and Links

### environment

A list of strings that define the Service's environment variables.

A pair like `FOO=bar` creates an environment variable named `FOO` with a default value of `bar`.

Defining a name without a value like `HOST` will require that variable to be set in the application's environment to deploy successfully.

You should not configure secrets here, as they would be recorded in version control. For secrets, specify the variable name, then set the actual value using the CLI `convox env set` command.

Only environment variables that are listed here will be provided to the Service at runtime.

### links

Set up links between Services on the same App.

Declaring `links: [<service>]` auto-injects a `<SERVICE>_URL` environment variable into this Service. The variable name is the linked service name uppercased with hyphens converted to underscores (for example `api` becomes `API_URL`). Its value is the linked Service's load-balancer URL `https://<app>-<service>.<RouterHost>`, or `<RouterInternalHost>` when the linked Service is internal. This is distinct from [resources](#resources), which inject `<NAME>_URL`, `_USER`, `_PASS`, `_HOST`, `_PORT`, and `_NAME`.

#### Example

```yaml
links:
  - web
```

This would add a `WEB_URL` environment variable that points to the load balancer of the `web` service on the same App.

### resources

The resources enumerated in the `resources` section that will be available to the Service as environment variables. For a resource named `foo`, Convox injects `FOO_URL`, `FOO_USER`, `FOO_PASS`, `FOO_HOST`, `FOO_PORT`, and `FOO_NAME` (the name is uppercased with hyphens converted to underscores). The network endpoint is `FOO_URL`.

## Health and Lifecycle

### health

See [Health Checks](/application/health-checks)

### deployment

Control the ECS minimum and maximum healthy-percent values that govern rolling deployments. Defaults are `minimum: 50` and `maximum: 200`, so up to 200% of desired capacity can run during a deploy, and at least 50% must remain healthy. Set both to `100` for strict in-place replacement (never more than desired, never less than desired - 1), or tune up for faster rollouts on Services that tolerate extra concurrent containers.

```yaml
services:
  web:
    deployment:
      minimum: 100
      maximum: 200
```

For [agent](#agent) and [singleton](#singleton) Services, the defaults are `minimum: 0` and `maximum: 100`, matching the replace-before-start behavior those modes require.

### drain

Specifies the timeout in seconds during which connections are allowed to drain for a Service before terminating during a rolling deploy. Defaults to `30`.

### termination

Sets the grace period after which a container will be forcefully killed if it does not gracefully exit during a shutdown.  Defaults to 30 seconds.

## Security and IAM

### policies

A list of ARN of IAM policies to attach to the Service's role. It must be created before the Service. It will create a new role dedicated to the Service, using only the specified policies. Overrides the App's [IamPolicy](/reference/app-parameters/IamPolicy) at the service level.

### privileged

Enabling this parameter results in the container being granted elevated privileges on the host container instance, similar to the root user. If the privileged parameter is set to true for a Service to which a timer is linked, the timer container will also be granted privileged access.

## Other

### tags

A map of key-value pairs to apply as AWS tags to the Service's resources.

```yaml
services:
  web:
    tags:
      team: platform
      environment: production
```

### test

Defines a command to be used when running `convox test` against an application.

### volumes

See [Volumes](/application/volumes)

## See Also

- [convox.yml Reference](/application/convox-yml)
- [Health Checks](/application/health-checks)
- [Resources](/application/resources)
- [Scaling](/scaling/scaling)
- [Service Tags](/management/service-tags)
