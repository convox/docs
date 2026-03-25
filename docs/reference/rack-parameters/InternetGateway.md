---
title: "InternetGateway"
description: "Specify the Internet Gateway ID when installing a Convox Rack into an existing VPC."
---

# InternetGateway

Internet Gateway ID for use with an existing VPC. Required when [ExistingVpc](/reference/rack-parameters/ExistingVpc) is set.

| Default value  | "" |

## Use Cases

- Required when using [ExistingVpc](/reference/rack-parameters/ExistingVpc) to install a Rack into a pre-existing VPC
- Connecting the Rack's routing tables to your VPC's existing internet gateway for outbound traffic

## Additional Information

This parameter is only needed when setting [ExistingVpc](/reference/rack-parameters/ExistingVpc). If you let Convox create its own VPC (the default), this parameter is not used and the Rack creates its own Internet Gateway automatically.

The value should be a valid AWS Internet Gateway ID, for example:

```bash
$ convox rack params set ExistingVpc=vpc-0abc1234 InternetGateway=igw-0def5678
```

## See Also

- [ExistingVpc](/reference/rack-parameters/ExistingVpc)
- [Private](/reference/rack-parameters/Private)
- [VPCCIDR](/reference/rack-parameters/VPCCIDR)
