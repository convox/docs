---
title: "Version"
description: "Specify the Convox release version running on the Rack."
---

# Version

(REQUIRED) Convox release version. This parameter is set automatically when a Rack is installed or updated and determines which version of the Convox software is running on the Rack.

| Minimum length | `1` |

## Use Cases

- Managed automatically during `convox rack update` — you typically do not set this directly
- Pin to a specific version to prevent unintended upgrades in production
- Roll back to a previous version if a Rack update introduces issues

## Additional Information

The `Version` parameter is managed by the Convox CLI and API during Rack installation and updates. You typically do not set this parameter directly. Instead, use the Convox CLI to update your Rack:

```bash
$ convox rack update
```

Or update to a specific version:

```bash
$ convox rack update <version>
```

See the [Rack Updates](/management/rack-updates) documentation for more information on managing Rack versions.

## See Also

- [Rack Updates](/management/rack-updates)
