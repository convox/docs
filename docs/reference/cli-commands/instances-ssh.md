---
title: "instances ssh"
description: "Run a shell on an instance, or execute a command."
---

# instances ssh

Run a shell on an instance, or execute a command. Provides direct access to the underlying EC2 instance for debugging and inspection. If no command is given, an interactive shell is opened.

## Syntax

```bash
$ convox instances ssh <instance> [command]
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox instances ssh i-0a1b2c3d
Last login: Wed Jan 15 12:00:00 2025
[ec2-user@ip-10-0-1-50 ~]$

$ convox instances ssh i-0a1b2c3d "df -h /"
Filesystem      Size  Used Avail Use% Mounted on
/dev/xvda1       50G   12G   38G  25% /
```

## See Also

- [instances](/reference/cli-commands/instances)
- [exec](/reference/cli-commands/exec)
- [Debugging](/management/debugging)
