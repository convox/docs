---
title: "Password"
description: "The required API HTTP password used to authenticate with the Convox Rack."
---

# Password

(REQUIRED) API HTTP password. This is the password used to authenticate API requests to the Rack. It is set during Rack installation and is required for all Rack operations.

| Minimum length | `1`  |
| Maximum length | `50` |

## Use Cases

- Set during Rack installation to secure API access
- Rotate periodically to comply with security policies
- Update when revoking access for former team members who knew the previous password

## Additional Information

The Password parameter is required and must be between 1 and 50 characters. It is used by the Convox CLI and API clients to authenticate with the Rack API endpoint.

This value is stored securely and is not displayed in CloudFormation outputs or parameter listings. When you run `convox rack params`, the password value is masked.

The password is typically set automatically during Rack installation. If you need to change it, be aware that this will require updating the credentials used by any CI/CD pipelines or CLI configurations that connect to this Rack.

```bash
$ convox rack params set Password=new-secure-password
```

## See Also

- [PrivateApi](/reference/rack-parameters/PrivateApi)
- [PrivateApiSecurityGroup](/reference/rack-parameters/PrivateApiSecurityGroup)
