---
title: "One-off Commands"
---

Convox allows you to execute individual commands on your containers in several different contexts.

<div class="block-callout block-show-callout type-info" markdown="1">

**To run a command on a container in your Convox rack...**

* in a new container: `convox run [service] [command]`
* in a running container: `convox exec [container ID] [command]`

**To run a command in a local container...**

* in a new container: `convox start [service] [command]`
* in a running container: There is no `convox` command for this. Use [`docker exec` instead](https://docs.docker.com/engine/reference/commandline/exec/).
</div>

We'll go into more detail about each of these commands below.


## convox run

**Syntax:** `convox run [service] [command]`
**Use:** Spawns a new container in your Convox rack to run the desired command.

**Examples:**

```
$ convox run web ls
Dockerfile  assets  bin  counter.py  requirements.txt  templates
```

You can spawn an interactive shell:

```
$ convox run web bash
root@eafc101d0980:/src#
```

Or run one-off administrative tasks:

```
$ convox run web bin/migrate
Migrating database... Done
```

<div class="block-callout block-show-callout type-info" markdown="1">
#### Detached processes

By default, `convox run` processes are run interactively ("attached"). This means you can interact with the running process via the terminal (stdin and stdout). Attached processes have a timeout of 1 hour.

If you don't need to run a process interactively, you can add the `--detach` option to your `convox run` command:

```
$ convox run web --detach bin/long_running_process
```

Detached processes have no timeout and their output is available in the application logs.
</div>


## convox exec

**Syntax:** `convox exec [container ID] [command]`
**Use:** Attaches to and executes the command in an existing container on your Convox rack.

<div class="block-callout block-show-callout type-info" markdown="1">
Use `convox ps` to get the container ID:

```
$ convox ps
ID            NAME  RELEASE      SIZE  STARTED     COMMAND
310481bf223f  web   RSPZQWVWGOP  256   5 days ago  bin/web
5e3c8576b942  web   RSPZQWVWGOP  256   4 days ago  bin/web
```
</div>

**Examples:**
You can spawn an interactive shell:

```
$ convox exec 5e3c8576b942 bash
/app #
```

or run a one-off command:

```
$ convox exec 0105c73ce736 ps
  PID TTY          TIME CMD
   12 ?        00:00:00 sh
   16 ?        00:00:00 ps
```

```
$ convox exec 0105c73ce736 env
REDIS_PORT=6379
RACK=python
HOSTNAME=13e67aec1985
FOO=bar
BAZ=qux
[snip]
```
