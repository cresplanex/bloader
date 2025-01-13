---
title: "Encryption Command"
parent: "Command"
nav_order: 1
---

# Encryption Commands

Used for simple encryption/decryption on command. This command can be used for both master and slave.

#### Encrypt (alias: `enc`)

Perform encryption or decryption using encryption methods defined in the configuration. Specify the `encryptID` with `-i` (or `--id`).

- **Encrypt Data**:
  ```bash
  bloader encrypt test -i encryptStaticCBC
  ```
  Output:
  ```bash
  Encrypted text: bltBHtsdAUUoHxPxRzyTZgN91TPmHP5rSKnIaAxIeM0=
  ```

- **Decrypt Data**:
  ```bash
  bloader encrypt bltBHtsdAUUoHxPxRzyTZgN91TPmHP5rSKnIaAxIeM0= -i encryptStaticCBC -d
  ```
  Output:
  ```bash
  Decrypted text: test
  ```