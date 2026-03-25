---
title: "InstanceBootCommand"
description: "Run a custom shell command early during EC2 instance boot on Convox Rack instances via cloud-init."
---

# InstanceBootCommand

A single line of shell script to run (as root) as a cloud-init command early during instance boot.

For more information about using cloud-init with EC2, see the AWS doc [Running Commands on Your Linux Instance at Launch](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html#user-data-cloud-init). For cloud-init specifics, see "bootcmd" in the doc [Run commands on first boot](https://cloudinit.readthedocs.io/en/latest/topics/examples.html#run-commands-on-first-boot).

The `InstanceBootCommand` will also apply to any [build instance](/reference/rack-parameters/BuildInstance) associated with the Rack.

| Default value  | "" |

## Use Cases

- Setting kernel parameters such as `vm.max_map_count` for workloads like Elasticsearch
- Configuring custom networking or DNS settings before the ECS agent starts
- Installing required kernel modules or drivers that must be loaded early in the boot process

## Additional Information

### Example: increase virtual memory

Tell the host EC2 instance to set operating system limits on mmap to `262144` (useful for Elasticsearch):

```bash
$ convox rack params set 'InstanceBootCommand="sysctl -w vm.max_map_count=262144"'
Updating parameters... OK
```

(Note the surrounding single quotes in the above command.)

The `InstanceBootCommand` runs during the `bootcmd` phase of cloud-init, which executes on every boot of the instance (not just the first boot). This makes it suitable for kernel tuning and other settings that need to persist across reboots.

## See Also

- [InstanceRunCommand](/reference/rack-parameters/InstanceRunCommand)
- [BuildInstance](/reference/rack-parameters/BuildInstance)
- [InstanceType](/reference/rack-parameters/InstanceType)
