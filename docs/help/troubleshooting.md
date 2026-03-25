---
title: "Troubleshooting"
description: "Common issues and solutions when deploying and running applications on Convox"
---

# Troubleshooting

## I got an error while installing Convox

Look at the AWS [Cloudformation Management Console](https://console.aws.amazon.com/cloudformation/home?region=us-east-1) and look for CREATE_FAILED errors on the "Events" tab.

Post in the new [Community Forum](https://community.convox.com/) with any errors you see in these events.

## I get an error when I deploy my app to Convox

During a deployment, CloudFormation will not complete an update until the ECS Services stabilize. If newly deployed processes crash or fail to pass health checks, eventually the update will time out and roll back. To figure out what's going wrong, you can look at the [app logs](/management/logs) via `convox logs` to check for crashes and [health check](/application/health-checks) failures.

ECS events can be found in the application logs as well. Use `convox logs --filter=ECS` to find them.

When you know there is an issue and want to stop a deployment, you can run the `convox apps cancel` command. This will trigger an immediate CloudFormation rollback so you can fix the problem and try another deployment.

### I get a `CREATE_FAILED BalancerListenerRule443Internal Priority 'xxxx' is currently in use`

As ALB routing rules must have a unique priority, Convox will generate a random one for each service in the range 1-50000 (the range allowed by AWS). This is generated from a checksum of the app name, service name, domain name, and a couple of other tweaks. It is very unlikely, but it is possible to have a collision between two separate services on the same rack. To solve this, amend your app or service name to generate a different checksum.

## My app deployed but I cannot access it

Run `convox services` to find the load balancer endpoints for your application.

Run `convox ps` to determine if your application is booting successfully.

Run `convox logs` to inspect your application logs and cluster events for problems placing your container, starting your app, or registering with the load balancer.

## My app stopped working and I want to restart it

You can perform a remote restart of an entire App (all running processes) from the CLI with:

```bash
$ convox restart -a app1
```

Or alternatively to restart the `web` service processes, you can perform:

```bash
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

## Not enough memory or CPU to place tasks

If you see errors like `not enough memory available to start process` or your deployments fail with ECS placement errors (`RESOURCE:MEMORY`, `RESOURCE:CPU`, or `RESOURCE:PORTS`), your cluster does not have enough capacity to run the requested tasks.

**Solutions:**

- Increase [InstanceCount](/reference/rack-parameters/InstanceCount) or use larger [InstanceType](/reference/rack-parameters/InstanceType)
- Reduce the `cpu` and `memory` values in your `convox.yml` service definitions
- Enable [Autoscale](/reference/rack-parameters/Autoscale) to allow the cluster to grow automatically
- Consider [FargateServices](/reference/app-parameters/FargateServices) to avoid instance capacity constraints entirely

## Build failures

Common build errors and their causes:

- **`no such file: convox.yml`** -- Your application directory does not contain a `convox.yml` (or `docker-compose.yml`). Ensure the manifest file exists in the root of your build context.
- **`reading manifest` or `loading manifest`** -- Your `convox.yml` has syntax errors. Validate the YAML before deploying.
- **ECR authentication failures** (`ecr auth token`, `empty authorization data`) -- AWS credentials are missing, expired, or lack ECR permissions. ECR tokens expire after 12 hours, so long-running build environments may encounter this.
- **Image pull failures** (`could not pull <image>`) -- The base image in your Dockerfile is unavailable. Verify the image name/tag and that your build environment has network access to the registry.

```bash
$ convox builds -a myapp
$ convox builds logs <build-id> -a myapp
```

## CloudFormation rollback

If a Rack update or app deployment triggers a CloudFormation rollback, check the stack events for the specific failure reason:

```bash
$ convox rack
```

Common rollback causes:

- **Resource limit exceeded** -- AWS account limits for ALB rules (100 per listener), security groups, or ECS services. Request a limit increase via the AWS console.
- **Invalid parameter values** -- A Rack or app parameter was set to a value that CloudFormation rejected.
- **Permission errors** -- The Rack's IAM role lacks permission for a new resource type. This can happen after a Rack update introduces new AWS resources.

If a rollback fails (stack stuck in `UPDATE_ROLLBACK_FAILED`), you may need to manually resolve the issue in the AWS CloudFormation console by continuing the rollback with resources to skip.

## Fargate

To have ECS use Fargate instead of regular EC2 you will need the following:

- have a rack installed in [a region where Fargate is available](https://aws.amazon.com/about-aws/global-infrastructure/regional-product-services/)
- specify `cpu` and `memory` in [convox.yml](/application/convox-yml) that is [compatible with Fargate](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-cpu-memory-error.html)
- either
   - run _all_ your services and/or tasks in Fargate by setting the relevant [App Parameters](/reference/app-parameters)
   - enable Fargate on a _per service_ basis via the ProcessName Formation parameter
      - e.g. `convox apps params set WebFormation=2,256,512,FARGATE` for running on Fargate
      - e.g. `convox apps params set WebFormation=2,256,512,FARGATE_SPOT` for running on Fargate Spot

## Fargate logging not available

If you see `cloudwatch disabled and ecs-exec not enabled; unable to stream logs for fargate task`, it means the Fargate task has no logging path configured. Ensure [LogDriver](/reference/rack-parameters/LogDriver) is set to `CloudWatch` (the default), or enable ECS Exec for the task.

## Private registry authentication failures

If builds fail with `unable to authenticate` when pulling from a private registry, verify your registry credentials are configured correctly:

```bash
$ convox registries
$ convox registries add <server> <username> <password>
```

## Environment variable decryption errors

If you see `failed decryption` errors, the Rack's KMS key cannot decrypt stored environment variables. This can happen if the KMS key was deleted or access was revoked. Check the Rack's KMS key status in the AWS KMS console.

## My CLI commands take a long time to return

If your local Kubernetes setup does not point to a valid cluster, that can slow down your Convox CLI operations as it tries to interrogate the invalid endpoint. In this case, you can set a local env var `$ export CONVOX_LOCAL=disable` to stop the CLI from doing this and speed up your commands.

## Migrating to Convox v3

For cloud-agnostic deployments across AWS (EKS), Google Cloud (GKE), Azure (AKS), and DigitalOcean (DOKS), consider migrating to [Convox v3](https://docs.convox.com/getting-started/introduction). See the [v2 to v3 migration guide](https://www.convox.com/blog/convox-v2-to-v3-migration-guide) for details.

## Still having trouble?

Some good places to search are:

- this site, via the search box in the sidebar
- visit our [Community Forum](https://community.convox.com/)

If you still need help, feel free to:

- post a question on the [Community Forum](https://community.convox.com/)
- open a ticket via the Support section [in the Convox web console](https://console.convox.com/)

## See Also

- [Logs](/management/logs)
- [Debugging](/management/debugging)
- [Health Checks](/application/health-checks)
- [Support Plans](/console/support-plans)
- [Convox v3 Documentation](https://docs.convox.com/getting-started/introduction)
