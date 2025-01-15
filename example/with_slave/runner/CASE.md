### Scenario

| Scenario                     | Link                              |
|------------------------------|-----------------------------------|
| Create UserProfile           | [SC1](#sc1-create-userprofile)   |
| Update UserPreference        | [SC2](#sc2-update-userpreference)|
| Create Organization          | [SC3](#sc3-create-organization)  |
| Add Organization User        | [SC4](#sc4-add-organization-user)|
| Create FileObject            | [SC5](#sc5-create-fileobject)    |
| Create Team                  | [SC6](#sc6-create-team)          |
| Add Team User                | [SC7](#sc7-add-team-user)        |
| Create Task                  | [SC8](#sc8-create-task)          |
| Update Task Status           | [SC9](#sc9-update-task-status)   |
| Find User                    | [SC10](#sc10-find-user)          |
| Get Users                    | [SC11](#sc11-get-users)          |
| Find UserPreference          | [SC12](#sc12-find-userpreference)|
| Get Users With Preference    | [SC13](#sc13-get-users-with-preference)|
| Find Team                    | [SC14](#sc14-find-team)          |
| Find Team With Users         | [SC15](#sc15-find-team-with-users)|
| Get Teams                    | [SC16](#sc16-get-teams)          |
| Get Teams With Users         | [SC17](#sc17-get-teams-with-users)|
| Find Organization            | [SC18](#sc18-find-organization)  |
| Find Organization With Users | [SC19](#sc19-find-organization-with-users)|
| Get Organizations            | [SC20](#sc20-get-organizations) |
| Get Organizations With Users | [SC21](#sc21-get-organizations-with-users)|
| Find FileObject              | [SC22](#sc22-find-fileobject)    |
| Get FileObjects              | [SC23](#sc23-get-fileobjects)    |
| Find Task                    | [SC24](#sc24-find-task)          |
| Find Task With Attachments   | [SC25](#sc25-find-task-with-attachments)|
| Get Tasks                    | [SC26](#sc26-get-tasks)          |
| Get Tasks With Attachments   | [SC27](#sc27-get-tasks-with-attachments)|

#### SC1: Create UserProfile
1. execute
    ``` sh
    bloader run -f sc1.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```
2. wait creating
3. retrieve users
    ``` sh
    bloader run -f store/store_user.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC2 Update UserPreference
1. execute
   ``` sh
   bloader run -f sc2.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
   ```

#### SC3 Create Organization
1. execute
    ``` sh
    bloader run -f sc3.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```
2. wait creating
3. retrieve users
    ``` sh
    bloader run -f store/store_organization.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC4 Add Organization User
1. execute
    ``` sh
    bloader run -f sc4.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC5 Create FileObject
1. execute
    ``` sh
    bloader run -f sc5.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```
2. wait creating
3. retrieve users
    ``` sh
    bloader run -f store/store_file_object.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC6 Create Team
1. execute
    ``` sh
    bloader run -f sc6.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```
2. wait creating
3. retrieve users
    ``` sh
    bloader run -f store/store_team.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC7 Add Team User
1. execute
    ``` sh
    bloader run -f sc7.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC8 Create Task
1. execute
    ``` sh
    bloader run -f sc8.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```
2. wait creating
3. retrieve users
    ``` sh
    bloader run -f store/store_task.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC9 Update Task Status
1. execute
    ``` sh
    bloader run -f sc9.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC10 Find User
1. execute
    ``` sh
    bloader run -f sc10.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC11 Get Users
1. execute
    ``` sh
    bloader run -f sc11.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC12 Find UserPreference
1. execute
    ``` sh
    bloader run -f sc12.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC13 Get Users With Preference
1. execute
    ``` sh
    bloader run -f sc13.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC14 Find Team
1. execute
    ``` sh
    bloader run -f sc14.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC15 Find Team With Users
1. execute
    ``` sh
    bloader run -f sc15.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC16 Get Teams
1. execute
    ``` sh
    bloader run -f sc16.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC17 Get Teams With Users
1. execute
    ``` sh
    bloader run -f sc17.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC18 Find Organization
1. execute
    ``` sh
    bloader run -f sc18.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC19 Find Organization With Users
1. execute
    ``` sh
    bloader run -f sc19.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC20 Get Organizations
1. execute
    ``` sh
    bloader run -f sc20.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC21 Get Organizations With Users
1. execute
    ``` sh
    bloader run -f sc21.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC22 Find FileObject
1. execute
    ``` sh
    bloader run -f sc22.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC23 Get FileObjects
1. execute
    ``` sh
    bloader run -f sc23.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC24 Find Task
1. execute
    ``` sh
    bloader run -f sc24.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC25 Find Task With Attachments
1. execute
    ``` sh
    bloader run -f sc25.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC26 Get Tasks
1. execute
    ``` sh
    bloader run -f sc26.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```

#### SC27 Get Tasks With Attachments
1. execute
    ``` sh
    bloader run -f sc27.yaml -d "SlaveCount=4:i" -d "ThreadPerSlaveCount=25:i" -d "RequestPerSlaveCount=3000:i"
    ```