// Called by the main and

package requests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func SendRequest(url, method string, authToken string) []byte  {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		//return nil, fmt.Errorf("Got error %s", err.Error())
	}

	req.Header.Set("user-agent", "golang application")
	req.Header.Add("Authorization", fmt.Sprintf("%s%s", "Bearer ", authToken))

	response, err := client.Do(req)
	if err != nil {
		//return fmt.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	return body
}
