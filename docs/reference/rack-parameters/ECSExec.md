---
title: "ECSExec"
description: "Enable interactive container access through ECS Exec over AWS SSM Session Manager on a Convox Rack."
---

# ECSExec

Routes interactive container access through ECS Exec instead of the Docker daemon. When enabled, `convox exec` opens an interactive session into a running container over AWS SSM Session Manager, and `convox run` enables the execute-command capability on each new task it starts.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Use Cases

- Reaching containers on Fargate, where there is no Docker daemon to connect to
- Auditing interactive access through CloudTrail and SSM Session Manager logging
- Removing the need for direct Docker socket access to inspect running Processes

## Additional Information

When set to `Yes`, the Rack calls the ECS `ExecuteCommand` API for `convox exec` and returns short-lived session credentials to the CLI. The CLI then launches `session-manager-plugin` locally to attach to the container, the same mechanism used by `aws ecs execute-command`. For `convox run`, the Rack sets the execute-command capability on each new task so those tasks can be reached the same way.

```bash
$ convox rack params set ECSExec=Yes
```

Because the credentials are brokered through SSM, users must have `session-manager-plugin` installed on their machine. If it is missing, the CLI reports that the plugin was not found in `PATH` and links to the AWS install instructions.

### Prerequisites

Install `session-manager-plugin` locally before running `convox exec` against an `ECSExec=Yes` Rack. See the [AWS Session Manager plugin install guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html).

Racks in fully private VPCs without a NAT gateway need VPC interface endpoints for `ssmmessages` and `ssm` so the container SSM agent and the local plugin can reach the SSM data channel.

### Fallback to Docker Exec

`convox exec` automatically falls back to Docker exec in three cases:

- Gen1 Apps, which always use Docker exec
- Racks with `ECSExec=No`
- Tasks that were started before ECS Exec was enabled

For the third case, the fallback session is local and is not recorded in CloudTrail. Redeploy the App to move its tasks onto ECS Exec:

```bash
$ convox deploy -a my-app
```

### Operational Notes

Changing this parameter updates the Docker labels on the Rack API task definition, which triggers a rolling restart of the Rack API as the new task definition rolls out.

Downgrading to a Rack version that predates this parameter requires setting `ECSExec=No` first. Apply that change, wait for the update to complete, then downgrade in a follow-up step.

## See Also

- [exec](/reference/cli-commands/exec)
- [One-off Commands](/management/one-off-commands)
- [Rack Parameters](/reference/rack-parameters)
