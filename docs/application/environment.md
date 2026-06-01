---
title: "Environment"
description: "Manage environment variables for Convox applications, including global, service-specific, and interpolated values."
---

# Environment

Convox applications are configured using environment variables.  You will need to [specify](#definition) your variables in your `convox.yml` to be available to your running application.  You can [set and unset](#setting-and-editing-your-environment-variables) them from the Convox CLI or web Console before or after any code deployments.  You can even control the [release](#promoting-your-environment-variable-changes) of your environmental changes.  The variables are stored securely and encrypted and only available to the specific application and services you require.

## Definition

All environment variables defined in your `convox.yml` are required to start up your application, but you can set a default or empty value in your configuration for those that are appropriate to do so.

### Global

Environment variables specified in the top level `environment:` section will be available to all Services in the application:

```yaml
environment:
  - ENABLED=true
  - LICENSE=
  - ENVIRONMENT
...
```

In this example the `ENABLED` variable has a default value of `true`, the `LICENSE` variable will be empty and the `ENVIRONMENT` variable must be defined using the CLI or Web console before the application will start.

### Service

You can also configure each Service individually to limit secrets to the Services that need them:

```yaml
services:
  web:
    build: .
    environment:
      - LOG_LEVEL=debug
      - DEFAULT_CLIENT=
      - MULTI_TENANT
```

In this example the `LOG_LEVEL` variable has a default value of `debug`, `DEFAULT_CLIENT` will be empty and the `MULTI_TENANT` variable must be defined using the CLI or Web console before the application will start.

Environment variables defined at the service level will not be shared or visible to other Services in the same App.  You can have a variable with the same name in separate Services with different declared values, but note that updating this variable after deployment will update the value in all Services with that variable name.  

### Wildcard

You can use the wildcard syntax to provide all available environment variables to a Service:

```yaml
services:
  web:
    build: .
    environment:
      - "*"
```

If you wish to declare and inject new environment variables on demand without declaring them in the `convox.yml` beforehand, you should use the wildcard syntax to pick them up.

## Setting and Editing your Environment Variables

You can configure your env vars through the Convox web console for your cloud-based Racks, or through the Convox CLI your local and cloud Racks:

```bash
$ convox env set FOO=bar
Setting FOO... OK
Release: RFMGESPZHC

$ convox env get FOO
bar

$ convox env unset FOO
Unsetting FOO... OK
Release: RZIDXFIJAC

$ convox env edit
Setting ... OK
Release: RRJQTMVKRS
```

`convox env edit` opens a terminal editor where you can interactively update your environment variables.

You can use `convox env edit` along with `convox releases` to restore an Environment Variable set from another release:

First run `convox releases -a <app-name>` to identify the correct `RELEASE-ID`

Next, examine the release's environment variables with `convox releases info <RELEASE-ID> -a <app-name>`

Copy the desired variables.

Finally run `convox env edit` paste and save the variables to create and promote a new `RELEASE` for your App with updated environment variables.

## Promoting your environment variable changes

If you update your env vars through the CLI, that will create a new release containing your changes.  By default this will not be promoted.
For this to take effect, you will need to promote that release:

```bash
$ convox env set FOO=bar
Setting FOO... OK
Release: RFMGESPZHC
$ convox releases promote RFMGESPZHC
Promoting RFMGESPZHC... OK
```

Alternatively you can pass the `--promote`/`-p` flag to the env commands to automatically promote the release:

```bash
$ convox env set FOO=bar --promote
Setting FOO... OK
Release: RYFPMIHLPKD
Promoting RYFPMIHLPKD... OK
```

Environment variable updates made through the Convox web console create releases which are promoted automatically for you.

## Environment variable comprehensive example

Given this example `convox.yml`:

```yaml
environment:
  - ENABLED=true
  - LICENSE=
  - ENVIRONMENT
services:
  web:
    ...
    environment:
      - WEB_LOGGING_LEVEL=info
      - WEB_CLIENT=
      - WEB_MULTI_TENANT
      - SHARED=web
    ...
  worker:
    ...
    environment:
      - WORKER_LOGGING_LEVEL=error
      - WORKER_CLIENT=
      - WORKER_MULTI_TENANT
      - SHARED=worker
    ...
  anotherservice:
    ...
    environment:
      - ANOTHER_LOGGING_LEVEL=trace
      - "*"
```

The `ENVIRONMENT`, `WEB_MULTI_TENANT` and `WORKER_MULTI_TENANT` variables will all need to be set before first deployment.  

Once deployed, the `web` service will have the `ENABLED, LICENSE, ENVIRONMENT, WEB_LOGGING_LEVEL, WEB_CLIENT, WEB_MULTI_TENANT, SHARED` environment variables all available to it, and `SHARED` will have the value of `web`.
The `worker` service will have the `ENABLED, LICENSE, ENVIRONMENT, WORKER_LOGGING_LEVEL, WORKER_CLIENT, WORKER_MULTI_TENANT, SHARED` environment variables all available to it, and `SHARED` will have the value of `worker`.
The `anotherservice` will have `ENABLED, LICENSE, ENVIRONMENT, ANOTHER_LOGGING_LEVEL` environment variables all available to it, as well as `WEB_MULTI_TENANT` and `WORKER_MULTI_TENANT` as I will have to have previously set those before deployment.

If I update any of the `ENABLED, LICENSE, ENVIRONMENT`, or `SHARED` environment variables and promote that update, those values will then be reflected in the `web` and `worker` services, as well as in `anotherservice`.  If I update any of the `WEB_*` or `WORKER_*` variables and promote that update, those values will then be reflected in their respective services, as well as in `anotherservice`.
If I set a value against a `NEW_VARIABLE` environment variable which wasn't originally declared, and promote that update, it will be reflected and available in the `anotherservice` only.  If I want that to be available to the other services, I should add that declaration to the `convox.yml` and redeploy, at which point they will then pick that up.

## Secrets Manager Injection

Convox controls how the environment variables for a Release are delivered into your running Processes. By default, the encrypted environment is stored in S3 and decrypted by each container at startup. You can switch a Rack (or an individual App) to resolve variables from AWS Secrets Manager instead, so values are injected by ECS when each task launches rather than fetched by the container itself.

### Default delivery (S3 and KMS)

When [SecretsManagerEnv](/reference/rack-parameters/SecretsManagerEnv) is `No` (the default), Convox writes the variables for the Release to an encrypted object in the Rack's S3 settings bucket. Each container receives `CONVOX_ENV_URL`, `CONVOX_ENV_KEY`, and `CONVOX_ENV_VARS` in its task definition. On startup the container fetches the encrypted object from S3 and decrypts it with the Rack's KMS key. The plaintext values never appear in the ECS task definition.

### Secrets Manager delivery

When [SecretsManagerEnv](/reference/rack-parameters/SecretsManagerEnv) is `Yes`, Convox maintains a per-App secret in AWS Secrets Manager named `{rack}/{app}`. The task definition omits the `CONVOX_ENV_URL`, `CONVOX_ENV_KEY`, and `CONVOX_ENV_VARS` variables and instead lists each variable as an ECS secret reference, so ECS resolves the values from the secret and injects them when the task launches. The task execution role is granted read access to that single secret only.

S3 remains the source of truth. Setting or unsetting variables, editing the environment, and restoring values from an earlier Release all behave exactly as described above. The Secrets Manager secret is refreshed during a [Release promote](/deployment/releases), so the secret reflects the environment of whatever Release you promote, not the variables you last set. Promote a Release after enabling the parameter so the secret is populated before any new tasks reference it.

```bash
$ convox rack params set SecretsManagerEnv=Yes
$ convox env set FOO=bar --promote
Setting FOO... OK
Release: RYFPMIHLPKD
Promoting RYFPMIHLPKD... OK
```

The parameter can be set on the Rack to cover every App, or on an individual App stack to opt a single App in. An App-level `Yes` enables Secrets Manager delivery for that App regardless of the Rack value, so you can roll the behavior out one App at a time before enabling it Rack-wide.

### Failure handling and fallback

If the Secrets Manager write fails during promote (for example, an API throttle or a permissions gap), Convox logs a warning and falls back to the default S3 and KMS delivery for that Release. The promote still succeeds and your Processes start with the correct environment. After three consecutive failures for an App, Convox sends a `rack:warning` event so you can investigate or set `SecretsManagerEnv=No`. Because both delivery paths read the same source of truth, switching the parameter off and promoting again returns the App to S3 and KMS delivery with no change to your variables.

Gen1 Apps are unaffected by this parameter. They follow a separate promote path that always uses S3 and KMS delivery.

## See Also

- [convox.yml Reference](/application/convox-yml)
- [Services](/application/services)
- [Multiple Environments](/application/multiple-environments)
- [Releases](/deployment/releases)
- [Running Migrations](/deployment/running-migrations)
