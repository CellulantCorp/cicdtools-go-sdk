## Project-Lookup

A simple reusable package which does the following :

#### 1. Check whether a project is in within the list of monitored projects

``` actions.CheckGroup(monitored_groups_names, group_name)  ``` <br>
       returns true / false <br>
        - monitored_groups_names : []string <br>
        - group_name              : string <br>

#### 2. Get a project's name

``` actions.GetProjectName(project_ID, "<git-token>") ```

#### 3. Get a group's name

``` actions.GetGroupName(group_ID, "<git-token>") ```
