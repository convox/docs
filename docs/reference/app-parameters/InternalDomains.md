---
title: "InternalDomains"
description: "Controls whether internal .convox.site and .convox domain names route to this application."
---

# InternalDomains

Internal domain name registration for the application. When set to `No`, internal `.convox.site` and `.convox` domain names are not created.
Use this parameter if you are running out of available rules on your load balancer.

| Default value  | `Yes` |
| Allowed values | `Yes`, `No` |

## Use Cases

- Set to `No` to free up load balancer listener rules when you are approaching the AWS limit
- Set to `No` for applications that do not need internal service-to-service routing via Convox-managed domain names
- Leave as `Yes` (default) to enable internal routing between services using `internal-<service>.<app>.<rack>.convox` domain names

## Additional Information

AWS Application Load Balancers have a limit on the number of listener rules (default 100). Each service with `InternalDomains=Yes` consumes additional listener rules for its internal domain entries. If you have many apps or services on a single Rack, you may need to disable internal domains on some applications to stay within this limit.

Disabling this parameter does not affect external domain routing or custom domains configured in `convox.yml`. It only prevents the automatic creation of the `convox.site` and `.convox` internal DNS records for the application.

```bash
$ convox apps params set InternalDomains=No
```

## See Also

- [Private](/reference/app-parameters/Private)
- [Isolate](/reference/app-parameters/Isolate)
- [RedirectHttps](/reference/app-parameters/RedirectHttps)
- [Rack Parameters](/reference/rack-parameters)
