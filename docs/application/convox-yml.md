---
title: convox.yml
order: 50
---

The `convox.yml` file is a manifest used to describe your application and all of its infrastructure needs. This file is similar to a [Docker Compose](https://docs.docker.com/compose/overview/) file but allows for detailed configuration options for your environments.

For a simple Rails application with a Postgres Database the convox.yml file might look like:
```
resources:
  database:
    type: postgres
services:
  web:
    build: .
    command: bundle exec rails server -b 0.0.0.0 -p 3000
    port: 3000
    resources:
      - database
```
The convox.yml file also support environment variable interpolation. This allows you to specify things like different domains or database options depending on which environment you deploy you application to. For example if you want to run different instance sizes for your database between staging and production. You can setup your convox.yml like

```
resources:
  database:
    type: postgres
    options:
      class: ${POSTGRES_INSTANCE_SIZE} 
``` 
You can then specify a different value for `POSTGRES_INSTANCE_SIZE` in your staging and production Racks using `convox env set POSTGRES_INSTANCE_SIZE=XXXX` and when your application is deployed to that Rack it will use that value to size the database.

For a complete set of options available in convox.yml can click on the various sections in the template below.

<pre>
<a href="/application/environment">environment</a>:
  - DEFAULT=value
<a href="/application/resources">resources</a>:
  database:
    type: postgres
    options:
      storage: 200
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
  worker:
    build: ./worker
    command: bin/worker
    <a href="/application/environment">environment</a>:
      - FOO
    links:
      - web
    <a href="/application/resources">resources</a>:
      - database
  metrics:
    agent: true
    image: awesome/metrics
<a href="/application/timers">timers</a>:
  cleanup:
    schedule: "0 3 * * ? *"
    command: bin/cleanup
    service: web
</pre>
