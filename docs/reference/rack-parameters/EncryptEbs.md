---
title: "EncryptEbs"
description: "Enables encryption at rest for EBS volumes attached to Convox Rack instances."
---

# EncryptEbs

Encryption at rest for EBS volumes attached to Rack instances.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Use Cases

- Meet compliance requirements (e.g., HIPAA, PCI-DSS, SOC 2) that mandate encryption at rest for all block storage.
- Protect data on instance volumes from unauthorized physical access to the underlying storage hardware.
- Enable encryption for racks that were initially provisioned without EBS encryption.

## Additional Information

When enabled, all EBS volumes attached to Rack instances are encrypted using the default AWS-managed KMS key. This includes the root volume and any additional data volumes. Encryption and decryption are handled transparently by AWS with minimal performance impact.

EBS encryption is applied at the volume level. Enabling this parameter on an existing Rack will encrypt new volumes as instances are replaced (e.g., during a scaling event or instance type change). Existing volumes on running instances are not retroactively encrypted until those instances are replaced.

```bash
$ convox rack params set EncryptEbs=Yes
```

## See Also

- [Encryption](/reference/rack-parameters/Encryption)
- [EnableSharedEFSVolumeEncryption](/reference/rack-parameters/EnableSharedEFSVolumeEncryption)
- [VolumeSize](/reference/rack-parameters/VolumeSize)
- [Rack Parameters](/reference/rack-parameters) for a full list of available parameters
