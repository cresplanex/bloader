# 🌈 Style Guide

> [!WARNING]
> Please check [Buf Usage Guide](buf.md) for notes on `proto` styles.

This style guide provides conventions and best practices for writing clean, maintainable, and consistent code in this Go project. Following this guide ensures that the codebase remains easy to read and contributes to a collaborative environment.

---

## 🔨 General Guidelines
1. **Use idiomatic Go**: Follow the [official Go code review comments](https://github.com/golang/go/wiki/CodeReviewComments).
2. **Write small, focused functions**: Keep functions short and focused on a single task.
3. **Use meaningful names**: Variables, functions, and package names should clearly describe their purpose.
4. **Avoid commented-out code**: Remove unused code instead of commenting it out.
5. **Add comments for exported identifiers**: Use `//` comments for all exported types, functions, and methods.

---

## 🛠️ Formatting

### ✏️ Tools
- Use `gofumpt -extra` for formatting.
- Install with:
  ```bash
  go install mvdan.cc/gofumpt@latest
  ```

### ▶️ Formatting Rules
1. **Indentation**: Use tabs for indentation.
2. **Imports**:
   - Group imports into three sections:
     ```go
     import (
         "standard/library"

         "external/library"

         "project/internal/package"
     )
     ```
   - Use `goimports -local github.com/cresplanex/bloader -w .` to organize imports automatically.

---

## 🌐 Naming Conventions

### ✍️ Packages
- Package names should be short and meaningful (e.g., `utils`, `auth`).
- Avoid using underscores or camel case in package names.

### ✍️ Variables
- Use camelCase for local variables and function arguments.
- Use meaningful, descriptive names for variables.
- Short variable names like `i`, `j`, or `x` are acceptable in small scopes (e.g., loops).

### ✍️ Constants
- Use `ALL_CAPS` for unexported constants (e.g., `defaultTimeout`).
- Use camelCase for exported constants (e.g., `MaxRetries`).

### ✍️ Functions
- Use camelCase for function and method names.
- Prefix test helper functions with `test` (e.g., `testValidateInput`).

### ✍️ Structs and Interfaces
- Use PascalCase for exported structs and interfaces.
- Use meaningful names that describe their purpose (e.g., `User`, `AuthService`).

---

## 🛡️ Testing
1. Write table-driven tests for functions with multiple input-output scenarios.
2. Use meaningful test names, starting with `Test` (e.g., `TestCalculateSum`).
3. Mock dependencies where necessary to isolate functionality.
4. Check for edge cases and invalid inputs.
5. Use `gotestsum` to run tests:
   ```bash
   gotestsum --format=short-verbose
   ```

---

## 🔢 Code Structure

### 🏢 Project Layout
- Follow the standard Go project layout:
  ```plaintext
  /cmd              # CLI commands
  /internal         # Private application code
  ```

---

## ⚠️ Avoid Common Pitfalls
1. **Global variables**: Minimize their use; prefer dependency injection.
2. **Error handling**:
   - Always check for errors.
   - Use meaningful error messages.
   - Wrap errors with context using `fmt.Errorf` or `errors.Join`.
3. **Panic**: Avoid using `panic` for error handling unless absolutely necessary.

---

## 🕐 Automation
- Set up your editor to automatically format code on save.
- We support `vscode` with `task` and `Makefile`. For more information, check the [Automation Guide](automation.md).
- Run `golangci-lint` to catch linting issues.
  ```bash
  golangci-lint run
  ```
- Include pre-commit hooks to ensure formatting and linting:
  ```bash
  gofumpt -extra -w . && \
  find . -name "*.go" -not -path "./gen/*" -exec goimports -w -local github.com/cresplanex/bloader {} + && \
  golangci-lint run
  ```

---

Adhering to this guide will ensure a consistent, maintainable, and high-quality codebase. Happy coding! 🎉
