---
title: "One-off Commands"
description: "Run one-off commands on Convox containers using convox run, exec, and cp for remote execution and file transfers."
---

# One-off Commands

Convox allows you to execute individual commands on your containers in several different contexts.

> **To run a command on a container in your Convox Rack...**
>
> * in a new container: `convox run [service] [command]`
> * in a running container: `convox exec [container ID] [command]`
>
> **To run a command in a local container...**
>
> * in a new container: `convox start [service] [command]`
> * in a running container: There is no `convox` command for this. Use [`docker exec` instead](https://docs.docker.com/engine/reference/commandline/exec/).

Each of these commands is described in more detail below.


## convox run

**Syntax:** `convox run [service] [command]`
**Use:** Spawns a new container in your Convox rack to run the desired command.

**Examples:**

```bash
$ convox run web ls
Dockerfile  assets  bin  counter.py  requirements.txt  templates
```

You can spawn an interactive shell:

```bash
$ convox run web bash
root@eafc101d0980:/src#
```

Or run one-off administrative tasks:

```bash
$ convox run web bin/migrate
Migrating database... Done
```

> **Note: Detached processes and timeouts**
>
> By default, `convox run` processes are run interactively ("attached"). This means you can interact with the running process via the terminal (stdin and stdout). Attached processes have a timeout of 1 hour. Use `--timeout [seconds]` to increase this:
>
> ```bash
> $ convox run web --timeout 7200 bin/long_running_process
> ```
>
> If you do not need to run a process interactively, you can add the `--detach` option to your `convox run` command:
>
> ```bash
> $ convox run web --detach bin/long_running_process
> ```
>
> Detached processes have no timeout and their output is available in the application logs.


## convox exec

**Syntax:** `convox exec [container ID] [command]`
**Use:** Attaches to and executes the command in an existing container on your Convox rack.

> **Note:** Use `convox ps` to get the container ID:
>
> ```bash
> $ convox ps
> ID            NAME  RELEASE      SIZE  STARTED     COMMAND
> 310481bf223f  web   RSPZQWVWGOP  256   5 days ago  bin/web
> 5e3c8576b942  web   RSPZQWVWGOP  256   4 days ago  bin/web
> ```

**Examples:**
You can spawn an interactive shell:

```bash
$ convox exec 5e3c8576b942 bash
/app #
```

or run a one-off command:

```bash
$ convox exec 0105c73ce736 ps
  PID TTY          TIME CMD
   12 ?        00:00:00 sh
   16 ?        00:00:00 ps
```

```bash
$ convox exec 0105c73ce736 env
REDIS_PORT=6379
RACK=python
HOSTNAME=13e67aec1985
FOO=bar
BAZ=qux
[snip]
```

## convox cp

**Syntax:** `convox cp [relative local file path] [container ID]:[remote absolute file path]`
**Use:** Copies from a local machine to a remote machine

> **Note:** Use `convox ps` to get the container ID:
>
> ```bash
> $ convox ps
> ID            NAME  RELEASE      SIZE  STARTED     COMMAND
> 310481bf223f  web   RSPZQWVWGOP  256   5 days ago  bin/web
> 5e3c8576b942  web   RSPZQWVWGOP  256   4 days ago  bin/web
> ```

 **Examples:**
Copy a file from a local machine to a remote container:

 ```bash
$ convox cp foo/bar.txt 5e3c8576b942:/app/foo/bar.txt
```

## See Also

- [Debugging](/management/debugging)
- [Running Migrations](/deployment/running-migrations)
- [Running Locally](/development/running-locally)
- [Logs](/management/logs)
