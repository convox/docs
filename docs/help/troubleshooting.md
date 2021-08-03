---
title: "Troubleshooting"
---

## I got an error while installing Convox

Look at the AWS [Cloudformation Management Console](https://console.aws.amazon.com/cloudformation/home?region=us-east-1) and look for CREATE_FAILED errors on the "Events" tab.

Post in the new [Community Forum](https://community.convox.com/) with any errors you see in these events.

## I get an error when I deploy my app to Convox

During a deployment, CloudFormation will not complete an update until the ECS Services stabilize. If newly deployed processes crash or fail to pass health checks, eventually the update will time out and roll back. To figure out what's going wrong, you can look at the [app logs](/management/logs) via `convox logs` to check for crashes and [health check](/application/health-checks) failures.

ECS events can be found in the application logs as well. Use `convox logs --filter=ECS` to find them.

When you know there is an issue and want to stop a deployment, you can run the `convox apps cancel` command. This will trigger an immediate CloudFormation rollback so you can fix the problem and try another deployment.

### I get a `CREATE_FAILED BalancerListenerRule443Internal Priority 'xxxx' is currently in use`

As ALB routing rules must have a unique priority, Convox will generate a random one for each service in the range 1-50000 (the range allowed by AWS).  This is generated from a checksum of the app name, service name, domain name, and a couple of other tweaks.  It is very unlikely, but it is possible to have a collision between two separate services on the same rack.  To solve this, simply slightly amend your app or service name to generate a different checksum.

## My app deployed but I cannot access it

Run `convox services` to find the load balancer endpoints for your application.

Run `convox ps` to determine if your application is booting successfully.

Run `convox logs` to inspect your application logs and cluster events for problems placing your container, starting your app, or registering with the load balancer.

## My app stopped working and I want to restart it

You can perform a remote restart of an entire App (all running processes) from the CLI with:

```
$ convox restart -a app1
```

Or alternatively to just restart the `web` service processes, you can perform:

```
$ convox services restart web -a app1
```

## ECS

If you're encountering problems like any of the following:

- deployments seem stuck in `UPDATE_IN_PROGRESS` CloudFormation state
- your Rack seems stuck in `converging`
- your deploys seem stuck with services continuously being killed via SIGKILL and restarted

check out the logs for your application. You will see messages from CloudFormation and ECS about the current rollout of your application.

If the listed reason is `Instance has failed at least the UnhealthyThreshold number of health checks consecutively`, this means your app is failing its [health checks](/deployment/rolling-updates#health-checks). Often this is because your app is not actually listening on the port(s) specified for the service in your `convox.yml`. Or your Rack might be missing an environment variable defined in your `convox.yml`. It might also be an error from your application itself (check the [app logs with `convox logs`](/management/debugging#convox-logs), or use [`convox instances ssh`](/management/debugging#convox-instances-ssh) to check docker logs).

Once you've found the reason, you can run `convox apps cancel -a app_name` to cancel this deployment. Once it rolls back, you can fix the error and try to deploy again.


## Fargate

To use have ECS use Fargate instead of regular EC2 you will need the following:

- have a rack installed in [a region where Fargate is available](https://aws.amazon.com/about-aws/global-infrastructure/regional-product-services/)
- specify `cpu` and `memory` in [convox.yml](/application/convox-yml) that is [compatible with Fargate](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-cpu-memory-error.html)
- either
   - run _all_ your services and/or tasks in Fargate by setting the relevant [App Parameters](/reference/app-parameters)
   - enable Fargate on a _per service_ basis via the ProcessName Formation parameter
      - e.g. `convox apps params set WebFormation=2,256,512,FARGATE` for running on Fargate
      - e.g. `convox apps params set WebFormation=2,256,512,FARGATE_SPOT` for running on Fargate Spot

   
## My CLI commands take a long time to return

If your local Kubernetes setup does not point to a valid cluster, that can slow down your Convox CLI operations as it tries to interrogate the invalid endpoint.  In this case, you can set a local env var `$ export CONVOX_LOCAL=disable` to stop the CLI from doing this and speed up your commands.

## Still having trouble?

Some good places to search are:

- this site, via the search box on in the sidebar
- visit our [Community Forum](https://community.convox.com/)

If you still need help, feel free to:

- post a question on the [Community Forum](https://community.convox.com/)
- open a ticket via the Support section [in the Convox web console](https://console.convox.com/)
