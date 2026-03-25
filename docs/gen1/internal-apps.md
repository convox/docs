---
title: "Internal Apps"
description: "Gen 1 (End of Life): How to configure Gen 1 Convox applications and Racks as internal, making them accessible only within the VPC."
---

# Internal Apps

> **This page documents Generation 1, which has reached End of Life.** Gen 1 apps use `docker-compose.yml`. For current documentation, see [Internal Services](/networking/internal-services).

An application can be made "internal" by setting its `Internal` parameter:

```bash
$ convox apps params set Internal=Yes
```

When an app is internal, its load balancer is only accessible from inside the Rack VPC. This is useful for creating non-public apps that should only be available via VPN, for example.

## Internal Racks

You can also set a Rack to be internal:

```bash
$ convox rack params set Internal=Yes
```

With this setting, all new apps created in the Rack will automatically be internal apps.

If you want to make an app in an internal Rack publicly accessible, you can explicitly set it as such:

```bash
$ convox apps params set Internal=No
```

## See Also

- [Internal Services (Gen 2)](/networking/internal-services)
- [App Parameters](/gen1/app-parameters)
- [Load Balancers](/gen1/load-balancers)
