---
title: "IMDSHttpPutResponseHopLimit"
description: "Set the maximum network hop count for EC2 instance metadata PUT responses on Convox Rack instances."
---

# IMDSHttpPutResponseHopLimit

Maximum number of network hops for IMDS PUT response packets. This parameter is particularly relevant when `IMDSHttpTokens` is set to `required`, ensuring enhanced security by enforcing IMDSv2.

| Default value  | `1`      |
| Allowed values | `1`-`64` |

## Use Cases

- Increasing to `2` when running containers that need to access instance metadata through IMDSv2
- Adjusting when applications running inside Docker containers cannot reach the IMDS endpoint with the default hop limit
- Setting to `2` when enabling IMDSv2 (`IMDSHttpTokens=required`) and experiencing connectivity issues with the metadata service

## Additional Information

> **Note:** When `IMDSHttpTokens` is set to `required`, some configurations may require increasing the `IMDSHttpPutResponseHopLimit` to ensure proper functionality. If encountering connectivity issues with applications requiring IMDSv2, consider setting `IMDSHttpPutResponseHopLimit` to `2`. This adjustment helps facilitate necessary communications with the instance metadata service.

> **Warning:** Adjusting the `IMDSHttpPutResponseHopLimit` above the default value should be done with understanding of your network topology and the security implications. Always verify that changes do not compromise your instance's security posture.

The default value of `1` means the metadata response cannot travel beyond the instance itself. When containers run inside Docker on EC2, the extra network hop from the container to the host requires a value of at least `2` for the metadata token to reach the container.

```bash
$ convox rack params set IMDSHttpPutResponseHopLimit=2
```

## See Also

- [IMDSHttpTokens](/reference/rack-parameters/IMDSHttpTokens)
- [InstanceType](/reference/rack-parameters/InstanceType)
- [AWS IMDS Documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-instance-metadata-options.html)
