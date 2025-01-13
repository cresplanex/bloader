---
title: Configuration Properties
parent: Configuration
nav_order: 1
---

# Configuration Props

### Visual Guide 🖼️
- **✅ Required**: Configuration must be defined for the feature to function.
- **❌ Not Required**: Optional and can be omitted without impacting functionality.

## General Props 🛠️

| **Item**            | **Description**                               | **Required**           | **Type**    |
|:--------------------|:----------------------------------------------|:----------------------:|:-----------:|
| `version`           | Specify the version of bloader. Currently only `v1` is valid.            | ✅                      | `string`    |
| `type`              | Master or slave configuration role            | ✅                      | `string`    |
| `env`               | Environment identifier (user-defined)         | ✅                      | `string`    |
| `loader`            | Loader settings for workload definitions      | ✅ (master) ❌ (slave)   | `object`    |
| `loader.base_path`  | Base path for the loader                      | ✅ (master) ❌ (slave)   | `string`    |

## Targets 🎯

| **Item**                 | **Description**                             | **Required**           | **Type**   |
|:-------------------------|:--------------------------------------------|:----------------------:|:-----------|
| `targets`                | Measurement targets                         | ✅ (master) ❌ (slave) | `[]object` |
| `targets[].id`           | Unique ID within the target array           | ✅                     | `string`   |
| `targets[].type`         | Type of measurement (e.g., `http`)          | ✅                     | `string`   |
| `targets[].values`       | Configuration for specific target types     | ✅                     | `[]object` |
| `targets[].values[].env` | Active environment                          | ✅                     | `string`   |
| `targets[].values[].url` | Target URL (when `type=http`)               | ✅                     | `string`   |

## Outputs 📤

| **Item**                       | **Description**                        | **Required**           | **Type**    |
|:-------------------------------|:---------------------------------------|:----------------------:|:-----------:|
| `outputs`                      | Output settings                        | ✅ (master) ❌ (slave) | `[]object`  |
| `outputs[].id`                 | Unique ID within the output array      | ✅                     | `string`    |
| `outputs[].values`             | Output-specific settings               | ✅                     | `[]object`  |
| `outputs[].values[].env`       | Active environment                     | ✅                     | `string`    |
| `outputs[].values[].type`      | Output type (e.g., `local`)            | ✅                     | `string`    |
| `outputs[].values[].format`    | Output format (e.g., `csv`)            | ✅                     | `string`    |
| `outputs[].values[].base_path` | Base path for output files             | ✅                     | `string`    |

## Store 🗄️

| **Item**             | **Description**                  | **Required**            | **Type**    |
|:---------------------|:---------------------------------|:-----------------------:|:-----------:|
| `store`              | Internal database settings       | ✅ (master) ❌ (slave)  | `object`    |
| `store.file`         | Database file settings           | ✅                      | `[]object`  |
| `store.file[].env`   | Active environment               | ✅                      | `string`    |
| `store.file[].path`  | Path to the database file        | ✅                      | `string`    |
| `store.buckets`      | List of predefined bucket names  | ✅                      | `[]string`  |

## Encryption 🔐

| **Item**                              | **Description**                                            | **Required**                 | **Type**    |
|:--------------------------------------|:-----------------------------------------------------------|:----------------------------:|:-----------:|
| `encrypts`                            | Encryption settings                                        | ✅                           | `[]object`  |
| `encrypts[].id`                       | Unique ID for the encryption                               | ✅                           | `string`    |
| `encrypts[].type`                     | Encryption type                                            | ✅                           | `string`    |
| `encrypts[].key`                      | Encryption key (recommended to override via external file) | ✅ (`type=static*`)       | `string`    |
| `encrypts[].store`                    | Key management settings for dynamic encryption             | ✅ (`type=dynamic*`)      | `object`    |
| `encrypts[].store.bucket_id`          | Bucket ID where the key is stored                          | ✅                           | `string`    |
| `encrypts[].store.key`                | Key identifier within the bucket                           | ✅                           | `string`    |
| `encrypts[].store.encrypt`            | Encryption settings for key management                     | ❌                           | `object`    |
| `encrypts[].store.encrypt.enabled`    | Enable encryption for key management                       | ❌                           | `boolean`   |
| `encrypts[].store.encrypt.encrypt_id` | Encryption ID for key management           

## Authentication 🔑

| **Item**                                            | **Description**                                                      | **Required**                                      | **Type**    |
|:----------------------------------------------------|:---------------------------------------------------------------------|:-------------------------------------------------------:|:------------:|
| `auth`                                              | Authentication settings                                              | ✅ (master) ❌ (slave)                                                   | `[]object`  |
| `auth[].id`                                         | Unique ID for the authentication configuration                       | ✅                                                       | `string`    |
| `auth[].default`                                    | Set as the default authentication configuration                      | ✅                                                       | `boolean`   |
| `auth[].type`                                       | Authentication type (`oauth2`, `apiKey`, `basic`, etc.)              | ✅                                                       | `string`    |
| `auth[].oauth2`                                     | OAuth2 configuration settings                                        | ✅ (`type=oauth2`)                                           | `object`     |
| `auth[].oauth2.grant_type`                          | OAuth2 grant type (`authorization_code`, `client_credentials`, etc.) | ✅                                                       | `string`    |
| `auth[].oauth2.client_id`                           | OAuth2 client ID                                                     | ✅                                                       | `string`    |
| `auth[].oauth2.scope`                               | OAuth2 scope                                                         | ✅                                                       | `[]string`  |
| `auth[].oauth2.client_secret`                       | OAuth2 client secret                                                 | ❌                                                       | `string`    |
| `auth[].oauth2.access_type`                         | OAuth2 access type (`online` or `offline`)                           | ✅ (`grant_type=authorization_code`)                         | `string`     |
| `auth[].oauth2.auth_url`                            | OAuth2 authorization endpoint                                        | ✅ (`grant_type=authorization_code`)                         | `string`     |
| `auth[].oauth2.token_url`                           | OAuth2 token endpoint                                                | ✅ (`grant_type=authorization_code` or `client_credentials`) | `string`     |
| `auth[].oauth2.username`                            | Username for OAuth2 password grant                                   | ✅ (`grant_type=password`)                                   | `string`     |
| `auth[].oauth2.password`                            | Password for OAuth2 password grant                                   | ✅ (`grant_type=password`)                                   | `string`     |
| `auth[].oauth2.credential`                          | Credential storage settings                                          | ✅                                                       | `object`     |
| `auth[].oauth2.credential.store`                    | Credential key storage                                               | ✅                                                       | `object`     |
| `auth[].oauth2.credential.store.bucket_id`          | Bucket ID for credential storage                                     | ✅                                                       | `string`     |
| `auth[].oauth2.credential.store.key`                | Credential key within the bucket                                     | ✅                                                       | `string`     |
| `auth[].oauth2.credential.store.encrypt`            | Encryption settings for credential storage                           | ❌                                                       | `object`     |
| `auth[].oauth2.credential.store.encrypt.enabled`    | Enable encryption for credential storage                             | ❌                                                       | `boolean`    |
| `auth[].oauth2.credential.store.encrypt.encrypt_id` | Encryption ID for credential storage                                 | ✅ (`encrypt.enabled=true`)                                           | `string`     |
| `auth[].api_key`                                    | API key configuration                                                | ✅ (`type=apiKey`)                                           | `object`     |
| `auth[].api_key.header_name`                        | Header name for the API key                                          | ✅                                                       | `string`     |
| `auth[].api_key.key`                                | API key value                                                        | ✅                                                       | `string`     |
| `auth[].basic`                                      | Basic authentication configuration                                   | ✅ (`type=basic`)                                            | `object`      |
| `auth[].basic.username`                             | Username for basic authentication                                    | ✅                                                       | `string`      |
| `auth[].basic.password`                             | Password for basic authentication                                    | ✅                                                       | `string`      |
| `auth[].jwt`                                        | JWT authentication configuration                                     | ✅ (`type=jwt`)                                              | `object`      |
| `auth[].jwt.credential`                             | JWT credential settings                                              | ✅                                                       | `object`      |
| `auth[].jwt.credential.store`                       | Storage settings for JWT credentials                                 | ✅                                                       | `object`      |
| `auth[].jwt.credential.store.bucket_id`             | Bucket ID for JWT credential storage                                 | ✅                                                       | `string`      |
| `auth[].jwt.credential.store.key`                   | JWT credential key within the bucket                                 | ✅                                                       | `string`      |
| `auth[].jwt.credential.store.encrypt`               | Encryption settings for JWT credential storage                       | ❌                                                       | `object`      |
| `auth[].jwt.credential.store.encrypt.enabled`       | Enable encryption for JWT credential storage                         | ❌                                                       | `boolean`     |
| `auth[].jwt.credential.store.encrypt.encrypt_id`    | Encryption ID for JWT credential storage                             | ✅ (`encrypt.enabled=true`)                                            | `string`      |

## Server Settings ⚙️

| **Item**                | **Description**                                                  | **Required**           | **Type**  |
|:------------------------|:-----------------------------------------------------------------|:----------------------:|:---------:|
| `server`                | Server-related configurations                                    | ✅ (master) ❌ (slave) | `object`  |
| `server.port`           | Port for the server                                              | ✅                     | `int`     |
| `server.redirect_port`  | Port for OAuth redirect (defaults to `server.port` if not set)   | ❌                     | `int`     |

## Slave Settings 🤝

| **Item**                               | **Description**                                         | **Required**           | **Type**   |
|:---------------------------------------|:--------------------------------------------------------|:----------------------:|:----------:|
| `slave_setting`                        | Configuration for slave mode                            | ❌ (master) ✅ (slave) | `object`   |
| `slave_setting.port`                   | gRPC server port for the slave                          | ✅                     | `int`      |
| `slave_setting.certificate`            | TLS certificate settings for secure communication       | ❌                     | `object`   |
| `slave_setting.certificate.enabled`    | Enable TLS communication for the slave                  | ❌                     | `boolean`  |
| `slave_setting.certificate.slave_cert` | Path to the TLS certificate for the slave               | ✅                     | `string`   |
| `slave_setting.certificate.slave_key`  | Path to the TLS private key for the slave               | ✅                     | `string`   |

## Logging 📋

| **Item**                    | **Description**                                                   | **Required**  | **Type**    |
|:-----------------------------|:-------------------------------------------------------------------|:--------------:|:------------:|
| `logging`                   | Logging-related configurations                                   | ✅           | `object`   |
| `logging.output`            | Output settings for logs                                         | ✅           | `[]object` |
| `logging.output[].type`     | Type of logging output (`stdout`, `file`, `tcp`)                 | ✅           | `string`   |
| `logging.output[].format`   | Format of log output (`text` or `json`)                          | ✅           | `string`   |
| `logging.output[].enabled_env` | List of environments where the log output is enabled          | ❌           | `[]string` |
| `logging.output[].level`    | Logging level (`debug`, `info`, `warn`, `error`)                | ✅           | `string`   |
| `logging.output[].filename` | File path for log output             | ✅ (logging.output[].type=file) | `string`   |
| `logging.output[].address`  | Address for log output                 | ✅ (logging.output[].type=tcp)  | `string`   |

## Clock Settings ⏰

| **Item**                    | **Description**                                                   | **Required**  | **Type**    |
|:-----------------------------|:-------------------------------------------------------------------|:--------------:|:------------:|
| `clock`                     | Clock-related configurations (currently unused)                  | ✅           | `object`   |
| `clock.fake`                | Fake clock settings for testing                                  | ❌           | `object`   |
| `clock.fake.enabled`        | Enable the fake clock                                            | ❌           | `boolean`  |
| `clock.fake.time`           | Fixed time for the fake clock                                    | ✅           | `string`   |
| `clock.format`              | Format for displaying time (default: `2006-01-02T15:04:05Z`)    | ✅           | `string`   |

## Language 🌐

| **Item**           | **Description**                              | **Required** | **Type**  |
|:--------------------|:----------------------------------------------|:-------------:|:----------:|
| `language`         | Language-related configurations (currently unused) | ✅         | `object` |
| `language.default` | Default language                            | ✅          | `string` |

## Overrides 🔄

| **Item**            | **Description**                                               | **Required** | **Type**    |
|:---------------------|:-------------------------------------------------------------|:-------------:|:------------:|
| `override`          | Override settings                                            | ✅           | `[]object` |
| `override[].type`   | Type of override (`file`, `static`)                          | ✅           | `string`   |
| `override[].file_type` | File type for overrides (`yaml`, `json`) | ✅ (override[].file_type=file)   | `string`   |
| `override[].path`   | Path to the override file                  | ✅ (override[].file_type=file)   | `string`   |
| `override[].partial`| Enable partial override for file type (default: `false`)     | ❌           | `boolean`  |
| `override[].vars`   | Variables for static override             | ✅ (override[].file_type=static) | `[]object` |
| `override[].vars[].key` | Key to override                                          | ✅ (override[].file_type=static) | `string`   |
| `override[].vars[].value` | Value to assign                                        | ✅ (override[].file_type=static) | `string`   |
| `override[].enabled_env` | List of environments where the override is enabled      | ❌           | `[]string` |
