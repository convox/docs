---
title: "Encryption"
description: "Controls whether Convox Rack secrets are encrypted using AWS KMS."
---

# Encryption

Encryption is always on. Convox uses an AWS KMS key to encrypt all Rack-managed secrets regardless of this parameter's value.

| Default value  | `Yes`       |
| Allowed values | `Yes`, `No` |

> **Note:** This parameter is retained for backward compatibility with existing CloudFormation stacks, but the ability to disable encryption has been removed. The KMS key is always created and secrets are always encrypted. Setting this parameter to `No` has no effect.

## Use Cases

- No action needed. KMS encryption of Rack secrets is always active.
- This parameter exists in the CloudFormation template for backward compatibility and cannot be removed without breaking existing stack updates.

## Additional Information

Convox uses an AWS KMS key to encrypt sensitive data such as environment variables, resource connection strings, and other secrets managed by the Rack. The KMS key is unconditionally created and used for all secret encryption operations. This ensures that secrets are always encrypted at rest in the backing DynamoDB table.

KMS usage incurs a small cost per API call. For most Racks, this cost is negligible.

```bash
$ convox rack params set Encryption=Yes
```

## See Also

- [EncryptEbs](/reference/rack-parameters/EncryptEbs)
- [EnableSharedEFSVolumeEncryption](/reference/rack-parameters/EnableSharedEFSVolumeEncryption)
- [DynamoDbTableDeletionProtectionEnabled](/reference/rack-parameters/DynamoDbTableDeletionProtectionEnabled)
- [Rack Parameters](/reference/rack-parameters) for a full list of available parameters
