---
title: "EnableS3Versioning"
description: "Controls S3 bucket versioning for all buckets created by a Convox Rack."
---

# EnableS3Versioning

S3 bucket versioning for Rack-managed buckets. This affects all the buckets created for this Rack.

| Default value  | `Suspended`              |
| Allowed values | `Enabled`, `Suspended`   |

## Use Cases

- Enable versioning to protect against accidental object deletion or overwrites in Rack-managed S3 buckets.
- Required by compliance frameworks (e.g., SOC 2, HIPAA) that mandate versioned storage for audit trails.
- Enables the ability to recover previous versions of build artifacts or other Rack-managed objects.

## Additional Information

When set to `Enabled`, S3 retains all versions of every object stored in the Rack's buckets. This includes previous versions of overwritten objects and delete markers for removed objects. Versioned objects consume additional storage, which will increase your S3 costs.

When set to `Suspended`, S3 stops creating new versions but preserves any existing versions that were created while versioning was enabled.

Note that once versioning has been enabled on a bucket, it cannot be fully disabled -- it can only be suspended. Existing object versions will remain until explicitly deleted.

```bash
$ convox rack params set EnableS3Versioning=Enabled
```

## See Also

- [LogBucket](/reference/rack-parameters/LogBucket)
- [EncryptEbs](/reference/rack-parameters/EncryptEbs)
- [Encryption](/reference/rack-parameters/Encryption)
- [Rack Parameters](/reference/rack-parameters) for a full list of available parameters
