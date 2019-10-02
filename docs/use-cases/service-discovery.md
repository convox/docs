---
title: "Service Discovery"
---

Convox makes it easy to build microservice architectures by enabling easy communication between services and applications whether they are closely related, unrelated, on the same rack, on different racks, and even access the raw nodes behind the load balancer if required.

Building a loosely coupled architecture relies on your different services being able to interrogate and discover, independently, where other services that they may need are deployed, without having to have specific IP addresses and ports defined.  In a cloud-centric world, these are transient and cannot be relied upon. 

Previously, you may have had to set up a Zookeeper or Consul instance and configured all or your internal services to report and interrogate with them.  Convox provides Service Discovery oppportunities out of the box without relying on yet another service to maintain!

<div class="block-callout block-show-callout type-info" markdown="1">
If you would like to make a service accessible only to other services on the same Rack, see [Internal Services](/deployment/internal-services).
</div>

### To Discover Services in the Same App

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

This would set an `WEB_URL` environment variable on the `worker` service pointing at the load balancer for the `web` service.  No matter where or how many `web` service instances are deployed, the `worker` instances will be able to access them via the load balancer.

### To Discover Apps on the Same Rack

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

This hostname would resolve to the load balancer of the `api` service of the `auth` app on the current Rack.  Using naming conventions like this allow for effective auto-discovery with no overhead.

### To Discover Apps on Different Racks

Use `convox services` to find the load balancer hostname of a given service and set an environment variable with the resulting hostname.

#### Example

```
$ convox services -a auth -r acme/utility
SERVICE  DOMAIN                                 PORTS
api      auth-api.router.us-east-1.convox.site  80:3000 443:3000
```

You can then extract the load balancer URL to inject into any other service you require...

```
$ AUTH_URL="$(convox services -a docs -r acme/test | sed -n 2p | awk '{print $2}')"
$ convox env set AUTH_URL=$AUTH_URL -a frontend -r acme/production
```

This would set an `AUTH_URL` environment variable on the `frontend` app on the `acme/production` Rack to point at the load balancer for the `api` service on the `auth` app on the `acme/utility` Rack.

### To Discover every Instance of a Service

You can interrogate the Convox API to retrieve the raw endpoints of the current container instances of a service or app.  The API will return JSON to describe the running processes for an app.

#### Example

```
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
    "started": "2019-09-25T09:32:51Z",
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
    "started": "2019-09-03T11:10:31Z",
    "status": "running"
  },
  {
    "app": "nodejs",
    "command": "",
    "cpu": 0,
    "host": "10.0.1.140",
    "id": "9bb2f258abdf",
    "image": "xxx.dkr.ecr.us-east-1.amazonaws.com/xxx:web.BIYKGJLHGIV",
    "instance": "i-0df6e9e3d56e8d1c7",
    "memory": 0,
    "name": "web",
    "ports": [
      "47773:3000"
    ],
    "release": "RLEUTDHBBKD",
    "started": "2019-09-25T09:32:51Z",
    "status": "running"
  }
]
```

We can use awesome tools like `jq` to parse the json and pull out the information we need, for instance to find all the `web` services running in the `nodejs` app we can do something like this:

```
$ convox api get /apps/nodejs/processes | jq '.[] | if .name == "web" then {node: .host, port: .ports[] | split(":")[0]} else "" end | join(":")'
"10.0.3.87:47770"
"10.0.1.140:37850"
```




