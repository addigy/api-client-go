package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// GET https://file-manager-dev.addigy.com/api/upload/url

func (addigy AddigyClient) GetFileUploadURL() (*string, error) {
	endpoint := "https://file-manager-dev.addigy.com/api/upload/url"
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		// Handle error from creating new request.
		return nil, fmt.Errorf("error occurred creating new request: %s", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("client-id", addigy.ClientID)
	req.Header.Add("client-secret", addigy.ClientSecret)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// Handle error from client performing HTTP request.
		return nil, fmt.Errorf("error occurred performing HTTP request: %s", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// Handler error from reading response.
		return nil, fmt.Errorf("error occurred reading response body: %s", err)
	}

	var url *string
	err = json.Unmarshal(body, &url)
	if err != nil {
		// Handle error from unmarshalling.
		return nil, fmt.Errorf("error occurred unmarshalling response body: %s", err)
	}

	return url, nil
}

// POST https://file-manager-dev.addigy.com/_ah/upload/#

func (addigy AddigyClient) UploadFile(uploadURL string, filePath string) (*Download, error) {
	file, err := os.Open(filePath)
	if err != nil {
		// Handle error from opening file.
		return nil, fmt.Errorf("error occurred opening file: %s", err)
	}

	defer file.Close()
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error occurred reading file: %s", err)
	}

	fi, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("error occurred getting file info: %s", err)
	}

	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	part, err := writer.CreateFormFile("file", fi.Name())
	if err != nil {
		return nil, fmt.Errorf("error occurred creating multipart form file: %s", err)
	}

	_, err = part.Write(fileContents)
	writer.Close()
	req, err := http.NewRequest("POST", uploadURL, buf)
	if err != nil {
		// Handle error from creating new request.
		return nil, fmt.Errorf("error occurred creating new request: %s", err)
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("client-id", addigy.ClientID)
	req.Header.Add("client-secret", addigy.ClientSecret)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// Handle error from client performing HTTP request.
		return nil, fmt.Errorf("error occurred performing HTTP request: %s", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// Handler error from reading response.
		return nil, fmt.Errorf("error occurred reading response body: %s", err)
	}

	var download *Download
	err = json.Unmarshal(body, &download)
	if err != nil {
		// Handle error from unmarshalling.
		return nil, fmt.Errorf("error occurred unmarshalling response body: %s", err)
	}

	return download, nil
}