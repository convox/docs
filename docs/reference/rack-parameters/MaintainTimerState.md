---
title: "MaintainTimerState"
description: "Preserve the enabled or disabled state of timer event rules across deployments."
---

# MaintainTimerState

Timer event rule state preservation across deployments. If you disable a timer in the AWS Console (by disabling its CloudWatch Event Rule), enabling this parameter will keep it disabled across subsequent deployments. By default, the timer state is not explicitly maintained and the behavior of state changes depends on CloudFormation's default handling.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Use Cases

- Temporarily disabling a cron job or scheduled task via the AWS Console and ensuring it stays disabled after the next deployment
- Maintaining manual overrides to timer states in environments where operators need to pause scheduled tasks without code changes
- Preventing CloudFormation from re-enabling timers that were intentionally disabled during an incident response

## Additional Information

When set to `No` (the default), CloudFormation manages the timer event rule state. If you disable a timer via the AWS Console, a subsequent Rack update or application deployment may re-enable it depending on CloudFormation's drift behavior.

When set to `Yes`, the Rack explicitly tracks the current state of each timer event rule before updates and preserves that state. This gives operators safe control to disable timers through the AWS Console without worrying about deployments overriding their changes.

```bash
$ convox rack params set MaintainTimerState=Yes
```

## See Also

- [EcsPollInterval](/reference/rack-parameters/EcsPollInterval)
