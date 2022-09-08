package actions

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gitlab.com/ronmachoka/project-lookup/requests"
)

var GroupIds []int
var GitlabUrl = "https://thanos.cellulant.africa/"

func GetGroupName(groupID int, authToken string) string {
	res := requests.SendRequest(fmt.Sprintf("%s%s%s", GitlabUrl, "api/v4/groups/", strconv.Itoa(groupID)), "GET", authToken)
	var payload interface{}

	json.Unmarshal(res, &payload)         // Convert JSON data into interface{} type
	m := payload.(map[string]interface{}) // To use the converted data,  convert it into a map[string]interface{}

	return fmt.Sprint(m["name"])
}

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
		}
	}
	return exists
}
