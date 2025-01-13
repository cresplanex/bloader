---
title: "Loaders"
nav_order: 5
---

# Loaders 🚦

Loaders define the actual load-testing logic in Bloader. Each loader type offers unique features to tailor your tests.

## Internal Database and Memory Stores

Bloader provides **three types of stores** for different use cases:

1. **Internal Database**: 
   - A persistent store using `Bolt`, suitable for data reuse even after the loader terminates. 
   - It supports encryption for storing sensitive information.

2. **Global Memory Store**:
   - A high-speed memory store where data is **ephemeral** and deleted after the load test completes. 
   - Faster than the internal database but lacks persistence.

3. **Thread Memory Store**:
   - Similar to the global memory store but only valid within the current loader thread. 
   - Ideal for managing temporary, thread-specific data.

---

## Loader Kinds

> Bloader supports multiple kinds of loaders, each designed for specific tasks:

| **Type**        | **Description**                                                                                  |
|------------------|--------------------------------------------------------------------------------------------------|
| **[StoreValue](storevalue.md)**     | Saves a value into the internal database.                                                       |
| **[MemoryValue](memoryvalue.md)**    | Stores a value in the memory store.                                                             |
| **[StoreImport](storeimport.md)**    | Imports values from the internal database into the memory store.                                |
| **[OneExecute](oneexecute.md)**     | Sends a request once. Can save data from responses to the memory store or database.             |
| **[MassExecute](massexecute.md)**    | Sends multiple requests simultaneously or at intervals. Only supports output, not data storage. |
| **[SlaveConnect](slaveconnect.md)**   | Establishes a connection with a Slave.                                                          |
| **[Flow](flow.md)**           | Defines workflows, allowing serial or parallel execution of multiple loaders. Enables complex processing. |

---

## Definition Format

The configuration files use the **Sprig template engine**, offering flexible and dynamic settings. Below are the common settings applicable to all loaders:

### Common Configuration Options

| **Item**                | **Description**                                                                            | **Required**                   | **Type**       |
|--------------------------|--------------------------------------------------------------------------------------------|--------------------------------|----------------|
| `kind`                  | Specifies the type of loader.                                                              | ✅                              | `string`       |
| `sleep`                 | Configures sleep settings.                                                                 | ❌                              | `object`       |
| `sleep.enabled`         | Enables or disables sleep. Default is `false`.                                             | ❌                              | `boolean`      |
| `sleep.values`          | Sleep timing and duration settings.                                                        | ✅ (`enabled=true`)          | `[]object`     |
| `sleep.values[].duration` | Sleep duration in time format (e.g., `10s`, `1s`).                                        | ✅                              | `string`       |
| `sleep.values[].after`  | Specifies when to sleep: `init`, `exec`, `failedExec`.                                      | ✅                              | `string`       |
| `store_import`          | Configures StoreImport settings.                                                           | ❌                              | `object`       |
| `store_import.enabled`  | Enables or disables StoreImport. Default is `false`.                                       | ❌                              | `boolean`      |
| `store_import.data`     | Defines the StoreImport data.                                                              | ✅ (`enabled=true`)          | `[]object`     |
| `store_import.data[].bucket_id` | The bucket ID to import from.                                                       | ✅                              | `string`       |
| `store_import.data[].store_key` | The key in the store to import.                                                     | ✅                              | `string`       |
| `store_import.data[].key` | The key to store imported data in the memory store.                                       | ✅                              | `string`       |
| `store_import.data[].thread_only` | If `true`, data is imported only into the thread memory store. Default is `false`. | ❌                              | `boolean`      |
| `store_import.data[].encrypt` | Configures encryption settings for the imported data.                                | ❌                              | `object`       |
| `store_import.data[].encrypt.enabled` | Enables or disables encryption for imported data. Default is `false`.         | ❌                              | `boolean`      |
| `store_import.data[].encrypt.encrypt_id` | The ID of the encryption used to decrypt the imported data.                | ✅ (`enabled=true`)          | `string`       |

---

{: .note }
> The `store_import` setting here can be similar to Kind's `StoreImport` in the loader, but since it does not exist at the time of initial loading, it must be checked for its existence as follows.

{% raw %}
``` yaml
kind: MassExecute
type: http
store_import:
  enabled: true
  data:
    - key: "usersList"
      bucket_id: "bucketForApp"
      thread_only: true
      store_key: "users"
      encrypt:
        enabled: true
        encrypt_id: encryptDynamicCBC
output:
  enabled: true
  ids: 
    - outputLocalCSV
auth: 
  enabled: true
  auth_id: authForWeb
requests:
# In the executable file, 
# variables that are to be loaded in the `storeImport` statement 
# must be checked to make sure that they are not null, 
# since they are loaded once as null and then loaded again after re-import.
{{- if .ThreadValues.usersList }}
{{- range slice .ThreadValues.usersList 0 3 }}
  - target_id: "apiServer"
    endpoint: "/user-profiles"
    method: POST
    interval: 500ms
    await_prev_response: false
    headers: {}
    query_param: {}
    path_variables: {}
    body_type: json
    body:
      userId: "{{ .uuid }}"
      name: "{{ .name }}"
      email: "{{ .email }}"
    success_break:
      - count
      - time
      - statusCode/badRequest
    break:
      time: 2m
      count: 5
      sys_error: true
      parse_error: true
      write_error: true
      status_code:
        - id: internalServerError
          op: eq
          value: 500
        - id: badRequest
          op: eq
          value: 400
    response_type: json
    data:
      - key: "resultType"
        jmes_path: "data.resultType"
      - key: "memory_active"
        jmes_path: "data.result[*].value"
        on_nil: "ignore" # cancel or ignore(default)
      - key: "invalid_key"
        jmes_path: "data.result[*].invalid_key"
        on_nil: "ignore"
{{- end }}
{{- end }}
```
{% endraw %}

## Template Variables 📋

Below is the list of variables available for use in templates:

```yaml
SlaveValues:          # Data from the global memory store
ThreadValues:         # Data from the thread memory store
Dynamic:
  OutputRoot:         # Current output root directory
  LoopCount:          # Counter incremented for each loop in Flow when `count` is specified
  CallCount:          # Counter incremented for each nested Flow
  RequestLoopCount:   # Counter incremented sequentially for each request in MassExecute
SlaveValues:
  SlaveID:            # The SlaveID defined in SlaveConnect
  Index:              # Index of executors in the Slave, incremented from the top
```

## Load Event

The loader has events, and each loader can start processing or notify the user according to the events issued by the loader.

{: .note }
> These functions are still developing and many of them are not yet implemented.

This functionality is available through the [Event](. /event.md) section.

---

Explore the individual loader pages for detailed examples and configurations.
