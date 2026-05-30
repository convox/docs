---
title: "SecretsManagerEnv"
description: "Inject App environment variables into ECS tasks from AWS Secrets Manager instead of decrypting from S3 at runtime."
---

# SecretsManagerEnv

Resolve App environment variables from AWS Secrets Manager at ECS task launch. When enabled, each task reads its environment from a per-App secret rather than fetching the encrypted env from S3 and decrypting it with KMS at startup. S3 remains the source of truth (this is a write-through model), and the Secrets Manager secret is populated during `convox releases promote`.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Use Cases

- Keep environment values out of the Rack S3 bucket as the runtime delivery path, letting ECS pull them directly from Secrets Manager
- Centralize App secrets behind Secrets Manager IAM policies, rotation tooling, and CloudTrail audit logging
- Reduce per-task startup decryption work by replacing the S3 plus KMS fetch with native ECS secret injection

## Additional Information

The Rack stores each App's environment in a Secrets Manager secret named `{rack}/{app}` (for example `myrack/myapp`). The secret is tagged with `convox:rack` and `convox:app`. During promote, the Rack writes the current environment to this secret, then renders each Service and Timer task definition with `secrets` entries that reference the secret using the `arn:key::` format, so ECS resolves each variable by JSON key at launch.

Enable the parameter for the whole Rack:

```bash
$ convox rack params set SecretsManagerEnv=Yes
```

Enabling the parameter does not change running tasks. The secret is created and populated on the next promote, so deploy or promote a new [Release](/deployment/builds) for each App after turning it on:

```bash
$ convox deploy -a myapp
```

When a task starts, the Rack looks up the secret ARN and builds the task definition `secrets` block from it. With Secrets Manager active, the S3 environment delivery variables (`CONVOX_ENV_KEY`, `CONVOX_ENV_URL`, `CONVOX_ENV_VARS`) are omitted from the task because ECS supplies the values directly.

### Precedence

Three levels control whether an App uses Secrets Manager, from lowest to highest priority:

1. The Rack parameter `SecretsManagerEnv` sets the default for every App on the Rack.
2. An App parameter opts a single App in even when the Rack default is `No`. Only `Yes` is honored at this level. Setting `SecretsManagerEnv=No` as an App parameter does not turn the feature off if the Rack default is `Yes`.
3. A `params` block in `convox.yml` overrides both the Rack and App parameters. This is the way to disable Secrets Manager for one App while the Rack default stays `Yes`.

Opt a single App in with an App parameter:

```bash
$ convox apps params set SecretsManagerEnv=Yes -a myapp
```

Disable Secrets Manager for one App while the Rack default is `Yes` by setting it in that App's `convox.yml`:

```yaml
params:
  SecretsManagerEnv: "No"
```

### Failure Handling

Secrets Manager problems never block a deploy. If the write to Secrets Manager fails during promote, the Rack logs a warning and falls back to S3 injection for that Release, so the deploy still completes. After three consecutive write failures for the same App, the Rack emits a warning event suggesting you set `SecretsManagerEnv=No` or investigate Secrets Manager availability.

### Deletion

When an App is deleted, its `{rack}/{app}` secret is scheduled for deletion with the standard 30-day recovery window. The secret can be restored within that window if the App is recreated before the window closes.

### Gen1 Limitation

Gen1 Apps use a separate promote path, so Secrets Manager injection does not apply to them. Enabling the parameter has no effect on gen1 Apps, which continue to use the S3 plus KMS environment delivery path.

### Downgrade

A Rack version that does not declare `SecretsManagerEnv` treats the parameter as `No`, so the feature is safe across version downgrades. Tasks fall back to the S3 plus KMS environment delivery path.

## See Also

- [Environment](/application/environment)
- [Encryption](/reference/rack-parameters/Encryption)
- [Rack Parameters](/reference/rack-parameters)
