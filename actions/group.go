package actions

import (
	"encoding/json"
	"fmt"
	"strconv"
	"thanos.cellulant.africa/cellulant-public/ci-cd-tools/packages/go/project-lookup/requests"
)

var GitlabUrl="https://thanos.cellulant.africa/"

func GetGroupName(groupID int, authToken string) string {
	res := requests.SendRequest(fmt.Sprintf("%s%s%s", GitlabUrl, "api/v4/groups/", strconv.Itoa(groupID)), "GET", authToken)
	var payload interface{}

	json.Unmarshal(res, &payload) // Convert JSON data into interface{} type
	m := payload.(map[string]interface{}) // To use the converted data,  convert it into a map[string]interface{}

	return fmt.Sprint(m["name"])
}
