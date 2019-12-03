---
title: "Environment"
---

Convox applications are configured using environment variables.  You will need to [specify](#definition) your variables in your `convox.yml` to be available to your running application.  You can [set and unset](#setting-and-editing-your-environment-variables) them from the Convox CLI or web Console before or after any code deployments.  You can even control the [release](#promoting-your-environment-variable-changes) of your environmental changes.  The variables are stored securely and encrypted and only available to the specific application and services you require.

## Definition

All environment variables defined in your `convox.yml` are required to start up your application, but you can easily set a default or empty value in your configuration for those that are appropriate to do so.

#### Global

Environment variables specified in the top level `environment:` section will be available to all services in the application:

```yml
environment:
  - ENABLED=true
  - LICENSE=
  - ENVIRONMENT
...
```

In this example the `ENABLED` variable has a default value of `true`, the `LICENSE` variable will just be empty and the `ENVIRONMENT` variable must be defined using the CLI or Web console before the application will start.

#### Service

You can also configure each service individually to limit secrets to the services that need them:

```yml
services:
  web:
    build: .
    environment:
      - LOG_LEVEL=debug
      - DEFAULT_CLIENT=
      - MULTI_TENANT
```

In this example the `LOG_LEVEL` variable has a default value of `debug`, `DEFAULT_CLIENT` will just be empty and the `MULTI_TENANT` variable must be defined using the CLI or Web console before the application will start.

Environment variables defined at the service level will not be shared or visible to other services in the same app.  You can have a variable with the same name in separate services with different declared values, but note that updating this variable after deployment will update the value in all services with that variable name.  

#### Wildcard

You can use the wildcard syntax to provide all available environment variables to a service:

```yml
services:
  web:
    build: .
    environment:
      - "*"
```

If you wish to declare and inject new environment variables on demand without declaring them in the `convox.yml` beforehand, you should use the wildcard syntax to pick them up.

## Setting and Editing your Environment Variables

You can configure your env vars through the Convox web console for your cloud-based Racks, or through the Convox CLI your local and cloud Racks:

```sh
$ convox env set FOO=bar
Setting FOO... OK
Release: RFMGESPZHC

$ convox env get FOO
bar

$ convox env unset FOO
Unsetting FOO... OK
Release: RZIDXFIJAC

$ convox env edit
Setting ... OK
Release: RRJQTMVKRS
```

`convox env edit` allows you to interactively update your environment variables in a terminal editor 😊

## Promoting your environment variable changes

If you update your env vars through the CLI, that will create a new release containing your changes.  By default this will not be promoted.
For this to take effect, you will need to promote that release:

```sh
$ convox env set FOO=bar
Setting FOO... OK
Release: RFMGESPZHC
$ convox releases promote RFMGESPZHC
Promoting RFMGESPZHC... OK
```

Alternatively you can pass the `--promote`/`-p` flag to the env commands to automatically promote the release:

```sh
$ convox env set FOO=bar --promote
Setting FOO... OK
Release: RYFPMIHLPKD
Promoting RYFPMIHLPKD... OK
```

Environment variable updates made through the Convox web console create releases which are promoted automatically for you.

## Environment variable comprehensive example

Given this example `convox.yml`:

```yml
environment:
  - ENABLED=true
  - LICENSE=
  - ENVIRONMENT
services:
  web:
    ...
    environment:
      - WEB_LOGGING_LEVEL=info
      - WEB_CLIENT=
      - WEB_MULTI_TENANT
      - SHARED=web
    ...
  worker:
    ...
    environment:
      - WORKER_LOGGING_LEVEL=error
      - WORKER_CLIENT=
      - WORKER_MULTI_TENANT
      - SHARED=worker
    ...
  anotherservice:
    ...
    environment:
      - ANOTHER_LOGGING_LEVEL=trace
      - "*"
```

The `ENVIRONMENT`, `WEB_MULTI_TENANT` and `WORKER_MULTI_TENANT` variables will all need to be set before first deployment.  

Once deployed, the `web` service will have the `ENABLED, LICENSE, ENVIRONMENT, WEB_LOGGING_LEVEL, WEB_CLIENT, WEB_MULTI_TENANT, SHARED` environment variables all available to it, and `SHARED` will have the value of `web`.
The `worker` service will have the `ENABLED, LICENSE, ENVIRONMENT, WORKER_LOGGING_LEVEL, WORKER_CLIENT, WORKER_MULTI_TENANT, SHARED` environment variables all available to it, and `SHARED` will have the value of `worker`.
The `anotherservice` will have `ENABLED, LICENSE, ENVIRONMENT, ANOTHER_LOGGING_LEVEL` environment variables all available to it, as well as `WEB_MULTI_TENANT` and `WORKER_MULTI_TENANT` as I will have to have previously set those before deployment.

If I update any of the `ENABLED, LICENSE, ENVIRONMENT`, or `SHARED` environment variables and promote that update, those values will then be reflected in the `web` and `worker` services, as well as in `anotherservice`.  If I update any of the `WEB_*` or `WORKER_*` variables and promote that update, those values will then be reflected in their respective services, as well as in `anotherservice`.
If I set a value against a `NEW_VARIABLE` environment variable which wasn't originally declared, and promote that update, it will be reflected and available in the `anotherservice` only.  If I want that to be available to the other services, I should add that declaration to the `convox.yml` and redeploy, at which point they will then pick that up.
