---
title: "RackUrl"
description: "Injects a RACK_URL environment variable into containers for programmatic Rack API access."
---

# RackUrl

Rack API URL injection as an environment variable. When set to `Yes`, a `RACK_URL` environment variable is injected into all Service and Timer containers. This variable contains the authenticated URL for the Rack API, allowing applications to communicate with the Rack control plane programmatically using the Convox SDK.

| Default value  | `No` |
| Allowed values | `Yes`, `No` |

## Use Cases

- Enable when your application needs to interact with the Convox Rack API at runtime (e.g., to trigger builds, inspect resources, or manage deployments)
- Enable when using the Convox SDK within your application code
- Keep disabled for standard applications that do not need direct Rack API access

## Additional Information

> Do not change this parameter unless you know what you are doing or are directed to by Convox support. Misconfiguring this value can affect your application's ability to communicate with the Rack API.

The `RACK_URL` environment variable contains a fully authenticated URL including credentials. Treat it as sensitive data. It is injected into both service containers and timer containers.

```bash
$ convox apps params set RackUrl=Yes
```

## See Also

- [Private](/reference/app-parameters/Private)
- [Rack Parameters](/reference/rack-parameters)
