---
title: Configuration Sample
parent: Configuration
nav_order: 3
---
# Sample Configuration

## Master Sample Configuration (YAML)

``` yaml
version: v1
type: master
env: "production"
loader:
  base_path: "loader"
targets:
  # The id is required, and it must be unique.
  - id: "apiServer"
    type: "http"
    values:
      - env: "local"
        url: "http://localhost:8080"
      - env: "production"
        url: "https://web.state.api.cresplanex.org"
  - id: "metricsServer"
    type: "http"
    values:
      - env: "local"
        url: "http://localhost:10090/api/v1/query"
      - env: "production"
        url: "https://prometheus.state.api.cresplanex.org"
  - id: "testServer"
    type: "http"
    values:
      - env: "local"
        url: "https://jsonplaceholder.typicode.com"
outputs:
    # The id is required, and it must be unique.
  - id: "outputLocalCSV"
    values:
      - env: "local"
        # The type is required.
        # Supported types are `local`.
        type: "local"
        # The format is required.
        # Supported formats are `csv`.
        format: "csv"
        base_path: "outputs/local-csv"
      - env: "production"
        type: "local"
        format: "csv"
        base_path: "outputs/prod-csv"
store:
  file:
    - env: "local"
      path: "local_store/local_store.db"
    - env: "production"
      path: "local_store/prod_store.db"
  buckets:
    - "bucketForEncrypt"
    - "bucketForCredential"
    - "bucketForApp"
    - "bucketForSlave"
encrypts:
    # The id is required, and it must be unique.
  - id: "encryptStaticCBC"
    # Supported types are `staticCBC`, `staticCFB`, `staticCTR`,
    # `dynamicCBC`, `dynamicCFB`, `dynamicCTR`.
    type: "staticCBC"
    key: "must override"
  - id: "encryptDynamicCBC"
    type: "dynamicCBC"
    store:
      bucket_id: "bucketForEncrypt"
      key: "encryptKeyForDynamicStore"
      encrypt:
        enabled: false
auth:
    # The id is required, and it must be unique.
  - id: "authForWeb"
    default: true
    type: "oauth2"
    oauth2:
      grant_type: "authorization_code"
      client_id: "register_client"
      client_secret: "register_client_secret"
      auth_url: "https://auth.state.api.cresplanex.org/oauth2/authorize"
      token_url: "https://auth.state.api.cresplanex.org/oauth2/token"
      access_type: "offline"
      scope:
        - "openid"
        - "email"
      credential:
        store:
          bucket_id: "bucketForCredential"
          key: "webAuthCredential"
          encrypt:
            enabled: true
            encrypt_id: "encryptDynamicCBC"
  - id: "authForMobile"
    default: false
    type: "oauth2"
    oauth2:
      grant_type: "client_credentials"
      client_id: "mobile_client"
      client_secret: "mobile_client_secret"
      token_url: "https://auth.state.api.cresplanex.org/oauth2/token"
      scope:
        - "openid"
        - "email"
      credential:
        store:
          bucket_id: "bucketForCredential"
          key: "mobileAuthCredential"
          encrypt:
            enabled: false
  - id: "authForApi"
    default: false
    type: "apiKey"
    api_key:
      header_name: "X-API-KEY"
      key: "You must override this value"
server:
  port: 9800
  # The server will be started on the `server.port` by default.
  # If you want to start the oauth redirect server on a different port, you can set the redirect_port.
  # For redirect URL, we use a fixed value of /auth/callback.
  # So, if you set the `server.redirect_port`` to 10800, 
  # the redirect URL will be http://localhost:10800/auth/callback.
  # redirect_port: 10800
logging:
  output:
    - type: "stdout"
      format: "text"
      level: "warn"
      # If enabled_env is not set, it will be enabled all environments.
      enabled_env:
        - "production"
    # - type: "file"
    #   filename: "logs/app.log"
    #   format: "text"
    #   level: "debug"
    # - type: "tcp"
    #   address: "127.0.0.1:5000"
    #   format: "json"
    #   level: "warn"
clock:
  fake: 
    # If enabled_env is not set, it will be enabled all environments.
    enabled: true
    time: "2021-01-01T00:00:00Z"
  # The default value is `2006-01-02T15:04:05Z`.
  format: "2006-01-02T15:04:05Z"
language:
  default: "en"
# The one below takes precedence over the one above.
# The bottom one takes precedence.
override:
  - type: "file"
    file_type: "yaml"
    path: "bloader/static_encrypt.yaml"
    partial: true
    vars:
      - key: encrypts[0].key
        value: "encrypt_key"
```

## Slave Sample Configuration (YAML)
``` yaml
version: v1
type: slave
env: "production" # Must be set to the same value as the master.
slave_setting:
  port: 50051
  certificate:
    enabled: true
    slave_cert: "certs/slave.crt"
    slave_key: "certs/slave.key"
encrypts:
    # The id is required, and it must be unique.
  - id: "encryptStaticCBC"
    # Supported types are `staticCBC`, `staticCFB`, `staticCTR`.
    type: "staticCBC"
    key: "must override"
logging:
  output:
    # - type: "stdout"
    #   format: "text"
    #   level: "warn"
    #   # If enabled_env is not set, it will be enabled all environments.
    #   enabled_env:
    #     - "production"
    - type: "file"
      filename: "logs/app.log"
      format: "text"
      level: "warn"
    # - type: "tcp"
    #   address: "127.0.0.1:5000"
    #   format: "json"
    #   level: "warn"
clock:
  fake: 
    # If enabled_env is not set, it will be enabled all environments.
    enabled: true
    time: "2021-01-01T00:00:00Z"
  # The default value is `2006-01-02T15:04:05Z`.
  format: "2006-01-02T15:04:05Z"
language:
  default: "en"
# The one below takes precedence over the one above.
# The bottom one takes precedence.
override:
  - type: "file"
    # The file_type is required to load the file, supported types are `yaml`, `json`.
    file_type: "yaml"
    path: "bloader/static_encrypt.yaml"
    partial: true
    vars:
      - key: encrypts[0].key
        value: "encrypt_key"
  # - type: "file"
  #   file_type: "yaml"
  #   path: "bloader/local_override.yaml"
  #   # If enabled_env is not set, it will be enabled all environments.
  #   enabled_env:
  #     - "local"
  # - type: "file"
  #   file_type: "yaml"
  #   path: "bloader/production_override.yaml"
  #   enabled_env:
  #     - "production"
  # - type: "static"
  #   key: "auth[0].oauth2.auth_url"
  #   value: "http://localhost:8080/oauth2/authorize"
  #   enabled_env:
  #     - "local"
  # - type: "static"
  #   key: "auth[0].oauth2.token_url"
  #   value: "http://localhost:8080/oauth2/token"
  #   enabled_env:
  #     - "local"
```