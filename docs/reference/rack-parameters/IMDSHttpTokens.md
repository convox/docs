---
title: "IMDSHttpTokens"
description: "Control whether Convox Rack EC2 instances require IMDSv2 tokens for instance metadata access."
---

# IMDSHttpTokens

Instance Metadata Service (IMDS) token requirement for Rack instances. You can set EC2 instances to use only IMDSv2 by setting `IMDSHttpTokens` to `required`. See [Configure the instance metadata options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-instance-metadata-options.html#configuring-IMDS-new-instances) in the AWS docs.

| Default value  | `optional`             |
| Allowed values | `optional`, `required` |

## Use Cases

- Setting to `required` to enforce IMDSv2 across all Rack instances for improved security posture
- Keeping as `optional` when applications or agents on your instances still depend on IMDSv1
- Enforcing `required` to meet compliance requirements that mandate IMDSv2

## Additional Information

When set to `optional`, both IMDSv1 and IMDSv2 requests are accepted. When set to `required`, only IMDSv2 requests (which use session tokens) are allowed, and IMDSv1 requests are rejected.

IMDSv2 adds a session-based authentication layer that protects against several classes of vulnerabilities, including SSRF (Server-Side Request Forgery) attacks. AWS recommends migrating to IMDSv2 for all workloads.

When setting this to `required`, you may also need to increase [IMDSHttpPutResponseHopLimit](/reference/rack-parameters/IMDSHttpPutResponseHopLimit) to `2` if your containers need to access instance metadata from within Docker.

```bash
$ convox rack params set IMDSHttpTokens=required
```

## See Also

- [IMDSHttpPutResponseHopLimit](/reference/rack-parameters/IMDSHttpPutResponseHopLimit)
- [InstanceType](/reference/rack-parameters/InstanceType)
- [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup)
- [AWS IMDS Documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-instance-metadata-options.html)
