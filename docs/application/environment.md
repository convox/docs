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

```
environment:
  - DEFAULT=true
  - REQUIRED
```

In this example the `DEFAULT` variable has a default value of `true` and the `REQUIRED` variable must be defined before the application will start.

#### Service

You can also configure each service individually to limit secrets to the services that need them:

```
services:
  web:
    build: .
    environment:
      - FOO=bar
      - BAZ
```

#### Wildcard

You can use the wildcard syntax to provide all available environment variables to a service:

```
services:
  web:
    build: .
    environment:
      - "*"
```

In this example the `FOO` variable has a default value of `bar` and the `BAZ` variable must be defined before the application will start.
