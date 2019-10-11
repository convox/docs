---
title: "Load Balancing"
---

## High Level Overview

For an interesting primer to Load Balancing in general, please read our [Blog post](https://convox.com/blog/AWS_ELB) on the topic!

When creating a Rack in your cloud infrastructure, Convox will automatically create a Load Balancer to route traffic to the apps that you will deploy there.  In AWS, an ALB will be created for this purpose (alongside a classic ELB to handle traffic to the Rack API exclusively).

If you have any [internal services](/deployment/internal-services) running - and have enabled it - an extra Load Balancer will be created exclusively for the internal traffic.

When you first deploy your app, Convox will use your configuration provided in the `convox.yml` to set up the Load Balancer rules to route traffic to your app, based on the originating domain.

The Load Balancer will also check the health of your services, based on the [health check](/application/health-checks) information you provide.

#### Traffic Balancing and Scaling

The Load Balancer will distribute incoming requests amongst the currently healthy instances of your service evenly.  If you have [automatic scaling rules](/deployment/scaling#service-autoscaling) configured for your app, then as the load grows on your backend instances, and scaling thresholds are reached, then new instances will be automatically spun up and, once healthy, brought into the pool of available instances for the Load Balancer to direct traffic to.  During times when load decreases, then unnecessary instances will be spun down and traffic redirected as appropriate.

#### Sticky Sessions

If your client supports sticky sessions and wishes to utilise it, sticky sessions are enabled by default.  The ALB will generate an `AWSALB` named cookie that is returned with the first response from a backend app to the client.  If that cookie is passed back by the client on subsequent requests, the ALB will attempt to route the request to the same backed instance.  If it unable to, then the ALB will attempt to route the request to another available backend instance.  If you wish to disable sticky sessions, you can do so within your [`convox.yml`](/application/services#sticky).

#### Rolling Deployments

The Load Balancer will continue to route traffic to healthy instances throughout a [rolling deployment](/deployment/rolling-updates), ensuring zero downtime and continued service to the client.  As new instances with a fresh release are brought online, the Load Balancer will monitor their health, and once passing the [health check](/application/health-checks), will bring them into the pool of available instances to route traffic to.  Older instances with the previous release will then be retired and traffic no longer routed to them.

- The very first time you deploy your app, if all instances of a service fail their health check, Load Balancer behaviour is to keep trying to route traffic to the instances and mark the deployment as complete.  This only applies to the first time you deploy a new app.  After this, if your new instances fail their health checks, the deployment would fail and be rolled back.

#### Limitations

- Currently, AWS ALB's have a 100 rule limit.  Because of this, there is a limit to the number of routable apps that should be deployed on any one rack.   If you run out of available rules on your load balancer, you can simply spread your apps across multiple racks to overcome this limitation, or alternative change the [`InternalDomains`](/reference/app-parameters#internaldomains) app parameter to disable the internal domains from utilising some of those rules.
