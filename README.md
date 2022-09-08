## Project-Lookup

A simple reusable package which does the following :

#### 1. Check whether a project is in within the list of monitored projects

``` CheckGroup(monitored_groups_ids, group_ID)  ``` <br>
       returns true / false <br>
        - monitored_groups_ids : []int <br>
        - group_ID              : int <br>

#### 2. Get a group name

``` actions.GetProjectName(project_ID, "<git-token>") ```

#### 3. Get a group's name

``` actions.GetGroupName(group_ID, "<git-token>") ```
