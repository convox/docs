---
title: "PruneOlderImagesCronRunFreq"
description: "Configure how frequently Docker image pruning runs on Convox Rack instances."
---

# PruneOlderImagesCronRunFreq

Cron frequency to prune older Docker images on Rack instances. This automated cleanup prevents disk space exhaustion by removing unused Docker images on a recurring schedule.

| Default value  | `daily`                         |
| Allowed values | `hourly`, `daily`, `weekly` |

## Use Cases

- Setting to `hourly` on Racks with frequent deployments that generate many Docker images quickly
- Keeping as `daily` (default) for most production Racks with moderate deployment frequency
- Setting to `weekly` for stable Racks with infrequent deployments where disk space is not a concern

## Additional Information

This parameter works together with [PruneOlderImagesInHour](/reference/rack-parameters/PruneOlderImagesInHour) to control Docker image cleanup. `PruneOlderImagesCronRunFreq` determines how often the cleanup job runs, while `PruneOlderImagesInHour` determines the age threshold for images to be pruned.

For Racks with high deployment velocity, consider setting this to `hourly` and reducing [PruneOlderImagesInHour](/reference/rack-parameters/PruneOlderImagesInHour) to keep disk usage under control.

```bash
$ convox rack params set PruneOlderImagesCronRunFreq=hourly
```

## See Also

- [PruneOlderImagesInHour](/reference/rack-parameters/PruneOlderImagesInHour)
- [VolumeSize](/reference/rack-parameters/VolumeSize)
- [BuildVolumeSize](/reference/rack-parameters/BuildVolumeSize)
