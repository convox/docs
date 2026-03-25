---
title: "RouterMitigationMode"
description: "Configure how the Convox Rack load balancer handles potentially malicious HTTP desync requests."
---

# RouterMitigationMode

Security mitigation mode for the Rack's load balancer. This controls the AWS ALB desync mitigation mode, which protects against HTTP desync attacks. See [AWS desync mitigation mode announcement](https://aws.amazon.com/about-aws/whats-new/2020/08/application-and-classic-load-balancers-adding-defense-in-depth-with-introduction-of-desync-mitigation-mode/) for more information.

| Default value  | `defensive`                             |
| Allowed values | `defensive`, `monitor`, `strictest` |

## Use Cases

- Keeping `defensive` (default) for most production workloads, which blocks known desync threats while allowing ambiguous requests
- Setting to `strictest` for high-security environments that require RFC 7230 compliance, accepting that some legitimate but non-compliant clients may be rejected
- Setting to `monitor` temporarily to audit how your traffic would be affected before switching to a stricter mode

## Additional Information

The three modes behave as follows:

- **defensive** -- Blocks requests that are known to be unsafe, but allows ambiguous requests through. This is the recommended default for most applications.
- **monitor** -- Allows all requests through but logs those that would be blocked in `defensive` or `strictest` mode. Useful for auditing before tightening security.
- **strictest** -- Enforces strict RFC 7230 compliance. Any request that is ambiguous or non-compliant is rejected. This provides the highest security but may break clients that send non-standard HTTP requests.

```bash
$ convox rack params set RouterMitigationMode=strictest
```

## See Also

- [RouterSecurityGroup](/reference/rack-parameters/RouterSecurityGroup)
- [RouterInternalSecurityGroup](/reference/rack-parameters/RouterInternalSecurityGroup)
- [SslPolicy](/reference/rack-parameters/SslPolicy)
