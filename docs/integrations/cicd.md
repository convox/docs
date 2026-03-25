---
title: "CI/CD Integrations"
description: "Configure continuous integration and deployment pipelines with Convox using GitHub Actions, CircleCI, or deploy keys."
---

# CI/CD Integrations

Convox integrates with popular CI/CD platforms to automate your Build and deployment pipeline. You can trigger Builds, promote Releases, and manage deployment directly from your CI/CD workflow.

## Available Integrations

| Integration | Description |
|-------------|-------------|
| [GitHub Actions](/integrations/github-actions) | Deploy to Convox from GitHub Actions workflows |
| [CircleCI](/integrations/circleci) | Deploy to Convox from CircleCI pipelines |
| [Deploy Keys](/console/deploy-keys) | Generate API keys for CI/CD authentication |

## Authentication

All CI/CD integrations authenticate using deploy keys. A deploy key is a Convox API key scoped to specific Rack and App permissions. See [Deploy Keys](/console/deploy-keys) for setup instructions.

## See Also

- [Integrations Overview](/integrations/integrations)
- [Builds](/deployment/builds)
- [Releases](/deployment/releases)
