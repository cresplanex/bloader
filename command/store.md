---
title: "Store Command"
parent: "Command"
nav_order: 3
---

# Store Commands

Used for simple handling of internal databases from commands. This command is valid only for the Master node.

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