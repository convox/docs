---
title: convox.yml
order: 50
---

The `convox.yml` file is a manifest used to describe your application and all of its infrastructure needs. This file is similar to a [Docker Compose](https://docs.docker.com/compose/overview/) file but allows for detailed configuration options for your environments.

For a simple Rails application with a Postgres Database the convox.yml file might look like:

```
resources:
  postgres:
    type: postgres
services:
  web:
    build: .
    command: bundle exec rails server -b 0.0.0.0 -p 3000
    port: 3000
    resources:
      - postgres
```

## Accessing Resources

You can access defined resources from services with environment variables.
In the above example, the `postgres` resource provides a `POSTGRES_URL` variable that is accessible from the `web` service.

The environment variable name is the resource name converted to all-caps, with a `_URL` suffix.

This would contain the entire connection string you would need, ie:

```
POSTGRES_URL=postgres://username:password@host.com:5432/databaseName
```

## Complete `convox.yml` options

For a complete set of options available in `convox.yml` can click on the various sections in the template below.

<pre>
<a href="/application/environment">environment</a>:
  - DEFAULT=value
<a href="/application/resources">resources</a>:
  database:
    type: postgres
    options:
      storage: 200
  queue:
    type: redis
<a href="/application/services">services</a>:
  web:
    build: .
    command: bin/web
    <a href="/application/environment">environment</a>:
      - FOO
      - BAR=baz
    <a href="/application/health-checks">health</a>: /health
    internal: true
    <a href="/application/port">port</a>: 3000
    <a href="/application/resources">resources</a>:
      - database
    test: make test
    <a href="/application/volumes">volumes</a>:
      - /tmp/something
  worker:
    build: ./worker
    command: bin/worker
    <a href="/application/environment">environment</a>:
      - FOO
    links:
      - web
    <a href="/application/resources">resources</a>:
      - database
      - queue
  metrics:
    agent: true
    image: awesome/metrics
<a href="/application/timers">timers</a>:
  cleanup:
    schedule: "0 3 * * ? *"
    command: bin/cleanup
    service: web
</pre>
