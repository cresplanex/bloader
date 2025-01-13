---
title: Store Value
parent: Loaders
nav_order: 1
---

# Store Value
> The preservation of data in the internal database, which is persistent, is inferior to memory restore in terms of performance, but it is effective in that it can be reused even after the flow has ended.

### Property

| **Item**               | **Description**                                                                 | **Required**             | **Type**    |
|------------------------|---------------------------------------------------------------------------------|--------------------------|-------------|
| `data`                | Values to be stored in the database                                             | ✅                       | `[]object`  |
| `data[].bucket_id`    | Bucket ID for the stored data                                                   | ✅                       | `string`    |
| `data[].key`          | Key for the stored data                                                        | ✅                       | `string`    |
| `data[].value`        | Data to be stored. Supports any type that can be expressed in YAML.             | ✅                       | `any`       |
| `data[].encrypt`      | Encryption settings for the stored data. Disabled by default.                  | ❌                       | `object`    |
| `data[].encrypt.enabled` | Enable encryption for the stored data. Defaults to `false`.                  | ❌                       | `boolean`   |
| `data[].encrypt.encrypt_id` | Encryption ID for the stored data.   | ✅ (`enabled=true`)   | `string`    |

### Sample

{% raw %}
``` yaml
{{- $mailDomainList := list "example.com" "test.com" "demo.org" -}}
kind: StoreValue
data:
  - bucket_id: "bucketForApp"
    key: "users"
    value:
    {{- range $index, $element := until .Values.DataCount }}
    {{- $randOn3 := randInt 0 3 }}
    {{- $randOn4to6 := randInt 4 7 }}
    {{- $rand2On4to6 := randInt 4 7 }}
    {{- $randOn6to8 := randInt 6 9 }}
      - id: {{ add $index 1 }}
        uuid: {{ uuidv4 }}
        name: {{ randAlpha $randOn4to6 | title }} {{ randAlpha $rand2On4to6 | title }}
        email: {{ randAlpha $randOn6to8 | lower }}@{{ index $mailDomainList $randOn3 }}
        phone: {{ randInt 1000000000 9999999999 }}
    {{- end }}
    encrypt:
      enabled: true
      encrypt_id: encryptDynamicCBC
```
{% endraw %}