---
title: "Tags"
description: "Add custom AWS resource tags to all resources created by the Convox Rack."
---

# Tags

Custom tags to add to AWS resources created by the Rack. Tags are applied to all taggable resources managed by the Rack's CloudFormation stack.

| Setting | Value |
|:--|:--|
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

Each set adds or changes the keys it lists. Keys left out keep their current values, and a Rack tag cannot be removed through Convox.

> **Note:** AWS has a limit of 50 tags per resource. Convox uses some tags internally, so keep Rack and App custom keys together under 40.

### Rack Tags on Apps

Rack `Tags` values reach a Generation 2 App when the App deploys. Generation 1 Apps do not receive them.

| Case | Result |
|:--|:--|
| An App deploys after a Rack key is set | The App copies the Rack value for that key |
| The Rack value changes for a key an App already carries | The App keeps the value it carries |
| An App sets its own value with the [Tags](/reference/app-parameters/Tags) App parameter | The App's value overrides the Rack value on that App, across later deploys and Rack `Tags` changes |

To move an App to a new Rack value for a key it already carries, set that key on the App, for example `convox apps params set Tags=CostCenter=12345 -a myapp`. Use CLI version 3.25.10 or newer, or 20261005214736 or newer. Older CLI versions apply the value but can print an error (a date-versioned CLI only with `--wait`), because the App's value then equals the Rack's.

> **Note:** Do not change only the case of a Rack `Tags` key, for example from `CostCenter` to `costcenter`. A Rack tag cannot be removed, so the Rack stack then carries both spellings. Tag keys are case-insensitive on resources such as IAM roles, which the Rack and its Generation 2 Apps tag, so the two spellings conflict.

## See Also

- [Tenancy](/reference/rack-parameters/Tenancy)
- [Private](/reference/rack-parameters/Private)
- [Tags App Parameter](/reference/app-parameters/Tags)
- [Service Tags](/management/service-tags)
