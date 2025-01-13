---
title: Configuration Override
parent: Configuration
nav_order: 2
---
# Override Settings

The **override feature** enables dynamic configuration changes for different environments, making it more flexible and environment-specific. Items can be overridden using **file** or **static** methods, with priority given to the last-defined override.

## Key Features 🔑
- Environment-specific overrides via `enabled_env`.
- Overrides applied later in the configuration file take precedence.
- Supported types: **file** and **static**.

---

## Override Types 🛠️

### Static
Static overrides allow fine-grained control over configuration, such as targeting specific array elements.

```yaml
- type: "static"
  key: "auth[0].oauth2.token_url"
  value: "http://localhost:8080/oauth2/token"
  enabled_env:
    - "local"
```

**Key Points:**
- Use `key` to specify the target configuration.
- Apply overrides for specific environments using `enabled_env`.
- Flexible and suitable for small, specific changes.

---

### File 📂
File-based overrides are divided into two patterns: **full file overrides** and **partial file overrides**. 

#### 1. Full File Override
In this method, the entire configuration is replaced with the contents of the specified file. 

```yaml
- type: "file"
  file_type: "yaml"
  path: "bloader/local_override.yaml"
  partial: false # Default
  # If enabled_env is not set, it will be enabled for all environments.
  enabled_env:
    - "local"
```

**Example File (`bloader/local_override.yaml`):**
```yaml
server:
  port: 8080
clock:
  fake: 
    enabled: false
  time: "2021-01-01T00:00:00Z"
  format: "2006-01-02T15:04:05Z"
```

**Key Points:**
- The file must follow the **Bloader config format**.
- Ideal for replacing entire configurations in bulk.
- Automatically applies for all environments if `enabled_env` is not specified.

---

#### 2. Partial File Override 🧩
This method updates specific configuration items without requiring the file to follow the full config format.

```yaml
- type: "file"
  file_type: "yaml"
  path: "bloader/static_encrypt.yaml"
  partial: true
  vars:
    - key: encrypts[0].key
      value: "encrypt_key"
```

**Example File (`bloader/static_encrypt.yaml`):**
```yaml
encrypt_key: "y8sF2gVz4MwqYLn3RtJxNk7P"
```

**Key Points:**
- Focuses on specific items, similar to `static`.
- Great for small adjustments without needing full config compliance.
- Scalable for environments with diverse override requirements.

---

Use override settings to create flexible and efficient configurations tailored to your application’s needs. 🎉
