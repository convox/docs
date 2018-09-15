---
title: "Custom Domains"
---

You can easily map a custom domain to a Convox application by creating a `CNAME` to your load balancer hostname.

## Balancer Hostname

You can find the load balancer hostname(s) for your application using `convox services`:

    $ convox services
    SERVICE  DOMAIN                                                  PORTS
    web      docs-web-R72RMTP-326048479.us-east-1.elb.amazonaws.com  80

## Configuring DNS

Create an appropriate DNS entry to map your desired custom domain to your Convox app. In the example above one might create the following DNS entry:

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
    <td><code>docs-web-R72RMTP-326048479.us-east-1.elb.amazonaws.com</code></td>
  </tr>
  <tr>
    <th>TTL</th>
    <td><code>60</code></td>
  </tr>
</table>

&nbsp;

<div class="block-callout block-show-callout type-info" markdown="1">
To set up DNS for a root domain you should use the `ALIAS` type with Route 53 or the equivalent with your DNS provider.
</div>
