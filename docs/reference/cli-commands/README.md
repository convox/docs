---
title: "CLI Commands"
description: "Complete reference of Convox CLI commands for managing Racks, Apps, Builds, Resources, and infrastructure."
---

# CLI Commands

The Convox CLI provides commands for managing your Racks, Apps, Builds, Services, Resources, and infrastructure. Run `convox help` to see all available commands, or `convox <command> -h` for help on a specific command.

## Authentication

| Command | Description |
|:--------|:------------|
| [login](/reference/cli-commands/login) | Authenticate with a Rack |
| [switch](/reference/cli-commands/switch) | Switch current Rack |

## App Management

| Command | Description |
|:--------|:------------|
| [apps](/reference/cli-commands/apps) | List Apps |
| [apps cancel](/reference/cli-commands/apps-cancel) | Cancel an App update |
| [apps create](/reference/cli-commands/apps-create) | Create an App |
| [apps delete](/reference/cli-commands/apps-delete) | Delete an App |
| [apps export](/reference/cli-commands/apps-export) | Export an App |
| [apps import](/reference/cli-commands/apps-import) | Import an App |
| [apps info](/reference/cli-commands/apps-info) | Get information about an App |
| [apps lock](/reference/cli-commands/apps-lock) | Enable termination protection |
| [apps params](/reference/cli-commands/apps-params) | Display App parameters |
| [apps params set](/reference/cli-commands/apps-params-set) | Set App parameters |
| [apps unlock](/reference/cli-commands/apps-unlock) | Disable termination protection |
| [apps wait](/reference/cli-commands/apps-wait) | Wait for an App to finish updating |

## Builds

| Command | Description |
|:--------|:------------|
| [build](/reference/cli-commands/build) | Create a Build |
| [builds](/reference/cli-commands/builds) | List Builds |
| [builds export](/reference/cli-commands/builds-export) | Export a Build |
| [builds import](/reference/cli-commands/builds-import) | Import a Build |
| [builds info](/reference/cli-commands/builds-info) | Get information about a Build |
| [builds logs](/reference/cli-commands/builds-logs) | Get logs for a Build |

## Deployment

| Command | Description |
|:--------|:------------|
| [deploy](/reference/cli-commands/deploy) | Create and promote a Build |
| [releases](/reference/cli-commands/releases) | List Releases for an App |
| [releases info](/reference/cli-commands/releases-info) | Get information about a Release |
| [releases manifest](/reference/cli-commands/releases-manifest) | Get manifest for a Release |
| [releases promote](/reference/cli-commands/releases-promote) | Promote a Release |
| [releases rollback](/reference/cli-commands/releases-rollback) | Copy an old Release forward and promote it |

## Environment Variables

| Command | Description |
|:--------|:------------|
| [env](/reference/cli-commands/env) | List env vars |
| [env edit](/reference/cli-commands/env-edit) | Edit env vars interactively |
| [env get](/reference/cli-commands/env-get) | Get an env var |
| [env set](/reference/cli-commands/env-set) | Set env var(s) |
| [env unset](/reference/cli-commands/env-unset) | Unset env var(s) |

## Processes

| Command | Description |
|:--------|:------------|
| [ps](/reference/cli-commands/ps) | List App processes |
| [ps info](/reference/cli-commands/ps-info) | Get information about a process |
| [ps stop](/reference/cli-commands/ps-stop) | Stop a process |
| [exec](/reference/cli-commands/exec) | Execute a command in a running process |
| [run](/reference/cli-commands/run) | Execute a command in a new process |
| [cp](/reference/cli-commands/cp) | Copy files to/from a process |

## Services

| Command | Description |
|:--------|:------------|
| [services](/reference/cli-commands/services) | List Services for an App |
| [services restart](/reference/cli-commands/services-restart) | Restart a Service |
| [restart](/reference/cli-commands/restart) | Restart all processes for an App |
| [scale](/reference/cli-commands/scale) | Scale a Service |
| [start](/reference/cli-commands/start) | Start an App for local development |

## Certificates

| Command | Description |
|:--------|:------------|
| [certs](/reference/cli-commands/certs) | List certificates |
| [certs delete](/reference/cli-commands/certs-delete) | Delete a certificate |
| [certs generate](/reference/cli-commands/certs-generate) | Generate a certificate |
| [certs import](/reference/cli-commands/certs-import) | Import a certificate |
| [ssl](/reference/cli-commands/ssl) | List certificate associations for an App |
| [ssl update](/reference/cli-commands/ssl-update) | Update certificate for an App |

## Logs

| Command | Description |
|:--------|:------------|
| [logs](/reference/cli-commands/logs) | Get logs for an App |
| [rack logs](/reference/cli-commands/rack-logs) | Get logs for the Rack |

## Instances

| Command | Description |
|:--------|:------------|
| [instances](/reference/cli-commands/instances) | List instances |
| [instances keyroll](/reference/cli-commands/instances-keyroll) | Roll SSH key on instances |
| [instances ssh](/reference/cli-commands/instances-ssh) | Run a shell on an instance |
| [instances terminate](/reference/cli-commands/instances-terminate) | Terminate an instance |

## Rack Management

| Command | Description |
|:--------|:------------|
| [rack](/reference/cli-commands/rack) | Get information about the Rack |
| [rack install](/reference/cli-commands/rack-install) | Install a Rack |
| [rack params](/reference/cli-commands/rack-params) | Display Rack parameters |
| [rack params set](/reference/cli-commands/rack-params-set) | Set Rack parameters |
| [rack ps](/reference/cli-commands/rack-ps) | List Rack processes |
| [rack releases](/reference/cli-commands/rack-releases) | List Rack version history |
| [rack scale](/reference/cli-commands/rack-scale) | Scale the Rack |
| [rack sync](/reference/cli-commands/rack-sync) | Sync v2 Rack API URL |
| [rack uninstall](/reference/cli-commands/rack-uninstall) | Uninstall a Rack |
| [rack update](/reference/cli-commands/rack-update) | Update the Rack |
| [rack wait](/reference/cli-commands/rack-wait) | Wait for Rack to finish updating |
| [racks](/reference/cli-commands/racks) | List available Racks |

## Rack Access

| Command | Description |
|:--------|:------------|
| [rack access](/reference/cli-commands/rack-access) | Generate Rack access credentials |
| [rack access key rotate](/reference/cli-commands/rack-access-key-rotate) | Rotate Rack access key |
| [rack runtimes](/reference/cli-commands/rack-runtimes) | List attachable runtime integrations |
| [rack runtime attach](/reference/cli-commands/rack-runtime-attach) | Attach a runtime integration to the Rack |
| [rack sync whitelist instances ip](/reference/cli-commands/rack-sync-whitelist-instances-ip) | Sync whitelisted instance IPs in security group |

## App Resources

| Command | Description |
|:--------|:------------|
| [resources](/reference/cli-commands/resources) | List Resources for an App |
| [resources info](/reference/cli-commands/resources-info) | Get information about a Resource |
| [resources proxy](/reference/cli-commands/resources-proxy) | Proxy a local port to a Resource |
| [resources url](/reference/cli-commands/resources-url) | Get URL for a Resource |

## Rack Resources

| Command | Description |
|:--------|:------------|
| [rack resources](/reference/cli-commands/rack-resources) | List Rack Resources |
| [rack resources create](/reference/cli-commands/rack-resources-create) | Create a Resource |
| [rack resources delete](/reference/cli-commands/rack-resources-delete) | Delete a Resource |
| [rack resources info](/reference/cli-commands/rack-resources-info) | Get information about a Resource |
| [rack resources link](/reference/cli-commands/rack-resources-link) | Link a Resource to an App |
| [rack resources options](/reference/cli-commands/rack-resources-options) | List options for a Resource type |
| [rack resources proxy](/reference/cli-commands/rack-resources-proxy) | Proxy a local port to a Rack Resource |
| [rack resources types](/reference/cli-commands/rack-resources-types) | List Resource types |
| [rack resources unlink](/reference/cli-commands/rack-resources-unlink) | Unlink a Resource from an App |
| [rack resources update](/reference/cli-commands/rack-resources-update) | Update Resource options |
| [rack resources url](/reference/cli-commands/rack-resources-url) | Get URL for a Resource |

## Registries

| Command | Description |
|:--------|:------------|
| [registries](/reference/cli-commands/registries) | List private registries |
| [registries add](/reference/cli-commands/registries-add) | Add a private registry |
| [registries remove](/reference/cli-commands/registries-remove) | Remove a private registry |

## Networking

| Command | Description |
|:--------|:------------|
| [proxy](/reference/cli-commands/proxy) | Proxy a connection inside the Rack |

## Utilities

| Command | Description |
|:--------|:------------|
| [api get](/reference/cli-commands/api-get) | Query the Rack API |
| [encrypt](/reference/cli-commands/encrypt) | Encrypt data using a symmetric key |
| [decrypt](/reference/cli-commands/decrypt) | Decrypt data using a symmetric key |
| [test](/reference/cli-commands/test) | Run tests |
| [update](/reference/cli-commands/update) | Update the CLI |
| [version](/reference/cli-commands/version) | Display version information |

## See Also

- [CLI Installation](/introduction/installation)
- [Rack Parameters](/reference/rack-parameters)
- [App Parameters](/reference/app-parameters)
- [Getting Started](/introduction/getting-started)
