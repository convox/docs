---
title: "Debugging"
description: "Debug running Convox applications using logs, exec, SSH access, and service restarts."
---

# Debugging

Convox provides several tools to debug your running applications.

### convox logs

You can view all of the output of your application's processes using `convox logs`:

```bash
$ convox logs
2025-01-15 14:32:00 i-0a34d285 service/web:RSPZQWVWGOP/5e3c8576b942 : 10.0.1.242 - - [15/Jan/2025:14:32:00 +0000] "GET / HTTP/1.1" 200 70 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"
2025-01-15 14:32:00 i-0a34d285 service/web:RSPZQWVWGOP/5e3c8576b942 : 10.0.1.242 - - [15/Jan/2025:14:32:00 +0000] "GET / HTTP/1.0" 200 70 0.0019
```

### convox exec

You can get an active shell into a running process using `convox exec`. First, find the ID of the process:

```bash
$ convox ps
ID            NAME  RELEASE      SIZE  STARTED     COMMAND
310481bf223f  web   RSPZQWVWGOP  256   5 days ago  bin/web
5e3c8576b942  web   RSPZQWVWGOP  256   4 days ago  bin/web
```

```bash
$ convox exec 5e3c8576b942 bash
/app #
```
    
You can use this shell to examine your running application in detail.

### convox instances ssh

You can also gain access to the individual instances that make up your Rack. First, find the ID of the instance:

```bash
$ convox instances
ID          AGENT  STATUS  STARTED      PS  CPU    MEM
i-0234d285  on     active  6 days ago   5   0.00%  12.36%
i-14cd0a89  on     active  7 hours ago  2   0.00%  5.18%
i-d8740c43  on     active  5 days ago   4   0.00%  21.04%
```

```bash
$ convox instances ssh i-0234d285
[ec2-user@ip-10-0-3-209 ~]$
```

> **Note:** Before you can use `convox instances ssh` the first time, you will need to run `convox instances keyroll` to load an SSH key onto your Rack instances.

### convox restart

If all else fails and you wish to restart all running instances of your app you can perform a remote restart from the CLI with:

```bash
$ convox restart -a app1
```

Or to restart only the `web` service processes, run:

```bash
$ convox services restart web -a app1
```

## See Also

- [Logs](/management/logs)
- [Logging Integrations](/integrations/logging)
- [One-off Commands](/management/one-off-commands)
- [Running Locally](/development/running-locally)
- [Application Monitoring](/management/application)
