---
title: "test"
description: "Run tests defined in the convox.yml manifest."
---

# test

Run the tests defined by the `test:` directive on each Service in `convox.yml`. A development Build is created (or an existing Release is used), and the test command for each Service is executed in sequence. Each command runs in a new Process of its Service. The CLI stops at the first command that exits non-zero, prints `ERROR: exit <code>`, and exits with status 1. When stdin is not a terminal, for example in a CI job, output follows the same rules as [exec](/reference/cli-commands/exec#running-without-a-terminal).

## Syntax

```bash
$ convox test [dir]
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--description` | `-d` | Build description |
| `--rack` | `-r` | Rack name |
| `--release` | | Use existing Release to run tests |
| `--timeout` | `-t` | Timeout in seconds |

## Example Usage

```bash
$ convox test -a myapp
Packaging source... OK
Uploading source... OK
Starting build... OK
Building: .
...
Running bin/test on web
12 examples, 0 failures
Running bin/test on worker
5 examples, 0 failures
```

## See Also

- [build](/reference/cli-commands/build)
- [deploy](/reference/cli-commands/deploy)
- [start](/reference/cli-commands/start)
- [convox.yml](/application/convox-yml)
