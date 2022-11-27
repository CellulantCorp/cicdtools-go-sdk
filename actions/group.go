package actions

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/CellulantCorp/cicdtools-go-sdk/requests"
	"github.com/sirupsen/logrus"
)

func GetGroupName(groupID int, authToken string) string {
	_, res := requests.SendGetRequest(fmt.Sprintf("%s%s%s", config.GitlabUrl, "api/v4/groups/", strconv.Itoa(groupID)))
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
	logrus.Infoln("Group: ", group)
	return group
}

func ToggleCICD(projectID int, enabled bool) {
	if enabled {
		// Enable CICD
		reqBody := strings.NewReader("{\"jobs_enabled\": true}")
		requests.SendPutRequest(config.GitlabUrl+"api/"+config.GitlabapiVersion+"/projects/"+strconv.Itoa(projectID), reqBody)
		fmt.Println("CI/CD Enabled for project ", projectID)

	} else {
		// Disable CICD
		reqBody := strings.NewReader("{\"jobs_enabled\": false}")
		requests.SendPutRequest(config.GitlabUrl+"api/"+config.GitlabapiVersion+"/projects/"+strconv.Itoa(projectID), reqBody)
		fmt.Println("CI/CD Disabled for project ", projectID)
	}

}
