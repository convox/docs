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

## Passing a Command

Everything after the Service name is joined into a single command string, which the new Process runs with `sh -c`. Quoting is not preserved across that join, so a command that contains shell syntax must be passed as one quoted argument:

```bash
$ convox run web 'bin/migrate && bin/seed' -a myapp
Migrating database... Done
Seeding database... Done
```

A command with no shell syntax needs no quoting, which is why the examples above pass `bin/rails db:migrate` directly.

If your command takes flags of its own, put `--` after the `convox` flags so the CLI stops parsing:

```bash
$ convox run web -a myapp -- bin/report --verbose
```

Without `--`, the CLI claims the flag for itself:

```bash
$ convox run web bin/report --verbose -a myapp
ERROR: unknown flag: --verbose
```

Do not add an `sh -c` of your own. The command string is already run through a shell, so an inner `sh -c` nests a second shell and produces confusing output rather than an error. See [exec](/reference/cli-commands/exec) for a worked example.

## See Also

- [exec](/reference/cli-commands/exec)
- [cp](/reference/cli-commands/cp)
- [ps](/reference/cli-commands/ps)
- [One-off Commands](/management/one-off-commands)
- [Debugging](/management/debugging)
