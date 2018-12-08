---
title: "Running Migrations"
order: 500
---

You can easily run migrations or other administrative tasks using [one-off commands](/docs/one-off-commands/).

    $ convox run web rake db:migrate
    
You can also include a migration command in the ***Before Promote*** section of a [workflow](/console/workflows) which will cause the migration to be run as a [one-off command](/docs/one-off-commands/) against your new release prior to promoting it.
