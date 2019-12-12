---
title: "Deploy Keys"
---

Deploy keys are limited scope API keys that allow you to run the following commands from your remote environment:

- `build`
- `builds`
- `builds export`
- `builds import`
- `deploy`
- `env set --replace`
- `logs`
- `rack`
- `racks`
- `run`

## Creating a Deploy Key

To generate a deploy key, log into your account at [console.convox.com](https://console.convox.com) and click on the **Settings** tab on the left. 

Go to the **Deploy Keys** section, give your deploy key a name, and click on **Create**.

<div class="block-callout block-show-callout type-info" markdown="1">
Deploy keys are specific to the organization they are created within.
</div>

## Using a Deploy Key

In your CI environment, download the latest version of the [Convox CLI](/introduction/installation) and use the deploy key like these examples:

```sh
$ env CONVOX_HOST=console.convox.com CONVOX_PASSWORD=<key> convox deploy
$ env CONVOX_HOST=console.convox.com CONVOX_PASSWORD=<key> convox run web bin/migrate
$ env CONVOX_HOST=console.convox.com CONVOX_PASSWORD=<key> convox env set NODE_ENV=production FOO=bar ... --replace
$ env CONVOX_HOST=console.convox.com CONVOX_PASSWORD=<key> convox builds export <build ID> -a <app1> -r <rack1> | convox builds import -a <app2> -r <rack2>
```
