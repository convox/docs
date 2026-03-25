---
title: "InternalRouterSuffix"
description: "Customize the DNS suffix appended to the internal router domain on a Convox Rack."
---

# InternalRouterSuffix

Suffix appended to the internal router domain name. This controls how internal service endpoints are named within the Rack.

| Default value  | `-rti` |

## Use Cases

- Changing the suffix to a custom value when integrating with existing internal DNS naming conventions
- Avoiding naming collisions if running multiple Racks in the same VPC with internal routing enabled

## Additional Information

The internal router suffix is appended to the Rack's base domain to form the internal ALB's DNS name. The default value of `-rti` (Router Internal) is used by Convox when constructing internal service URLs.

This parameter only has an effect when [Internal](/reference/rack-parameters/Internal) or [InternalOnly](/reference/rack-parameters/InternalOnly) is enabled.

```bash
$ convox rack params set InternalRouterSuffix=-internal
```

## See Also

- [Internal](/reference/rack-parameters/Internal)
- [InternalOnly](/reference/rack-parameters/InternalOnly)
- [Private](/reference/rack-parameters/Private)
