---
title: "EnableContainerReadonlyRootFilesystem"
description: "Controls whether ECS containers run with a read-only root filesystem for enhanced security."
---

# EnableContainerReadonlyRootFilesystem

Read-only root filesystem enforcement for the application's ECS containers, enhancing security by preventing modifications to critical system files.

| Default value  | `No` |
| Allowed values | `Yes`, `No` |

- `Yes`: Activates **read-only mode**, restricting write operations to the root filesystem. Applications must use external storage or explicitly writable paths.
- `No` (default): Containers have standard write access to the root filesystem.

## Use Cases

- Enable to harden container security in compliance-sensitive environments (e.g., PCI DSS, HIPAA)
- Enable to prevent containers from writing malicious files to the root filesystem in the event of a compromise
- Keep disabled when your application writes temporary files, logs, or caches to the root filesystem and cannot be easily reconfigured

## Additional Information

**Potential Breaking Change:**
Enabling this parameter may **disrupt applications** that write to the root filesystem. Ensure your application is compatible before enabling it. Use **external storage solutions** or **designated writable paths** to prevent unexpected failures.

Applications that need to write data at runtime should use mounted volumes (such as EFS resources defined in `convox.yml`) or write to paths that are explicitly configured as writable via `tmpfs` mounts. Common directories that applications write to include `/tmp`, `/var/run`, and application-specific cache directories.

Test this change in a staging environment before applying it to production.

```bash
$ convox apps params set EnableContainerReadonlyRootFilesystem=Yes
```

## See Also

- [FargateServices](/reference/app-parameters/FargateServices)
- [Private](/reference/app-parameters/Private)
- [Rack Parameters](/reference/rack-parameters)
