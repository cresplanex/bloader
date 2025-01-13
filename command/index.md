---
title: Command
nav_order: 4
---

# Available Commands 📤

#### Common Options

Bloader explores configuration files in the following order by default: `./bloader`, `$HOME/bloader`, `/etc/bloader`. It searches for `config.*` files (supported formats align with Viper's supported formats). To override this behavior, use the `-c` (or `--config`) option:

```bash
bloader {command} -c $HOME/custom/bloader/config.yaml
```

You can also display help for any command using the `-h` (or `--help`) option.

---

### Common Commands

#### Display Help
```bash
bloader help
```

#### Generate Completion Scripts
Generate shell completion scripts for supported shells (`bash`, `zsh`, `fish`, `powershell`):
```bash
bloader completion bash
```

#### View Current Configuration
Display the current configuration after file and environment variable overrides:
```bash
bloader config
```

---

### Encryption Commands

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

---

### Master-Only Commands

#### Authentication (alias: `auth`)
Manage authentication for OAuth or JWT-based authentication methods. Tokens are stored in the internal database.

- **Login or Refresh Authentication**:
  ```bash
  bloader auth login
  bloader auth refresh
  ```
  Use `-i` (or `--id`) to specify an `authID`. If omitted, the default `auth` is used.

---

#### Store Commands (alias: `st`)
Operate on the internal database. These commands affect only the database of the current environment.

##### List Buckets (alias: `ls`)
List all buckets:
```bash
bloader store list
```

##### Manage Objects (alias: `obj`)
Operate on objects within a specified bucket using `-b` (or `--bucket`).

- **Delete an Object**:
  ```bash
  bloader store object delete --bucket $BUCKET_ID objectKey
  ```

- **Retrieve an Object**:
  ```bash
  bloader store object get --bucket $BUCKET_ID objectKey
  ```

- **List All Keys in a Bucket**:
  ```bash
  bloader store object list --bucket $BUCKET_ID
  ```

- **Store a String Object**:
  ```bash
  bloader store object put --bucket $BUCKET_ID objectKey objectValue
  ```

For `get` and `put`, use `-e` (or `--encrypt`) to specify encryption settings.

##### Clear All Data
```bash
bloader store clear
```

---

#### Output Commands (alias: `out`)

##### Clear Outputs
Clear specific outputs using `-i` (or `--id`), or clear all outputs using `-A` (or `--all`):
```bash
bloader output clear -A
```

---

#### Run Load Test
Run load tests using loader files.

- **Interactive File Input**:
  ```bash
  bloader run
  ```
  You'll be prompted for the file path:
  ```bash
  ✔ Enter the file to run the load test:
  ```

- **Specify File and Data Inline**:
  ```bash
  bloader run -f main.yaml -d IntData=10:i -d StrData=test:s
  ```

---

### Slave-Only Commands

#### Slave Commands (alias: `sl`)

##### Run Slave Server
Start the gRPC server for the slave:
```bash
bloader slave run -c bloader/slave_config.yaml
```

---

These commands provide a flexible and comprehensive way to manage and execute load testing tasks with Bloader. 



