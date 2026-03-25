---
title: CLI Commands
description: Complete reference of Convox CLI commands for managing Racks, Apps, Builds, Resources, and infrastructure.
---

# CLI Commands

The Convox CLI provides commands for managing your Racks, Apps, Builds, Services, Resources, and infrastructure. Run `convox help` to see all available commands, or `convox <command> -h` for help on a specific command.

## Authentication

| Command | Description |
|---------|-------------|
| `login` | Authenticate with a Rack |
| `switch` | Switch current Rack |

## App Management

| Command | Description |
|---------|-------------|
| `apps` | List Apps |
| `apps cancel` | Cancel an App update |
| `apps create` | Create an App |
| `apps delete` | Delete an App |
| `apps export` | Export an App |
| `apps import` | Import an App |
| `apps info` | Get information about an App |
| `apps lock` | Enable termination protection |
| `apps params` | Display App parameters |
| `apps params set` | Set App parameters |
| `apps unlock` | Disable termination protection |
| `apps wait` | Wait for an App to finish updating |

## Builds

| Command | Description |
|---------|-------------|
| `build` | Create a Build |
| `builds` | List Builds |
| `builds export` | Export a Build |
| `builds import` | Import a Build |
| `builds info` | Get information about a Build |
| `builds logs` | Get logs for a Build |

## Deployment

| Command | Description |
|---------|-------------|
| `deploy` | Create and promote a Build |
| `releases` | List Releases for an App |
| `releases info` | Get information about a Release |
| `releases manifest` | Get manifest for a Release |
| `releases promote` | Promote a Release |
| `releases rollback` | Copy an old Release forward and promote it |

## Environment Variables

| Command | Description |
|---------|-------------|
| `env` | List env vars |
| `env edit` | Edit env vars interactively |
| `env get` | Get an env var |
| `env set` | Set env var(s) |
| `env unset` | Unset env var(s) |

## Processes

| Command | Description |
|---------|-------------|
| `ps` | List App processes |
| `ps info` | Get information about a process |
| `ps stop` | Stop a process |
| `exec` | Execute a command in a running process |
| `run` | Execute a command in a new process |
| `cp` | Copy files to/from a process |

## Services

| Command | Description |
|---------|-------------|
| `services` | List Services for an App |
| `services restart` | Restart a Service |
| `restart` | Restart all processes for an App |
| `scale` | Scale a Service |
| `start` | Start an application for local development |

## Certificates

| Command | Description |
|---------|-------------|
| `certs` | List certificates |
| `certs delete` | Delete a certificate |
| `certs generate` | Generate a certificate |
| `certs import` | Import a certificate |
| `ssl` | List certificate associations for an App |
| `ssl update` | Update certificate for an App |

## Logs

| Command | Description |
|---------|-------------|
| `logs` | Get logs for an App |
| `rack logs` | Get logs for the Rack |

## Instances

| Command | Description |
|---------|-------------|
| `instances` | List instances |
| `instances keyroll` | Roll SSH key on instances |
| `instances ssh` | Run a shell on an instance |
| `instances terminate` | Terminate an instance |

## Rack Management

| Command | Description |
|---------|-------------|
| `rack` | Get information about the Rack |
| `rack install` | Install a Rack |
| `rack params` | Display Rack parameters |
| `rack params set` | Set Rack parameters |
| `rack ps` | List Rack processes |
| `rack releases` | List Rack version history |
| `rack scale` | Scale the Rack |
| `rack sync` | Sync v2 Rack API URL |
| `rack uninstall` | Uninstall a Rack |
| `rack update` | Update the Rack |
| `rack wait` | Wait for Rack to finish updating |
| `racks` | List available Racks |

## Rack Access

| Command | Description |
|---------|-------------|
| `rack access` | Generate Rack access credentials |
| `rack access key rotate` | Rotate Rack access key |
| `rack runtimes` | List attachable runtime integrations |
| `rack runtime attach` | Attach a runtime integration to the Rack |
| `rack sync whitelist instances ip` | Sync whitelisted instance IPs in security group |

## App Resources

| Command | Description |
|---------|-------------|
| `resources` | List Resources for an App |
| `resources info` | Get information about a Resource |
| `resources proxy` | Proxy a local port to a Resource |
| `resources url` | Get URL for a Resource |

## Rack Resources

| Command | Description |
|---------|-------------|
| `rack resources` | List Rack Resources |
| `rack resources create` | Create a Resource |
| `rack resources delete` | Delete a Resource |
| `rack resources info` | Get information about a Resource |
| `rack resources link` | Link a Resource to an App |
| `rack resources options` | List options for a Resource type |
| `rack resources proxy` | Proxy a local port to a Rack Resource |
| `rack resources types` | List Resource types |
| `rack resources unlink` | Unlink a Resource from an App |
| `rack resources update` | Update Resource options |
| `rack resources url` | Get URL for a Resource |

## Registries

| Command | Description |
|---------|-------------|
| `registries` | List private registries |
| `registries add` | Add a private registry |
| `registries remove` | Remove a private registry |

## Networking

| Command | Description |
|---------|-------------|
| `proxy` | Proxy a connection inside the Rack |

## Utilities

| Command | Description |
|---------|-------------|
| `api get` | Query the Rack API |
| `test` | Run tests |
| `update` | Update the CLI |
| `version` | Display version information |

## Command Details

### `deploy`

```bash
$ convox deploy [dir]
```

Create a Build from the specified directory (default: `.`) and promote the resulting Release.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--build-args` | | Docker build arguments (can be specified multiple times) |
| `--description` | `-d` | Build description |
| `--development` | | Create a development Build |
| `--id` | | Send logs to stderr, release ID to stdout |
| `--manifest` | `-m` | Path to manifest file |
| `--no-cache` | | Build without Docker cache |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |
| `--wildcard-domain` | | Enable wildcard domain |

### `build`

```bash
$ convox build [dir]
```

Create a Build from the specified directory (default: `.`) without promoting.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--build-args` | | Docker build arguments (can be specified multiple times) |
| `--description` | `-d` | Build description |
| `--development` | | Create a development Build |
| `--id` | | Send logs to stderr, release ID to stdout |
| `--manifest` | `-m` | Path to manifest file |
| `--no-cache` | | Build without Docker cache |
| `--rack` | `-r` | Rack name |
| `--wildcard-domain` | | Enable wildcard domain |

### `builds`

```bash
$ convox builds
```

List Builds for an App.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--limit` | `-l` | Number of Builds to list |
| `--rack` | `-r` | Rack name |

### `run`

```bash
$ convox run <service> <command>
```

Execute a command in a new Process.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--detach` | `-d` | Run Process in the background |
| `--entrypoint` | `-e` | Use the container entrypoint (default: `true`) |
| `--rack` | `-r` | Rack name |
| `--release` | | Run against a specific Release |
| `--timeout` | `-t` | Timeout in seconds |

### `env`

```bash
$ convox env
```

List environment variables for an App.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

### `env edit`

```bash
$ convox env edit
```

Edit environment variables interactively using `$EDITOR`.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--promote` | `-p` | Promote the resulting Release |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

### `env get`

```bash
$ convox env get <var>
```

Get the value of a single environment variable.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

### `env set`

```bash
$ convox env set <key=value> [key=value]...
```

Set one or more environment variables. Values can also be piped via stdin.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--id` | | Send logs to stderr, release ID to stdout |
| `--promote` | `-p` | Promote the resulting Release |
| `--rack` | `-r` | Rack name |
| `--replace` | | Replace all environment variables with the given ones |
| `--wait` | `-w` | Wait for completion |

### `env unset`

```bash
$ convox env unset <key> [key]...
```

Remove one or more environment variables.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--id` | | Send logs to stderr, release ID to stdout |
| `--promote` | `-p` | Promote the resulting Release |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

### `rack`

```bash
$ convox rack
```

Display information about the current Rack (name, provider, region, router, status, version).

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `apps`

```bash
$ convox apps
```

List all Apps on the current Rack.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `resources`

```bash
$ convox resources
```

List Resources for an App.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

### `services`

```bash
$ convox services
```

List Services for an App.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

### `releases`

```bash
$ convox releases
```

List Releases for an App.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--limit` | `-l` | Number of Releases to list |
| `--rack` | `-r` | Rack name |

### `logs`

```bash
$ convox logs
```

Stream logs for an App.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--filter` | | Filter logs by a pattern |
| `--no-follow` | | Do not follow log output |
| `--rack` | `-r` | Rack name |
| `--since` | | Show logs since a duration (default: `2m`) |

### `ps`

```bash
$ convox ps
```

List running Processes for an App.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |
| `--release` | | Filter by Release |
| `--service` | `-s` | Filter by Service |

### `scale`

```bash
$ convox scale [service]
```

Display or update scaling parameters for Services. When called without `--count`, `--cpu`, or `--memory`, displays current scaling for all Services.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--count` | | Desired Process count |
| `--cpu` | | CPU units (1024 = 1 vCPU) |
| `--memory` | | Memory in MB |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

### `login`

```bash
$ convox login [hostname]
```

Authenticate with a Rack or Console. Defaults to `console.convox.com` if no hostname is specified.

| Flag | Short | Description |
|------|-------|-------------|
| `--password` | `-p` | Password |

### `switch`

```bash
$ convox switch [rack]
```

Switch the current Rack. When called without an argument, displays the current Rack name.

*This command has no flags.*

### `apps cancel`

```bash
$ convox apps cancel [app]
```

Cancel an in-progress App update and roll back to the last active Release.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

### `apps create`

```bash
$ convox apps create [name]
```

Create a new App on the current Rack.

| Flag | Short | Description |
|------|-------|-------------|
| `--generation` | `-g` | App generation (default: `2`) |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

### `apps delete`

```bash
$ convox apps delete <app>
```

Delete an App.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

### `apps export`

```bash
$ convox apps export [app]
```

Export an App (including Build and environment) to a tarball.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--file` | `-f` | Export to file |
| `--rack` | `-r` | Rack name |

### `apps import`

```bash
$ convox apps import [app]
```

Import an App from a previously exported tarball.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--file` | `-f` | Import from file |
| `--rack` | `-r` | Rack name |

### `apps info`

```bash
$ convox apps info [app]
```

Display information about an App (name, status, generation, locked, release, router).

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

### `apps lock`

```bash
$ convox apps lock [app]
```

Enable termination protection for an App.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

### `apps unlock`

```bash
$ convox apps unlock [app]
```

Disable termination protection for an App.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

### `apps params`

```bash
$ convox apps params [app]
```

Display App parameters and their current values.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

### `apps params set`

```bash
$ convox apps params set <Key=Value> [Key=Value]...
```

Set one or more App parameters.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

### `apps wait`

```bash
$ convox apps wait [app]
```

Wait for an App to finish updating.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

### `builds export`

```bash
$ convox builds export <build>
```

Export a Build to a tarball.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--file` | `-f` | Export to file |
| `--rack` | `-r` | Rack name |

### `builds import`

```bash
$ convox builds import
```

Import a Build from a previously exported tarball. Reads from stdin or a file.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--file` | `-f` | Import from file |
| `--id` | | Send logs to stderr, release ID to stdout |
| `--rack` | `-r` | Rack name |

### `builds info`

```bash
$ convox builds info <build>
```

Display information about a Build (ID, status, release, description, started, elapsed).

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

### `builds logs`

```bash
$ convox builds logs <build>
```

Get logs for a Build.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

### `releases info`

```bash
$ convox releases info <release>
```

Display information about a Release (ID, build, created, description, env).

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

### `releases manifest`

```bash
$ convox releases manifest <release>
```

Display the manifest for a Release.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

### `releases promote`

```bash
$ convox releases promote [release]
```

Promote a Release. If no Release is specified, promotes the most recent Release.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

### `releases rollback`

```bash
$ convox releases rollback <release>
```

Copy an old Release forward (preserving its Build and environment) and promote the new Release.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--id` | | Send logs to stderr, release ID to stdout |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

### `exec`

```bash
$ convox exec <pid> <command>
```

Execute a command in a running Process.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

### `ps info`

```bash
$ convox ps info <pid>
```

Display information about a Process (ID, app, command, instance, release, service, started, status).

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

### `ps stop`

```bash
$ convox ps stop <pid>
```

Stop a running Process.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

### `cp`

```bash
$ convox cp <[pid:]src> <[pid:]dst>
```

Copy files to or from a Process. Prefix a path with `pid:` to reference a running Process. Process paths must be absolute.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

### `services restart`

```bash
$ convox services restart <service>
```

Restart all Processes for a specific Service.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

### `restart`

```bash
$ convox restart
```

Restart all Processes for all Services of an App.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

### `start`

```bash
$ convox start [service] [service...]
```

Start an application for local development.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--generation` | `-g` | Generation |
| `--manifest` | `-m` | Manifest file |
| `--no-build` | | Skip build |
| `--no-cache` | | Build without layer cache |
| `--no-sync` | | Do not sync local changes into running containers |
| `--rack` | `-r` | Rack name |
| `--shift` | `-s` | Shift local port numbers (generation 1 only) |

### `certs`

```bash
$ convox certs
```

List certificates on the Rack.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `certs delete`

```bash
$ convox certs delete <cert>
```

Delete a certificate.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `certs generate`

```bash
$ convox certs generate <domain> [domain...]
```

Generate a certificate for one or more domains.

| Flag | Short | Description |
|------|-------|-------------|
| `--id` | | Send logs to stderr, certificate ID to stdout |
| `--rack` | `-r` | Rack name |

### `certs import`

```bash
$ convox certs import <pub> <key>
```

Import a certificate from public certificate and private key files.

| Flag | Short | Description |
|------|-------|-------------|
| `--chain` | | Intermediate certificate chain file |
| `--id` | | Send logs to stderr, certificate ID to stdout |
| `--rack` | `-r` | Rack name |

### `ssl`

```bash
$ convox ssl
```

List certificate associations for an App.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

### `ssl update`

```bash
$ convox ssl update <process:port> <certificate>
```

Update the certificate for an App endpoint (generation 1 only).

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

### `rack logs`

```bash
$ convox rack logs
```

Stream logs for the Rack.

| Flag | Short | Description |
|------|-------|-------------|
| `--filter` | | Filter logs by a pattern |
| `--no-follow` | | Do not follow log output |
| `--rack` | `-r` | Rack name |
| `--since` | | Show logs since a duration (default: `2m`) |

### `instances`

```bash
$ convox instances
```

List Rack instances.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `instances keyroll`

```bash
$ convox instances keyroll
```

Roll the SSH key on all instances.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

### `instances ssh`

```bash
$ convox instances ssh <instance> [command]
```

Run a shell on an instance, or execute a command.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `instances terminate`

```bash
$ convox instances terminate <instance>
```

Terminate an instance.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `rack install`

```bash
$ convox rack install <type> [Parameter=Value]...
```

Install a new Rack.

| Flag | Short | Description |
|------|-------|-------------|
| `--name` | `-n` | Rack name |
| `--raw` | | Raw output |
| `--version` | `-v` | Rack version |

### `rack params`

```bash
$ convox rack params
```

Display Rack parameters and their current values.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `rack params set`

```bash
$ convox rack params set <Key=Value> [Key=Value]...
```

Set one or more Rack parameters.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

### `rack ps`

```bash
$ convox rack ps
```

List Rack Processes.

| Flag | Short | Description |
|------|-------|-------------|
| `--all` | `-a` | Show all Processes |
| `--rack` | `-r` | Rack name |

### `rack releases`

```bash
$ convox rack releases
```

List Rack version history.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `rack scale`

```bash
$ convox rack scale
```

Display or update Rack scaling parameters. When called without flags, displays current scaling (autoscale, count, status, type).

| Flag | Short | Description |
|------|-------|-------------|
| `--count` | `-c` | Instance count |
| `--rack` | `-r` | Rack name |
| `--type` | `-t` | Instance type |

### `rack sync`

```bash
$ convox rack sync
```

Sync v2 Rack API URL.

| Flag | Short | Description |
|------|-------|-------------|
| `--name` | `-n` | Rack name (displays the Rack API URL) |
| `--rack` | `-r` | Rack name |

### `rack uninstall`

```bash
$ convox rack uninstall <type> <name>
```

Uninstall a Rack.

| Flag | Short | Description |
|------|-------|-------------|
| `--force` | `-f` | Force uninstall (required for non-interactive use) |

### `rack update`

```bash
$ convox rack update [version]
```

Update the Rack to the next version, or to a specific version if specified.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

### `rack wait`

```bash
$ convox rack wait
```

Wait for the Rack to finish updating.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `rack access`

```bash
$ convox rack access
```

Generate Rack access credentials.

| Flag | Short | Description |
|------|-------|-------------|
| `--duration-in-hours` | | Duration in hours |
| `--rack` | `-r` | Rack name |
| `--role` | | Access role: `read` or `write` |

### `rack access key rotate`

```bash
$ convox rack access key rotate
```

Rotate the Rack access key.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `rack runtimes`

```bash
$ convox rack runtimes
```

List attachable runtime integrations.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `rack runtime attach`

```bash
$ convox rack runtime attach <runtime>
```

Attach a runtime integration to the Rack.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `rack sync whitelist instances ip`

```bash
$ convox rack sync whitelist instances ip
```

Sync whitelisted instance IPs in the security group.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `resources info`

```bash
$ convox resources info <resource>
```

Get information about an App Resource (name, type, URL).

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

### `resources proxy`

```bash
$ convox resources proxy <resource>
```

Proxy a local port to an App Resource.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--port` | `-p` | Local port |
| `--rack` | `-r` | Rack name |
| `--tls` | `-t` | Wrap connection in TLS |

### `resources url`

```bash
$ convox resources url <resource>
```

Get the URL for an App Resource.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

### `rack resources`

```bash
$ convox rack resources
```

List Rack Resources.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `rack resources create`

```bash
$ convox rack resources create <type> [Option=Value]...
```

Create a Rack Resource.

| Flag | Short | Description |
|------|-------|-------------|
| `--name` | `-n` | Resource name |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

### `rack resources delete`

```bash
$ convox rack resources delete <name>
```

Delete a Rack Resource.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

### `rack resources info`

```bash
$ convox rack resources info <resource>
```

Get information about a Rack Resource (name, type, status, options, URL, linked apps).

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `rack resources link`

```bash
$ convox rack resources link <resource>
```

Link a Rack Resource to an App.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

### `rack resources options`

```bash
$ convox rack resources options <type>
```

List available options for a Resource type.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `rack resources proxy`

```bash
$ convox rack resources proxy <resource>
```

Proxy a local port to a Rack Resource.

| Flag | Short | Description |
|------|-------|-------------|
| `--port` | `-p` | Local port |
| `--rack` | `-r` | Rack name |
| `--tls` | `-t` | Wrap connection in TLS |

### `rack resources types`

```bash
$ convox rack resources types
```

List available Resource types.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `rack resources unlink`

```bash
$ convox rack resources unlink <resource>
```

Unlink a Rack Resource from an App.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

### `rack resources update`

```bash
$ convox rack resources update <name> [Option=Value]...
```

Update options for a Rack Resource.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

### `rack resources url`

```bash
$ convox rack resources url <resource>
```

Get the URL for a Rack Resource.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `registries`

```bash
$ convox registries
```

List private registries configured on the Rack.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `registries add`

```bash
$ convox registries add <server> <username> <password>
```

Add a private registry.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `registries remove`

```bash
$ convox registries remove <server>
```

Remove a private registry.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `proxy`

```bash
$ convox proxy <[port:]host:hostport> [[port:]host:hostport]...
```

Proxy one or more local ports to hosts inside the Rack. If no local port is specified, the host port is used.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |
| `--tls` | `-t` | Wrap connection in TLS |

### `api get`

```bash
$ convox api get <path>
```

Query the Rack API directly and return JSON output.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `test`

```bash
$ convox test [dir]
```

Run tests defined in the manifest (`test:` directive on Services). Creates a development Build (or uses an existing Release) and executes each Service's test command.

| Flag | Short | Description |
|------|-------|-------------|
| `--app` | `-a` | App name |
| `--description` | `-d` | Build description |
| `--rack` | `-r` | Rack name |
| `--release` | | Use existing Release to run tests |
| `--timeout` | `-t` | Timeout in seconds |

### `update`

```bash
$ convox update [version]
```

Update the CLI to the latest version, or to a specific version if specified.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `version`

```bash
$ convox version
```

Display client and server version information.

| Flag | Short | Description |
|------|-------|-------------|
| `--rack` | `-r` | Rack name |

### `racks`

```bash
$ convox racks
```

List all available Racks.

*This command has no flags.*

## See Also

- [Rack Parameters](/reference/rack-parameters)
- [App Parameters](/reference/app-parameters)
- [Getting Started](/introduction/getting-started)
