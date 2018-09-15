---
title: "Deploy Keys"
---

Deploy keys are special, limited scope API keys that allow you to run the `build` and `deploy` commands from remote servers for the purposes of continuous integration.

You should give a deploy key to a CI service like [Travis CI](/docs/travis-ci/), [Circle CI](/docs/circle-ci/) or [Datadog](/docs/datadog/) so it can deploy code but not access or modify any other Rack resources. For details, see the **Integrations** section in the list of topics on the left.

![Deploy Keys](/assets/images/docs/rbac/deploy-keys.png){: .screenshot } *List of deploy keys*


## Creating deploy keys

To generate a **deploy key**, log into your account at [console.convox.com](https://console.convox.com) and click on the **Organization** section on the left. Click on the **Deploy Keys** tab to create a new key.
