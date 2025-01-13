---
title: Configuration
nav_order: 3
---

# Configuration 🎛️

Bloader's configuration system is designed for simplicity and flexibility, leveraging **YAML** and **Sprig** templates to suit a variety of needs.

## Features
- Any format supported by Viper, such as yaml, json, toml, etc., can be read. 
- The configuration file for master and slave is different from each other. 
- Environment variables starting with `BLOADER_` can be overridden. Override is preferred [Override Mechanism](override.md).

## Configuration Options
1. **[Configurable Properties](prop.md)**: Properties that should be set as properties
2. **[Override Properties](override.md)**: Overwriting of configuration files by files and constants
3. **[Sample Configuration](sample.md)**: Indicates a configuration that serves as a configuration sample
4. **[Communication between Master and Slave](communication.md)**: Notes on communication between Master and Slave

Explore each section to customize Bloader to your needs.



