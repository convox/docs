---
title: "Services"
description: "Configure application services in convox.yml, including build settings, scaling, environment, ports, health checks, and deployment options."
---

# Services

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

### agent

The `agent` attribute may be used to define that this service should start one container on every instance.

This is useful for services that gather metrics or perform other instance-level behaviors.

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

### build

Configuration options that define the build context and Dockerfile used.

Can be defined as either a string containing a path to use to build this service:

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

### command

Override the default command for this service.

### deployment

Control the ECS minimum and maximum healthy-percent values that govern rolling deployments. Defaults are `minimum: 50` and `maximum: 200` — up to 200% of desired capacity can run during a deploy, and at least 50% must remain healthy. Set both to `100` for strict in-place replacement (never more than desired, never less than desired - 1), or tune up for faster rollouts on services that tolerate extra concurrent containers.

```yaml
services:
  web:
    deployment:
      minimum: 100
      maximum: 200
```

For [agent](#agent) and [singleton](#singleton) services, the defaults are `minimum: 0` and `maximum: 100` — matching the replace-before-start behavior those modes require.

### domain

See [Custom Domains](/deployment/custom-domains)

### drain

Specifies the timeout in seconds during which connections are allowed to drain for a service before terminating during a rolling deploy. Defaults to `30`.

### environment

A list of strings that define the service's environment variables.

A pair like `FOO=bar` creates an environment variable named `FOO` with a default value of `bar`.

Defining a name without a value like `HOST` will require that variable to be set in the application's environment to deploy successfully.

You should not configure secrets here, as they would be recorded in version control. For secrets, specify the variable name, then set the actual value using the CLI `convox env set` command.

Only environment variables that are listed here will be provided to the service at runtime.

### health

See [Health Checks](/application/health-checks)

### image

Use an external Docker image to back this service.

### init

Use a Docker-provided pid1 for intracontainer process management.

### internal

Flag a service as internal, preventing access to it from outside your VPC. Defaults to `false`.

Your Rack must have the `Internal` param set to `Yes` to deploy internal services. You can set it with:

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

### links

Set up links between services on the same app.

#### Example

```yaml
links:
  - web
```

This would add a `WEB_URL` environment variable that points to the load balancer of the `web` service on the same app.

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

- `port` (required, integer 1..65535) — the port the NLB listens on (what clients connect to). Within a single App, port values must be unique across all `nlb:` entries regardless of scheme; across Apps, the same port can be used on the public and internal NLB without conflict (they are separate AWS load balancers). See [Network Load Balancing](/networking/nlb#architecture) for cross-App implications.
- `protocol` — `tcp` (default) or `tls`. `tls` terminates TLS at the NLB and forwards plaintext TCP to the container.
- `containerPort` — optional. The container port that receives traffic from this listener. Defaults to `port` when omitted.
- `scheme` — `public` (routes to the public Rack NLB) or `internal` (routes to the internal Rack NLB). Defaults to `public`.
- `certificate` — required when `protocol: tls`. Accepts a full ACM ARN (`arn:aws:acm:<region>:<account>:certificate/<uuid>`) or an IAM server-certificate ARN (`arn:aws:iam::<account>:server-certificate/<name>`). Must be in the same region and account as the Rack. ACM certificates must be in `ISSUED` state — the release is rejected at promote time otherwise. `convox certs` lists certificates but displays Convox short IDs (`acm-...`) or IAM server-certificate names, not ARNs — copy the full ARN from the AWS Console or `aws acm list-certificates --region <rack-region>`. See [Network Load Balancing](/networking/nlb#tls-termination-at-the-nlb) for the full ARN lookup procedure.

TLS listeners use `ELBSecurityPolicy-TLS13-1-2-2021-06` (TLS 1.2 and 1.3, ECDHE only). The target group protocol stays TCP — backends never see TLS traffic.

Services can combine `nlb:` with `port:` to expose both HTTP (via the ALB) and raw TCP or TLS (via the NLB) on the same `containerPort`. Services with `agent: enabled: true` cannot use `nlb:` — agent mode pins one task per instance, which is incompatible with horizontally scaled NLB target groups.

See [Network Load Balancing](/networking/nlb) for setup, TLS details, and limitations.

### port

Defines the port on which an HTTP service is listening.

If you'd like to use end-to-end encryption, have your application listenin on HTTPS (self-signed certificates are fine) and prefix the port with `https:`

If you'd like to run the GRPC service, then prefix the port with `grpc:` for insecure grpc and `secure-grpc:` for secure grpc

#### Examples

* `port: 3000`
* `port: https:3001`
* `port: grpc:50051`
* `port: secure-grpc:50051`

### policies

A list of ARN of IAM policies to attach to the service's role. It must be created before the service. It will create a new role dedicated to the service, using only the specified policies. Overrides the App's [IamPolicy](/reference/app-parameters/IamPolicy) at the service level.

### privileged

Enabling this parameter results in the container being granted elevated privileges on the host container instance, similar to the root user. If the privileged parameter is set to true for a service to which a timer is linked, the timer container will also be granted privileged access.

### resources

The resources enumerated in the `resources` section that will be available to the service as environment variables. The network endpoint for a resource named `foo` would be `FOO_URL`.

### scale

Set the initial scale parameters for this service. Default CPU is `256` (0.25 vCPU) and default memory is `512` (MiB).

### singleton

Controls deployment behavior. When set to true existing containers for this service will be stopped before new containers are deployed.

### sticky

Toggle load balancer stickiness (using a cookie to keep a user associated with a single container) which helps some applications maintain consistency during rolling deploys. Defaults to `true`.

### termination

Sets the grace period after which a container will be forcefully killed if it does not gracefully exit during a shutdown.  Defaults to 30 seconds.

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
