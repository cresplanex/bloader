---
title: Master Slave Communication
parent: Configuration
nav_order: 4
---

# Master and Slave Overview

### Output Handling 📤
- All output operations are handled **entirely on the Master side**. For example, when using an output type such as `type=local`, the output is **not written to the Slave's local file**. Instead, the data is **transmitted to the Master**, where it is written to the **Master's local file system**.

### Communication Security 🔒

{: .warning }
> It is recommended to read this section with caution, as it is not only about the bloader, but also about the security of the measurement target.

- The Slave **does not retain loaders or authentication information**. All necessary data is securely transmitted from the Master to the Slave **prior to starting the load test**. 
- To ensure safe communication between the Master and Slave, enabling **TLS communication** is strongly recommended. TLS provides **encryption for data transmission**, offering a secure and reliable communication channel. 