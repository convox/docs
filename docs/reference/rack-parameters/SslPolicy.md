---
title: "SslPolicy"
description: "Specify the SSL/TLS security policy for the Convox Rack load balancer."
---

# SslPolicy

SSL/TLS policy for the primary Rack load balancer. This controls the SSL/TLS protocols and ciphers that the load balancer supports when terminating HTTPS connections.

| Default value  | "" |
| Allowed values | [ELB SSL Policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#describe-ssl-policies) |

## Use Cases

- Setting a stricter policy (e.g., `ELBSecurityPolicy-TLS13-1-2-2021-06`) for compliance requirements that mandate TLS 1.2 or higher
- Using a custom policy to support older clients that require TLS 1.0 or specific cipher suites
- Applying a Forward Secrecy policy (e.g., `ELBSecurityPolicy-FS-1-2-Res-2020-10`) for enhanced security

## Additional Information

When this parameter is blank, AWS uses its default SSL policy for Application Load Balancers, which supports a broad range of TLS versions and cipher suites. Setting a specific policy allows you to enforce minimum TLS version requirements and restrict which cipher suites are available.

Common policy choices include:

- **ELBSecurityPolicy-TLS13-1-2-2021-06** -- Supports only TLS 1.2 and TLS 1.3
- **ELBSecurityPolicy-FS-1-2-Res-2020-10** -- Requires forward secrecy with TLS 1.2 minimum
- **ELBSecurityPolicy-2016-08** -- The default policy, supporting TLS 1.0 through 1.2

See the [AWS ELB SSL Policies documentation](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#describe-ssl-policies) for the full list of available policies and their supported protocols and ciphers.

```bash
$ convox rack params set SslPolicy=ELBSecurityPolicy-TLS13-1-2-2021-06
```

## See Also

- [RouterSecurityGroup](/reference/rack-parameters/RouterSecurityGroup)
- [RouterMitigationMode](/reference/rack-parameters/RouterMitigationMode)
- [LoadBalancerIdleTimeout](/reference/rack-parameters/LoadBalancerIdleTimeout)
