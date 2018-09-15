---
title: "Docker Compose File"
---

[Docker Compose](https://docs.docker.com/compose/overview/) makes it easier to configure and run applications made up of multiple containers. For the uninitiated, imagine being able to define three containers&mdash;one running a web app, another running postgres, and a third running redis&mdash;all in one YAML file and then running those three connected containers with a single command. That file is `docker-compose.yml`, and when using Convox locally, the command is `convox start`.

## Service Names

Please note that service names should not include underscores.

    services:
      foo_bar: # will not work
        ...

Dashed service names are allowed.

    services:
      foo-bar: # will work
        ...

## Supported Docker Compose Configuration Options

Though there are [50+ configuration options](https://docs.docker.com/compose/compose-file/) supported by Docker Compose, Convox curates a smaller list of options to keep configuration simple and to ensure that the majority of users can get their multi-container applications up and running quickly.

Select a key for more information and example usage.

<pre>
  version: '2'
  services:
    web:
      <a href="#build">build</a>: .
      <a href="#build">build</a>:
        <a href="#context">context</a>: .
        <a href="#dockerfile">dockerfile</a>: Dockerfile.alternate
      <a href="#command">command</a>: bin/web
      <a href="#cpu-shares">cpu_shares</a>: 73
      <a href="#entrypoint">entrypoint</a>: /bin/entrypoint
      <a href="#environment">environment</a>:
        - RACK_ENV=development
        - SECRET_KEY
        - FOO=
      <a href="#image">image</a>: convox/rails
      <a href="#labels">labels</a>:
        - convox.port.443.protocol=tls
        - convox.port.443.proxy=true
      <a href="#links">links</a>:
        - database
      <a href="#memory-limit">mem_limit</a>: 1024MB
      <a href="#ports">ports</a>:
        - 80:4000
        - 443:4001
      <a href="#privileged">privileged</a>: true
    database:
      <a href="#image">image</a>: convox/postgres
      <a href="#ports">ports</a>:
        - 5432
      <a href="#volumes">volumes</a>:
        - /var/lib/postgresql/data
  <a href="#networks">networks</a>:
    outside:
      external:
        name: foo-bar
</pre>

<div class="block-callout block-show-callout type-info" markdown="1">
  There are two versions of the Docker Compose file format:<br />
  <br />
  <strong>Version 1</strong> is specified by omitting a version key at the root of the YAML.<br />
  <strong>Version 2</strong> is specified with a <code class="highlighter-rouge">version: '2'</code> entry at the root of the YAML.<br />
  <br />
  Convox recommends using the version 2 format, but supports the legacy version 1 format as well.
</div>

### Build

Specify the path to the Dockerfile.

    build: .

    build: ./another/dir

If you are using the Docker Compose v2 file format, you may also need to use the nested build configuration options `context` and `dockerfile` when using a Dockerfile with a non-standard name. This nested structure is not supported by the Docker Compose v1 file format.

    build:
      context: .
      dockerfile: Dockerfile.alternate

### Command

Override the default command.

    command: bin/web

### Context

Supported by the Docker Compose v2 file format only. Specify a path to a directory containing a Dockerfile. Must be nested under the `build:` directive.

    build:
      context: .

In the v1 and v2 file formats, the following syntax is equivalent to the nested example above.

    build: .

The `context` directive is required if you are using the Docker Compose v2 file format and specifying a non-standard name for a Dockerfile with the `dockerfile:` directive.

    build:
      context: .
      dockerfile: Dockerfile.alternate

### CPU Shares

CPU shares (relative weight).

Corresponds to the `-c` or `--cpu-shares` flag of the `docker run` command.

This can be specified for the initial deployment in `docker-compose.yml`:

    cpu_shares: 512

Or by running `convox scale` (see also [Scaling](/docs/gen1/scaling/)):

    convox scale <process> --cpu=512

The default Convox value is `0`. When set to `0`, Docker will ignore the value and use the default of `1024` instead.

Note:

- With `convox deploy`: `cpu_shares` is read and taken into account from `docker-compose.yml` only during the initial deployment. If you need to change it after deploying for the first time, use [`convox scale`](/docs/gen1/scaling/).
- With `convox start`: When running locally via `convox start`, `cpu_shares` is always read from `docker-compose.yml`.

You can view the current CPU shares allocation with either of the following commands:

- `convox api get /apps/<app name>/formation`
- `convox scale <process name>`

See also [CPU share constraint](https://docs.docker.com/engine/reference/run/#/cpu-share-constraint "docs.docker.com") in the Docker documentation.

### Dockerfile

Specify an alternate name if not named `Dockerfile`. Note that the Docker Compose v1 file format differs from v2.

Docker Compose v1 file format:

    build: .
    dockerfile: Dockerfile.alternate

Docker Compose v2 file format:

    build:
      context: .
      dockerfile: Dockerfile.alternate

### Entrypoint

Override the default entrypoint.

    entrypoint: /bin/entrypoint

### Environment

Set environment variables, or allow them to be set, when the container is started.

    environment:
      - RACK_ENV=development
      - SECRET_KEY
      - FOO=

Variables can be:

- **defined directly** in the Compose file, as `RACK_ENV` above (can be optionally overridden by `.env`)
- **required, but not defined** in the Compose file, as `SECRET_KEY` above (must therefore be defined in `.env`)
- **set to an empty value**, as `FOO` above (can be optionally overridden by `.env`)

See our [environment documentation](/docs/gen1/environment) for more.

### Image

Specify the image used when starting the container.

    image: postgres
    image: ubuntu:16.04
    image: convox/rails

### Labels

Add metadata to containers using Docker labels. Convox has several custom labels, which are listed below.

    labels:
      - convox.cron.<task name>
      - convox.deployment.maximum
      - convox.deployment.minimum
      - convox.health.path
      - convox.health.port
      - convox.health.timeout
      - convox.idle.timeout
      - convox.port.<number>.protocol
      - convox.port.<number>.proxy
      - convox.port.<number>.secure
      - convox.start.shift

For more details about how to use these labels, visit our [Docker Compose Labels doc](/docs/gen1/docker-compose-labels).

### Links

Specifying a link to another container instructs Convox to provide the linking container with environment variables that allow it to connect to the target container. In the example below, Convox would set `DATABASE_URL=postgres://postgres:password@172.17.0.1:5432/app` in the `web` container, allowing an application there to reference `DATABASE_URL` to easily connect to the Postgres server running in the `database` container.

    services:
      web:
        build: .
        links:
          - database
      database:
        image: convox/postgres

See our [linking documentation](/docs/gen1/linking) for more.

### Networks

Convox currently supports a single use (via `convox start`) of the `networks` configuration option: connecting all services found in your docker-compose.yml to an existing external user-defined network. In the example below, services would be connected to the network named `foo-bar`. To be more specific, the container of each service would be started with a `--net foo-bar` option being passed to `docker run`.

    networks:
      outside:
        external:
          name: foo-bar

### Memory Limit

Amount of memory, in bytes, available to the specified process type.

This can be specified for the initial deployment in `docker-compose.yml`:

    mem_limit: 256MB

Or by running `convox scale` (see also [Scaling](/docs/gen1/scaling/)):

    convox scale <process> --memory=256

Minimum is 4 MB. If no unit label (e.g. `MB`) is provided, the amount is interpreted in bytes.

Note:

- With `convox deploy`: `mem_limit` is read and taken into account from `docker-compose.yml` only during the initial deployment. If you need to change it after deploying for the first time, use [`convox scale`](/docs/gen1/scaling/).
- With `convox start`: When running locally via `convox start`, `mem_limit` is always read from `docker-compose.yml`.

You can view the current memory limit with either of the following commands:

- `convox api get /apps/<app name>/formation`
- `convox scale <process name>`

### Ports

Define the ports on which the process should listen.

    ports:
      - 5000
      - 80:5000

See our [port mapping documentation](/docs/gen1/port-mapping) for more.

### Privileged

Give extended privileges to the container, including access to host devices.

    privileged: true

### Volumes

Share data between Processes of the same type by mounting volumes from the host or a network filesystem (EFS) inside the container.

    volumes:
      - /var/lib/postgresql/data

See our [volumes documentation](/docs/volumes) for more.

## Missing Configuration Options

Didn't find what you were looking for? If you encounter critical Docker Compose configuration options that we have not implemented, we invite you to [submit an issue](https://github.com/convox/rack/issues) describing your use case.
