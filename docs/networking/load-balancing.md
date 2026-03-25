---
title: "Load Balancing"
description: "How Convox uses AWS ALB to route traffic, distribute load, enable sticky sessions, and support rolling deployments."
---

# Load Balancing

## High Level Overview

For background on load balancing concepts, see the [Convox blog post on AWS ELB](https://convox.com/blog/AWS_ELB).

When creating a Rack in your cloud infrastructure, Convox will automatically create a Load Balancer to route traffic to the apps that you will deploy there. In AWS, an ALB will be created for this purpose, along with a separate ALB dedicated to handling traffic to the Rack API.

If you have any [internal services](/networking/internal-services) running - and have enabled it - an extra Load Balancer will be created exclusively for the internal traffic.

When you first deploy your app, Convox will use your configuration provided in the `convox.yml` to set up the Load Balancer rules to route traffic to your app, based on the originating domain.

The Load Balancer will also check the health of your services, based on the [health check](/application/health-checks) information you provide.

### Traffic Balancing and Scaling

The Load Balancer will distribute incoming requests amongst the currently healthy instances of your service evenly.  If you have [automatic scaling rules](/scaling/scaling#service-autoscaling) configured for your app, then as the load grows on your backend instances, and scaling thresholds are reached, then new instances will be automatically spun up and, once healthy, brought into the pool of available instances for the Load Balancer to direct traffic to.  During times when load decreases, then unnecessary instances will be spun down and traffic redirected as appropriate.

### Sticky Sessions

If your client supports sticky sessions and wishes to utilise it, sticky sessions are enabled by default.  The ALB will generate an `AWSALB` named cookie that is returned with the first response from a backend app to the client.  If that cookie is passed back by the client on subsequent requests, the ALB will attempt to route the request to the same backend instance.  If it unable to, then the ALB will attempt to route the request to another available backend instance.  If you wish to disable sticky sessions, you can do so within your [`convox.yml`](/application/services#sticky).

### Rolling Deployments

The Load Balancer will continue to route traffic to healthy instances throughout a [rolling deployment](/deployment/rolling-updates), ensuring zero downtime and continued service to the client.  As new instances with a fresh release are brought online, the Load Balancer will monitor their health, and once passing the [health check](/application/health-checks), will bring them into the pool of available instances to route traffic to.  Older instances with the previous release will then be retired and traffic no longer routed to them.

> **Note:** The very first time you deploy your App, if all instances of a Service fail their health check, Load Balancer behavior is to keep trying to route traffic to the instances and mark the deployment as complete. This only applies to the first time you deploy a new App. After this, if your new instances fail their health checks, the deployment would fail and be rolled back.

### Limitations

Currently, AWS ALBs have a 100 rule limit. Because of this, there is a limit to the number of routable apps that should be deployed on any one Rack. If you run out of available rules on your load balancer, spread your apps across multiple Racks to overcome this limitation. Alternatively change the [InternalDomains](/reference/app-parameters/InternalDomains) app parameter to disable the internal domains from utilising some of those rules.

## Finding Your Load Balancer Hostname

You can find the load balancer hostname(s) for your application using `convox apps info`:

```bash
$ convox apps info
Name        myapp
Status      running
Generation  2
Locked      false
Release     RABCDEFGHIJ
Router      myapp-Router-ABC123-123456789.us-east-1.elb.amazonaws.com
```

Use the hostname from the `Router` output to configure DNS records (e.g., a CNAME record) for your [custom domains](/deployment/custom-domains).

## Protocol and TLS Termination

The Rack ALB handles TLS termination by default — HTTPS connections from clients are decrypted at the load balancer, and traffic is forwarded to your containers over the Rack's internal network. You do not need to configure TLS certificates within your application.

Convox automatically provisions SSL certificates via [AWS ACM](/deployment/ssl). To control which TLS protocol versions and ciphers the load balancer accepts, use the [SslPolicy](/reference/rack-parameters/SslPolicy) rack parameter.

HTTP traffic is automatically redirected to HTTPS by default. To allow plain HTTP connections, set the [RedirectHttps](/reference/app-parameters/RedirectHttps) app parameter to `No`.

## Internal Load Balancers

For services that should only be reachable within your VPC (not from the public internet), Convox supports internal load balancers:

- Set the [Internal](/reference/rack-parameters/Internal) rack parameter to `Yes` to enable an internal load balancer alongside the public one.
- Set [InternalOnly](/reference/rack-parameters/InternalOnly) to `Yes` if you want the Rack to only support internal applications (no public-facing load balancer).

Internal services are configured in your `convox.yml` using the `internal: true` directive. See [Internal Services](/networking/internal-services) for details.

## Health Check Integration

The ALB continuously monitors the health of your service containers using the [health check](/application/health-checks) path and interval defined in your `convox.yml`. Only containers passing health checks receive traffic from the load balancer. During [rolling deployments](/deployment/rolling-updates), new containers must pass their health check before the ALB routes traffic to them.

## Access Control

You can restrict which IP addresses or networks can reach your load balancer:

- Use the [WhiteList](/reference/rack-parameters/WhiteList) rack parameter to restrict access to the Rack API, build instances, and workload instances by CIDR range.
- Use [RouterSecurityGroup](/reference/rack-parameters/RouterSecurityGroup) to attach a custom security group to the public router ALB, allowing you to define fine-grained ingress rules (e.g., restrict to your office VPN).
- Use [RouterInternalSecurityGroup](/reference/rack-parameters/RouterInternalSecurityGroup) to control access to the internal router.

## Idle Timeout

The ALB idle timeout controls how long the load balancer waits for activity on a connection before closing it. The default is 3600 seconds (1 hour). If your application uses WebSockets, long-polling, or other long-lived connections, you may need to adjust this via the [LoadBalancerIdleTimeout](/reference/rack-parameters/LoadBalancerIdleTimeout) rack parameter.

```bash
$ convox rack params set LoadBalancerIdleTimeout=300
```

## Load Balancing Algorithm

By default, the ALB uses round-robin routing to distribute requests across healthy targets. For workloads where backend response times vary, you can switch to least-outstanding-requests routing using the [LoadBalancerAlgorithm](/reference/app-parameters/LoadBalancerAlgorithm) app parameter:

```bash
$ convox apps params set LoadBalancerAlgorithm=least_outstanding_requests
```

## See Also

- [Internal Services](/networking/internal-services)
- [Health Checks](/application/health-checks)
- [Custom Domains](/deployment/custom-domains)
- [SSL](/deployment/ssl)
- [Scaling](/scaling/scaling)
- [Rolling Updates](/deployment/rolling-updates)
- [SslPolicy](/reference/rack-parameters/SslPolicy)
- [RouterSecurityGroup](/reference/rack-parameters/RouterSecurityGroup)
- [WhiteList](/reference/rack-parameters/WhiteList)
- [LoadBalancerIdleTimeout](/reference/rack-parameters/LoadBalancerIdleTimeout)
- [LoadBalancerAlgorithm](/reference/app-parameters/LoadBalancerAlgorithm)
