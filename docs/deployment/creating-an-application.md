---
title: "Creating an Application"
description: "Create and configure a Convox Application to deploy your codebase."
---

# Creating an Application

An Application represents a single codebase deployed to a Convox Rack. Each App has its own set of Services, environment variables, and resources defined by a [convox.yml](/application/convox-yml) manifest.

## Creating an App

Select the Rack where you want the Application to live, then create it:

```bash
$ convox switch myorg/production
$ convox apps create myapp
```

Wait for the Application to finish creating:

```bash
$ convox apps wait -a myapp
```

The `create` command provisions the underlying AWS infrastructure (ECR repository, IAM roles, S3 settings bucket, and CloudWatch log group) via CloudFormation. This typically takes 1-2 minutes.

### Options

| Flag | Description |
|------|-------------|
| `-g`, `--generation` | App generation: `1` (legacy) or `2` (default). Generation 2 uses Fargate and ALB. |
| `--wait` | Block until the App is fully created before returning. |

Example creating a Generation 1 App:

```bash
$ convox apps create myapp -g 1
```

### Naming Rules

App names must follow these constraints:

- Lowercase letters, numbers, hyphens, and underscores only
- No uppercase characters
- Cannot be the same as the Rack name

## Viewing App Information

Use `convox apps info` to see the current state of an Application:

```bash
$ convox apps info -a myapp
Name        myapp
Status      running
Generation  2
Locked      false
Release     RABCDEF1234
```

| Field | Description |
|-------|-------------|
| **Name** | Application name |
| **Status** | Current state: `creating`, `running`, `updating`, or `deleting` |
| **Generation** | App generation (`1` or `2`) |
| **Locked** | Whether [termination protection](/reference/app-parameters) is enabled |
| **Release** | Currently promoted Release ID |

## Listing Apps

```bash
$ convox apps
APP    STATUS
myapp  running
api    running
```

## Deleting an App

```bash
$ convox apps delete myapp
```

> Deleting an App removes all associated AWS resources, including the ECR repository, load balancer rules, and ECS services. This action cannot be undone.

If the App has termination protection enabled (`Locked=true`), you must unlock it first:

```bash
$ convox apps unlock myapp
$ convox apps delete myapp
```

## Next Steps

After creating an App, you typically:

1. Set [environment variables](/application/environment)
2. Deploy your code with `convox deploy`
3. Monitor with `convox apps info` and `convox logs`

## See Also

- [Builds](/deployment/builds)
- [convox.yml](/application/convox-yml)
- [Environment](/application/environment)
- [App Parameters](/reference/app-parameters)
- [App Statuses](/reference/app-statuses)
