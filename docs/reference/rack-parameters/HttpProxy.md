---
title: "HttpProxy"
description: "Configure an outbound HTTP proxy for network-restricted Convox Rack instances."
---

# HttpProxy

HTTP proxy for outbound HTTP connections (for network-restricted Racks).

Set this value to the hostname (or IP address) and port number of an HTTP proxy for all outbound connections from Rack instances. This configures the `http_proxy`, `https_proxy`, `HTTP_PROXY`, and `HTTPS_PROXY` environment variables on each instance.

| Default value  | "" |

## Use Cases

- Routing Rack instance traffic through a corporate proxy for compliance or security auditing
- Operating a Rack in a network-restricted environment where direct internet access is not available
- Funneling outbound traffic through a proxy for centralized logging or inspection

## Additional Information

For more information, see [HTTP Proxy Configuration](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/http_proxy_config.html) in the AWS docs.

```bash
$ convox rack params set HttpProxy=proxy.example.com:8080
```

The `HttpProxy` parameter will not change how your apps access the internet -- app traffic will not go through the proxy. Only the instances' own requests will use the proxy for outbound connections. If you want your apps to use the proxy, you must configure it in your `convox.yml`:

```yaml
services:
  web:
    build: .
    port: 3000
    environment:
      - http_proxy=10.0.1.124:8888
      - https_proxy=10.0.1.124:8888
      - HTTP_PROXY=10.0.1.124:8888
      - HTTPS_PROXY=10.0.1.124:8888
      - NO_PROXY=169.254.170.2
```

## See Also

- [Private](/reference/rack-parameters/Private)
- [InstanceBootCommand](/reference/rack-parameters/InstanceBootCommand)
- [AWS HTTP Proxy Configuration](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/http_proxy_config.html)
