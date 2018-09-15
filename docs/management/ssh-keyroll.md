---
title: "SSH Keyroll"
---

Keyroll generates and replaces the EC2 keypair used for SSH. Under the hood, this replaces an entire cluster with new instances and new SSH keys.

Note that `convox instances ssh` won't work at all until you do a successful keyroll.


### Roll Rack SSH key 
 
`convox instances keyroll` creates (or regenerates, if one exists already) the SSH key behind `convox instances ssh`. 

This sets up a keypair between the Rack API and the EC2 instances for SSH access. You can then use the `convox instances ssh <instanceID>` command to get SSH access to the instances.


## Under the hood

Keyroll triggers a CloudFormation update that performs a rolling replacement of all of the Rack's instances.

The replacement takes advantage of [ECS container instance draining](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-instance-draining.html); therefore, this operation should not cause any downtime.

## SSH Keyroll FAQ

### Can I provide my own SSH keys to be installed on the instances?

No, for a couple of reasons.

The instances on which a Rack is running are terminated frequently and replaced, so instead of ssh-ing directly to the instances, Convox supplies a proxied connection via `convox instances ssh`.

Furthermore, we consider it a best practice to consider your infrastructure ephemeral and immutable. Therefore, we discourage relying on direct SSH access to your instances because in a well-designed system, it shouldn't be necessary.
