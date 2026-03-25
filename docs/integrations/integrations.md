---
title: Integrations
description: Third-party integrations for CI/CD pipelines, log forwarding, and private container registries with Convox.
---

# Integrations

Convox integrates with third-party services for continuous deployment, log aggregation, and private container image storage. Each integration connects to your Rack through standard protocols (webhooks, syslog, Docker registry credentials).

## CI/CD

Automate builds and deployments from your existing CI/CD pipeline.

| Integration | Description |
|-------------|-------------|
| [GitHub Actions](/integrations/github-actions) | Deploy to Convox using official GitHub Actions |
| [CircleCI](/integrations/circleci) | Deploy to Convox as part of a CircleCI workflow using the Convox Orb |
| [Deploy Keys](/console/deploy-keys) | Generate keys for automated deployment from any CI system |

## Logging

Forward application logs to third-party log aggregation services using syslog.

| Integration | Description |
|-------------|-------------|
| [Syslogs](/deployment/syslogs) | Configure syslog forwarding to any compatible provider |
| [Datadog](/integrations/datadog) | Send logs and metrics to Datadog for operational visibility |
| [Papertrail](/integrations/papertrail) | Forward logs to Papertrail via syslog |
| [Loggly](/integrations/loggly) | Forward logs to Loggly via syslog |
| [Mezmo](/integrations/logdna) | Forward logs to Mezmo (formerly LogDNA) via syslog |

## Private Registries

Pull container images from authenticated private registries.

| Integration | Description |
|-------------|-------------|
| [Private Registries](/integrations/private-registries) | Configure Docker credentials for private image repositories |

## See Also

- [Console Integrations](/console/integrations)
- [CI/CD Integrations](/integrations/cicd)
- [Logging Integrations](/integrations/logging)
- [Syslogs](/deployment/syslogs)
- [Environment](/application/environment)
