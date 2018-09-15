---
title: "Advanced Build Options"
---

By default, `convox build` processes run in a Rack just like other application processes. This is designed to handle a normal number of builds at a minimal cost. Builds use spare capacity on a Rack instance, and will autoscale the Rack up and then back down if any addition build capacity is needed.

Convox offers advanced options that can improve the speed, throughput and security of builds.

### Dedicated Build Instance

If you'd like to segregate builds from apps or just need better build performance, you can configure a dedicated build instance that will give more CPU, memory and caching to builds.

You can modify the instance type of your build instance by setting the [`BuildInstance`](/docs/rack-parameters/#buildinstance) Rack parameter:

```
$ convox rack params set BuildInstance=c4.large
```

This will guarantee that:

* Builds get dedicated CPU and memory resources
* Builds do not affect app performance
* Builds do not affect app instance security
* A single Docker image cache is shared among every build

This option offers the best performance for very CPU or memory intensive builds. It also helps with compliance and system auditing requirements by moving all build operations off of application servers.

You will incur the cost of running a dedicated build instance 24/7.
