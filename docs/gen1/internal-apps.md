---
title: "Internal Apps"
---

An application can be made "internal" by setting its `Internal` parameter:

```
$ convox apps params set Internal=Yes
```

When an app is internal, its load balancer is only accessible from inside the Rack VPC. This is useful for creating non-public apps that should only be available via VPN, for example.

## Internal Racks

You can also set a Rack to be internal:

```
$ convox rack params set Internal=Yes
```

With this setting, all new apps created in the Rack will automatically be internal apps.

If you want to make an app in an internal Rack publicly accessible, you can explicitly set it as such:

```
$ convox apps params set Internal=No
```
