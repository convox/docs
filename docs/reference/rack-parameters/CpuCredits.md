---
title: "CpuCredits"
description: "Controls the CPU credit option for burstable T-family EC2 instances in a Convox Rack."
---

# CpuCredits

The credit option for CPU usage of a T instance. AWS T-family instances (e.g., `t3.small`, `t3.medium`) use a CPU credit model that governs how burst performance is billed. This parameter lets you choose between `standard` mode (where you earn and spend credits) and `unlimited` mode (where you can burst beyond your credit balance at an additional cost).

When left blank, the instance uses the AWS account-level default for T instance CPU credits.

| Default value  | "" |
| Allowed values | `standard`, `unlimited`, "" |

## Use Cases

- Set to `unlimited` when your workloads have unpredictable CPU spikes and you want to avoid throttling, accepting the additional cost for sustained burst usage.
- Set to `standard` to cap CPU burst to your earned credit balance, keeping costs predictable.
- Leave blank to inherit the AWS account-level default setting.

## Additional Information

In `standard` mode, a T instance earns CPU credits when idle and spends them during bursts. If credits are exhausted, the instance is limited to its baseline CPU performance. In `unlimited` mode, the instance can burst indefinitely but surplus credits are charged at a flat rate.

For more details, see the AWS documentation on [Burstable performance instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html).

```bash
$ convox rack params set CpuCredits=unlimited
```

## See Also

- [InstanceType](/reference/rack-parameters/InstanceType)
- [Rack Parameters](/reference/rack-parameters) for a full list of available parameters
