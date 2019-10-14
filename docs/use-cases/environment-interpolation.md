---
title: "Environment Interpolation"
---

Convox allows you to use environment variable interpolation in your `convox.yml` for the purposes of maintaining the best practice of consistent infrastructure environment configuration.  You can have one version-controlled `convox.yml` across your development, staging, test, production, etc environments, and change any appropriate settings through the use of environment variables.

As your `convox.yml` is version controlled and environment changes are recorded through releases, this allows you to be able to exactly replicate any specific environment setup at will.  Very useful in a Disaster Recovery or debug situation!  

As an example, where we want to inform all services of which environment they are in, we want to change the Database options to be appropriate to the environment, we want to tweak the startup command for our main service, and the domain it sits behind, the value of an env var we are passing to the service itself, and we want to change when our cleanup timer runs depending on the environment:

```

environment:
  - ENVIRONMENT=${GLOBAL_ENVIRONMENT}
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
      - FOOBAR=${FOOBAR}
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

In this case, you can set your env vars for different environments something along the lines of:

```
$ convox env set GLOBAL_ENVIRONMENT=dev DB_STORAGE_SIZE=10 DB_DURABLE=false WEB_COMMAND="bin/web --logging=debug" HOST=mylocal.dev TIMER_SCHEDULE="*/15 * * * *" FOOBAR=devbar --rack local/convox
$ convox env set GLOBAL_ENVIRONMENT=staging DB_STORAGE_SIZE=50 DB_DURABLE=false WEB_COMMAND="bin/web --logging=error --threads 2" HOST=staging.mydomain.com TIMER_SCHEDULE="0 0 * * *" FOOBAR=stagebar --rack acme/staging
$ convox env set GLOBAL_ENVIRONMENT=prod DB_STORAGE_SIZE=200 DB_DURABLE=true WEB_COMMAND="bin/web --logging=warn --threads 10" HOST=www.mydomain.com TIMER_SCHEDULE="0 * * * *" FOOBAR=prodbar --rack acme/production
```

Or use `convox env edit` to edit them interactively.

Documentation around the environment variables that your application code will see and have access to is [here](/application/environment).

