## Example Guideline 📑

> Here is an example of a benchmark test using `Bloader`. These are just examples, but should be sufficient for the use of this tool. Create a scenario test that works for you, referring to the one closest to your use case.

### Scenario

| Dir | Description | Kind to use |
|:---:|:------------|:------------|
| [with_slave](#with-slave) | This is a use case test using Slave, which is a load test while measuring metrics. | `MassExecute`, `OneExecute`, `Flow`, `StoreValue`, `MemoryValue`, `StoreImport`, `SlaveConnect` |

#### With Slave

It is defined using many Kinds, and as for Slave, it is set up using `terraform`, `AWS`.

0. `Go to project route`
    ```bash
    cd example/with_slave
    ```
1. `Initialization of terraform`:
    ```bash
    terraform -chdir=slave_setup/ec2 init
    ```
2. `Apply terraform`
    ```bash
    terraform -chdir=slave_setup/ec2 apply
    ```
3. Rewriting the address of SLAVE
   Change the definition in `with_slave/runner/slave/memory.yaml` to the confirmed IP address.
4. `SSH Connect`(Optional)
    Save the IP address output when `terraform apply` is done to `$IP_ADDRESS`.
    ```bash
    ssh -i ssh_keys/slave.id_rsa ec2-user@${IP_ADDRESS}
    ```

> [!NOTE]  
> For actual testing, change [bloader.yaml](with_slave/bloader.yaml) and each request definition to your own, and check [scenario execution example](with_slave/runner/CASE.md).

##### Destroy

```bash
terraform -chdir=slave_setup/ec2 destroy
```

