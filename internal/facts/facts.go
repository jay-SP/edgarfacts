package facts

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Load Company Facts
func LoadFacts(cik, name, organization, email string) ([]byte, error) {
	//Define URL for API
	url := fmt.Sprintf("http://data.sec.gov/api/xbrl/companyfacts/CIK%s.json", cik)

	//Create Client
	client := &http.Client{}

	// Prepare Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	//Set Custom User-Agent
	userAgent := fmt.Sprintf("%s %s %s", organization, name, email)
	req.Header.Set("User-Agent", userAgent)

	//Make Request
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	//Check response code
	if response.StatusCode != http.StatusOK {
		errorStatus := errors.New(fmt.Sprintf("Status Code != OK: %v", response.StatusCode))
		return nil, errorStatus
	}

	//Read Response Body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	//Return Final Response
	return body, nil

}
