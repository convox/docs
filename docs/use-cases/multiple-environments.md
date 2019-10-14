---
title: "Multiple Environments"
---

Convox allows you to use environment variable interpolation in your `convox.yml` for the purposes of maintaining the best practice of consistent infrastructure environment configuration.  You can have one version-controlled `convox.yml` across your development, staging, test, production, etc environments, and change any appropriate settings through the use of environment variables.

As your `convox.yml` is version controlled and environment changes are recorded through releases, this allows you to be able to exactly replicate any specific environment setup at will.  Very useful in a Disaster Recovery or debug situation!  

As an example, where we want to inform all services of which environment they are in, we want to change the Database options to be appropriate to the environment, we want to tweak the startup command for our main service, and the domain it sits behind, the value of an env var we are passing to the service itself, and we want to change when our cleanup timer runs depending on the environment:

```
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

As we must have a schedule for the timer, and command for our service, we have added those to the environment list as required variables to ensure they are set.

In this case, you can set your env vars for different environments something along the lines of:

```
$ convox env set ENVIRONMENT=dev WEB_COMMAND="bin/web --logging=debug" HOST=mylocal.dev TIMER_SCHEDULE="*/15 * * * *" FOO=devfoo BAR=devbar FOOBAR=devfoobar --rack local/convox
```

We are creating a smaller database just using the default values, starting our service with a lot of logging for dev purposes, using a dev domain, and running our cleanup every 15 minutes.  For the `web` service, we are specifying more env vars, `FOO` is required, `BAR` has a default value which we are overriding here, and `FOOBAR` would be set to an empty string unless we pass in our value as we are here.

```
$ convox env set ENVIRONMENT=staging DB_STORAGE_SIZE=50 WEB_COMMAND="bin/web --logging=error --threads 2" HOST=staging.mydomain.com TIMER_SCHEDULE="0 0 * * *" FOO=stagefoo FOOBAR=stagefoobar --rack acme/staging
```

We are creating a slightly larger database, starting our service with error logging and a couple of threads, using our staging domain, and running our cleanup once a day.  For the `web` service, we are specifying more env vars, `FOO` is required, `BAR` has a default value which we are leaving as is, and `FOOBAR` would be set to an empty string unless we pass in our value as we are here.

```
$ convox env set ENVIRONMENT=prod DB_STORAGE_SIZE=200 DB_DURABLE=true WEB_COMMAND="bin/web --logging=warn --threads 10" HOST=www.mydomain.com TIMER_SCHEDULE="0 * * * *" FOO=prodfoo FOOBAR=prodbar --rack acme/production
```

We are creating a large, robust database, starting our service with appropriate logging and thread capacity for production usage, using our actual domain, and running our cleanup once an hour.  For the `web` service, we are specifying more env vars, `FOO` is required, `BAR` has a default value which we are leaving as is, and `FOOBAR` would be set to an empty string unless we pass in our value as we are here.

You can also use `convox env edit` to edit the variables interactively.

<div class="block-callout block-show-callout type-info" markdown="1">
If you don't set a variable to be interpolated, it will simply be treated as an empty string when it comes to injecting them into the `convox.yml`.  
</div>

Documentation around the environment variables that your application code will see and have access to is [here](/application/environment).

