# 📚 Buf Usage Guide

Working with gRPC definitions between master and slave? Here’s how to properly use `buf` CLI to manage your `.proto` files efficiently. 💡

---

## 🛠️ Getting Started with Buf CLI

1. **Install Buf CLI**:
   - Follow the [official installation guide](https://buf.build/docs/installation/).
   - Ensure Buf CLI is properly installed before contributing.

---

## 🔧 Pre-Commit Best Practices

Before committing your changes, follow these steps:

1. **Format Your `.proto` Files**:
   ```bash
   buf format -w
   ```

2. **Check for Breaking Changes**:
   Run the following command to detect breaking changes against the branch you're working on:
   ```bash
   buf breaking proto --against '.git#subdir=proto'
   ```

3. **Accidental Commit? No Problem!**  
   If you committed without running the checks, refer to the [Buf Breaking Changes Overview](https://buf.build/docs/breaking/overview/) to identify issues against the base branch.

4. **Local Checks Are Recommended**:  
   While these checks are performed in the CI workflow, running them locally helps catch issues earlier.

---

## 📋 Pull Request Guidelines

1. If your changes include **breaking changes**, clearly mention this in your PR description.
2. Follow the project’s formatting and breaking change detection practices.

---

## 🏗️ Generating Code with Buf

After defining your `.proto` files:
1. Run `buf generate` to generate code:
   ```bash
   buf generate
   ```
2. Review the generated files before committing.

---

## 🔄 Managing Dependencies

Buf CLI allows you to update dependencies as needed. Ensure you verify all changes and resolve conflicts, if any.

---

## 🚀 CI Workflow

Our CI workflow automatically publishes the latest gRPC definitions to [Buf Schema Registry (BSR)](https://buf.build/cresplanex/bloader). Stay aligned with these standards for smooth integration.

---

Let’s make those `.proto` files efficient and reliable! 🚀
