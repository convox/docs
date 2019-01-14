---
title: "Deploy Keys"
---

Deploy keys are special, limited scope API keys that allow you to run the `build` and `deploy` commands from remote servers for the purposes of continuous integration.

![](/assets/images/docs/console/deploy-keys.png)

## Creating deploy keys

To generate a **deploy key**, log into your account at [console.convox.com](https://console.convox.com) and click on the **Settings** section on the left. Click on create in the **Deploy Keys** section to create a new key.

## Using a deploy key

```
$ env CONVOX_PASSWORD=<key> convox deploy
```
