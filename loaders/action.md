---
title: Load Action
parent: Loaders
nav_order: 9
---

# Load Action 📆

> The current implementation includes only a limited number of actions per Kind. However, more actions are planned for future updates.

### Common Actions
- **`sys:term_with_err`**: Exit with error. This error is usually contagious to other runners, so all runners exit with the error.
- **`sys:term_without_err`**: Exit without error. It is useful to control the end timing of asynchronous execution, since only that runner is terminated without contagion to others.

### SlaveConnect Actions
- **`slaveConnect:disconnect`**: Disconnects the connection to the specified Slave.