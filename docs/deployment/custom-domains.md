---
title: "Custom Domains"
---

## Definition

You can specify that your service should listen on a custom domain:

```yaml
services:
  web:
    domain: myapp.example.org
    port: 3000
```

You can also specify multiple domains using this syntax:

```yaml
services:
  web:
    domain:
      - myapp.example.org
      - "*.example.org"
    port: 3000
```

<div class="block-callout block-show-callout type-warning" markdown="1">
  Using a custom domain requires a valid SSL certificate for the domains being specified.
  <p>If a single certificate does not already exist in the same regions as the rack is installed and that matches all the domains you specify for your service, one will be automatically created. The DNS owner will receive a validation email with a link that needs to be clicked for this process to complete.  This will happen the first time you deploy your service with your configuration</p>
</div>

<div class="block-callout block-show-callout type-info" markdown="1">
You can [pre-generate your SSL certificate](/deployment/ssl#pre-generate-your-certificate) ahead of deploy time if you wish.
</div>

#### Dynamic Configuration

You can use environment interpolation so that you don't have to hardcode the hostname in your `convox.yml`:

```yaml
services:
  web:
    domain: ${HOST}
    port: 3000
```

## Configuring DNS

Run `convox rack` and find the `Router` value. Configure your custom domain as a `CNAME` to this domain.

#### Example

<table>
  <tr>
    <th>Name</th>
    <td><code>docs.convox.com</code></td>
  </tr>
  <tr>
    <th>Type</th>
    <td><code>CNAME</code></td>
  </tr>
  <tr>
    <th>Value</th>
    <td><code>produ-Route-1ABCDEFGHIJK-01234569.us-east-1.elb.amazonaws.com</code></td>
  </tr>
  <tr>
    <th>TTL</th>
    <td><code>60</code></td>
  </tr>
</table>

<div class="block-callout block-show-callout type-info" markdown="1">
To set up DNS for a root domain you should use the an **Alias** type with Route 53 or the equivalent with your DNS provider.
</div>
