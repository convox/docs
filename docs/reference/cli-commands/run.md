---
title: "run"
description: "Execute a command in a new Process."
---

# run

Execute a command in a new Process. Unlike `exec`, this starts a fresh container from the current Release. Use `--detach` to run the Process in the background, which is useful for long-running tasks like database migrations.

## Syntax

```bash
$ convox run <service> <command>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--detach` | `-d` | Run Process in the background |
| `--entrypoint` | `-e` | Use the container entrypoint (default: `true`) |
| `--rack` | `-r` | Rack name |
| `--release` | | Run against a specific Release |
| `--timeout` | `-t` | Timeout in seconds |

## Example Usage

```bash
$ convox run web bin/rails db:migrate -a myapp
== 20250115120000 AddUsersTable: migrating ====================================
-- create_table(:users)
   -> 0.0123s
== 20250115120000 AddUsersTable: migrated (0.0124s) ===========================
```

```bash
$ convox run web bin/report --detach -a myapp
Running detached process... OK, web-pqr3456-stu7
```

## See Also

- [exec](/reference/cli-commands/exec)
- [cp](/reference/cli-commands/cp)
- [ps](/reference/cli-commands/ps)
- [One-off Commands](/management/one-off-commands)
- [Debugging](/management/debugging)
