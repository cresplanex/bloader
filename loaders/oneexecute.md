---
title: One Execute
parent: Loaders
nav_order: 4
---

# One Execute
> Instead of being able to make only one request, data can be extracted from the response body and stored in an internal database or memory store.

### Property

| **Field**                    | **Description**                                                                                                                                                                             | **Required**                         | **Type**      |
|-------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------|---------------|
| `type`                       | Execution target type. Currently, only `http` is supported.                                                                                                                                | ✅                                   | `string`      |
| `output`                     | Output settings for execution. Defaults to disabled.                                                                                                                                       | ❌                                   | `object`      |
| `output.enabled`             | Enable output for execution. Defaults to `false`.                                                                                                                                          | ❌                                   | `boolean`     |
| `output.ids`                 | Specify the outputs to enable. Defaults to an empty array.                                                                                                                                 | ❌                                   | `[]string`    |
| `auth`                       | Authentication settings. Defaults to disabled.                                                                                                                                             | ❌                                   | `object`      |
| `auth.enabled`               | Enable authentication. Defaults to `false`.                                                                                                                                               | ❌                                   | `boolean`     |
| `auth.auth_id`               | ID of the authentication to enable. Defaults to the default authentication if not specified.                                                                                               | ❌                                   | `string`      |
| `request`                    | The request to be sent.                                                                                                                                                                   | ✅ (`type=http`)                  | `object`      |
| `request.target_id`          | Target ID for the request.                                                                                                                                                                | ✅                                   | `string`      |
| `request.endpoint`           | Endpoint relative to the target base. Bracketed variables `{var}` can be replaced using `path_variables`.                                                                                  | ✅                                   | `string`      |
| `request.method`             | HTTP method for the request. Supported values: `OPTIONS`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `TRACE`, `CONNECT`.                                                                     | ✅                                   | `string`      |
| `request.query_param`        | Query parameters. Arrays can be used for multiple values under the same key.                                                                                                               | ❌                                   | `map[string]any` |
| `request.path_variables`     | Path variables to replace bracketed variables in the endpoint.                                                                                                                            | ❌                                   | `map[string]string` |
| `request.headers`            | Request headers. Arrays can be used for multiple values under the same key.                                                                                                                | ❌                                   | `map[string]any` |
| `request.body_type`          | Body type for the request. Supported values: `json`, `form`, `multipart`.                                                                                                                  | ✅                                   | `string`      |
| `request.body`               | Request body. The type varies depending on `body_type`.                                                                                                                                    | ❌                                   | `any`         |
| `request.response_type`      | Response body type. Supported values: `json`, `xml`, `yaml`, `text`, `html`.                                                                                                               | ✅                                   | `string`      |
| `request.data`               | Data to include in the output. Default keys include `success`, `sendDatetime`, `receivedDatetime`, `Count`, `ResponseTime`, `StatusCode`. Extracted data from the body can also be included. | ❌                                   | `[]object`    |
| `request.data[].key`         | Key for the output data.                                                                                                                                                                   | ✅                                   | `string`      |
| `request.data[].extractor`   | Extractor for the output data.                                                                                                                                                            | ✅                                   | `object`      |
| `request.data[].extractor.type` | Type of extractor. Currently, only `jmesPath` is supported (used for `json`, `xml`, `yaml` response types).                                                                               | ✅                                   | `string`      |
| `request.data[].extractor.jmes_path` | Extraction rule specified using JMESPath. Required if `extractor.type=jmesPath`.                                                                                                     | ✅ (`type=jmesPath`)              | `string`      |
| `request.data[].extractor.on_nil` | Behavior when extraction fails. Options: `empty`, `null`, `error`. Defaults to `null`.                                                                                               | ❌                                   | `string`      |
| `request.memory_data`        | Data to store in the global memory store.                                                                                                                                                 | ❌                                   | `[]object`    |
| `request.memory_data[].key`  | Key for the memory store data.                                                                                                                                                           | ✅                                   | `string`      |
| `request.memory_data[].extractor` | Extractor for the memory store data.                                                                                                                                                  | ✅                                   | `object`      |
| `request.memory_data[].extractor.type` | Type of extractor. Currently, only `jmesPath` is supported (used for `json`, `xml`, `yaml` response types).                                                                         | ✅                                   | `string`      |
| `request.memory_data[].extractor.jmes_path` | Extraction rule specified using JMESPath. Required if `extractor.type=jmesPath`.                                                                                                 | ✅ (`type=jmesPath`)              | `string`      |
| `request.memory_data[].extractor.on_nil` | Behavior when extraction fails. Options: `empty`, `null`, `error`. Defaults to `null`.                                                                                           | ❌                                   | `string`      |
| `request.store_data`         | Data to store in the internal database.                                                                                                                                                   | ❌                                   | `[]object`    |
| `request.store_data[].bucket_id` | Bucket ID for the database entry.                                                                                                                                                      | ✅                                   | `string`      |
| `request.store_data[].store_key` | Store key for the database entry.                                                                                                                                                      | ✅                                   | `string`      |
| `request.store_data[].encrypt` | Encryption settings for the database entry. Defaults to disabled.                                                                                                                       | ❌                                   | `object`      |
| `request.store_data[].encrypt.enabled` | Enable encryption for the database entry. Defaults to `false`.                                                                                                                    | ❌                                   | `boolean`     |
| `request.store_data[].encrypt.encrypt_id` | Encryption ID for the database entry. Required if `encrypt.enabled=true`.                                                                                                        | ✅ (`encrypt.enabled=true`)       | `string`      |
| `request.store_data[].extractor` | Extractor for the database entry data.                                                                                                                                                | ✅                                   | `object`      |
| `request.store_data[].extractor.type` | Type of extractor. Currently, only `jmesPath` is supported (used for `json`, `xml`, `yaml` response types).                                                                         | ✅                                   | `string`      |
| `request.store_data[].extractor.jmes_path` | Extraction rule specified using JMESPath. Required if `extractor.type=jmesPath`.                                                                                                 | ✅ (`type=jmesPath`)              | `string`      |
| `request.store_data[].extractor.on_nil` | Behavior when extraction fails. Options: `empty`, `null`, `error`. Defaults to `null`.                                                                                           | ❌                                   | `string`      |

### Sample

#### HTTP Query Sample 

{% raw %}
``` yaml
kind: OneExecute
type: http
output:
  enabled: true
  ids: 
    - outputLocalCSV
output:
  enabled: false
auth:
  enabled: true
  auth_id: authForWeb
request:
  target_id: "testServer"
  endpoint: "/todos/{todo_id}"
  method: GET
  query_param: {}
  path_variables:
    todo_id: "1"
  response_type: "json"
  data:
    - key: "UserID"
      extractor:
        type: "jmesPath"
        jmes_path: "userId"
        # on_nil: "error" # error or null(default) or empty
    - key: "Title"
      extractor:
        type: "jmesPath"
        jmes_path: "title"
    - key: "Completed"
      extractor:
        type: "jmesPath"
        jmes_path: "completed"
  memory_data:
    - key: "userId"
      extractor:
        type: "jmesPath"
        jmes_path: "userId"
        on_nil: "error" # error or null(default) or empty
  store_data:
    - bucket_id: "bucketForApp"
      store_key: "user"
      extractor:
        type: "jmesPath"
        jmes_path: "@"
        on_nil: "error" # error or null(default) or empty
      encrypt:
        enabled: true
        encrypt_id: encryptDynamicCBC
```
{% endraw %}

#### HTTP Command Sample 

{% raw %}
``` yaml
kind: OneExecute
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
type: http
output:
  enabled: true
  ids: 
    - outputLocalCSV
output:
  enabled: false
auth:
  enabled: true
  auth_id: authForWeb
request:
  target_id: "testServer"
  endpoint: "/posts"
  method: POST
  query_param: {}
  path_variables: {}
  body_type: "json"
  body:
    i: {{ .Dynamic.LoopCount }}
{{- if .ThreadValues.usersList }}
{{- with index .ThreadValues.usersList .Dynamic.LoopCount }}
    userId: "{{ .uuid }}"
    name: "{{ .name }}"
    email: "{{ .email }}"
{{- end }}
{{- end }}
  response_type: "json"
  data:
    - key: "ID"
      extractor:
        type: "jmesPath"
        jmes_path: "id"
        # on_nil: "error" # error or null(default) or empty
    - key: "UserId"
      extractor:
        type: "jmesPath"
        jmes_path: "userId"
        # on_nil: "error" # error or null(default) or empty
    - key: "Name"
      extractor:
        type: "jmesPath"
        jmes_path: "name"
    - key: "Email"
      extractor:
        type: "jmesPath"
        jmes_path: "email"
    - key: "Body"
      extractor:
        type: "jmesPath"
        jmes_path: "@"
```
{% endraw %}