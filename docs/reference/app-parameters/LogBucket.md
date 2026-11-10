---
title: "LogBucket"
description: "The S3 bucket that receives access logs for an App's settings bucket, set from the Rack's log bucket."
---

# LogBucket

S3 bucket that receives the access logs of the App's settings bucket. Convox sets it to the Rack's log bucket: the bucket in the Rack's [LogBucket](/reference/rack-parameters/LogBucket) parameter, or the log bucket the Rack creates when that parameter is blank.

| Setting | Value |
|:--|:--|
| Default value  | The Rack's log bucket |

## Additional Information

Log files are written with the prefix `convox/logs/<stack-name>/s3/`.

Convox sets this parameter when the App is created and at every deploy or promote, unless `params:` in `convox.yml` sets it. On rack version 20261107025028 or newer, `convox apps params set` does not apply a value that differs from the Rack's. To send these logs to another bucket, set the Rack's [LogBucket](/reference/rack-parameters/LogBucket) parameter, which also receives the Rack's load balancer access logs.

## See Also

- [LogBucket Rack Parameter](/reference/rack-parameters/LogBucket)
- [LogDriver](/reference/app-parameters/LogDriver)
- [LogRetention](/reference/app-parameters/LogRetention)
- [Rack Parameters](/reference/rack-parameters)
