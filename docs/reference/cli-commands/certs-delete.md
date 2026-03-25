---
title: "certs delete"
description: "Delete a certificate from the Rack."
---

# certs delete

Delete a certificate from the Rack. The certificate must not be in use by any App. If the certificate is currently associated with a Service endpoint, remove the association first using `ssl update`.

## Syntax

```bash
$ convox certs delete <cert>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox certs delete cert-ab12cd
Deleting certificate cert-ab12cd... OK
```

## See Also

- [certs](/reference/cli-commands/certs)
- [certs generate](/reference/cli-commands/certs-generate)
- [certs import](/reference/cli-commands/certs-import)
- [SSL](/deployment/ssl)
