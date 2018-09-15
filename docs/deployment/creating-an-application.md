---
title: "Creating an Application"
order: 150
---

An application represents a single codebase deployed to Convox. You can create an application from the CLI:

**Select the Rack where you want the application to live:**

    $ convox switch <org>/<rack>

**Create an application:**

    $ convox apps create <name>

**Wait for the application to finish creating:**

    $ convox apps wait -a <name>
