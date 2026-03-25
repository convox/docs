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

## See Also

- [run](/reference/cli-commands/run)
- [cp](/reference/cli-commands/cp)
- [ps](/reference/cli-commands/ps)
- [One-off Commands](/management/one-off-commands)
- [Debugging](/management/debugging)
