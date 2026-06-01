---
title: "ECSExec"
description: "Enable AWS ECS Exec for interactive shell access to this application's containers."
---

# ECSExec

Routes interactive container access for this application through [ECS Exec](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-exec.html) over AWS SSM Session Manager. When enabled, `convox exec` and `convox run` open sessions into running containers over SSM instead of the Docker daemon. On every deploy this app-level value is reset from the [rack-level ECSExec parameter](/reference/rack-parameters/ECSExec), so the durable per-App control is the `params` block in `convox.yml`.

| Default value  | `No` |
| Allowed values | `Yes`, `No` |

## Use Cases

- Set to `Yes` when you need interactive shell access into Fargate tasks, where the Docker exec path is unavailable
- Set to `Yes` when you want exec sessions logged to AWS CloudTrail for audit purposes
- Keep as `No` for applications that do not require interactive container access

## Additional Information

ECS Exec only applies to Generation 2 applications. For gen1 apps the parameter has no effect and exec falls back to the Docker daemon. When set to `Yes`, the application's ECS services are created with execute-command enabled and the task role receives the SSM messages permissions required for the exec channel automatically, so no manual IAM changes are needed.

Enabling ECSExec affects only tasks launched after the change. Tasks started before it was enabled fall back to Docker exec and will not appear in CloudTrail, so redeploy the application after enabling so the setting applies to all running tasks. If a session reports that the SSM agent is not ready, wait a few seconds and retry while the in-container agent connects.

There are three ways to set this parameter, with different durability:

- `convox rack params set ECSExec=Yes` enables ECS Exec rack-wide for every App on the Rack. This is the value the App stack reads on each deploy, so it persists.
- A `params` block in `convox.yml` pins the value for a single application and survives every deploy, overriding the rack-level value.
- `convox apps params set ECSExec=Yes -a my-app` updates the App stack immediately, but the next deploy resets it to the rack-level value unless `convox.yml` also pins it.

```bash
$ convox apps params set ECSExec=Yes
```

After enabling, redeploy the application so existing tasks pick up the change. To set ECS Exec for every App on the Rack instead, use the [rack-level ECSExec parameter](/reference/rack-parameters/ECSExec).

## See Also

- [EnableContainerReadonlyRootFilesystem](/reference/app-parameters/EnableContainerReadonlyRootFilesystem)
- [FargateServices](/reference/app-parameters/FargateServices)
- [exec](/reference/cli-commands/exec)
- [Rack Parameters](/reference/rack-parameters)
