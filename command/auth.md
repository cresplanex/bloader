---
title: "Authentication Command"
parent: "Command"
nav_order: 2
---

# Authentication Commands

Use when you want to do something authentication related. This command is valid only for the Master node.

#### Authentication (alias: `auth`)

Manage authentication for OAuth or JWT-based authentication methods. Tokens are stored in the internal database.

- **Login**:
  ```bash
  bloader auth login
  ```
  Use `-i` (or `--id`) to specify an `authID`. If omitted, the default `auth` is used.

- **Refresh Authentication**:
  ```bash
  bloader auth refresh
  ```
  Use `-i` (or `--id`) to specify an `authID`. If omitted, the default `auth` is used.