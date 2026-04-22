---
title: "rack params"
description: "Display Rack parameters and their current values."
---

# rack params

Display Rack parameters and their current values. Parameters control infrastructure, networking, scaling, and security settings for the Rack.

## Syntax

```bash
$ convox rack params
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--group` | `-g` | Filter output to one parameter group (see [Groups](#groups) below) |
| `--rack` | `-r` | Rack name |
| `--reveal` | | Show the real value of masked parameters on a TTY |

## Example Usage

```bash
$ convox rack params
BuildInstance
CertDuration          2160
HighAvailability      true
HttpProxy             **********
IdleTimeout           3600
InstanceCount         3
InstanceType          t3.medium
Internal              false
NodeDisk              20
Password              **********
Ssl                   true
```

Two parameters — [Password](/reference/rack-parameters/Password) and [HttpProxy](/reference/rack-parameters/HttpProxy) — render as `**********` on a terminal to avoid accidental exposure during screen shares or pasted bug reports. Empty values render as empty (never as `**********`), so a parameter that has never been set is still distinguishable from one that is masked.

The mask is a display-only wrapper. Piped output bypasses it so `grep`, `awk`, and backup automation keep working:

```bash
$ convox rack params | grep -E '^Password|^HttpProxy'
HttpProxy  http://user:secret@corp-proxy.example.com:3128
Password   correct-horse-battery-staple
```

Pass `--reveal` to show the real values on a terminal:

```bash
$ convox rack params --reveal
...
HttpProxy             http://user:secret@corp-proxy.example.com:3128
...
Password              correct-horse-battery-staple
...
```

### Groups

The 110-parameter Rack surface is partitioned into ten logical groups. Filter to one with `-g` / `--group`:

```bash
$ convox rack params -g nlb
NLB                            Yes
NLBAllowCIDR                   0.0.0.0/0
NLBCrossZone                   No
NLBDeletionProtection          No
NLBInternal                    Yes
NLBInternalAllowCIDR
NLBInternalCrossZone           No
NLBInternalDeletionProtection  No
NLBInternalPreserveClientIP    No
NLBPreserveClientIP            No
```

Unique prefixes are accepted — `-g net` resolves to `network`, `-g nlb` resolves to `nlb`. Ambiguous prefixes emit an error that names the collisions and suggests shortest-unique alternatives:

```text
$ convox rack params -g n
ERROR: group 'n' matches multiple groups: network, nlb (use 'net' or 'nlb')
  available groups:
  api        Rack API web process config, router, ingress toggles
  build      Build method, build instance, Fargate build, image pruning
  ...
```

The listing is truncated here; `convox rack params -g <unknown>` emits all ten group names with their descriptions.

| Group | Scope |
|:------|:------|
| `api` | Rack API web process config, router, ingress toggles |
| `build` | Build method, build instance, Fargate build, image pruning |
| `instances` | AMI, instance type, boot/run commands, IAM policy, volumes, tenancy |
| `logging` | Logs destination, retention, syslog format |
| `meta` | Version, development mode, telemetry, client ID, ECS tuning |
| `network` | VPC, subnets, gateways, connectivity, proxy |
| `nlb` | Network Load Balancer: listeners, cross-zone, allow-CIDR, preserve-client-IP |
| `scaling` | Autoscaling, spot fleet, instance counts, schedules, HA |
| `security` | Credentials, allowlists, SGs, SSL, IMDS, container hardening |
| `storage` | S3 versioning, DynamoDB protection, EFS encryption |

A parameter may belong to more than one group when it spans two operator concerns (for example, [HttpProxy](/reference/rack-parameters/HttpProxy) appears in both `network` and `security`).

If the selected group has no matching parameters set on the Rack, `convox rack params -g <group>` prints `NOTICE: no params in group '<group>' for this rack` to stderr and exits cleanly.

## See Also

- [rack params set](/reference/cli-commands/rack-params-set)
- [apps params](/reference/cli-commands/apps-params)
- [Rack Parameters](/reference/rack-parameters)
- [rack](/reference/cli-commands/rack)
