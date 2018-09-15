---
title: "Environment"
---

Convox applications are configured using environment variables. Environment variables must be specified _both_ at the individual service level or on the application level for all services, for more information see the [convox.yml](/docs/convox-yml) reference.

### `convox.yml`

The required environment variables for your service are defined like this:

```
services:
  web:
    build: .
    environment:
      - FOO=bar
      - BAZ
```

In this example the `FOO` variable has a default value of `bar` and the `BAZ` variable must be defined before the application can be deployed.

<div class="block-callout block-show-callout type-warning" markdown="1">
  Only environment variables that are listed will be provided to the service at runtime.
</div>

<div class="block-callout block-show-callout type-warning" markdown="1">
  You can provide all available environment variables to a service by specifying `- "*"`
</div>
