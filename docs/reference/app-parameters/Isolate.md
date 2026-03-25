---
title: "Isolate"
description: "Run application services in an isolated network configuration when the app is in a Private Rack."
---

# Isolate

Isolated network configuration for application services. When set to `Yes`, services run on dedicated Fargate tasks with network-level isolation. Only takes effect when the app is running in a Private Rack.

| Default value  | `No` |
| Allowed values | `Yes`, `No` |

## Use Cases

- Enable for workloads that require strict network isolation from other applications on the same Rack
- Enable when compliance requirements mandate that services run in dedicated network segments
- Keep disabled for typical applications that can share network infrastructure with other apps on the Rack

## Additional Information

This parameter only takes effect when the [Private](/reference/app-parameters/Private) app parameter is set to `Yes`. If `Private=No`, setting `Isolate=Yes` has no effect.

When both `Private=Yes` and `Isolate=Yes` are set, services use `awsvpc` network mode with IP-based target groups to ensure they do not share underlying network infrastructure with other tasks. This provides an additional layer of isolation beyond private subnet placement.

Note that services using [FargateServices](/reference/app-parameters/FargateServices) always run with `awsvpc` network mode and IP-based target groups, regardless of the `Isolate` setting.

```bash
$ convox apps params set Isolate=Yes
```

## See Also

- [Private](/reference/app-parameters/Private)
- [FargateServices](/reference/app-parameters/FargateServices)
- [Rack Parameters](/reference/rack-parameters)
- [Private Networking](/networking/private-networking)
