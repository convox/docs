---
title: "SecretsManagerEnv"
description: "Inject this application's environment variables into ECS tasks from AWS Secrets Manager instead of the default S3 delivery."
---

# SecretsManagerEnv

Source this application's [environment variables](/application/environment) from [AWS Secrets Manager](https://aws.amazon.com/secrets-manager/) when ECS tasks launch, rather than the default S3-based delivery. This parameter applies to gen2 applications only and changes only the delivery mechanism, so you continue to manage variables normally with `convox env set`.

| Default value  | `No` |
| Allowed values | `Yes`, `No` |

## Use Cases

- Set to `Yes` to keep this application's environment out of the Rack S3 bucket as the runtime delivery path and let ECS read it directly from Secrets Manager
- Set to `Yes` to centralize this application's variables behind Secrets Manager IAM policies and CloudTrail audit logging
- Keep as `No` for gen1 applications, where the parameter has no effect, or when the default S3 delivery is sufficient

## Additional Information

When set to `Yes`, the Rack writes this application's full environment to a single Secrets Manager secret named `<rack>/<app>` during release promote and threads the secret ARN into the application stack. The gen2 task definition then builds one ECS secret per environment key, referencing each value as `<arn>:<key>::`, and the S3 delivery variables (`CONVOX_ENV_KEY`, `CONVOX_ENV_URL`, `CONVOX_ENV_VARS`) are omitted from the container. No user-side Secrets Manager setup is required: the Rack creates and manages the secret, and the per-application task execution role is granted a scoped `secretsmanager:GetSecretValue` permission limited to that one secret.

This parameter sits alongside a Rack-level parameter of the same name. The Rack value sets the default for every application, and setting this application parameter to `Yes` opts a single application in even when the Rack default is `No`. The effective value is enabled when either the Rack or the application is set to `Yes`.

A few operational notes:

- The marshaled environment must stay under the 64KB Secrets Manager limit, or the write fails.
- Secrets Manager problems never block a deploy: on a write or lookup failure the Rack logs a warning and falls back to S3 delivery for that release. After three consecutive failures for an application, the Rack emits a `rack:warning` event suggesting you set `SecretsManagerEnv=No` or investigate Secrets Manager availability.
- Enabling the parameter does not change running tasks. The secret is written during promote, so deploy or promote a new release after toggling it.

```bash
$ convox apps params set SecretsManagerEnv=Yes
```

To set a default for every application on the Rack instead of a single one, use the Rack-level [SecretsManagerEnv](/reference/rack-parameters/SecretsManagerEnv) parameter.

## See Also

- [Environment](/application/environment)
- [SecretsManagerEnv](/reference/rack-parameters/SecretsManagerEnv)
- [Rack Parameters](/reference/rack-parameters)
