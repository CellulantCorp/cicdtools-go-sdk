package actions

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"thanos.cellulant.africa/cellulant-public/ci-cd-tools/packages/go/project-lookup/requests"
)

// var GroupIds []int
var GitlabUrl = "https://thanos.cellulant.africa/"

func GetGroupName(groupID int, authToken string) string {
	res := requests.SendRequest(fmt.Sprintf("%s%s%s", GitlabUrl, "api/v4/groups/", strconv.Itoa(groupID)), "GET", authToken)
	var payload interface{}

	json.Unmarshal(res, &payload)         // Convert JSON data into interface{} type
	m := payload.(map[string]interface{}) // To use the converted data,  convert it into a map[string]interface{}

	return fmt.Sprint(m["name"])
}

// // Get group ids and store them in the global variable GroupIds
// func retrieveGroupIds(groups []int) {
// 	for _, group := range groups {
// 		GroupIds = append(GroupIds, group)
// 	}
// }

// Check whether the group is among the ones being monitored or not
// If it is, return true

func CheckGroup(groups []string, groupNamespace string) bool {
	// retrieveGroupIds(groups)
	group := splitNamespace(groupNamespace)
	exists := false
	for _, x := range groups {
		if x == group {
			exists = true
		}
	}
	return exists
}

func splitNamespace(namespace string) string {
	// Split the namespace into group and project
	splitNamespace := strings.Split(namespace, "/")
	group := splitNamespace[1]
	return group
}
