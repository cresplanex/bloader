# 🔀 Pull Request Guidelines

Pull Requests (PRs) are the lifeblood of open source. Here’s how to create a great one:

## 📋 Checklist Before Submitting
- [ ] The code is clean and follows our [style guide](style_guide.md).
- [ ] Code formatting has been applied using `gofumpt -extra`.
- [ ] Import formatting has been applied using `goimports -local github.com/cresplanex/bloader`.
- [ ] Linting has been performed using `golangci-lint run`.
- [ ] Tests have been executed with `gotestsum` and all pass.
- [ ] Documentation has been updated (if applicable).
- [ ] CI checks pass successfully.

## 📝 PR Template
Use this format for your PR description:

```markdown
### Description
A brief overview of the changes.

### Related Issue
Closes #<Issue Number>

### How Has This Been Tested?
Explain how you verified your changes.

### Checklist
- [ ] I have run `gofumpt -extra`.
- [ ] I have run `goimports -local github.com/cresplanex/bloader`.
- [ ] I have run `golangci-lint`.
- [ ] I have tested the code with `gotestsum`.
- [ ] I have updated documentation if necessary.
```

## 🚀 Get Ready for Feedback
Don’t worry if you get a few comments! It’s all part of the process. Let’s make it better together! 💪