---
title: "test"
description: "Run tests defined in the convox.yml manifest."
---

# test

Run the tests defined by the `test:` directive on each Service in `convox.yml`. A development Build is created (or an existing Release is used), and the test command for each Service is executed in sequence.

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
Running tests for web...
  ✓ 12 tests passed
  ✗ 0 tests failed
Running tests for worker...
  ✓ 5 tests passed
  ✗ 0 tests failed
OK
```

## See Also

- [build](/reference/cli-commands/build)
- [deploy](/reference/cli-commands/deploy)
- [start](/reference/cli-commands/start)
- [convox.yml](/application/convox-yml)
