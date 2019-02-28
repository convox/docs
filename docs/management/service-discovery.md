---
title: "Service Discovery"
---

Convox makes it easy to build microservice architectures by enabling easy communication between services and applications.

<div class="block-callout block-show-callout type-info" markdown="1">
If you would like to make a service accessible only to other services on the same Rack, see [Internal Services](/deployment/internal-services).
</div>

### Between Services in the Same App

You can link services in the same app using the [links](/application/services#links) section of [convox.yml](/application/convox-yml).

#### Example

```
services:
  web:
    port: 3000
  worker:
    links:
      - web
```

This would set an `WEB_URL` environment variable on the `worker` service pointing at the load balancer for the `web` service.

### Between Apps on the Same Rack

Convox sets up internal DNS on a Rack such that the following hostname format will resolve to a specific service:

```
[service].[app].[rack].convox
```

<div class="block-callout block-show-callout type-info" markdown="1">
The Rack's name is available as the environment variable `RACK`.
</div>

#### Example

```
$ curl https://api.auth.$RACK.convox
```

This hostname would resolve to the load balancer of the `api` service of the `auth` app on the current Rack.

### Between Apps on Different Racks

Use `convox services` to find the load balancer hostname of a given service and set an environment variable with the resulting hostname.

#### Example

```
$ convox services -a auth -r acme/utility
SERVICE  DOMAIN                                 PORTS
api      auth-api.router.us-east-1.convox.site  80:3000 443:3000

$ convox env set AUTH_URL=https://auth-api.router.us-east-1.convox.site -a frontend -r acme/production
```

This would set an `AUTH_URL` environment variable on the `frotend` app on the `acme/production` Rack to point at the load balancer for the `api` service on the `auth` app on the `acme/utility` Rack.
