---
title: "Get Started"
nav_order: 2
---

# Get Started 🚀

A great way to start is to follow this document and learn the rudimentary usage, then go on to the detailed customization in the respective sections!

## How to Install

We offer many installation options to work on many platforms.
Make sure you can execute the following commands in accordance with [the installation guidelines](installation.md).

- Linux or Mac
  ```bash
  bloader version
  ```
- Windows
  ```powershell
  bloader.exe version
  ```

## Init On Your Project

`bloader` requires installation files, `$(PWD)/bloader.yaml`, `${HOME}/bloader.yaml`, and `/etc/bloader/bloader.yaml`.

{: .note }
> If you want to draw test scenarios for each project, it is recommended to prepare `bloader.yaml` directly under the project.

[configuration guideline](../configuration/index.md), please manually describe the configuration you wish to test.

{: .highlight }
In the future, we plan to provide a `bloader init` command for easy creation.

## Slave Setup (Optional)

You will need to do this if you want to use Slave to perform load testing in a distributed fashion.

After understanding the guidelines for Slave Setup, you may want to try one type of Loader Commad, [Slave Connect](../loaders/slaveconnect.md).

---

Once this is done, you are ready for most load testing. Now let's get to testing.
