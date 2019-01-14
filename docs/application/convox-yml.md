---
title: convox.yml
order: 50
---

The `convox.yml` file is a configuration file used to describe your application and all of its infrastructure needs.

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
