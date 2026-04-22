---
title: "rack"
description: "Display information about the current Rack."
---

# rack

Display information about the current Rack, including its name, provider, region, router endpoint, status, and version. When [NLB](/reference/rack-parameters/NLB) or [NLBInternal](/reference/rack-parameters/NLBInternal) is enabled, the Network Load Balancer DNS name(s) and any allocated Elastic IPs are also shown.

## Syntax

```bash
$ convox rack
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox rack
Name      production
Provider  aws
Region    us-east-1
Router    router.0a1b2c3d4e5f.convox.cloud
Status    running
Version   20260421192651
```

With NLB enabled:

```bash
$ convox rack
Name          production
Provider      aws
Region        us-east-1
Router        router.0a1b2c3d4e5f.convox.cloud
NLB           production-nlb-abc123.elb.us-east-1.amazonaws.com (52.1.2.3, 52.4.5.6, 52.7.8.9)
NLB Internal  production-nlb-internal-xyz789.elb.us-east-1.amazonaws.com
Status        running
Version       20260421192651
```

The parenthesized values after the public NLB are the Elastic IPs bound to the NLB — one per Availability Zone. Use these for clients that connect by IP (e.g., allowlisted firewalls) and the DNS name for clients that resolve hostnames.

## See Also

- [rack params](/reference/cli-commands/rack-params)
- [rack update](/reference/cli-commands/rack-update)
- [Network Load Balancing](/networking/nlb)
- [Rack Statuses](/reference/rack-statuses)
