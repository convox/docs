---
title: "login"
description: "Authenticate the Convox CLI with a Rack or Console."
---

# login

Authenticate the Convox CLI with a Rack or Console. If no hostname is specified, the CLI defaults to `console.convox.com`. You can also log in directly to a self-hosted Rack by providing its hostname and a password.

## Syntax

```bash
$ convox login [hostname]
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--password` | `-p` | Password |

## Example Usage

```bash
$ convox login
Authenticating with console.convox.com... OK
```

```bash
$ convox login rack.example.com -p mysecretpassword
Authenticating with rack.example.com... OK
```

## See Also

- [switch](/reference/cli-commands/switch)
- [racks](/reference/cli-commands/racks)
- [Getting Started](/introduction/getting-started)
- [CLI Installation](/introduction/installation)
