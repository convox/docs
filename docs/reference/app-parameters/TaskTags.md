---
title: "TaskTags"
description: "Enable ECS tag propagation to the task level for cost allocation and resource tracking."
---

# TaskTags

ECS tag propagation to the task level. When enabled, service-level tags are propagated to individual tasks.

| Default value  | `No` |
| Allowed values | `Yes`, `No` |

## Use Cases

- Enable when you need per-task cost allocation using AWS Cost Explorer tags
- Enable when you use tag-based IAM policies or resource tracking that require tags on individual ECS tasks
- Keep disabled when you do not need task-level tagging and want to avoid any additional API overhead from tag propagation

## Additional Information

When enabled, tags defined on the ECS service (such as `App`, `Name`, and `Type`) are propagated to each individual ECS task launched by that service. This allows you to identify and track individual tasks by their tags in the AWS console, CLI, or through the ECS API.

Task-level tags are useful for AWS Cost Explorer when you want to attribute costs to specific applications or services at the task granularity rather than the cluster or service level.

```bash
$ convox apps params set TaskTags=Yes
```

## See Also

- [FargateServices](/reference/app-parameters/FargateServices)
- [Rack Parameters](/reference/rack-parameters)
