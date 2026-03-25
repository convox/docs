---
title: "Service Discovery"
description: "Discover and connect to Convox services across apps and Racks using internal DNS, links, and the Convox API."
---

# Service Discovery

Convox enables communication between Services and applications whether they are closely related, unrelated, on the same Rack, on different Racks, or accessing the raw nodes behind the load balancer.

Building a loosely coupled architecture relies on your different services being able to interrogate and discover, independently, where other services that they may need are deployed, without having to have specific IP addresses and ports defined.  In a cloud-centric world, these are transient and cannot be relied upon. 

Previously, you may have had to set up a Zookeeper or Consul instance and configured all of your internal services to report and interrogate with them.  Convox provides Service Discovery opportunities out of the box without relying on yet another service to maintain.

> **Note:** If you would like to make a Service accessible only to other Services on the same Rack, see [Internal Services](/networking/internal-services).

## To Discover Services in the Same App

You can link services in the same app using the [links](/application/services#links) section of [convox.yml](/application/convox-yml).

### Example

```yaml
services:
  web:
    port: 3000
  worker:
    links:
      - web
```

This would set an `WEB_URL` environment variable on the `worker` service pointing at the load balancer for the `web` service.  No matter where or how many `web` service instances are deployed, the `worker` instances will be able to access them via the load balancer.

## To Discover Services/Apps on the Same Rack

Convox sets up internal DNS on a Rack such that the following hostname format will resolve to a specific service:

```text
[service].[app].[rack].convox
```

> **Note:** The Rack's name is available as the environment variable `RACK`.

### Example

```bash
$ curl https://api.auth.$RACK.convox
```

This hostname would resolve to the load balancer of the `api` service of the `auth` app on the current Rack.  Using naming conventions like this allow for effective auto-discovery with no overhead.

## To Discover Services/Apps on Different Racks

Use `convox services` to find the load balancer hostname of a given service and set an environment variable with the resulting hostname.

### Example

```bash
$ convox services -a auth -r acme/utility
SERVICE  DOMAIN                                 PORTS
api      auth-api.router.us-east-1.convox.site  80:3000 443:3000
```

You can then extract the load balancer URL to inject into any other service you require...

```bash
$ AUTH_URL="$(convox services -a auth -r acme/utility | sed -n 2p | awk '{print $2}')"
$ convox env set AUTH_URL=$AUTH_URL -a frontend -r acme/production
```

This would set an `AUTH_URL` environment variable on the `frontend` app on the `acme/production` Rack to point at the load balancer for the `api` service on the `auth` app on the `acme/utility` Rack.

## To Discover every Instance of a Service

You can interrogate the Convox API to retrieve the raw endpoints of the current container instances of a service or app.  The API will return JSON to describe the running processes for an app.

### Example

```bash
$ convox api get /apps/nodejs/processes
[
  {
    "app": "nodejs",
    "command": "",
    "cpu": 0,
    "host": "10.0.2.10",
    "id": "3485248b1d18",
    "image": "xxx.dkr.ecr.us-east-1.amazonaws.com/xxx:web.BIYKGJLHGIV",
    "instance": "i-092941b05e75008ff",
    "memory": 0,
    "name": "web",
    "ports": [
      "47889:3000"
    ],
    "release": "RLEUTDHBBKD",
    "started": "2025-01-15T09:32:51Z",
    "status": "running"
  },
  {
    "app": "nodejs",
    "command": "",
    "cpu": 0,
    "host": "10.0.3.87",
    "id": "8db3b9378db0",
    "image": "xxx.dkr.ecr.us-east-1.amazonaws.com/xxx:web.BIYKGJLHGIV",
    "instance": "i-02cf094037d64f5e5",
    "memory": 0,
    "name": "web",
    "ports": [
      "32768:3000"
    ],
    "release": "RLEUTDHBBKD",
    "started": "2025-01-15T11:10:31Z",
    "status": "running"
  }
]
```

> **Note:** You can use `jq` to parse the JSON and extract specific information. For example, to find all the `web` services running in the `nodejs` app:

```bash
$ convox api get /apps/nodejs/processes | jq '.[] | if .name == "web" then {node: .host, port: .ports[] | split(":")[0]} else "" end | join(":")'
"10.0.3.87:32768"
"10.0.2.10:47889"
```

## See Also

- [Internal Services](/networking/internal-services)
- [Services Configuration](/application/services)
- [Load Balancing](/networking/load-balancing)
- [Environment Variables](/application/environment)
