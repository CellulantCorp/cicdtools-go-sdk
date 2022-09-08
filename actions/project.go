// retrieving group name from group ID
// retrieveing projecct name from actions ID

package actions

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gitlab.com/ronmachoka/project-lookup/requests"
)

func GetProjectName(projectID int, authToken string) string {
	res := requests.SendRequest(fmt.Sprintf("%s%s%s", GitlabUrl, "api/v4/projects/", strconv.Itoa(projectID)), "GET", authToken)
	var payload interface{}

	json.Unmarshal(res, &payload)         // Convert JSON data into interface{} type
	m := payload.(map[string]interface{}) // To use the converted data,  convert it into a map[string]interface{}

	return fmt.Sprint(m["name"])
}
