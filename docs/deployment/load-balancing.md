---
title: "Load Balancing"
---

## High Level overview

For an interesting primer to Load Balancing in general, please read our [Blog post](https://convox.com/blog/AWS_ELB) on the topic!

When creating a Rack in your cloud infrastructure, Convox will automatically create a Load Balancer to route traffic to the apps that you will deploy there.  In AWS, an ALB will be created for this purpose (alongside a classic ELB to handle traffic to the Rack API exclusively).

If you have any [internal services](/deployment/internal-services) running - and have enabled it - an extra Load Balancer will be created exclusively for the internal traffic.

When you first deploy your app, Convox will use your configuration provided in the `convox.yml` to set up the Load Balancer rules to route traffic to your app, based on the originating domain.

The Load Balancer will also check the health of your services, based on the [health check](/application/health-checks) information you provide.

#### Traffic Balancing

The Load Balancer will distribute incoming requests amongst the currently healthy instances of your service evenly.  


#### Sticky Sessions

If your client supports sticky sessions and wishes to utilise it, sticky sessions are enabled by default.  The ALB will generate an `AWSALB` named cookie that is returned with the first response from a backend app to the client.  If that cookie is passed back by the client on subsequent requests, the ALB will attempt to route the request to the same backed instance.  If it unable to, then the ALB will attempt to route the request to another available backend instance.  If you wish to disable sticky sessions, you can do so within your [`convox.yml`](/application/services#sticky).


#### Rolling deployments




#### Limitations

- Currently, AWS ALB's have a 100 rule limit.  Because of this, no more than 50 routable apps should be deployed on any one rack - 2 rules could be created for each app, one for the custom Convox domain, and one for your own domains.  

- If, on your very first deployment of an app, all instances of a service are unhealthy, the default Load Balancer behaviour is to keep trying to route traffic to the instances and mark the deployment as complete.  Otherwise, there is no previous deployment to rollback to.  Once one successful deployment has been made, then normal processes apply.

