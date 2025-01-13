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
| `data[].bucket_id`| Bucket ID for the stored data.   | ✅           | `string`       |
| `data[].store_key`| Store Key for the stored data.   | ✅           | `string`       |
| `data[].encrypt`      | Encryption settings for the stored data. Disabled by default.                  | ❌                       | `object`    |
| `data[].encrypt.enabled` | Enable encryption for the stored data. Defaults to `false`.                  | ❌                       | `boolean`   |
| `data[].encrypt.encrypt_id` | Encryption ID for the stored data.   | ✅ (`enabled=true`)   | `string`    |


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