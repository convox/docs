---
title: "EnableContainerReadonlyRootFilesystem"
description: "Enables a read-only root filesystem for containers running in a Convox Rack."
---

# EnableContainerReadonlyRootFilesystem

Read-only root filesystem enforcement for ECS containers on the Rack. Enabling this will remove write access to the root filesystem.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Use Cases

- Harden container security by preventing runtime modifications to the filesystem, reducing the attack surface for container escape or malware injection.
- Meet compliance requirements (e.g., CIS Benchmarks, PCI-DSS) that mandate immutable container filesystems.
- Detect applications that incorrectly write to the root filesystem instead of using designated writable volumes.

## Additional Information

When enabled, the container's root filesystem is mounted as read-only. Any attempt by the application to write to the filesystem (outside of explicitly mounted volumes) will fail with a permission error.

Applications that need to write temporary files should use writable [volumes](/application/volumes) or the `/tmp` directory (if mounted as a `tmpfs`). Review your application's file I/O patterns before enabling this parameter to avoid unexpected failures.

This setting is configured at the Rack level and pushed to all application stacks during deployment. It cannot be selectively enabled for individual services within an app.

```bash
$ convox rack params set EnableContainerReadonlyRootFilesystem=Yes
```

## See Also

- [EnableSharedEFSVolumeEncryption](/reference/rack-parameters/EnableSharedEFSVolumeEncryption)
- [EncryptEbs](/reference/rack-parameters/EncryptEbs)
- [Rack Parameters](/reference/rack-parameters) for a full list of available parameters
