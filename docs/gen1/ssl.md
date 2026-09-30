---
title: "SSL"
description: "Gen 1 (End of Life): How to configure TLS/SSL certificates for Gen 1 Convox applications, including certificate management and load balancer setup."
---

# SSL

> **This page documents Generation 1, which has reached End of Life.** Gen 1 apps use `docker-compose.yml`. For current documentation, see [SSL](/deployment/ssl).

You can easily secure traffic to your application using TLS (SSL).

## Add a Secure Port to Your Manifest

Edit your app's `docker-compose.yml` file to create a port mapping for your secure traffic. For most web applications this will be port 443, the standard for HTTPS.

You'll also need to set the protocol for the port using the `convox.port.<port>.protocol` label. Use `https` as the value if you want to get HTTP headers and don't need to support websockets. Otherwise use `tls`. For example:

```yaml
web:
  labels:
    - convox.port.443.protocol=https
  ports:
    - 80:3000
    - 443:3000
```

When you're done editing, redeploy your application.

```bash
$ convox deploy
```

Your app is now configured to serve encrypted traffic with a self-signed certificate on port 443. To use a real certificate, you will need to acquire an SSL Certificate and apply it to your SSL endpoint. See the following sections for more information.

## Acquire an SSL Certificate

### Generate a Certificate

You can request an SSL certificate for any domain you control using `convox certs generate`:

```bash
$ convox certs generate foo.example.org
Generating certificate... OK, acm-c28a4f6e0b37
```

AWS Certificate Manager (ACM) issues the certificate once the domain is validated. `convox certs` does not list it until then, so keep the id the command prints. Each run requests a new certificate.

#### Wildcard certificates

You can generate a wildcard certificate with `*`, e.g. `convox certs generate "*.example.com"`. However, note that the wildcard only covers that level of the domain and not the bare domain. For instance, `*.example.com` will cover `www.example.com`, `mail.example.com` and so on, but not `example.com` itself.

### Purchase a Certificate

You can also purchase an SSL certificate from most registrars and DNS providers. Convox is a fan of [Gandi](https://www.gandi.net/ssl).

Import your certificate and private key using `convox certs import`:

```bash
$ convox certs import example.org.pub example.org.key
Importing certificate... OK, acm-5d0c7e2b91f4
```

### Apply the Certificate

You can then apply a certificate to your load balancer with `convox ssl update`:

```bash
$ convox ssl update web:443 acm-5d0c7e2b91f4
Updating certificate... OK
```

### Inspect SSL Configuration

You can use the Convox CLI to view SSL configuration for an app.

```bash
$ convox ssl
ENDPOINT  CERTIFICATE       DOMAIN       EXPIRES
web:443   acm-5d0c7e2b91f4  example.org  2 months from now
```

## Managing Certificates

The Convox CLI includes commands that let you list, update, and remove SSL certificates.

### Listing Certificates

You can see the IAM server certificates in your AWS account and the issued ACM certificates in the Rack's region with `convox certs`:

```bash
$ convox certs
ID                                DOMAIN                 EXPIRES
acm-5d0c7e2b91f4                  example.org            2 months from now
acm-c28a4f6e0b37                  foo.example.org        1 year from now
cert-production-1758802512-04821  *.*.elb.amazonaws.com  5 days ago
```

Certificates imported with `convox certs import` or generated with `convox certs generate` have ids like `acm-c28a4f6e0b37`. The self-signed certificate the Rack uploads for `https` and `tls` ports is an IAM server certificate named `cert-<rack>-<unix time>-<number>`, and other IAM server certificates are listed by name.

### Updating Your SSL Certificate

When it's time to update your SSL certificate, import your new certificate and use `convox ssl update` again:

```bash
$ convox certs import example.org.pub example.org.key
Importing certificate... OK, acm-9b1e60d2c8f7

$ convox ssl update web:443 acm-9b1e60d2c8f7
Updating certificate... OK
```

### Removing Old Certificates

You can remove old certificates that you are no longer using. Pass the full id.

```bash
$ convox certs delete acm-5d0c7e2b91f4
Deleting certificate acm-5d0c7e2b91f4... OK
```

## See Also

- [SSL (Gen 2)](/deployment/ssl)
- [Custom Domains](/gen1/custom-domains)
- [Load Balancers](/gen1/load-balancers)
