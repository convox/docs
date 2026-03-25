---
title: "decrypt"
description: "Decrypt data using a symmetric key."
---

# decrypt

Decrypt data that was previously encrypted using the `encrypt` command. Provide the same key that was used for encryption.

## Syntax

```bash
$ convox decrypt --key <key> --data <data>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--key` | | Decryption key (must match the key used to encrypt) |
| `--data` | | Encrypted data to decrypt |

## Example Usage

```bash
$ convox decrypt --key mySecretKey123 --data "aGVsbG8gd29ybGQ..."
s3cr3t-api-token
```

## See Also

- [encrypt](/reference/cli-commands/encrypt)
- [env get](/reference/cli-commands/env-get)
- [Environment](/application/environment)
