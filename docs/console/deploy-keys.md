---
title: "Deploy Keys"
---

Deploy keys are special, limited scope API keys that allow you to run the `build` and `deploy` commands from remote servers for the purposes of continuous integration.

You should give a deploy key to a CI service like [Travis CI](/integrations/travis-ci), [Circle CI](/integrations/circle-ci) or [Datadog](/integrations/datadog) so it can deploy code but not access or modify any other Rack resources. For details, see the **Integrations** section in the list of topics on the left.

![](/assets/images/docs/console/deploy-keys.png)


## Creating deploy keys

To generate a **deploy key**, log into your account at [console.convox.com](https://console.convox.com) and click on the **Settings** section on the left. Click on create in the **Deploy Keys** section to create a new key.
