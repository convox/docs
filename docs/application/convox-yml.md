---
title: convox.yml
description: The application manifest that defines Services, Resources, Environment, Timers, and other infrastructure for a Convox App.
---

# convox.yml

The `convox.yml` file is the application manifest that describes your App and its infrastructure. Convox reads this file to build, configure, and deploy your application on AWS.

## Top-Level Sections

A `convox.yml` file supports five top-level sections:

| Section | Purpose |
|---------|---------|
| [environment](/application/environment) | Global environment variables available to all Services |
| [params](/reference/app-parameters) | App-level parameters that control infrastructure behavior |
| [resources](/application/resources) | Backing infrastructure such as databases and caches |
| [services](/application/services) | The containers that make up your application |
| [timers](/application/timers) | Scheduled tasks that run on a cron schedule |

Service-level configuration also supports [health checks](/application/health-checks), [port](/application/port) mapping, and [volumes](/application/volumes).

## Example

A web application with a Postgres database:

```yaml
environment:
  - RACK_ENV=production
resources:
  database:
    type: postgres
services:
  web:
    build: .
    command: bundle exec rails server -b 0.0.0.0 -p 3000
    environment:
      - SECRET_KEY_BASE
    health: /health
    port: 3000
    resources:
      - database
  worker:
    build: .
    command: bin/worker
    resources:
      - database
timers:
  cleanup:
    schedule: "0 3 * * ? *"
    command: bin/cleanup
    service: web
```

## Accessing Resources

When a Service lists a Resource in its `resources` array, Convox injects a connection URL as an environment variable. The variable name is the resource name converted to uppercase with a `_URL` suffix.

For the example above, the `database` resource provides:

```text
DATABASE_URL=postgres://username:password@host.com:5432/databaseName
```

## See Also

- [Services](/application/services)
- [Resources](/application/resources)
- [Environment](/application/environment)
- [Multiple Environments](/application/multiple-environments)
- [App Parameters](/reference/app-parameters)
