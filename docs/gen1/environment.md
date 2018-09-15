---
title: "Environment"
---

Convox applications are configured using environment variables. Environment management differs depending on whether you are running your applications locally or in the cloud. See the sections below for details.

1. [Local](#local)
1. [Deployed](#deployed)
1. [Scope](#scope)
1. [Sensitive Information](#sensitive-information)
1. [Under the hood](#under-the-hood)
1. [Additional Security](#additional-security)
1. [Further Reading](#further-reading)

## Local

<div class="block-callout block-show-callout type-info" markdown="1">
The priority for evaluating env values is:

1. [`.env`](/docs/gen1/environment/#env)
2. [host environment](/docs/gen1/environment/#host-environment)
3. [`docker-compose.yml`](/docs/gen1/environment/#docker-composeyml)
</div>

### `.env`

When running your application with `convox start` you should set values for your application's environment in a `.env` file:

```
SECRET_KEY=xyzzy
FOO=bar
```

### Host environment

Variables defined in your local environment will be taken into account as long as:

- they are declared in `docker-compose.yml` (in either `KEY` or `KEY=VALUE` format), and
- are not defined in `.env`.

```
$ DEVELOPMENT=true convox start
```

### `docker-compose.yml`

`convox start` will always read `.env`, but the environment variables set there will not be passed along to your application's containers unless you declare them in `docker-compose.yml`.

```
services:
  web:
    build: .
    environment:
      - SECRET_KEY
      - FOO
```

The `environment` section of `docker-compose.yml` plays a couple of roles locally:

1. _It serves as a list of required environment variables._ `convox start` will refuse to run your application if it doesn't find values for each of your required env vars in `.env`.
1. _It allows you to set default values for environment variables._ Values in `.env` will override these defaults:

```
    environment:
      - SECRET_KEY
      - FOO=default
```

## Deployed

### `convox env`

When dealing with a deployed Convox application, use the `convox env` commands to manage your environment:

```
$ convox env                       # retrieve all env vars
$ convox env set FOO=bar BAZ=qux   # set one or more env vars
$ convox env get FOO               # get the value of a single env var
$ convox env unset FOO             # delete an env var
```

To save time, you can also pipe the contents of an environment to `convox env set`:

```
$ cat .env | convox env set
Updating environment... OK
To deploy these changes run `convox releases promote RLGUFIKSJFY`
```

#### Flags

The following flags can optionally be provided to `convox env set`:

* `--promote`: automatically promote the release.
* `--id`: send `env set` logs to stderr; send only the resulting release ID to stdout (default behavior is to send all logs to stdout).

### docker-compose.yml

For deployed Convox applications, `docker-compose.yml` can only be used to set default values.

```
services:
  web:
    build: .
    environment:
      - SECRET_KEY    # ignored by deployed apps; only used locally
      - FOO=default   # override with `convox env set FOO=newvalue`
```

## Linking

Convox links containers by injecting environment variables. For example, if your `docker-compose.yml` links a service named `web` to a service named `database`, Convox will add environment variables in the form `DATABASE_URL`, `DATABASE_SCHEME`, etc to the `web` container environment. **This will override any environment variables you may have previously defined by the same name.** For details, see [Defining links](/docs/gen1/linking#defining-links).

## Scope

Environment variables set via `convox env` are considered app-level configuration.

There isn't currently a way to set an environment variable for a whole cluster rather than just a single app.

You can set environment variables at the [service](/docs/definitions/#service) level in `docker-compose.yml`.

## Sensitive Information

Since `docker-compose.yml` will typically be under version control, we recommend not setting sensitive default values there.

We also recommend leaving your `.env` file out of version control and any images you build of your application to protect against leaking sensitive information. That might mean adding `.env` to `.gitignore` and `.dockerignore`, for example.

## Under the hood

When you use the `convox env` commands, you're actually interacting with a KMS-encrypted file stored in an S3 bucket in your AWS account. For example, `convox env` fetches that file, decrypts it, and then prints the environment variable pairs it contains. Likewise, `convox env get` fetches and decrypts that file, but then prints only the requested environment variable. As you might expect, `convox env set` fetches and decrypts the env file from S3, then appends the env var pairs you pass as arguments to that file, before re-encrypting the file and putting it back in the S3 bucket. Finally, `convox env unset` removes the specified variable from the file, rather than adding or updating it like `convox env set`.

The environment of each service is stored in its ECS Task Definition so that ECS can set the environment of each task (i.e. container) appropriately. As a result, you can view the decrypted environment data in the app's ECS Task Definitions. If this is a concern for you, please see the following Additional Security section.

## Additional Security

You can prevent your environment variables from being visible in ECS Task Definitions by setting the `convox.environment.secure=true` label on each relevant service.

    version: "2"

    services:
      web:
        labels:
          - convox.environment.secure=true

When this label is set the environment variables will not be applied to that service's Task Definitions. Instead, the application itself will need to download and decrypt the environment file. To facilitate this, two environment variables will be available at runtime:

- `SECURE_ENVIRONMENT_URL` - The URL of the S3 file that contains your environment variables
- `SECURE_ENVIRONMENT_KEY` - The ARN of the KMS key needed to decrypt the environment file

Although it's possible to handle this directly with the AWS SDK, Convox has provided a binary called `secure-environment` to make things easier. You can download the latest binary from the [releases page](https://github.com/convox/secure-environment/releases). It can be used with a Docker ENTRYPOINT script to download, decrypt, and source your environment variables at container runtime.

See the example app at [https://github.com/convox-examples/secure-env-example](https://github.com/convox-examples/secure-env-example) for more details.

## Further Reading

Convox embraces the strict separation of configuration from code advanced by the [Twelve-Factor App](http://12factor.net/config) methodology.
