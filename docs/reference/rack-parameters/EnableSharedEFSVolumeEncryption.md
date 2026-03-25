---
title: "EnableSharedEFSVolumeEncryption"
description: "Enables AWS KMS encryption on the shared EFS volume used by Convox Rack applications."
---

# EnableSharedEFSVolumeEncryption

AWS KMS encryption for the default shared EFS volume used for application [volumes](/application/volumes).

| Default value  | `false`         |
| Allowed values | `true`, `false` |

> **Warning:** Enabling `EnableSharedEFSVolumeEncryption` will recreate the EFS volume and **all** application shared volume data will be lost. To preserve data, follow these steps:
>
> - **Backup:** Use AWS Backup or a similar tool to create a snapshot of the existing Amazon EFS volume, ensuring all current data is securely copied.
> - **Restore:** After enabling encryption, restore your data from the backup snapshot to the new encrypted EFS volume.

## Use Cases

- Meet compliance requirements (e.g., HIPAA, PCI-DSS) that mandate encryption at rest for all persistent storage.
- Protect sensitive application data stored on shared volumes from unauthorized access at the storage layer.
- Enable encryption for racks that were initially created without EFS encryption.

## Additional Information

When enabled, the shared EFS volume is encrypted at rest using AWS Key Management Service (KMS). All data written to the volume is automatically encrypted, and all reads are automatically decrypted. There is no performance impact noticeable by most workloads.

Because AWS does not support enabling encryption on an existing EFS file system, Convox must delete and recreate the EFS volume when this parameter is changed. This is a **destructive operation** -- all data on the existing shared volume will be permanently lost unless backed up beforehand.

After enabling encryption, verify that your applications are functioning correctly and that any restored data is intact.

```bash
$ convox rack params set EnableSharedEFSVolumeEncryption=true
```

## See Also

- [EnableContainerReadonlyRootFilesystem](/reference/rack-parameters/EnableContainerReadonlyRootFilesystem)
- [EncryptEbs](/reference/rack-parameters/EncryptEbs)
- [Encryption](/reference/rack-parameters/Encryption)
- [Rack Parameters](/reference/rack-parameters) for a full list of available parameters
