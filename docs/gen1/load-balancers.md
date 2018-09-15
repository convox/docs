---
title: "Load Balancers"
---

Once you have containers running, the next step is to allow them to be accessed from the Internet. Convox automatically sets up and configures load balancers appropriately to route traffic to your containers.

### Configuration

Load balancers will be automatically created for any ports listed in your `docker-compose.yml`.

```
web:
  build: .
  command: bin/web
  ports:
    - 80:5000
worker:
  build: .
  command: bin/worker
```

In this example, Convox will create a load balancer in front of the `web` process. This load balancer will accept traffic from the Internet on port 80 and forward it to the `web` containers on port `5000`.

<div class="block-callout block-show-callout type-warning" markdown="1">
Convox will only create a load balancer for ports in your `docker-compose.yml` file, not in your `Dockerfile`.
</div>

### Balancer Hostname

You can find the load balancer hostname(s) for your application using `convox apps info`:

    $ convox apps info
    Name       docs
    Status     running
    Release    RHUFNNNVEAP
    Processes  web
    Endpoints  docs-web-R72RMTP-326048479.us-east-1.elb.amazonaws.com:80 (web)

### Advanced Options

#### Internal Load Balancers

You can create a load balancer that is only accessible inside your Rack by specifying a single port:

```
web:
  ports:
    - "5000"
```

<div class="block-callout block-show-callout type-info" markdown="1">
**Note: Convox creates only one load balancer per service.** If you specify both internal and external ports, only an internal load balancer will be created.
This is due to the fact that while an ELB can have listeners on multiple ports, an ELB itself can only be either internal or external.
</div>

#### Balancer Protocol

You can specify one of four protocol types for a load balancer port in your `docker-compose.yml`:

```
web:
  labels:
    - convox.port.443.protocol=https
  ports:
    - "443:5000"
```

<table>
  <tr>
    <th>Protocol</th>
    <th>Notes</th>
  </tr>
  <tr>
    <td><code>http</code></td>
    <td>Unencrypted HTTP <em><strong>(includes common HTTP headers but does not support websockets)</strong></em></td>
  </tr>
  <tr>
    <td><code>https</code></td>
    <td>Encrypted HTTP <em><strong>(includes common HTTP headers but does not support websockets)</strong></em></td>
  </tr>
  <tr>
    <td><code>tcp</code></td>
    <td>Unencrypted TCP <em><strong>(arbitrary TCP including websockets, no HTTP header injection)</strong></em></td>
  </tr>
  <tr>
    <td><code>tls</code></td>
    <td>Encrypted TCP <em><strong>(arbitrary TCP including websockets, no HTTP header injection)</strong></em></td>
  </tr>
</table>

If no protocol label is specified, the default of `tcp` will be used.

#### Health Check Options

By default Convox will set up a `tcp` health check to your application. For more information, see [Health Checks](/docs/gen1/health-checks).

#### End-to-end encryption

By default, HTTPS/TLS is terminated at the load balancer, and the resulting data is transmitted unencrypted to your application. This is OK, because traffic between your load balancer and your application happens entirely on your Rack's internal network. However, for extra security you can encrypt the traffic between your load balancer and application by setting the `convox.port.<port>.secure` label.

```
web:
  labels:
    - convox.port.443.secure=true
    - convox.port.443.protocol=https
  ports:
    - "443:5001"
```

When you use this option you will need to terminate HTTPS or TLS directly inside your application or with a reverse proxy like nginx or haproxy.

#### PROXY protocol

When using the `tcp` or `tls` protocols, standard proxy HTTP headers like `X-Forwarded-For` are not injected. You can get access to information about the remote endpoint using the [PROXY protocol](http://www.haproxy.org/download/1.5/doc/proxy-protocol.txt). Once you configure your application to accept this extra header you can configure your load balancer to send it in your `docker-compose.yml`:

```
web:
  labels:
    - convox.port.443.protocol=tls
    - convox.port.443.proxy=true
  ports:
    - "443:5000"
```

#### Limited Application Access

For security reasons, access to an application might need to be limited. To achieve this, an existing security group can be applied to an application's load balancer. For example, within said security group, access can be granted only to an office VPN.

This is done via an [application parameter](/docs/gen1/app-parameters/#securitygroup) with a known security group ID:

```
convox apps params --app <name> set SecurityGroup=sg-123456
```

Validate the setting has been applied by running:

```
convox apps params --app <name>
```

For further reading on security groups, check out the AWS [documentation](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html) and [CLI reference](http://docs.aws.amazon.com/cli/latest/userguide/cli-ec2-sg.html).

## See also

- [Port mapping](/docs/gen1/port-mapping/)
