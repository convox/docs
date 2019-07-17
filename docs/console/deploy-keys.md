---
title: "Deploy Keys"
---

Deploy keys are limited scope API keys that allow you to run the `build` and `deploy` commands from remote environments.

## Creating a Deploy Key

To generate a deploy key, log into your account at [console.convox.com](https://console.convox.com) and click on the **Settings** tab on the left. 

Go to the **Deploy Keys** section, give your deploy key a name, and click on **Create**.

## Using a Deploy Key

In your CI environment, download the latest version of the [Convox CLI](/introduction/installation) and use the deploy key like this:

```
$ env CONVOX_HOST=console.convox.com CONVOX_PASSWORD=<key> convox deploy
```
