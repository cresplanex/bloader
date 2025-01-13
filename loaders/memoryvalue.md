---
title: Memory Value
parent: Loaders
nav_order: 2
---

# Memory Value
> The use of an on-memory store allows for the fast transfer of reusable data into and out of the current flow for load testing.

### Property

| **Item**       | **Description**                                                          | **Required** | **Type**    |
|----------------|--------------------------------------------------------------------------|--------------|-------------|
| `data`        | Values to be stored in memory                                            | ✅           | `[]object`  |
| `data[].key`  | Key for storing the data in memory                                       | ✅           | `string`    |
| `data[].value`| Data to be stored in memory. Supports any type that can be expressed in YAML. | ✅       | `any`       |

### Sample

{% raw %}
``` yaml
{{- $mailDomainList := list "example.com" "test.com" "demo.org" -}}
kind: MemoryValue
data:
  - key: "usersList"
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
```
{% endraw %}