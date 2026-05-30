---
title: "exec"
description: "Execute a command in a running Process."
---

# exec

Execute a command in a running Process. Attaches stdin/stdout for interactive use, making it suitable for debugging sessions and one-off tasks. Use `run` to execute a command in a new Process instead.

## Syntax

```bash
$ convox exec <pid> <command>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox exec web-abc1234-def5 bash -a myapp
root@web-abc1234-def5:/app# ls
Gemfile  Gemfile.lock  Rakefile  app  config  db  lib  public
root@web-abc1234-def5:/app# exit
$
```

## ECS Exec

When the Rack has [ECSExec](/reference/rack-parameters/ECSExec)=`Yes`, `convox exec` tunnels the session through AWS SSM Session Manager (the ECS Exec feature) instead of connecting to the Docker daemon on the host instance. SSM brokers the connection to the container, so you can exec into Fargate tasks and into tasks running on EC2 instances the Rack API cannot reach directly.

`ECSExec` defaults to `No`. With the default, `convox exec` continues to use Docker exec against the host instance, and the behavior of this command is unchanged.

ECS Exec requires the AWS `session-manager-plugin` binary on the machine running the `convox` CLI. The CLI launches the plugin to carry the interactive stream, and reports an error with installation instructions if the plugin is not found on your `PATH`. Install it from the [AWS Session Manager plugin guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html).

```bash
$ convox rack params set ECSExec=Yes
```

After enabling, redeploy each App so its tasks register with ECS Exec.

### Fallback to Docker exec

`convox exec` falls back to the Docker exec path in three cases:

- The Rack has `ECSExec=No`.
- The App is a gen1 App. Gen1 Apps always use Docker exec regardless of the Rack setting.
- The task started before ECS Exec was enabled. These tasks register with ECS Exec only after a redeploy, so until then the command falls back to Docker exec and prints a warning. Redeploy the App to route every task through SSM.

The Docker exec path requires the CLI to reach the host instance. On Fargate, where there is no reachable host, exec is only available with `ECSExec=Yes`. Attempting to exec into a Fargate task on an `ECSExec=No` Rack returns an error directing you to set `ECSExec=Yes` and redeploy.

## See Also

- [run](/reference/cli-commands/run)
- [cp](/reference/cli-commands/cp)
- [ps](/reference/cli-commands/ps)
- [One-off Commands](/management/one-off-commands)
- [Debugging](/management/debugging)
