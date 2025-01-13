# ✏️ Commit Message Guide

We follow a structured format for commit messages to maintain clarity and consistency across the project. This guide outlines the required prefixes and provides examples.

## ✨ Valid Commit Prefixes
Use one of the following prefixes in your commit messages:

- **feat**: A new feature
- **fix**: A bug fix
- **docs**: Documentation only changes
- **style**: Changes that do not affect the meaning of the code (e.g., white-space, formatting, missing semi-colons)
- **refactor**: A code change that neither fixes a bug nor adds a feature
- **perf**: A code change that improves performance
- **test**: Adding missing or correcting existing tests
- **chore**: Changes to the build process or auxiliary tools and libraries (e.g., documentation generation)

---

## 🖋️ Commit Message Format
Each commit message should have the following format:

```plaintext
<type>: <short description>

<detailed explanation (if necessary)>
```

### Example
```plaintext
feat: add support for new CLI command

Added a new CLI command `app run` to execute user-defined tasks. This feature includes options for customization and integrates with existing modules.
```

---

## 🔗 Why Use This Format?
1. **Consistency**: Ensures all commit messages are easy to understand.
2. **Automation**: Helps tools like changelogs and release scripts parse commits effectively.
3. **Collaboration**: Makes it easier for team members to understand changes.

---

For questions or suggestions, feel free to reach out or open an issue. Happy committing! 🚀
