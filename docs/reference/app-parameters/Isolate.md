---
title: "Isolate"
description: "Run application services in an isolated network configuration when the app is in a Private Rack."
---

# Isolate

Isolated network configuration for application Services. Only takes effect when the App is running in a Private Rack.

| Setting | Value |
|:--|:--|
| Default value  | `No` |
| Allowed values | `Yes`, `No` |

## Use Cases

- Enable for workloads that require strict network isolation from other applications on the same Rack
- Enable when compliance requirements mandate that Services run in dedicated network segments
- Keep disabled for typical applications that can share network infrastructure with other Apps on the Rack

## Additional Information

This parameter only takes effect when the [Private](/reference/app-parameters/Private) app parameter is set to `Yes`. If `Private=No`, setting `Isolate=Yes` has no effect.

When both `Private=Yes` and `Isolate=Yes` are set, Services use `awsvpc` network mode with IP-based target groups to ensure they do not share underlying network infrastructure with other tasks. This provides an additional layer of isolation beyond private subnet placement.

Note that Services using [FargateServices](/reference/app-parameters/FargateServices) always run with `awsvpc` network mode and IP-based target groups, regardless of the `Isolate` setting.

```bash
$ convox apps params set Isolate=Yes
```

Before moving the Rack to a version older than 20261103122112, set `Isolate=No` on any App that has `Isolate=Yes` and an `internalAndExternal` Service. Otherwise the App's first deploy or promote on the older version removes the external router's access to that Service's processes. Its external domain stops answering, and the deploy does not finish for hours unless you cancel it with `convox apps cancel -a myapp`.

## See Also

- [Private](/reference/app-parameters/Private)
- [FargateServices](/reference/app-parameters/FargateServices)
- [Rack Parameters](/reference/rack-parameters)
- [Private Networking](/networking/private-networking)
