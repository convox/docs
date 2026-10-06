---
title: "rack resources link"
description: "Link a Rack Resource to an App."
---

# rack resources link

Link a `syslog` Rack Resource to an App so the App's logs are sent to the syslog destination. Other Resource types cannot be linked.

## Syntax

```bash
$ convox rack resources link <resource>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

## Example Usage

```bash
$ convox rack resources link syslog-3785 -a myapp --wait
Linking to myapp... OK
```

## See Also

- [rack resources unlink](/reference/cli-commands/rack-resources-unlink)
- [rack resources info](/reference/cli-commands/rack-resources-info)
- [rack resources](/reference/cli-commands/rack-resources)
- [Rack and External Resources](/management/rack-resources)
