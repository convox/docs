---
title: "InstanceRunCommand"
description: "Run a custom shell command late during EC2 instance boot on Convox Rack instances via cloud-init."
---

# InstanceRunCommand

A single line of shell script to run as a cloud-init command late during instance boot.

For more information about using cloud-init with EC2, see the AWS doc [Running Commands on Your Linux Instance at Launch](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html#user-data-cloud-init). For cloud-init specifics, see "runcmd" in the doc [Run commands on first boot](https://cloudinit.readthedocs.io/en/latest/topics/examples.html#run-commands-on-first-boot).

The `InstanceRunCommand` will also apply to any [build instance](/reference/rack-parameters/BuildInstance) associated with the Rack.

| Default value  | "" |

## Use Cases

- Installing additional packages or agents after the instance has fully booted (e.g., monitoring or security agents)
- Running configuration scripts that depend on services started during boot
- Executing setup commands that need networking to be fully available

## Additional Information

Unlike [InstanceBootCommand](/reference/rack-parameters/InstanceBootCommand), which runs during the `bootcmd` phase (early, on every boot), the `InstanceRunCommand` runs during the `runcmd` phase. The `runcmd` phase executes only on the first boot of the instance, after all other cloud-init modules have run.

This makes `InstanceRunCommand` suitable for one-time setup tasks such as installing software packages or configuring services that persist across reboots.

```bash
$ convox rack params set InstanceRunCommand="sudo yum install -y awscli"
```

## See Also

- [InstanceBootCommand](/reference/rack-parameters/InstanceBootCommand)
- [BuildInstance](/reference/rack-parameters/BuildInstance)
- [InstanceType](/reference/rack-parameters/InstanceType)
