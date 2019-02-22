---
title: Services
---

## Definition

```yaml
services:
  web:
    build: .
    command: bin/web
    domain: ${HOST}
    environment:
      - FOO=bar
      - HOST
    health: /health
    image: ubuntu:16.04
    init: true
    port: 3000
    resources:
      - database
    scale:
      count: 2
      memory: 1024
      cpu: 512
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

If you don't specify a build path then `.` is used by default.

### command

Override the default command for this service.

### domain

See [Custom Domains](/deployment/custom-domains)

### drain

Specifies the timeout in seconds during which connections are allowed to drain for a service before terminating during a rolling deploy.

### environment

A list of strings that define the service's environment variables.

A pair like `FOO=bar` creates an environment variable named `FOO` with a default value of `bar`.

Defining a name without a value like `HOST` will require that variable to be set in the application's environment to deploy successfully.

You should not configure secrets here, as they would be recorded in version control. For secrets, simply specify the variable name, then set the actual value using the CLI `convox env set` command.

Only environment variables that are listed here will be provided to the service at runtime.

### health

See [Health Checks](/application/health-checks)

### image

Use an external Docker image to back this service.

### init

Use a Docker-provided pid1 for intracontainer process management.

### internal

Flag a service as internal, preventing access to it from outside your VPC. Defaults to `false`.

Your rack must have the `Internal` param set to Yes to deploy internal services. You can set it with:

```shell
$ convox rack params set Internal=Yes
```

### port

Defines the port on which an HTTP service is listening.

If you'd like to use end-to-end encryption, have your application listenin on HTTPS (self-signed certificates are fine) and prefix the port with `https:`

#### Examples

* `port: 3000`
* `port: https:3001`

### resources

The resources enumerated in the `resources` section that will be available to the service as environment variables. The network endpoint for a resource named `foo` would be `FOO_URL`.

### scale

Set the initial scale parameters for this service.

### singleton

Controls deployment behavior. When set to true existing containers for this service will be stopped before new containers are deployed.

### sticky

Toggle load balancer stickiness (using a cookie to keep a user associated with a single container) which helps some applications maintain consistency during rolling deploys. Defaults to `true`.

### test

Defines a command to be used when running `convox test` against an application.

### volumes

See [Volumes](/application/volumes)
