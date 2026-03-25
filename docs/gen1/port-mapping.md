---
title: "Port Mapping"
description: "Gen 1 (End of Life): How to define external, internal, and UDP port mappings for Gen 1 Convox application processes."
---

# Port Mapping

> **This page documents Generation 1, which has reached End of Life.** Gen 1 apps use `docker-compose.yml`. For current documentation, see [Port](/application/port).

You can define the ports on which your processes will listen in the manifest for your application.

## External Ports

External ports are open to the Internet. You define an external port using a port pair in the `ports:` section of your `docker-compose.yml`:

```yaml
ports:
  - 80:5000
```

This example configuration would listen to port `80` on an Internet-accessible load balancer and forward connections to port `5000` on the Process.

## Internal Ports

Internal ports are only accessible to other apps and services on the same Rack. You define an internal port using a single port in the `ports:` section of your `docker-compose.yml`:

```yaml
ports:
  - 5000
```

This example configuration would listen to port `5000` on an internal-only load balancer and forward connections to port `5000` on the Process.

If you want to make all of an application's ports internal, regardless of port definition, you can set the Internal app parameter.

```bash
$ convox apps params set Internal=Yes
```

## TCP/UDP Protocols

By default, exposed ports are for the TCP protocol, but you can explicitly define the protocol to be used:

```yaml
ports:
  - 80:5000
  - 443:5000/tcp
  - 514:6000/udp
```

This example configuration would listen to ports `80` and `443` **on an Internet-accessible load balancer** and forward TCP connections to port `5000` on the Process, while it listens to port `514` **on the host** and forwards UDP packets to port `6000` on the Process. This difference between where the port is exposed is due to Elastic Load Balancers only supporting TCP connections. As a result, UDP ports are always internal-only. Additionally, since the host isn't using a random port for communicating with a load balancer, ensuring there are no port conflicts between processes is an exercise left up to the user.

## See Also

- [Load Balancers](/gen1/load-balancers)
- [Docker Compose File](/gen1/docker-compose-file)
