---
title: "Tags"
description: "Add or change custom AWS tags on a single Generation 2 App's stack and resources."
---

# Tags

Custom AWS tags for a single App. The tags are added to the App's CloudFormation stack and the resources in it, and a key set on the App overrides the [Tags](/reference/rack-parameters/Tags) Rack parameter value for that key on this App. This parameter applies to Generation 2 Apps only.

| Setting | Value |
|:--|:--|
| Default value  | "" |
| Format        | `<key>=<val>,<key>=<val>` |

Requires rack version 20261005214736 or newer.

## Use Cases

- Attributing one App's AWS costs to its own team or cost center in AWS Cost Explorer
- Applying governance or compliance tags that only one App needs
- Overriding a Rack-wide tag value, such as `Environment`, on a single App

## Additional Information

Set one or more keys with `convox apps params set`:

```bash
$ convox apps params set Tags=CostCenter=abc,Team=web -a myapp
Updating parameters... OK
```

Each set adds or changes the keys it lists, and keys left out keep their current values. A tag cannot be removed through this parameter.

```bash
$ convox apps params set Tags=Team=api -a myapp
Updating parameters... OK
$ convox apps params -a myapp
...
Tags                                   CostCenter=abc,Team=api
...
```

Unlike App parameters that a Rack update adds, `Tags` does not need a deploy first: it can be set on any Generation 2 App as soon as the Rack is updated. `Tags` is set with the CLI only. A `Tags` entry under `params:` in `convox.yml` is ignored.

### When the Tags Apply

| Resource | When the tags apply |
|:--|:--|
| App stack and its resources: ECR repository, log group, settings bucket, IAM roles, timer launcher Lambda function, ACM certificates | During the `convox apps params set` update, with no new Release and no restart |
| Service, timer and resource stacks and their resources: ECS services, task definitions, target groups, timer rules, databases and caches | During the same update |
| ECS tasks of Services | Only with [TaskTags](/reference/app-parameters/TaskTags) set to `Yes`, on tasks started after the update |

Running tasks keep the tags they started with. To tag them, wait for the update to finish and restart the App:

```bash
$ convox apps wait myapp
Waiting for app... OK
$ convox restart -a myapp
Restarting web... OK
```

This parameter does not tag Processes started with `convox run` or their task definitions, timer tasks, Build tasks, the Secrets Manager secret used by [SecretsManagerEnv](/reference/app-parameters/SecretsManagerEnv), or Rack instances and their EBS volumes.

### Rules

| Rule | Behavior |
|:--|:--|
| App generation | Generation 2 Apps only. A Generation 1 App returns `Tags is only supported on generation 2 apps` |
| Format | `key=value` pairs separated by commas. The first `=` in a pair ends the key, so a value can contain `=` and a key cannot |
| Characters | Letters, numbers, spaces and `_ . : / = + - @`. Keys up to 128 characters, values up to 256. Quote the argument when it contains spaces |
| Reserved keys | `App`, `Generation`, `Name`, `Rack`, `System`, `Type` and `Version` are rejected, in any case |
| `aws:` prefix | Rejected on keys and values, in any case |
| Key case | A key that differs only in case from another key in the set, or from a key the App or the Rack already has, is rejected: `Tags key costcenter conflicts with CostCenter; tag keys are case-insensitive` |
| Empty values | `Tags=` and empty values such as `Tags=CostCenter=` are rejected |
| Duplicate keys | The first occurrence wins: `Tags=Team=a,Team=b` sets `Team=a` |

When `Tags` is rejected, no other parameter in the same command is applied.

### Rack and App Values

| Case | Result |
|:--|:--|
| The App sets its own value for a key | The App's value overrides the Rack value on this App, across later deploys and Rack `Tags` changes |
| The App deploys after a Rack `Tags` key is set | The App copies the Rack value for that key |
| The Rack value changes for a key the App already carries | The App keeps the value it carries. Set the key on the App to change it |

`convox apps params` shows a `Tags` row with the App's values that differ from the Rack's current `Tags` value, sorted by key. Values the App carries unchanged from the Rack are not listed, so an App can show a `Tags` row without a `Tags` set when it carries a Rack value that the Rack later changed.

`convox apps export` includes the `Tags` row with the App's other parameters, and `convox apps import` sets those values on the new App.

### Older Rack Versions

On a Rack older than 20261005214736, `convox apps params set` does not apply `Tags`, `convox apps params` shows no `Tags` row, and an import does not apply exported `Tags` values. After a downgrade to such a version, tags set earlier stay on the App's stack and keep applying on each deploy, but they can no longer be changed.

> **Note:** AWS allows 50 tags per resource, and Convox adds several of its own. Keep Rack and App custom keys together under 40. Convox does not check this limit.

> **Note:** An AWS Organizations tag policy with enforcement can reject a tag value on some resource types, and the App update then rolls back.

> **Note:** Use CLI version 3.25.10 or newer, or 20261005214736 or newer (`sudo convox update`). Older CLI versions apply the tags but can print an error and exit with a non-zero status when a set has more than one key, sets a value equal to the Rack's, or targets an App that already has its own `Tags` values.

## See Also

- [Tags Rack Parameter](/reference/rack-parameters/Tags)
- [TaskTags](/reference/app-parameters/TaskTags)
- [Service Tags](/management/service-tags)
- [apps params set](/reference/cli-commands/apps-params-set)
