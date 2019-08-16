package sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type AddigyClient struct {
	ClientID     string
	ClientSecret string
	BaseURL		 string
}

func NewAddigyClient(clientID string, clientSecret string, realm string) *AddigyClient {
	return &AddigyClient{
		ClientID: clientID,
		ClientSecret: clientSecret,
		BaseURL: fmt.Sprintf("https://%s.addigy.com", realm),
	}
}

func (addigy AddigyClient) do (req *http.Request, responseObj interface{}) error {
	req.Header.Add("client-id", addigy.ClientID)
	req.Header.Add("client-secret", addigy.ClientSecret)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// Handle error from client performing HTTP request.
		return fmt.Errorf("error occurred performing HTTP request: %s", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// Handler error from reading response.
		return fmt.Errorf("error occurred reading response body: %s", err)
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("response from server: %s", string(body[:]))
	}

	if responseObj != nil {
		err = json.Unmarshal(body, &responseObj)
		if err != nil {
			// Handle error from unmarshalling.
			return fmt.Errorf("error occurred unmarshalling response body: %s", err)
		}
	}

	return nil
}

func (addigy AddigyClient) buildURL (endpoint string, params map[string]interface{}) string {
	paramString := "?"
	for k, v := range params {
		paramString += fmt.Sprintf("%s=%v&", k, v)
	}
	paramString = strings.TrimRight(paramString, "&")

	return addigy.BaseURL + endpoint + paramString
}