---
title: "Running Migrations"
description: "Run database migrations and other administrative tasks during deployment."
---

# Running Migrations

Database migrations and other administrative tasks can be run as [one-off commands](/management/one-off-commands) against your Convox App. There are two approaches: running migrations manually from the CLI, or automating them as part of a deployment workflow.

## Running Migrations Manually

Use `convox run` to execute a migration command in a new container using the current Release:

```bash
$ convox run web rake db:migrate
```

This spawns a new container with the same environment variables and configuration as your running Service, executes the command, and exits.

### Common Examples

Rails:

```bash
$ convox run web rake db:migrate
```

Django:

```bash
$ convox run web python manage.py migrate
```

Node.js (Sequelize):

```bash
$ convox run web npx sequelize-cli db:migrate
```

### Running Against a Specific Release

By default, `convox run` uses the currently promoted Release. To run a migration against a different Release (for example, a newly built but not-yet-promoted Release):

```bash
$ convox run web --release RABCDEF1234 rake db:migrate
```

This is useful when you want to run migrations before promoting a new Release.

## Automating Migrations with Workflows

[Workflows](/console/workflows) support **Before Promote** and **After Promote** hooks that run one-off commands as part of the deployment pipeline.

To run migrations automatically before each deployment is promoted:

1. Navigate to your Workflow in the Console
2. In the **Before Promote** section, specify the Service and command:
   - **Service**: `web`
   - **Command**: `rake db:migrate`

The workflow will build your code, run the migration command against the new Release, and then promote the Release only if the migration succeeds. If the migration command fails (non-zero exit code), the promotion is halted.

**After Promote** hooks work the same way but execute after the new Release is live. Use these for tasks like cache warming or notification scripts.

## Best Practices

**Run migrations before promoting.** Use the `--release` flag or a Before Promote workflow hook to ensure migrations complete successfully before the new code starts serving traffic.

**Keep migrations backward-compatible.** During a rolling deployment, old and new code may run simultaneously. Migrations should not break the currently running Release. For example, add new columns before deploying code that uses them, and drop old columns only after all processes have updated.

**Use detached mode for long-running migrations.** If a migration takes longer than the default 1-hour timeout, run it in detached mode:

```bash
$ convox run web --detach rake db:migrate
```

Detached processes have no timeout. Monitor progress in the application logs.

**Test migrations in a staging environment first.** Use a separate Rack or App for staging to verify migrations apply cleanly before running them in production.

## See Also

- [One-off Commands](/management/one-off-commands)
- [Workflows](/console/workflows)
- [Builds](/deployment/builds)
- [Rolling Updates](/deployment/rolling-updates)
