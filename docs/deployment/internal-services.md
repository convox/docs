---
title: "Internal Services"
---

A service can be made accessible only on the Rack's internal network by setting its `internal` attribute:

```yaml
services:
  web:
    internal: true
```

Before you can set any services to *internal* you must enable the internal load balancer on your Rack:

```shell
$ convox rack params set Internal=Yes
```
