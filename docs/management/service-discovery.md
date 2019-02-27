---
title: "Service Discovery"
---

Convox makes it easy to connect your services.

#### Linking Services Within a Single App

If you have multiple services defined in your application you can link those services by defining a `links` section in your [convox.yml](/application/convox-yml).
For example if you have a service called `web` and a service called `api` and you would like `web` to be able to call `api` your convox.yml might look like:

```
services:
  web:
    build: .
    port: 3000
    resources:
      - database
    links:
      - api
  api:
    build: .
    port: 3000
```

This will cause Convox to automatically inject an environment variable called `API_URL` in all containers running the `web` service.

#### Communicating Between Apps 

All services can be accessed by other services running on the same Rack via a URL with the following format
```
[servicename].[appname].[rackname].convox
```
For example if we have a Rack named `production` with an app named `myapp` and a service named `api` the URL would be:
```
https://api.myapp.production.convox
```

We recommend that you store this URL as an [environment variable](/application/environment) which will then be made available to all the apps in your Rack.

This URL is only accessible by other services running on the same Rack. If you need to access a service from another Rack, or other external source, you will need to use the public URL available with `convox services`
    
#### Internal Services

If you have a service that you want to be internal only, meaning it is only accessible by other services running in the same Rack, you can create an [internal service](/deployment/internal-services)

First make sure your Rack has the internal load balancer enabled with 
```
$ convox rack params set Internal=Yes
```
and then modify the service definition in your convox.yml by adding `internal: true`

ex:
```
services:
  web:
    build: .
    port: 3000
    resources:
      - database
    links:
      - api
  api:
    build: .
    port: 3000
    internal: true
```

<div class="block-callout block-show-callout type-warning" markdown="1">
Note: Services that are marked as internal are only accesible by other services running in the same Rack.
</div>
