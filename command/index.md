---
title: Command
nav_order: 4
---

# Commands 📤

The `bloader` provides many commands in addition to running load tests. Most commands use [Config file](../configuration/index.md) is required for most commands.

{: .highlight }
However, some commands listed in the [Base Commands](#base-commands) section do not require a configuration file.

{: .warning }
> In addition, there are commands available only for slave and commands available only for master. The appropriate configuration file must be selected.

If you want to specify a configuration file on a per-command basis, check the [Command Options](#common-options) section.

#### Common Options

- `--config`
  Bloader explores configuration files in the following order by default: `$PWD`, `$HOME`, `/etc/bloader`. It searches for `bloader.*` files (supported formats align with [Viper's supported formats](https://github.com/spf13/viper?tab=readme-ov-file#reading-config-files)). To override this behavior, use the `-c` (or `--config`) option:

  ```bash
  bloader {command} -c bloader/config.yaml
  ```
- `--help`
  You can display help for any command using the `-h` (or `--help`) option. You can also learn about valid options for commands, etc.

---

### Base Commands

- `version`
  You can get the current version, the date and time it was released, etc. This command can be executed without a configuration file.
  ```bash
  bloader version
  ```
- `help`
  The entire help can be displayed. This command can be executed without the need for a configuration file.
  ```bash
  bloader help
  ```
- `completion`
  Generate a completion script. For details, see [Add Completion Script](../setup/installation.md#add-completion-script) for details. This command can be executed without a configuration file.
- `config`
  Display the current configuration after file and environment variable overrides
  ```bash
  bloader config
  ```

---

### Application Commands
1. **[Encryption](encryption.md)**: Used for simple encryption/decryption on command.
2. **[Authentication](auth.md)**: Use when you want to do something authentication related.
3. **[Store](store.md)**: Used for simple handling of internal databases from commands.
4. **[Output](output.md)**: Use when you want to perform output-related operations.
5. **[Load](load.md)**: Used for load testing.
6. **[Slave](slave.md)**: Commands for slave.

---

These commands provide a flexible and comprehensive way to manage and execute load testing tasks with Bloader. 



