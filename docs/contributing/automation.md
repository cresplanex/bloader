# 🚀 Project Automation Guide

Welcome to the comprehensive guide for automating your development workflow with VSCode tasks and Makefile. This document provides a detailed overview of all available tasks, their purpose, and how to use them effectively.

---

## 🔧 Automation Overview

This project is configured with:
1. **VSCode Tasks**: Predefined tasks for streamlined development.
2. **Makefile**: Command-line automation for developers who prefer terminal workflows.

Both options allow for individual task execution or grouped workflows to ensure consistency and efficiency.

---

## 🔧 Available Tasks

### **Go-related Tasks**

1. **`go mod tidy`**: Clean up and synchronize your `go.mod` file with your codebase.
   ```bash
   # VSCode
   Run task: "Run go mod tidy"
   
   # Makefile
   make go-mod-tidy
   ```

2. **`gofumpt`**: Format Go files with extra strictness.
   ```bash
   # VSCode
   Run task: "Run gofumpt"
   
   # Makefile
   make gofumpt
   ```

3. **`goimports`**: Organize Go imports while excluding `./gen` directory.
   ```bash
   # VSCode
   Run task: "Run goimports"
   
   # Makefile
   make goimports
   ```

4. **`golangci-lint`**: Run the linter to catch potential issues.
   ```bash
   # VSCode
   Run task: "Run golangci-lint"
   
   # Makefile
   make golangci-lint
   ```

5. **`gotestsum`**: Execute tests and display results in a short, verbose format.
   ```bash
   # VSCode
   Run task: "Run gotestsum"
   
   # Makefile
   make gotestsum
   ```

---

### **Buf-related Tasks**

1. **`buf generate`**: Generate code from `.proto` definitions.
   ```bash
   # VSCode
   Run task: "Run buf generate"
   
   # Makefile
   make buf-generate
   ```

2. **`buf lint`**: Lint `.proto` files to ensure they follow best practices.
   ```bash
   # VSCode
   Run task: "Run buf lint"
   
   # Makefile
   make buf-lint
   ```

3. **`buf format`**: Format `.proto` files in place.
   ```bash
   # VSCode
   Run task: "Run buf format"
   
   # Makefile
   make buf-format
   ```

4. **`buf breaking`**: Detect breaking changes in `.proto` files against the repository.
   ```bash
   # VSCode
   Run task: "Run buf breaking"
   
   # Makefile
   make buf-breaking
   ```

---

## 🔄 Workflow Automation

### **Go Workflows**

1. **Rewrite Flow**:
   - Execute `gofumpt` and `goimports` tasks that may cause go rewriting.
   ```bash
   # VSCode
   Run task: "Run go rewrite flow"
   
   # Makefile
   make go-rewrite-flow
   ```

2. **Test Flow**:
   - Runs `golangci-lint` and `gotestsum` to ensure code quality.
   ```bash
   # VSCode
   Run task: "Run go test flow"
   
   # Makefile
   make go-test-flow
   ```

3. **CI Flow**:
   - Combines `rewrite and test workflows` for a complete Go CI process.
   ```bash
   # VSCode
   Run task: "Run go CI flow"
   
   # Makefile
   make go-ci-flow
   ```

4. **All Go Tasks**:
   - Runs all Go-related tasks, including `go mod tidy`.
   ```bash
   # VSCode
   Run task: "Run all go tasks"
   
   # Makefile
   make all-go-tasks
   ```

---

### **Buf Workflows**

1. **Test Flow**:
   - Lints `.proto` files and detects breaking changes.
   ```bash
   # VSCode
   Run task: "Run buf test flow"
   
   # Makefile
   make buf-test-flow
   ```

2. **CI Flow**:
   - Formats `.proto` files and runs the test flow.
   ```bash
   # VSCode
   Run task: "Run buf CI flow"
   
   # Makefile
   make buf-ci-flow
   ```

3. **All Buf Tasks**:
   - Runs code generation, formatting, linting, and breaking change detection.
   ```bash
   # VSCode
   Run task: "Run all buf tasks"
   
   # Makefile
   make all-buf-tasks
   ```

---

### **Complete Workflow**

To execute all tasks, including Go and Buf-related processes:
```bash
# VSCode
Run task: "Run all tasks"

# Makefile
make all-tasks
```

---

## 🛠️ Usage Scenarios

1. **Quick Formatting**:
   - Use `go rewrite flow` or `buf format` to ensure code and `.proto` files are formatted.

2. **Pre-Commit Check**:
   - Run `all tasks` to validate everything before committing.

---

## 📚 Best Practices

- Always run **all tasks** locally before pushing to the repository to catch errors early.
- Review breaking changes detected by `buf breaking` and document them in your PR.
- Use Makefile for command-line workflows and VSCode tasks for IDE-based workflows.

---

Let’s make development smoother and more efficient with these automation tools! 💪

