---
title: Flow
parent: Loaders
nav_order: 7
---

# Flow
> All loader definitions can be stitched together, allowing for very sophisticated load testing using internal memory stores.

### Property

| **Field**                           | **Description**                                                                                                                                                                      | **Required**                                      | **Type**   |
|-------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------|------------|
| `step`                              | Definition of the flow step.                                                                                                                                                        | ✅                                              | `object`   |
| `step.concurrency`                  | Maximum concurrency for execution. `-1` runs all flows simultaneously, `0` ensures sequential execution on the main thread.                                                         | ✅                                              | `int`      |
| `step.flows`                        | Definition of flows to execute.                                                                                                                                                     | ❌                                              | `[]Flow`   |
| `step.flows[].id`                   | Unique flow ID within the file. Note: the uniqueness applies across the file, not just parallel flows.                                                                               | ✅                                              | `string`   |
| `step.flows[].depends_on`           | Dependencies that must resolve before starting, irrespective of `concurrency`.                                                                                                      | ❌                                              | `[]object` |
| `step.flows[].depends_on[].flow`    | Flow ID dependency. Must be within the same file.                                                                                                                                    | ✅                                              | `string`   |
| `step.flows[].depends_on[].event`   | Event triggering the dependency resolution. Supported events are listed under [Events](./event.md).                                                                                    | ✅                                              | `string`   |
| `step.flows[].type`                 | Type of flow to execute. Supported types: `file`, `flow`, `slaveCmd`.                                                                                                                | ✅                                              | `string`   |
| `step.flows[].file`                 | Path to the loader file to execute.                                                                                                                                                | ✅ (`type=file` or `type=slaveCmd`)          | `string`   |
| `step.flows[].mkdir`                | Creates a new directory for output related to the flow, useful for separating concurrent requests to different services. Defaults to `false`.                                         | ❌                                              | `boolean`  |
| `step.flows[].count`                | Number of executions for `type=file` flows. Defaults to `1`.                                                                                                                        | ❌                                              | `int`      |
| `step.flows[].values`               | Data stored in the global memory store. Accessible after execution but not during the same file's parsing.                                                                           | ❌                                              | `[]object` |
| `step.flows[].values.key`           | Key for the memory store data.                                                                                                                                                      | ❌                                              | `string`   |
| `step.flows[].values.value`         | Value for the memory store data.                                                                                                                                                    | ❌                                              | `any`      |
| `step.flows[].thread_values`        | Data stored in the thread-local memory store, valid only within the flow.                                                                                                            | ❌                                              | `[]object` |
| `step.flows[].thread_values.key`    | Key for the thread memory store data.                                                                                                                                               | ❌                                              | `string`   |
| `step.flows[].thread_values.value`  | Value for the thread memory store data.                                                                                                                                             | ❌                                              | `any`      |
| `step.flows[].flows`                | Nested flow definitions. Valid if `type=flow`.                                                                                                                                      | ❌                                              | `[]Flow`   |
| `step.flows[].concurrency`          | Maximum concurrency for nested flows. `-1` runs all flows simultaneously, `0` ensures sequential execution.                                                                          | ✅ (`type=flow`)                             | `int`      |
| `step.flows[].executors`            | Slave configuration for execution.                                                                                                                                                  | ❌                                              | `[]object` |
| `step.flows[].executors.slave_id`   | Slave ID as defined in [SlaveConnect](./slaveconnect.md).                                                                                                                 | ✅                                              | `string`   |
| `step.flows[].executors.output`     | Output configuration for individual slaves. Defaults to disabled.                                                                                                                   | ❌                                              | `object`   |
| `step.flows[].executors.output.enabled` | Enable output for individual slaves. Defaults to `false`.                                                                                                                       | ❌                                              | `boolean`  |
| `step.flows[].executors.output.root_path` | Root directory for slave-specific output. Required if `output.enabled=true`.                                                                                                  | ✅ (`output.enabled=true`)                  | `string`   |
| `step.flows[].executors.inherit_values` | Inherit global memory store from the master during execution. Defaults to `false`.                                                                                              | ❌                                              | `boolean`  |
| `step.flows[].executors.additional_values` | Data to store in the slave's global memory store, allowing slave-specific values.                                                                                                | ❌                                              | `[]object` |
| `step.flows[].executors.additional_values.key` | Key for the slave memory store data.                                                                                                                                             | ❌                                              | `string`   |
| `step.flows[].executors.additional_values.value` | Value for the slave memory store data.                                                                                                                                           | ❌                                              | `any`      |
| `step.flows[].executors.additional_thread_values` | Data stored in the slave's thread-local memory store, valid only within the flow, for slave-specific values.                                                                    | ❌                                              | `[]object` |
| `step.flows[].executors.additional_thread_values.key` | Key for the slave thread memory store data.                                                                                                                                     | ❌                                              | `string`   |
| `step.flows[].executors.additional_thread_values.value` | Value for the slave thread memory store data.                                                                                                                                   | ❌                                              | `any`      |

### Sample

{% raw %}
``` yaml
kind: Flow
sleep:
  enabled: true
  values:
    - duration: "5s"
      after: init
step:
  concurrency: -1
  flows:
    - id: "slaveConnect"
      type: file
      mkdir: false
      file: "sc/slave/connect.yaml"
      values: []
      thread_only_values: []
    - id: "metrics"
      type: file
      mkdir: true
      file: "sc/metrics/main.yaml"
      values:
        - key: "MetricsInterval"
          value: "5s"
        - key: "MetricsBreakTime"
          value: "10m"
      thread_only_values: []
    - id: "request"
      depends_on:
        - flow: slaveConnect
          event: slaveConnect:connected
      type: slaveCmd
      mkdir: true
      thread_only_values:
        - key: "Interval"
          value: "100ms"
        - key: "BreakTime"
          value: "5m"
      executors:
        {{- range slice .Values.slaveLists 0 .Values.SlaveCount }}
        - slave_id: "{{ .id }}"
          output:
            enabled: true
            root_path: "{{ .id }}"
          inherit_values: true
          additional_values: []
          additional_thread_only_values: []
        {{- end }}
      file: "sc/sc1/request.yaml"
```
{% endraw %}