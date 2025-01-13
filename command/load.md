---
title: "Load Command"
parent: "Command"
nav_order: 5
---

# Load Commands

Used for load testing. This command is valid only for the Master node.

#### Run Load Test
Run load tests using loader files.

- **Interactive File Input**:
  ```bash
  bloader run
  ```
  You'll be prompted for the file path:
  ```bash
  ✔ Enter the file to run the load test:
  ```

- **Specify File Inline**:
  ```bash
  bloader run -f main.yaml
  ```

##### Pass data to memory store inline
The `--data(-d)` option can be used to pass data to the global memory store inline.

``` bash
bloader run -f main.yaml -d Username=test
```

If you set up as above, you can access the following.

{% raw %}
``` yaml
kind: OneExecute
type: http
output:
  enabled: false
auth:
  enabled: true
  auth_id: authForWeb
request:
  target_id: "testServer"
  endpoint: "/posts"
  method: POST
  body_type: "json"
  body:
    username: "{{ .Values.Username }}"
  response_type: "json"
```
{% endraw %}

However, by default, it is passed as a string. It can be passed as various types by doing the following.

```bash
bloader run -f main.yaml -d IntData=10:i -d StrData=test:s -d IntArrayData=1,2,3:ai
```

Those supported include.

| Type | Notation | Example |
|:-----:|:-------:|:--------|
| `int` | `i` | 1 |
| `string` | `s` | test |
| `bool` | `b` | true |
| `float` | `f` | 12.3 |
| `uint` | `u` | 1 |
| `[]int` | `ai` | 1,-2,3 |
| `[]string` | `as` | test1,test2 |
| `[]bool` | `ab` | true,false,true |
| `[]float` | `af` | 12.3,24.5 |
| `[]uint` | `au` | 1,2,3 |
