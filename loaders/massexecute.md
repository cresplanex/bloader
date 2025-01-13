---
title: Mass Execute
parent: Loaders
nav_order: 5
---

# Mass Execute
> It is possible to send a large number of requests simultaneously and to test the behavior under high loads. It is also very flexible in terms of termination conditions, etc., allowing scenario testing that meets your requirements.

### Property

#### General Settings

| **Field**               | **Description**                                                                                                      | **Required** | **Type**      |
|-------------------------|----------------------------------------------------------------------------------------------------------------------|--------------|---------------|
| `type`                 | Type of execution target. Currently, only `http` is supported.                                                       | ✅           | `string`      |
| `output`               | Output settings for execution. Default is disabled.                                                                 | ❌           | `object`      |
| `output.enabled`       | Enable output settings for execution. Default is `false`.                                                            | ❌           | `boolean`     |
| `output.ids`           | Outputs to enable. Defaults to an empty array.                                                                       | ❌           | `[]string`    |
| `auth`                 | Authentication settings. Default is disabled.                                                                        | ❌           | `object`      |
| `auth.enabled`         | Enable authentication settings. Default is `false`.                                                                  | ❌           | `boolean`     |
| `auth.auth_id`         | Specify the authentication ID to enable. If not specified, the default enabled authentication is used.               | ❌           | `string`      |

---

#### Requests Settings

| **Field**               | **Description**                                                                                                      | **Required**                        | **Type**           |
|-------------------------|----------------------------------------------------------------------------------------------------------------------|--------------------------------------|--------------------|
| `requests`             | Requests to be sent. Multiple requests can be executed concurrently.                                                 | ✅ (`type=http`)                  | `[]object`         |
| `requests[].target_id` | Target ID for the request.                                                                                            | ✅                                  | `string`           |
| `requests[].endpoint`  | Endpoint to append to the target. Placeholders like `{var}` can use values from `path_variables`.                     | ✅                                  | `string`           |
| `requests[].method`    | HTTP method for the request. Supports `OPTIONS`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `TRACE`, `CONNECT`.          | ✅                                  | `string`           |
| `requests[].headers`   | Headers for the request. Multiple values can be assigned to a single key as an array.                                 | ❌                                  | `map[string]any`   |

---

#### Body and Response

| **Field**                 | **Description**                                                                                                   | **Required**                | **Type**       |
|---------------------------|-------------------------------------------------------------------------------------------------------------------|----------------------------|----------------|
| `requests[].body_type`   | Type of request body. Valid values are `json`, `form`, and `multipart`.                                           | ✅                          | `string`      |
| `requests[].body`        | Request body. Format depends on `body_type`.                                                                      | ❌                          | `any`         |
| `requests[].response_type`| Response body type. Valid values are `json`, `xml`, `yaml`, `text`, and `html`.                                  | ✅                          | `string`      |
| `requests[].data`         | Output data settings for the response.                                                                           | ❌                          | `[]object`    |
| `requests[].data`             | List of data extraction configurations. Each configuration specifies how to extract and store data from the response. | ❌                                | `[]object`     |
| `requests[].data[].key`       | Key name for the extracted data in the output.                                                                      | ✅                                | `string`       |
| `requests[].data[].extractor` | Configuration for extracting the data.                                                                              | ✅                                | `object`       |
| `requests[].data[].extractor.type` | Type of extractor used for data extraction. Supported: `jmesPath` (for JSON, XML, YAML response types).             | ✅                                | `string`       |
| `requests[].data[].extractor.jmes_path` | JMESPath expression for extracting data from the response body.                                                   | ✅ (type is `jmesPath`)        | `string`       |
| `requests[].data[].extractor.on_nil` | Behavior when data extraction fails. Options: `empty` (use empty value), `null` (use null), `error` (terminate). | ❌ (default: `null`)              | `string`       |

---

#### Advanced Features

| **Field**                             | **Description**                                                                                          | **Required**                                  | **Type**       |
|---------------------------------------|----------------------------------------------------------------------------------------------------------|----------------------------------------------|----------------|
| `interval`                           | Interval between requests. Format: `10s`, `1s`, etc.                                                    | ✅                                            | `string`       |
| `await_prev_response`                | Whether to wait for the previous request's response before sending the next. Default is `false`.         | ❌                                            | `boolean`      |
| `break`                              | Break conditions for request termination.                                                               | ❌                                            | `object`       |
| `break.time`                       | Time-based termination condition. Format: `10s`, `1s`, etc.                                              | ❌                                            | `string`       |
| `break.count`                      | Number of requests after which the execution will stop.                                                   | ❌                                            | `int`          |
| `break.sys_error`                  | Terminate on system errors during execution. Default is `false`.                                          | ❌                                            | `boolean`      |
| `break.parse_error`                | Terminate on response parsing errors. Default is `false`.                                                 | ❌                                            | `boolean`      |
| `break.write_error`                | Terminate on data writing errors. Default is `false`.                                                     | ❌                                            | `boolean`      |
| `break.status_code`                | Terminate based on status code conditions. Each condition can include operators and values.               | ❌                                            | `[]object`     |
| `break.status_code[].id`           | Unique ID for the status code filter.                                                                     | ✅                                            | `string`       |
| `break.status_code[].op`           | Operator for the status code filter (e.g., `equals`, `notEquals`).                                        | ✅                                            | `string`       |
| `break.status_code[].value`        | Value to match the operator against for the status code filter.                                           | ✅                                            | `any`          |
| `break.response_body`              | Terminate based on response body content.                                                                | ❌                                            | `[]object`     |
| `break.response_body[].id`         | Unique ID for the response body filter.                                                                  | ✅                                            | `string`       |
| `break.response_body[].extractor`  | Extractor settings for response body filtering.                                                          | ✅                                            | `object`       |
| `break.response_body[].extractor.type` | Type of extractor. Currently supports `jmesPath` for JSON, XML, and YAML response types.                | ✅                                            | `string`       |
| `break.response_body[].extractor.jmes_path` | JMESPath for extracting data from the response body.                                                   | ✅ (extractor type is `jmesPath`)         | `string`       |
| `break.response_body[].extractor.on_nil` | Behavior when extraction fails. Options: `empty`, `null`, `error`. Default: `null`.                     | ❌                                            | `string`       |

#### Success Break Conditions

| **Field**          | **Description**                                                                                                           | **Required**   | **Type**       |
|--------------------|---------------------------------------------------------------------------------------------------------------------------|----------------|----------------|
| `success_break`   | Conditions under which the execution will be considered successful and stopped. Format: `terminateType/param1,param2` or `terminateType`    | ❌             | `[]string`     |

**Details:**
- Supported `terminateType` values: `context`, `count`, `sysError`, `createRequestError`, `parseError`, `writeError`, `responseBody`, `statusCode`, etc.
- If only `terminateType` is specified, it terminates when only the `terminateType` matches. 
- The possible types of “param” are `responseBodyWriteFilterError`, `responseBodyBreakFilterError`, `responseBody`, and `statusCode`, and the filter ID for each is specified. 
- If more than one param is specified, success is judged when any of them is matched.

---

#### Filters

| **Field**                              | **Description**                                                                                         | **Required**                                  | **Type**       |
|----------------------------------------|---------------------------------------------------------------------------------------------------------|----------------------------------------------|----------------|
| `record_exclude_filter`               | Filter for excluding specific data records from output.                                                | ❌                                            | `object`       |
| `record_exclude_filter.count`              | Filters records based on execution count.                                                              | ❌                                            | `[]object`     |
| `record_exclude_filter.count[].id`         | Unique ID for the count filter.                                                                        | ✅                                            | `string`       |
| `record_exclude_filter.count[].op`         | Operator for the count filter (e.g., `equals`, `greaterThan`). Check the [Count Filter](#count-filter) section for details.                                         | ✅                                            | `string`       |
| `record_exclude_filter.count[].value`      | Value for the count filter. Check the [Count Filter](#count-filter) section for details.                                                                              | ✅                                            | `any`          |
| `record_exclude_filter.status_code`        | Filters records based on status code.                                                                  | ❌                                            | `[]object`     |
| `record_exclude_filter.status_code[].id`   | Unique ID for the status code filter.                                                                  | ✅                                            | `string`       |
| `record_exclude_filter.status_code[].op`   | Operator for the status code filter (e.g., `equals`, `notEquals`). Check the [Status Code Filter](#status-code-filter) section for details.                                         | ✅                                            | `string`       |
| `record_exclude_filter.status_code[].value`| Value for the status code filter.  Check the [Status Code Filter](#status-code-filter) section for details.                                                                     | ✅                                            | `any`          |
| `record_exclude_filter.response_body`      | Filters records based on response body.                                                                | ❌                                            | `[]object`     |
| `record_exclude_filter.response_body[].id` | Unique ID for the response body filter.                                                                | ✅                                            | `string`       |
| `record_exclude_filter.response_body[].extractor` | Extractor settings for the response body filter.                                                      | ✅                                            | `object`       |
| `record_exclude_filter.response_body[].extractor.type` | Type of extractor (e.g., `jmesPath`). Supported for JSON, XML, YAML responses.                        | ✅                                            | `string`       |
| `record_exclude_filter.response_body[].extractor.jmes_path` | JMESPath for extracting data from the response body.                                                  | ✅ (extractor type is `jmesPath`)         | `string`       |
| `record_exclude_filter.response_body[].extractor.on_nil` | Behavior when extraction fails: `empty`, `null`, `error`. Default is `null`.                           | ❌                                            | `string`       |

### Filter

{: .info }
> Filters that require complex settings are listed below. These allow for flexible settings.

#### Count Filter
Count filters allow for various conditions to be applied to numerical values. Below are the supported filters and their behaviors:

- **`none`**: Always returns `false`.  
- **`eq`**: The value must be an `int` and returns `true` only if it equals the specified value.  
- **`ne`**: The value must be an `int` and returns `true` only if it does not equal the specified value.  
- **`lt`**: The value must be an `int` and returns `true` only if it is less than the specified value.  
- **`le`**: The value must be an `int` and returns `true` only if it is less than or equal to the specified value.  
- **`gt`**: The value must be an `int` and returns `true` only if it is greater than the specified value.  
- **`ge`**: The value must be an `int` and returns `true` only if it is greater than or equal to the specified value.  
- **`in`**: The value must be a `[]int` and returns `true` only if it matches any value in the list.  
- **`nin`**: The value must be a `[]int` and returns `true` only if it does not match any value in the list.  
- **`between`**: The value must be a `map[string]int` and returns `true` only if it lies between `key=min` and `key=max`.  
- **`notBetween`**: The value must be a `map[string]int` and returns `true` only if it does not lie between `key=min` and `key=max`.  
- **`mod`**: The value must be an `int` and returns `true` only if it is divisible by the specified value.  
- **`notMod`**: The value must be an `int` and returns `true` only if it is not divisible by the specified value.  
- **`regex`**: The value must be a `string` and returns `true` only if it matches the specified regular expression.

---

#### Status Code Filter
Status code filters operate similarly to count filters but with some limitations:  

- All filters from [Count Filters](#count-filter) are supported **except** for `mod`.  
These filters enable precise control over HTTP status code conditions.

### Sample

{% raw %}
``` yaml
kind: MassExecute
type: http
output:
  enabled: true
  ids: 
    - outputLocalCSV
auth: 
  enabled: true
  auth_id: authForWeb
requests:
{{- range $index := until 3 }}
  - target_id: "testServer"
    endpoint: "/todos/{todo_id}"
    method: GET
    query_param: {}
    path_variables:
      todo_id: "1"
    interval: 1s
    await_prev_response: false
    headers: {}
    body_type: json
    body: {}
    success_break:
      - count
      - time
      - statusCode/badRequest,internalServerError
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
      response_body:
        - id: code
          extractor:
            type: jmesPath
            jmes_path: "code == '0002:0003'"
    response_type: json
    record_exclude_filter:
      # count:
      #   - id: odd
      #     op: mod
      #     value: 2
      #   - id: initial
      #     op: lt
      #     value: 5
      status_code:
        - id: expected
          op: in
          value: [401, 403, 404]
      response_body:
        - id: "code"
          extractor:
            type: "jmesPath"
            jmes_path: "code == '0002:0000'"
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
{{- end }}
```
{% endraw %}