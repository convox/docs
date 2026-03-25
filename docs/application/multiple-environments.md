---
title: "Multiple Environments"
description: "Use environment variable interpolation in convox.yml to manage dev, staging, and production environments from a single configuration file."
---

# Multiple Environments

Convox allows you to use [environment variable](/application/environment) interpolation in your `convox.yml` to maintain consistent infrastructure configuration across development, staging, production, and other environments. A single version-controlled `convox.yml` adapts to each environment through environment variables set per Rack.

Because `convox.yml` is version controlled and environment changes are recorded through [Releases](/deployment/releases), you can replicate any specific environment setup at will — useful for disaster recovery or debugging.

## Example Configuration

This `convox.yml` uses interpolation to vary the database size, service command, domain, and timer schedule by environment:

```yaml
environment:
  - ENVIRONMENT
  - TIMER_SCHEDULE
  - WEB_COMMAND
resources:
  database:
    type: postgres
    options:
      storage: ${DB_STORAGE_SIZE}
      durable: ${DB_DURABLE}
services:
  web:
    build: .
    command: ${WEB_COMMAND}
    domain: ${HOST}
    environment:
      - FOO
      - BAR=baz
      - FOOBAR=
    health: /health
    port: 3000
    resources:
      - database
timers:
  cleanup:
    schedule: ${TIMER_SCHEDULE}
    command: bin/cleanup
    service: web
```

Variables listed in the top-level `environment` section (`TIMER_SCHEDULE`, `WEB_COMMAND`) are required — the App will not start unless they are set.

## Setting Variables Per Environment

**Development** — small database (defaults), verbose logging, frequent cleanup:

```bash
$ convox env set ENVIRONMENT=dev WEB_COMMAND="bin/web --logging=debug" HOST=mylocal.dev TIMER_SCHEDULE="*/15 * * * *" FOO=devfoo BAR=devbar FOOBAR=devfoobar --rack local/convox
```

**Staging** — moderate database, error-level logging, daily cleanup:

```bash
$ convox env set ENVIRONMENT=staging DB_STORAGE_SIZE=50 WEB_COMMAND="bin/web --logging=error --threads 2" HOST=staging.mydomain.com TIMER_SCHEDULE="0 0 * * *" FOO=stagefoo FOOBAR=stagefoobar --rack acme/staging
```

**Production** — large durable database, production logging, hourly cleanup:

```bash
$ convox env set ENVIRONMENT=prod DB_STORAGE_SIZE=200 DB_DURABLE=true WEB_COMMAND="bin/web --logging=warn --threads 10" HOST=www.mydomain.com TIMER_SCHEDULE="0 * * * *" FOO=prodfoo FOOBAR=prodbar --rack acme/production
```

You can also use `convox env edit` to update variables interactively.

> If an interpolated variable is not set, it is treated as an empty string when injected into `convox.yml`.

## See Also

- [Environment](/application/environment) — environment variable definition, scoping, and promotion
- [Resources](/application/resources) — dynamic resource configuration with interpolation
- [Releases](/deployment/releases) — how environment changes create and promote Releases
