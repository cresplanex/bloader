---
title: Installation
parent: "Get Started"
nav_order: 1
---

# Installation Guide 🛠️

## Prerequisites
- Go version: `>=1.23`
- OS: Linux, macOS, Windows

---

## 1. Download Binary Directly (macOS/Linux)

### Prerequisites

Ensure the following environment variables are set before proceeding:

#### Required Environment Variables
- **`VERSION`**: Specify the version of `bloader` you want to install (e.g., `0.1.2`).
- **`BLOADER_BIN_PATH`**: Install Path Prefix (e.g., `/usr/local/bin`).

#### Example Setup

``` sh
export VERSION=0.1.2
export BLOADER_BIN_PATH="/usr/local/bin"
```

### Download the Release File
Download the appropriate archive file for your environment from the [GitHub Releases page](https://github.com/cresplanex/bloader/releases):

Use `curl` or `wget` to download the `.tar.gz` file:
``` sh
curl -LO "https://github.com/cresplanex/bloader/releases/download/v${VERSION}/bloader_${VERSION}_$(uname -s)_$(uname -m).tar.gz"
```

### Extract the Archive

Extract the `.tar.gz` file:
``` sh
tar -xzf "bloader_${VERSION}_$(uname -s)_$(uname -m).tar.gz"
```

### Move the Binary to a Directory in `PATH`

``` sh
sudo mv bloader ${BLOADER_BIN_PATH}
```

### Confirmation of installation completion

``` sh
bloader version
```

## 2. Download Binary Directly (Windows)

### Prerequisites

Ensure the following environment variables are set before proceeding:

#### Required Environment Variables
- **`ARCH`**: Set to your system architecture. Valid values are:
  - `x86_64`
  - `arm64`
- **`VERSION`**: Specify the version of `bloader` you want to install (e.g., `v1.2.3`).

#### Example Setup
For `x86_64` architecture with version `0.1.2`:

``` powershell
$env:ARCH="x86_64"
$env:VERSION="0.1.2"
```

### Download the Release File
Download the appropriate archive file for your environment from the [GitHub Releases page](https://github.com/cresplanex/bloader/releases):

Use `Invoke-WebRequest` to download the `.zip` file:
``` powershell
Invoke-WebRequest -Uri "https://github.com/cresplanex/bloader/releases/download/v$env:VERSION/bloader_$env:VERSION_Windows_$env:ARCH.zip" -OutFile bloader.zip
```

### Extract the Archive

Extract the `.zip` file:
``` powershell
Expand-Archive -Path bloader.zip -DestinationPath .
```

### Move the Binary to a Directory in `PATH`
Move the `bloader` binary to a directory in your `PATH`, such as a system path (Windows):

``` powershell
New-Item -ItemType Directory -Path "C:\Program Files\bloader"
Move-Item -Path bloader.exe -Destination "C:\Program Files\bloader\bloader.exe"
$env:Path += ";C:\Program Files\bloader"
```

### Confirmation of installation completion

``` sh
bloader.exe version
```

## 3. Using `go install`

If Go is installed, you can easily install the library using:

1. Run the following command:

   ```bash
   go install github.com/cresplanex/bloader@latest
   ```

   - This command places the binary in `$GOPATH/bin` or `$HOME/go/bin`.

2. If the Go binary directory is not in your `$PATH`, add it:

   ```bash
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

3. Confirm that the installation has been completed with the following command:

   ``` bash
   bloader version
   ```

## 4. Homebrew (macOS/Linux)

You can use Homebrew for easy installation:

```bash
brew install cresplanex/tap/bloader
```

Confirm that the installation has been completed with the following command.

``` bash
bloader version
```

## 5. Debian-based Linux (APT)

``` sh
export BLOADER_VERSION=0.1.2 # must update what you need
```

Installation is possible using APT:

1. Prerequire Install

   ``` sh
   sudo apt update
   sudo apt install wget
   ```

2. Download deb file

   ```bash
   wget "https://github.com/cresplanex/bloader/releases/download/v${BLOADER_VERSION}/bloader_${BLOADER_VERSION}_$(uname -s)_$(uname -m).deb"
   ```

3. Install the package:

   ```bash
   sudo apt install ./bloader_${BLOADER_VERSION}_$(uname -s)_$(uname -m).deb
   ```

## 6. RedHat-based Linux (RPM)

> Note that the test has not been performed.(TODO)

You can install using the RPM package:

1. Prerequire Install

   ``` sh
   sudo yum install wget
   ```

2. Download deb file

   ```bash
   wget "https://github.com/cresplanex/bloader/releases/download/v${BLOADER_VERSION}/bloader_${BLOADER_VERSION}_$(uname -s)_$(uname -m).rpm"
   ```

3. Install the package:

   If use `yum`...

   ```bash
   sudo yum localinstall bloader_${BLOADER_VERSION}_$(uname -s)_$(uname -m).rpm
   ```

   If use `dnf`...

   ``` bash
   sudo dnf install bloader_${BLOADER_VERSION}_$(uname -s)_$(uname -m).rpm
   ```

## 7. Build from Source

You can also build the library directly from the source using Go:

```bash
git clone https://github.com/cresplanex/bloader.git
cd bloader
go build -o bloader ./main.go
mv bloader /usr/local/bin/
```

---

# Advance Guide

## Add Completion Script
If necessary, a completion script can be generated so that commands can be used quickly and accurately.The currently supported complementary scripts are as follows.

- `bash`
- `fish`
- `powershell`
- `zsh`

For example, the following is an example of generating a completion script for bash.

``` sh
bloader completion bash > bloader-completion.bash
```

> Below is an example for linux, but it will work fine if the scripts are properly placed and loaded for each platform.

1. Adaptation to system
   If you want to enable this system-wide on linux, the following command is useful.

   ``` sh
   sudo mv bloader-completion.bash /etc/bash_completion.d/bloader
   ```

   For immediate use, execute the following command.
   ``` sh
   exec $SHELL
   ```

2. Adaptation to specific users
   Others can be enabled only for specific users as follows.

   ``` sh
   mv bloader-completion.bash ~/.bloader_bash_completion
   ```

   Include the following in your `.bashrc`, etc.
   ``` sh
   if [ -f ~/.bloader_bash_completion ]; then
      . ~/.bloader_bash_completion
   fi
   ```

   For immediate use, execute the following command.
   ``` sh
   source ~/.bashrc
   ```

## 🎨 Editor Configuration
If you are using **VS Code**, add the following configuration to your `.vscode/settings.json` file to ensure compatibility with YAML and other tools:

```json
{
    "yaml.validate": false,
    "yaml.customTags": [
        "!<tag:yaml.org,2002:map>",
        "!<tag:yaml.org,2002:str>",
        "!<tag:yaml.org,2002:int>",
        "{{.*}}"
    ]
}
```