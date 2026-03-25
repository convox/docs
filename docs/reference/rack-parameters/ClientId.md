---
title: "ClientId"
description: "Anonymous identifier used internally by the Convox Rack for tracking purposes."
---

# ClientId

Anonymous identifier. This parameter is used internally by the Rack and does not typically need to be changed by users.

| Default value  | "" |

## Use Cases

- This is an internal parameter that is automatically managed by Convox
- You should not need to set or modify this value under normal circumstances

## Additional Information

The `ClientId` is an anonymous identifier generated during Rack installation. It is used internally for tracking purposes and does not contain any personally identifiable information. Changing this value is not recommended.

This parameter is managed automatically and typically does not need to be set manually. You can view the current value with:

```bash
$ convox rack params
```

## See Also

- [Password](/reference/rack-parameters/Password)
- [Rack Parameters](/reference/rack-parameters) for a full list of available parameters
