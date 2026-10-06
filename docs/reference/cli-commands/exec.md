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
$ convox exec 5e3c8576b942 bash -a myapp
root@a1b2c3d4e5f6:/app# ls
Gemfile  Gemfile.lock  Rakefile  app  config  db  lib  public
root@a1b2c3d4e5f6:/app# exit
$
```

## Passing a Command

Everything after the Process ID is joined into a single command string. On the default Docker exec path, the Rack runs that string in the container with `sh -c`.

Quoting is not preserved across the join, so a command that contains shell syntax must be passed as one quoted argument:

```bash
$ convox exec 5e3c8576b942 'echo A; echo B' -a myapp
A
B
```

A command with no shell syntax needs no quoting, which is why the Example Usage above passes `bash` directly.

If your command takes flags of its own, put `--` after the `convox` flags so the CLI stops parsing:

```bash
$ convox exec 5e3c8576b942 -a myapp -- ls -la /app
```

Without `--`, the CLI claims the flag for itself:

```bash
$ convox exec 5e3c8576b942 ls -la /app -a myapp
ERROR: unknown flag: -la
```

Do not add an `sh -c` of your own. The command string is already run through a shell, so an inner `sh -c` is flattened into that same string and produces confusing output rather than an error:

```bash
$ convox exec 5e3c8576b942 -a myapp -- sh -c 'echo A; echo B'

B
```

That arrives as the string `sh -c echo A; echo B`. The shell splits it on `;` and runs `sh -c echo A` as the first statement, where `echo` is the command and `A` becomes `$0`, so it prints a blank line. The shell then runs `echo B` itself and `A` is never printed.

## Running Without a Terminal

When stdin is not a terminal, for example in a pipe, with a redirect such as `< /dev/null`, or in a CI job, `convox exec` runs the command without a TTY. The command's stdout and stderr both arrive on the CLI's stdout, and the CLI exits with the command's exit code. When stdin is a terminal, output passes through a TTY, which ends each line with `\r\n`, so redirect stdin when a script captures the output:

```bash
$ convox exec 5e3c8576b942 'echo hello' -a myapp < /dev/null | od -c
0000000   h   e   l   l   o  \n
0000006
```

Requires rack version 20261005214736 or newer. Older Racks put an 8-byte stream header in front of each chunk of this output. Output that starts with a NUL byte prints with CLI version 3.25.10 or newer, or 20261005214736 or newer; CLI versions 3.24.8 to 3.25.9 and 20260531001024 to 20260929232050 print `ERROR: ECS Exec session ended before it was established; please retry` instead.

## ECS Exec

When the Rack has [ECSExec](/reference/rack-parameters/ECSExec)=`Yes`, `convox exec` tunnels the session through AWS SSM Session Manager (the ECS Exec feature) instead of connecting to the Docker daemon on the host instance. SSM brokers the connection to the container, so you can exec into Fargate tasks and into tasks running on EC2 instances the Rack API cannot reach directly.

`ECSExec` defaults to `No`. With the default, `convox exec` continues to use Docker exec against the host instance, and the behavior of this command is unchanged.

The command is still joined into a single string, but on this path the Rack hands it to the ECS `ExecuteCommand` API rather than running it through its own `sh -c` wrapper. If a command containing shell syntax does not behave as it does on the Docker exec path, start a shell with `convox exec <pid> sh -a <app>` and run the command inside the session.

ECS Exec requires the AWS `session-manager-plugin` binary on the machine running the `convox` CLI. The CLI launches the plugin to carry the interactive stream, and reports an error with installation instructions if the plugin is not found on your `PATH`. Install it from the [AWS Session Manager plugin guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html).

```bash
$ convox rack params set ECSExec=Yes
```

After enabling, redeploy each App so its tasks register with ECS Exec.

### Fallback to Docker exec

`convox exec` falls back to the Docker exec path in four cases:

- The Rack has `ECSExec=No`.
- The App is a gen1 App. Gen1 Apps always use Docker exec regardless of the Rack setting.
- The task started before ECS Exec was enabled. These tasks register with ECS Exec only after a redeploy, so until then the command falls back to Docker exec and prints a warning. Redeploy the App to route every task through SSM.
- stdin is not a terminal, as in a pipe, a redirect or a CI job. ECS Exec accepts only interactive sessions, so these sessions use Docker exec. For a gen2 App on a Rack with `ECSExec=Yes`, the CLI's stdout then starts with a blank line and the same `WARNING: task was started before ECS Exec was enabled ...` line as in the previous case, ahead of the command's output, and a redeploy does not remove it. The same applies to `convox run` and `convox test`.

The Docker exec path requires the Rack to reach the Docker daemon on the EC2 instance that runs the task. On Fargate, where there is no such instance, exec is only available with `ECSExec=Yes` and only when stdin is a terminal. Attempting to exec into a Fargate task on an `ECSExec=No` Rack, or without a terminal on any Rack, prints `ERROR: cannot exec into Fargate task without ECS Exec enabled; set ECSExec=Yes on the rack and redeploy`. Without a terminal, setting `ECSExec=Yes` does not help; run the command from a terminal. The CLI exits with status 0 after an error the Rack reports during the session, such as this one or an unknown Process id, so a script must check the output as well as the exit code.

## See Also

- [run](/reference/cli-commands/run)
- [cp](/reference/cli-commands/cp)
- [ps](/reference/cli-commands/ps)
- [One-off Commands](/management/one-off-commands)
- [Debugging](/management/debugging)
