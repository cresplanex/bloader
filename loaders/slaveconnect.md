---
title: Slave Connect
parent: Loaders
nav_order: 6
---

# Slave Connect
> It can be used to connect to Slave, but each Slave can have its own configuration, for example, TLS can be defined separately as enabled and disabled.

### Property

| **Field**                            | **Description**                                                                                                                                                                    | **Required**                      | **Type**   |
|--------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------|------------|
| `slaves`                             | Configuration settings for slaves.                                                                                                                                                | ✅                                | `[]object` |
| `slaves[].id`                        | Unique ID for the slave.                                                                                                                                                         | ✅                                | `string`   |
| `slaves[].uri`                       | Address of the slave, specified according to [gRPC naming conventions](https://github.com/grpc/grpc/blob/master/doc/naming.md).                                                  | ✅                                | `string`   |
| `slaves[].certificate`               | TLS settings for communication with the slave. Defaults to disabled.                                                                                                             | ❌                                | `object`   |
| `slaves[].certificate.enabled`       | Enable TLS for communication with the slave. Defaults to `false`.                                                                                                                | ❌                                | `boolean`  |
| `slaves[].certificate.ca_cert`       | Path to the CA certificate used for TLS. Required if `certificate.enabled=true`.                                                                                                 | ✅ (`certificate.enabled=true`) | `string`   |
| `slaves[].certificate.server_name_override` | Override for the server name used in TLS. Required if `certificate.enabled=true`.                                                                                               | ✅ (`certificate.enabled=true`) | `string`   |
| `slaves[].certificate.insecure_skip_verify` | Skip server name verification in TLS. Defaults to `false`.                                                                                                                     | ❌                                | `boolean`  |

### Sample

{% raw %}
``` yaml
kind: SlaveConnect
slaves:
{{- range slice .Values.slaveLists 0 .Values.SlaveCount }}
  - id: "{{ .id }}"
    uri: "dns:{{ .address }}:{{ .port }}" # support schema https://github.com/grpc/grpc/blob/master/doc/naming.md
    certificate:
      enabled: true
      ca_cert: "certs/ca.crt"
      server_name_override: "localhost"
      insecure_skip_verify: false
{{- end }}
```
{% endraw %}