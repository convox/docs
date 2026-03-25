---
title: "Tags"
description: "Add custom AWS resource tags to all resources created by the Convox Rack."
---

# Tags

Custom tags to add to AWS resources created by the Rack. Tags are applied to all taggable resources managed by the Rack's CloudFormation stack.

| Default value  | "" |
| Format        | `<key>=<val>,<key>=<val>` |

Example: `key1=val1,key2=val2`

## Use Cases

- Adding cost allocation tags (e.g., `Environment=production`, `Team=platform`) for AWS billing reports
- Applying compliance or governance tags required by your organization (e.g., `DataClassification=internal`)
- Tagging resources for use with AWS resource groups or automated management policies

## Additional Information

Tags are set as a comma-separated list of key-value pairs:

```bash
$ convox rack params set Tags="Environment=staging,Team=engineering,CostCenter=12345"
```

AWS tags are useful for organizing resources, controlling access via IAM policies, filtering in the AWS console, and tracking costs in AWS Cost Explorer. Convox automatically adds `Name` and `Rack` tags to resources; your custom tags are added in addition to these.

> **Note:** AWS has a limit of 50 tags per resource. Convox uses some tags internally, so ensure your custom tags do not exceed the remaining capacity.

## See Also

- [Tenancy](/reference/rack-parameters/Tenancy)
- [Private](/reference/rack-parameters/Private)
