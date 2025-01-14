# 🛠️ Local Setup Guide

Follow these steps to get the project running locally:

## 🖥️ Requirements
- **Go** 1.23+
- A terminal emulator (e.g., bash, zsh, PowerShell)
- Optional: Docker (for advanced testing)

## 📦 Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/cresplanex/bloader.git
   ```
2. Navigate into the project directory:
   ```bash
   cd bloader
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Format the code:
   ```bash
   gofumpt -extra -w .
   ```
5. Format package imports:
   ```bash
   find . -name "*.go" -not -path "./gen/*" -exec goimports -w -local github.com/cresplanex/bloader {} +
   ```
6. Run linting:
   ```bash
   golangci-lint run
   ```
7. Run the CLI:
   ```bash
   go run main.go
   ```

## 🚦 Testing
Run all tests using `gotestsum`:
```bash
gotestsum --format=short-verbose
```

## 🐢 Pro Tip
- Use `gofumpt`, `goimports`and `golangci-lint` integrations in your editor for a smooth workflow.
- If you also want to make changes to the `proto`, consider introducing [bufCLI](buf.md) commands into the editor integration flow.
- Use Docker if you need to test integrations or other advanced features.

You're all set! 🚀
