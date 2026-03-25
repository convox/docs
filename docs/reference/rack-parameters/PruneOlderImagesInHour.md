---
title: "PruneOlderImagesInHour"
description: "Set the age threshold in hours for pruning old Docker images on Convox Rack instances."
---

# PruneOlderImagesInHour

Age threshold (in hours) for pruning Docker images on Rack instances. Images that exceed this age threshold are removed during the next scheduled cleanup run, helping to reclaim disk space.

| Default value  | `96` |
| Minimum value  | `1`  |

## Use Cases

- Reducing from the default 96 hours (4 days) on Racks with limited disk space or very frequent deployments
- Increasing the value on Racks where you need to retain older images for rollback purposes
- Setting to a low value (e.g., `24`) on development Racks where old images are rarely needed

## Additional Information

This parameter works together with [PruneOlderImagesCronRunFreq](/reference/rack-parameters/PruneOlderImagesCronRunFreq) to control Docker image cleanup. `PruneOlderImagesInHour` sets the age threshold (in hours) for which images are eligible for pruning, while `PruneOlderImagesCronRunFreq` determines how often the cleanup job runs.

The default of `96` hours (4 days) provides a reasonable balance between disk space conservation and the ability to quickly roll back to a recent deployment.

```bash
$ convox rack params set PruneOlderImagesInHour=48
```

## See Also

- [PruneOlderImagesCronRunFreq](/reference/rack-parameters/PruneOlderImagesCronRunFreq)
- [VolumeSize](/reference/rack-parameters/VolumeSize)
- [BuildVolumeSize](/reference/rack-parameters/BuildVolumeSize)
