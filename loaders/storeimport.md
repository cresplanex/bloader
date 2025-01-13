---
title: Store Import
parent: Loaders
nav_order: 3
---

# Store Import
> By loading data from the persistent internal database into the memory store, data can be moved in and out of this flow at high speed from then on.

### Property

| **Item**       | **Description**                                                       | **Required** | **Type**    |
|----------------|-----------------------------------------------------------------------|--------------|-------------|
| `data`        | Values to be stored in memory                                         | ✅           | `[]object`  |
| `data[].key`  | Key where the data will be stored in memory                           | ✅           | `string`    |
| `data[].value`| Data to be stored in memory. Supports any type expressible in YAML.   | ✅           | `any`       |


### Sample

{% raw %}
``` yaml
kind: StoreImport
data:
  - key: "usersList"
    bucket_id: "bucketForApp"
    store_key: "users"
    encrypt:
      enabled: true
      encrypt_id: encryptDynamicCBC
```
{% endraw %}