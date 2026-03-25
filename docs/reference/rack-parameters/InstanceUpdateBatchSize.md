---
title: "InstanceUpdateBatchSize"
description: "Control how many Convox Rack EC2 instances are updated simultaneously during a rolling update."
---

# InstanceUpdateBatchSize

The number of instances to update in a batch during rolling updates.

| Default value  | `1` |
| Minimum value | `1` |

## Use Cases

- Increasing the batch size to speed up Rack updates on large clusters with many instances
- Keeping at `1` for production environments to minimize the impact of rolling updates on running services
- Setting to a higher value for non-production Racks where faster updates are preferred over zero-downtime guarantees

## Additional Information

When the Rack's CloudFormation stack is updated (for example, after changing an instance type or AMI), instances are replaced in batches of this size. A batch size of `1` means instances are replaced one at a time, which provides the safest rolling update but takes the longest.

Larger batch sizes reduce total update time but temporarily remove more instances from the cluster, which may affect available capacity for running services.

```bash
$ convox rack params set InstanceUpdateBatchSize=2
```

## See Also

- [InstanceCount](/reference/rack-parameters/InstanceCount)
- [InstanceType](/reference/rack-parameters/InstanceType)
- [HighAvailability](/reference/rack-parameters/HighAvailability)
