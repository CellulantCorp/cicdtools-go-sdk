// Obtain the slice of group ids
// Iterate over the slice of group ids while getting the group name and path from gitlab
// return a bool that indicates whether the group is among the ones being monitored or not

package main

import (
	"fmt"
)

var GroupIds []int


// Get group ids and store them in the global variable GroupIds
func retrieveGroupIds(groups []int) {
	for _, group := range groups {
		GroupIds = append(GroupIds, group)
	}
}

// Check whether the group is among the ones being monitored or not
// If it is, return true

func CheckGroup(groups []int, groupID int) bool {
	retrieveGroupIds(groups)
	exists := false
	for _, x := range GroupIds {
		if x == groupID {
			exists = true
			fmt.Println("Group is in list of monitored groups")
		}
	}
	return exists
}
