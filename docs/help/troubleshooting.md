---
title: "Troubleshooting"
---

## I got an error while installing Convox

Look at the AWS [Cloudformation Management Console](https://console.aws.amazon.com/cloudformation/home?region=us-east-1) and look for CREATE_FAILED errors on the "Events" tab.

Open a [Github issue](https://github.com/convox/rack/issues/new) with any errors you see in these events.

## I get an error when I deploy my app to Convox

During a deployment, CloudFormation will not complete an update until the ECS Services stabilize. If newly deployed processes crash or fail to pass health checks, eventually the update will time out and roll back. To figure out what's going wrong, you can look at the [app logs](/docs/logs) via `convox logs` to check for crashes and [health check](/docs/health-checks) failures.

ECS events can be found in the application logs as well. Use `convox logs --filter=ECS` to find them.

When you know there is an issue and want to stop a deployment, you can run the `convox apps cancel` command. This will trigger an immediate CloudFormation rollback so you can fix the problem and try another deployment.

## My app deployed but I cannot access it

Run `convox services` to find the load balancer endpoints for your application.

Run `convox ps` to determine if your application is booting successfully.

Run `convox logs` to inspect your application logs and cluster events for problems placing your container, starting your app, or registering with the load balancer.


## ECS

If you're encountering problems like any of the following:

- deployments seem stuck in `UPDATE_IN_PROGRESS` CloudFormation state
- your Rack seems stuck in `converging`
- your deploys seem stuck with services continuously being killed via SIGKILL and restarted

check out the logs for your application. You will see messages from CloudFormation and ECS about the current rollout of your application.

If the listed reason is `Instance has failed at least the UnhealthyThreshold number of health checks consecutively`, this means your app is failing its [health checks](/docs/rolling-updates/#health-checks). Often this is because your app is not actually listening on the port(s) specified for the service in your `convox.yml`. Or your Rack might be missing an environment variable defined in your `convox.yml`. It might also be an error from your application itself (check the [app logs with `convox logs`](/docs/debugging/#convox-logs), or use [`convox instances ssh`](/docs/debugging/#convox-instances-ssh) to check docker logs).

Once you've found the reason, you can run [`convox apps cancel -a app_name`](/docs/deploying-to-convox/#canceling-a-deployment) to cancel this deployment. Once it rolls back, you can fix the error and try to deploy again.


## Fargate

To use have ECS use Fargate instead of regular EC2 you will need the following:

- have a rack installed in [a region where Fargate is available](https://aws.amazon.com/about-aws/global-infrastructure/regional-product-services/)
- specify `cpu` and `memory` in [convox.yml](/docs/convox-yml) that is [compatible with Fargate](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-cpu-memory-error.html)
- either
   - run _all_ your services and/or tasks in Fargate by setting the relevant [App Parameters](/docs/app-parameters)
   - enable Fargate on a _per service_ basis via the [&lt;ProcessName&gt;Formation](/docs/gen1/app-parameters/) parameter
      - e.g. `convox apps params set WebFormation=2,256,512,FARGATE`


## Still having trouble?

Some good places to search are:

- this site, via [google](https://www.google.com/search?q=site%3Aconvox.com) or the search box above
- the history of our [public Slack channel](https://invite.convox.com/)
- [GitHub issues](https://github.com/convox/rack/issues)

If you still need help, feel free to:

- reach out [on Slack](https://invite.convox.com/)
- open a ticket via the Support section [in the Convox web console](https://console.convox.com/)
