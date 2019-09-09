---
title: "Service Tags"
---

Often, you may like to be able to facet your AWS billing data by things like rack, app, and service.

You can facilitate doing this with the existing tags attached to your resources by Convox with a couple of easy steps.

* Ensure you are opted in to the new ARN and resource ID format from AWS.  Amazon have an excellent blog post detailing this [here](https://aws.amazon.com/blogs/compute/migrating-your-amazon-ecs-deployment-to-the-new-arn-and-resource-id-format-2/).
* For all the apps you wish to track, set the app level `TaskTags` parameter to `Yes`.

#### Example

```
convox apps params set TaskTags=Yes
```

Once done, the tags will propogate through all your resources and then become available to filter by in the AWS billing reports.