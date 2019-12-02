---
title: "Environment"
---

Convox applications are configured using environment variables.

<div class="block-callout block-show-callout type-warning" markdown="1">
  Environment variables must be specified to be available to your running application.
</div>

<div class="block-callout block-show-callout type-warning" markdown="1">
  If an environment variable is specified without a default it **must** be defined before the application will start.
</div>

## Definition

#### Global

Environment variables specified in the top level `environment:` section will be available to all services in the application:

```yml
environment:
  - DEFAULT=true
  - STANDARD=
  - REQUIRED
```

In this example the `DEFAULT` variable has a default value of `true`, the `STANDARD` variable will just be empty and the `REQUIRED` variable must be defined before the application will start.

#### Service

You can also configure each service individually to limit secrets to the services that need them:

```yml
services:
  web:
    build: .
    environment:
      - FOO=bar
      - BOO=
      - BAZ
```

In this example the `FOO` variable has a default value of `bar`, `BOO` will just be empty and the `BAZ` variable must be defined before the application will start.

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

You can configure your env vars through the Convox web console for your cloud-based Racks, or through the Convox CLI for all your Racks:

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

<div class="block-callout block-show-callout type-info" markdown="1">
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
</div>

## Environment variable deeper dive

Given this example `convox.yml`:

```yml
environment:
  - DEFAULT=true
  - STANDARD=
  - REQUIRED
services:
  web:
    ...
    environment:
      - WEB_DEFAULT=bar
      - WEB_STANDARD=
      - WEB_REQUIRED
      - SHARED=web
    ...
  worker:
    ...
    environment:
      - WORKER_DEFAULT=foobar
      - WORKER_STANDARD=
      - WORKER_REQUIRED 
      - SHARED=worker
    ... 
  anotherservice:
    ...
    environment:
      - "*"
```

The `REQUIRED`, `WEB_REQUIRED` and `WORKER_REQUIRED` variables will all need to be set before first deployment.  

Once deployed, the `web` service will have the `DEFAULT, STANDARD, REQUIRED, WEB_DEFAULT, WEB_STANDARD, WEB_REQUIRED, SHARED` environment variables all available to it, and `SHARED` will have the value of `web`.
The `worker` service will have the `DEFAULT, STANDARD, REQUIRED, WORKER_DEFAULT, WORKER_STANDARD, WORKER_REQUIRED, SHARED` environment variables all available to it, and `SHARED` will have the value of `worker`.
The `anotherservice` will have `DEFAULT, STANDARD, REQUIRED` environment variables all available to it, as well as `WEB_REQUIRED` and `WORKER_REQUIRED` as I will have to have previously set those before deployment.

If I update any of the `DEFAULT, STANDARD, REQUIRED`, or `SHARED` environment variables and promote that update, those values will then be reflected in the `web` and `worker` services, as well as in `anotherservice`.  If I update any of the `WEB_*` or `WORKER_*` variables and promote that update, those values will then be reflected in their respective services, as well as in `anotherservice`.
If I set a value against a `NEW_VARIABLE` environment variable which wasn't originally declared, and promote that update, it will be reflected and available in the `anotherservice` only.