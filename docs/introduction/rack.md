---
title: "Convox Rack"
order: 300
---

Convox Rack is an [open source](https://github.com/convox/rack) deployment platform that is installed into your AWS account. A Rack creates and manages all of the underlying infrastructure needed to run and monitor your applications. A Rack is the unit of network isolation -- applications and services on a Rack can only communicate with other applications and services on the same Rack.

![](/assets/images/docs/what-is-a-rack/convox-rack-diagram.jpg)

### Dynamic Runtime

A Rack will start multiple identical servers on which it will containerize and run your applications. By using a homogenous runtime we can treat each individual server as disposable and recover easily from common failure scenarios.

### Private Network

Each Rack creates a private network inside which it runs its servers and services. All access from the internet comes through load balancers which are specifically configured to route traffic to your containers.
