---
title: Load Event
parent: Loaders
nav_order: 8
---

# Load Event 📆

> The current implementation includes only a limited number of events per Kind. However, more events are planned for future updates.

### Common Events
- **`sys:start`**: Triggered when the loader starts.
- **`sys:store:importing`**: Fired before the `store_import` process begins. This event is unrelated to the `StoreImport` Kind and is not triggered if the `store_import` field is not configured.
- **`sys:store:imported`**: Fired after the completion of the `store_import` process. Similar to `sys:store:importing`, it is unrelated to the `StoreImport` Kind and requires the `store_import` field to be configured.
- **`sys:validating`**: Triggered before the validation process begins.
- **`sys:validated`**: Fired after the validation process completes.
- **`sys:terminated`**: Triggered when the loader terminates.

### SlaveConnect Events
- **`slaveConnect:connecting`**: Fired before establishing a connection with a Slave.
- **`slaveConnect:connected`**: Triggered upon successfully connecting to a Slave.