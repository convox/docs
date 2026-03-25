---
title: "Internal Services"
description: "Configure Services accessible only within your VPC using the internal load balancer."
---

# Internal Services

An internal Service is accessible only from within the Rack's VPC. Traffic from the public internet cannot reach internal Services. This is useful for backend APIs, microservices, and any workload that should not be directly exposed.

## Prerequisites

Before marking any Service as internal, enable the internal load balancer on your Rack:

```bash
$ convox rack params set Internal=Yes
```

This creates a dedicated internal Application Load Balancer (ALB) within your VPC. The internal ALB accepts traffic only from within the VPC CIDR range on ports 80 and 443, with HTTP automatically redirecting to HTTPS.

## Configuring a Service as Internal

Set the `internal` attribute on a Service in your `convox.yml`:

```yaml
services:
  api:
    build: .
    port: 3000
    internal: true
```

Deploy the change:

```bash
$ convox deploy
```

## Internal DNS and Routing

Each internal Service gets a DNS record in the Rack's private Route53 hosted zone with the format:

```text
{service}.{app}.{rack}.convox
```

For example, a Service named `api` in an App named `backend` on a Rack named `production` is reachable at:

```text
api.backend.production.convox
```

Other Services within the same Rack can reach internal Services using this hostname over HTTPS on port 443. The internal ALB uses an auto-provisioned ACM wildcard certificate.

## Service-to-Service Communication

Use links to connect Services. When a Service links to an internal Service, Convox automatically resolves the link URL to the internal ALB. Define links in your [convox.yml](/application/convox-yml):

```yaml
services:
  web:
    build: .
    port: 3000
    links:
      - api
  api:
    build: ./api
    port: 4000
    internal: true
```

The `web` Service receives an `API_URL` environment variable pointing to the internal `api` Service endpoint.

## Internal and External Access

If a Service needs to be reachable from both inside the VPC and the public internet, use the `internalAndExternal` attribute instead:

```yaml
services:
  api:
    build: .
    port: 3000
    internalAndExternal: true
```

This registers the Service with both the internal and external ALBs, producing two endpoints.

## Rack Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `Internal` | `No` | Enable the internal load balancer for the Rack |
| `InternalOnly` | `No` | Restrict the Rack to only support internal Services (no external ALB) |

If `InternalOnly` is set to `Yes`, deploying a Service without `internal: true` will fail with an error.

## Use Cases

- **Backend APIs**: Keep API Services accessible only to frontend Services within the same Rack
- **Microservices**: Run internal microservices that communicate via service links without public exposure
- **Database Proxies**: Run connection poolers or proxy Services that should never face the internet
- **Admin Dashboards**: Host internal tools accessible only via VPN or bastion host within the VPC

## See Also

- [convox.yml](/application/convox-yml)
- [Private Networking](/networking/private-networking)
- [Service Discovery](/networking/service-discovery)
- [Rack Parameters](/reference/rack-parameters)
