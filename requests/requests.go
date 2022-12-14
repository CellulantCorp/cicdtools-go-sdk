// You can access the headers and body of the requests' response using the following code:
//		headers ,req := requests.SendGetRequest("https://thanos.cellulant.africa/api/v4/projects/2773")  	// Send the request
//		var Payload interface{}																			 	// Empty Payload interface
//		json.Unmarshal(req, &Payload)         	   															// Convert JSON data into interface{} type
//		m := Payload.(map[string]interface{}) 																// To use the converted data,  convert it into a map[string]interface{}
//		fmt.Println(m["name"])

package requests

import (
	"fmt"
	"gitlab-scanner-go/config"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

func SendGetRequest(url string) (*http.Response, []byte) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//return nil, fmt.Errorf("Got error %s", err.Error())
	}

	req.Header.Set("user-agent", "GitlabScannerAPI")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Keep-Alive", "timeout=30, max=60")
	req.Header.Add("Authorization", fmt.Sprintf("%s%s", "Bearer ", config.GitlabauthToken))

	response, err := client.Do(req)
	if err != nil {
		//return fmt.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()
	CheckRateLimit(response)

	body, err := ioutil.ReadAll(response.Body)
	// return response.Header, body
	return response, body
}

func SendPostRequest(url string, requestBody io.Reader) (*http.Response, []byte) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		//return nil, fmt.Errorf("Got error %s", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("user-agent", "GitlabScannerAPI")
	req.Header.Set("Accept", "*/*")
	//req.Header.Set("Keep-Alive", "timeout=30, max=60")
	req.Header.Add("Authorization", fmt.Sprintf("%s%s", "Bearer ", config.GitlabauthToken))

	response, err := client.Do(req)
	if err != nil {
		//return fmt.Errorf("Got error %s", err.Error())
		fmt.Println(" Error: ", err.Error())
		return nil, nil
	}
	defer response.Body.Close()
	CheckRateLimit(response)

	body, err := ioutil.ReadAll(response.Body)

	// Check for 200 response code before returning the body
	statusOK := response.StatusCode >= 200 && response.StatusCode < 300
	if statusOK {
		return response, body
	} else {
		logrus.Errorln(response.StatusCode, " - Error sending POST request. ", string(body))
		return response, nil
	}

}

func SendPutRequest(url string, requestBody io.Reader) (*http.Response, []byte) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("PUT", url, requestBody)
	if err != nil {
		//return nil, fmt.Errorf("Got error %s", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("user-agent", "GitlabScannerAPI")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Keep-Alive", "timeout=30, max=60")
	req.Header.Add("Authorization", fmt.Sprintf("%s%s", "Bearer ", config.GitlabauthToken))

	response, err := client.Do(req)
	if err != nil {
		//return fmt.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()
	CheckRateLimit(response)

	body, err := ioutil.ReadAll(response.Body)
// 	return response, body
	
	// Check for 200 response code before returning the body
	statusOK := response.StatusCode >= 200 && response.StatusCode < 300
	if statusOK {
		return response, body
	} else {
		// retry request
		http.NewRequest("PUT", url, requestBody)

		logrus.Errorln(response.StatusCode, " - Error sending POST request. ", string(body))
		return response, nil
	}
	
}

func SendDeleteRequest(url string, requestBody io.Reader) (*http.Response, []byte) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("DELETE", url, requestBody)
	if err != nil {
		//return nil, fmt.Errorf("Got error %s", err.Error())
	}

	req.Header.Set("user-agent", "GitlabScannerAPI")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Keep-Alive", "timeout=30, max=60")
	req.Header.Add("Authorization", fmt.Sprintf("%s%s", "Bearer ", config.GitlabauthToken))

	response, err := client.Do(req)
	if err != nil {
		//return fmt.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()
	CheckRateLimit(response)

	body, err := ioutil.ReadAll(response.Body)
	return response, body
}

// CheckRateLimit - Check for a 429 response code - which indicates a Rate Limit exceeded,
// then sleeps until the rate limit is reset
func CheckRateLimit(Response *http.Response) {
	if Response.StatusCode == 429 {
		SleepPeriod, _ := strconv.Atoi(Response.Header["Retry-After"][0])
		fmt.Println("Rate Limit Exceeded", "Backing Off for ", SleepPeriod, " seconds")
		time.Sleep((20 + (time.Duration(SleepPeriod))) * time.Second)
		fmt.Println("Attempting to Resume after rate limit ... ")
	}
}
