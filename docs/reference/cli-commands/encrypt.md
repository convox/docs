---
title: "encrypt"
description: "Encrypt data using a symmetric key for secure storage or transmission."
---

# encrypt

Encrypt arbitrary data using a symmetric key. Useful for securing sensitive values such as API keys, database passwords, or tokens before storing them in version control, sharing with team members, or embedding in deployment scripts.

## Syntax

```bash
$ convox encrypt --key <key> --data <data>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--key` | | Encryption key |
| `--data` | | Data to encrypt |

## Example Usage

```bash
$ convox encrypt --key mySecretKey123 --data "s3cr3t-api-token"
aGVsbG8gd29ybGQ...
```

## See Also

- [decrypt](/reference/cli-commands/decrypt)
- [env set](/reference/cli-commands/env-set)
- [Environment](/application/environment)
