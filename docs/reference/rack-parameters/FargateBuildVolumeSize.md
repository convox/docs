---
title: "FargateBuildVolumeSize"
description: "Set the build disk size in GiB for the Fargate builder in a Convox Rack."
---

# FargateBuildVolumeSize

Build disk size in GiB for the Fargate builder. Only used when [BuildMethod](/reference/rack-parameters/BuildMethod) is set to `fargate`.

| Setting | Value |
|:--|:--|
| Default value  | ""                                          |
| Allowed values | blank, or a whole number from `21` to `200` |

## Use Cases

- Increasing disk space for Fargate builds that fail with "No space left on device" errors
- Building images with large base layers or large build contexts that outgrow the Fargate default disk
- Sizing the builder disk on Racks that build several large images in one Build

## Additional Information

This parameter sets the ephemeral storage size on the Fargate build task. It is the Fargate counterpart to [BuildVolumeSize](/reference/rack-parameters/BuildVolumeSize), which sizes the EBS volume attached to the `ec2` build instance. On a Rack with [BuildMethod](/reference/rack-parameters/BuildMethod)=`ec2` the value is accepted and the Rack updates, but it has no effect on the build disk. Use [BuildVolumeSize](/reference/rack-parameters/BuildVolumeSize) there instead.

```bash
$ convox rack params set FargateBuildVolumeSize=50
```

A blank value is not sent to the task definition at all, so the build task runs on the Fargate default of 20 GiB. Blank does not mean zero.

`20` is not an accepted value. Fargate already provisions 20 GiB, and AWS accepts an explicit ephemeral storage size only from 21 GiB upward. Set `21` or higher to enlarge the disk, or leave the parameter blank to stay at 20 GiB.

Only whole numbers are accepted. A fractional value such as `50.5` is refused at parameter validation.

A value outside the allowed range is rejected during CloudFormation parameter validation, before any resource is updated, so the Rack is left exactly as it was:

```bash
$ convox rack params set FargateBuildVolumeSize=20
Updating parameters... ERROR: ValidationError: Parameter FargateBuildVolumeSize failed to satisfy constraint: must be blank or a whole number from 21 to 200
```

Each Fargate task includes 20 GiB of ephemeral storage at no additional cost, and AWS bills the storage you configure above that. See [Fargate task ephemeral storage](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/fargate-task-storage.html) and [AWS Fargate pricing](https://aws.amazon.com/fargate/pricing/).

The parameter exists only on Rack releases that declare it. On an older Rack it does not appear in `convox rack params`, and setting it is either dropped silently or rejected with `ValidationError: No updates are to be performed.` Update the Rack first, then set the value.

Downgrading a Rack to a release that predates this parameter drops it. Upgrading again brings it back blank rather than at the value you had set, so set it again after a downgrade and return trip.

## See Also

- [BuildVolumeSize](/reference/rack-parameters/BuildVolumeSize)
- [BuildMethod](/reference/rack-parameters/BuildMethod)
- [FargateBuildCpu](/reference/rack-parameters/FargateBuildCpu)
- [FargateBuildMemory](/reference/rack-parameters/FargateBuildMemory)
- [Rack Parameters](/reference/rack-parameters) for a full list of available parameters
