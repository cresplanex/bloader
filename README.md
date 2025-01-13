<h1 align="center">
  <a href="https://docs.bloader.cresplanex.org" style="display: inline-block; margin-bottom: 16px;">
    <picture>
      <source height="125" media="(prefers-color-scheme: dark)" srcset="docs/static/bloader_logo.png">
      <img height="125" alt="Bloader" src="docs/static/bloader_logo.png">
    </picture>
  </a>
  <br>
  <a href="https://pkg.go.dev/github.com/cresplanex/bloader">
    <img src="https://pkg.go.dev/badge/github.com/cresplanex/bloader.svg" alt="Go Reference">
  </a>
  <a href="https://github.com/cresplanex/bloader/actions/workflows/ci.yaml">
    <img src="https://github.com/cresplanex/bloader/actions/workflows/ci.yaml/badge.svg" alt="CI">
  </a>
  <a href="https://codecov.io/gh/cresplanex/bloader">
    <img src="https://codecov.io/gh/cresplanex/bloader/branch/main/graph/badge.svg" alt="Test Coverage">
  </a>
  <a href="https://goreportcard.com/report/github.com/cresplanex/bloader">
    <img src="https://goreportcard.com/badge/github.com/cresplanex/bloader" alt="Go Report Card">
  </a>
  <a href="https://github.com/cresplanex/bloader/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/cresplanex/bloader" alt="License">
  </a>
  <a href="https://github.com/cresplanex/bloader/releases">
    <img src="https://img.shields.io/github/v/release/cresplanex/bloader.svg" alt="Release">
  </a>
  <a href="https://github.com/cresplanex/bloader/commits/main">
    <img src="https://img.shields.io/github/last-commit/cresplanex/bloader.svg" alt="Last Commit">
  </a>
  <a href="https://golang.org/doc/devel/release.html">
    <img src="https://img.shields.io/github/go-mod/go-version/cresplanex/bloader.svg" alt="Go Version">
  </a>
</h1>
<p align="center">
  <em>Bloader is a cutting-edge load testing tool designed for <b>simplicity</b> and <b>flexibility</b>. While still in active development, it already offers robust features to execute and manage load tests effectively. We welcome contributions from the community to help shape its future. 💡</em>
</p>

---

## ✨ Key Features

✔️ **Internal Storage**: Bloader can store request data for enhanced testing capabilities.

✔️ **Master-Slave Architecture**: Avoid client-side bottlenecks by leveraging gRPC communication between Master and Slave nodes.

✔️ **Human-Friendly Configuration**: Load tests are defined using YAML, with support for Sprig's template engine (used in Helm), offering exceptional flexibility.

---

## 🛠️ Installation

Many installation methods are supported, details of which can be found at [Installation](docs/installation.md) for details.

## 📄 Configuration File Support

Bloader supports multiple formats, including YAML, JSON, and TOML, via the Viper library. Configuration items differ for Master and Slave nodes and can be overridden by environment variables prefixed with `BLOADER_`, which take precedence over file-based configurations.

### Example: Master Configuration
```yaml
type: master
env: "production"
loader:
  base_path: "loader"
targets:
  - id: "apiServer"
    type: "http"
    values:
      - env: "local"
        url: "http://localhost:8080"
      - env: "production"
        url: "https://api.example.org"
outputs:
  - id: "localOutput"
    values:
      - env: "production"
        type: "local"
        format: "csv"
        base_path: "outputs/prod"
store:
  file:
    - env: "production"
      path: "local_store/prod_store.db"
encrypts:
  - id: "dynamicEncrypt"
    type: "dynamicCBC"
    store:
      bucket_id: "encryptBucket"
      key: "dynamicKey"
server:
  port: 9800
logging:
  output:
    - type: "stdout"
      format: "text"
      level: "warn"
```

### Example: Slave Configuration
```yaml
type: slave
env: "production"
slave_setting:
  port: 50051
  certificate:
    enabled: true
    slave_cert: "certs/slave.crt"
    slave_key: "certs/slave.key"
logging:
  output:
    - type: "file"
      filename: "logs/slave.log"
      format: "text"
      level: "warn"
```

---

## ⚙️ Command-Line Interface

### Global Options

- **Config File**: Override the default configuration path.
  ```sh
  bloader {command} -c /path/to/config.yaml
  ```
- **Help**: Display help for any command.
  ```sh
  bloader help
  ```

### Common Commands

- **Show Current Config**: Displays the merged configuration after overrides.
  ```sh
  bloader config
  ```
- **Encrypt/Decrypt Data**: Secure your data using pre-configured encryption settings.
  ```sh
  bloader encrypt "test-data" -i dynamicEncrypt
  bloader encrypt "encrypted-data" -i dynamicEncrypt -d
  ```

### Master-Specific Commands

- **Run Load Test**: Execute a load test using a specified file.
  ```sh
  bloader run -f loader.yaml
  ```
- **Authenticate**: Manage authentication tokens.
  ```sh
  bloader auth login -i oauthAuth
  ```
- **Manage Data Store**:
  ```sh
  bloader store list
  bloader store object get --bucket encryptBucket keyName
  ```

### Slave-Specific Commands

- **Start Slave Node**: Initialize a Slave node for distributed testing.
  ```sh
  bloader slave run -c /path/to/slave_config.yaml
  ```

---

## 🎯 Supported Features

- **Load Test Definitions**: Define targets and parameters using YAML.
- **Internal and Memory Store**: Data can be persisted or temporarily stored for flexibility.
- **Extensible Encryption**: Supports dynamic and static encryption configurations.
- **Multi-Environment Support**: Easily switch between environments (e.g., `local`, `production`).

---

## 🚀 Planned Features

- Transition from BoltDB to a more actively maintained database.
- Integration with cloud services for configuration overrides.
- Enhanced analysis tools and visualization capabilities.
- Web-based UI for intuitive load test management.
- Support for gRPC-based targets.
- Plugin system for custom extensions.

---

## 🤝 Contributing

We welcome contributions of all kinds! If you're interested in improving Bloader, please:

1. Fork the repository.
2. Make your changes in a new branch.
3. Submit a pull request with a detailed description.

For more details, see [CONTRIBUTION.md](./docs/contributing/index.md).

---

## 🛠️ Related Tools

- [Sprig](https://masterminds.github.io/sprig/): Template engine for flexible configurations.
- [Cobra](https://github.com/spf13/cobra): CLI framework.
- [Viper](https://github.com/spf13/viper): Configuration management.
- [Bolt](https://github.com/boltdb/bolt): Lightweight internal database.

---

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
