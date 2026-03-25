---
title: "Example Apps"
description: "Sample applications demonstrating how to deploy common frameworks and patterns on Convox."
---

# Example Apps

Convox maintains a collection of example applications in the [convox-examples](https://github.com/convox-examples/) GitHub organization. Each example includes a `convox.yml` manifest and demonstrates the steps needed to go from a base application to running on Convox.

## Web Frameworks

| Example | Language | Repository |
|---------|----------|------------|
| Ruby on Rails | Ruby | [convox-examples/rails](https://github.com/convox-examples/rails) |
| Express | Node.js | [convox-examples/nodejs](https://github.com/convox-examples/nodejs) |
| Django | Python | [convox-examples/django](https://github.com/convox-examples/django) |
| PHP | PHP | [convox-examples/php](https://github.com/convox-examples/php) |
| .NET Core | C# | [convox-examples/dotnet-core](https://github.com/convox-examples/dotnet-core) |
| Deno | TypeScript | [convox-examples/deno](https://github.com/convox-examples/deno) |
| httpd | Static | [convox-examples/httpd](https://github.com/convox-examples/httpd) |

## Frontend Frameworks

| Example | Framework | Repository |
|---------|-----------|------------|
| Next.js | React | [convox-examples/nextjs](https://github.com/convox-examples/nextjs) |
| Angular | Angular | [convox-examples/Angular](https://github.com/convox-examples/Angular) |
| Svelte | SvelteKit | [convox-examples/Svelte](https://github.com/convox-examples/Svelte) |
| Gatsby | React | [convox-examples/gatsby](https://github.com/convox-examples/gatsby) |

## CMS and Applications

| Example | Description | Repository |
|---------|-------------|------------|
| WordPress | WordPress with Convox | [convox-examples/Wordpress](https://github.com/convox-examples/Wordpress) |
| Drupal 11 | Drupal 11 | [convox-examples/drupal-11](https://github.com/convox-examples/drupal-11) |
| Drupal 10 | Drupal 10 | [convox-examples/drupal-10](https://github.com/convox-examples/drupal-10) |
| n8n | Self-hosted workflow automation | [convox-examples/n8n](https://github.com/convox-examples/n8n) |
| Hasura | GraphQL engine | [convox-examples/hasura](https://github.com/convox-examples/hasura) |

## Architecture Patterns

| Example | Description | Repository |
|---------|-------------|------------|
| Frontend + Backend | Multi-service Node.js app | [convox-examples/node-frontend-backend](https://github.com/convox-examples/node-frontend-backend) |
| Multi-stage Build | Docker multi-stage build | [convox-examples/multi-stage-build](https://github.com/convox-examples/multi-stage-build) |
| gRPC | gRPC service in Go | [convox-examples/grpc](https://github.com/convox-examples/grpc) |
| Internal Services | Internal-only services | [convox-examples/internal-services](https://github.com/convox-examples/internal-services) |
| EFS Volumes | Persistent storage with EFS | [convox-examples/efs-volumes](https://github.com/convox-examples/efs-volumes) |
| Datadog | Datadog Agent sidecar | [convox-examples/datadog](https://github.com/convox-examples/datadog) |

## AI / Machine Learning

| Example | Description | Repository |
|---------|-------------|------------|
| LLM GPU API | GPU-accelerated LLM API | [convox-examples/llm-gpu-api](https://github.com/convox-examples/llm-gpu-api) |
| Llama 2 Chatbot | Llama 2 chatbot | [convox-examples/llama2-convox-chatbot](https://github.com/convox-examples/llama2-convox-chatbot) |

## Getting Started

To deploy any example, clone the repository and deploy to your Rack:

```bash
$ git clone https://github.com/convox-examples/rails.git
$ cd rails
$ convox apps create rails
$ convox deploy
```

Contributions of additional examples are welcome. Open a pull request in the relevant repository or request a specific example through [support](/help/support).

## See Also

- [Creating an Application](/deployment/creating-an-application)
- [convox.yml](/application/convox-yml)
- [Builds](/deployment/builds)
- [Rails Example](/example-apps/rails)
- [Django Example](/example-apps/django)
- [Node.js Example](/example-apps/nodejs)
