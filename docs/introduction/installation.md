---
title: "CLI Installation"
description: "Install the Convox CLI on macOS, Linux, or Windows"
---

# CLI Installation

The `convox` command line tool makes building, configuring, scaling, and securing your apps straightforward.

Here are some of the highlights:

* `convox start` - A single command to start your development environment
* `convox deploy` - A single command to deploy your application to a Rack
* `convox rack update` - A single command to deliver API and infrastructure improvements to a Rack

You can install it via the command line:

## macOS

### x86_64 / amd64

```bash
$ curl -L https://download.convox.com/cli/darwin/convox -o /tmp/convox
$ sudo mv /tmp/convox /usr/local/bin/convox
$ sudo chmod 755 /usr/local/bin/convox
```

### arm64

```bash
$ curl -L https://download.convox.com/cli/darwin/convox-arm64 -o /tmp/convox
$ sudo mv /tmp/convox /usr/local/bin/convox
$ sudo chmod 755 /usr/local/bin/convox
```

## Linux

### x86_64 / amd64

```bash
$ curl -L https://download.convox.com/cli/linux/convox -o /tmp/convox
$ sudo mv /tmp/convox /usr/local/bin/convox
$ sudo chmod 755 /usr/local/bin/convox
```

### arm64

```bash
$ curl -L https://download.convox.com/cli/linux/convox-arm64 -o /tmp/convox
$ sudo mv /tmp/convox /usr/local/bin/convox
$ sudo chmod 755 /usr/local/bin/convox
```

## Windows

```bash
$ curl -L https://download.convox.com/cli/windows/convox.exe -O
```

## Next steps

### Logging in to Console

After installing Convox, you'll need to `convox login`:

```bash
$ convox login console.convox.com
Password: <your Console API key>
Logged in successfully.
```

## Updating the CLI

To update the CLI you can run `convox update`:

```bash
$ convox update
Updating convox: OK
```

## See Also

- [Getting Started](/introduction/getting-started)
- [Reference: CLI Commands](/reference/cli-commands)
- [Core Concepts: Console](/core-concepts/console)
- [Frequently Asked Questions](/introduction/faq)
