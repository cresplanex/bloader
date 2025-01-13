# 💻 Code Contribution Guide

Want to dive into the codebase? Follow these steps to ensure your contributions fit in smoothly.

## 🚀 Steps to Contribute
1. **File an Issue**: Start by [opening an issue](https://github.com/cresplanex/bloader/issues) to discuss your idea.
2. **Create a Branch**: Use a meaningful branch name (e.g., `feature/add-command`).
3. **Write Code**: Adhere to the [style guide](../style_guide.md) and keep it clean!
4. **Format Your Code**: Use `gofumpt -extra` to ensure consistent formatting (details below).
5. **Import Format Your Code**: Use `goimports -local github.com/cresplanex/bloader` to ensure consistent import formatting (details below).
6. **Lint Your Code**: Run `golangci-lint` to catch potential issues (details below).
7. **Test Your Code**: Use `gotestsum` to run tests and ensure everything passes (details below).
8. **Submit a Pull Request**: Follow the [Pull Request Guide](pull_request.md) to submit your changes.

---

## ⚙️ Formatting Your Code
To maintain a consistent style across the project, we use **`gofumpt`** with the `-extra` flag. This ensures that all code adheres to Go standards while applying stricter formatting rules for readability and consistency.

### 🛠️ Install `gofumpt`
If you don’t already have `gofumpt`, install it with:
```bash
go install mvdan.cc/gofumpt@latest
```

### 📋 Run `gofumpt`
Before committing your changes, format your code:
```bash
gofumpt -extra -w .
```

### 🌐 Install `goimports`
Install `goimports` with:
```bash
go install golang.org/x/tools/cmd/goimports@latest
```

### ✏️ Run `goimports`

Use `goimports` to manage and organize imports with three groups:

1. **Standard library packages**
2. **Local project packages**
3. **Third-party packages**

- **`-local`**: Put imports beginning with this string after 3rd-party packages.
- **`-w`**: Write result to (source) file instead of stdout.

Configure `goimports` to use three groups by setting it up in your editor or running it manually:
```bash
goimports -local github.com/cresplanex/bloader -w .
```

### 🛡️ CI Enforcement
Our CI pipeline will check for proper formatting. Running `gofumpt` and `goimports` locally ensures your changes pass without issues.

---

## 🛠️ Linting Your Code
To ensure high-quality code, use **`golangci-lint`** to check for linting issues.

### 🛠️ Install `golangci-lint`
Install `golangci-lint` with:
```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### 📋 Run `golangci-lint`
Run the linter on your codebase:
```bash
golangci-lint run
```

### 🛡️ CI Enforcement
Our CI pipeline will fail if linting issues are detected. Ensure your code is lint-free before submitting a PR.

---

## ⚖️ Testing Your Code
We use **`gotestsum`** to enhance test output readability and manage test runs effectively.

### 🛠️ Install `gotestsum`
Install `gotestsum` with:
```bash
go install gotest.tools/gotestsum@latest
```

### 📋 Run Tests
Run tests with:
```bash
gotestsum --format=short-verbose
```

### ✅ Ensure Tests Pass
Make sure all tests pass locally before pushing your changes.

---

## ⚠️ Coding Standards
- Follow idiomatic Go practices.
- Use `gofumpt -extra` for formatting (it enforces stricter rules than `gofmt`).
- Use `goimports -local github.com/cresplanex/bloader` for package imports formatting.
- Run `golangci-lint` to catch linting issues.
- Add meaningful comments for exported functions and avoid commented-out code.

---

## 🕐 Pro Tip
To save time, configure your editor to run `gofumpt`, `goimports` and `golangci-lint` automatically on save. Many IDEs and editors (like VS Code, GoLand, or Vim) support integration with external formatting and linting tools.

Happy coding! 🎉
