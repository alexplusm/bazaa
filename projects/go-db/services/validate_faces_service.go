package services

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type ValidateFacesService struct {
}

/*
	source: https://gist.github.com/andrewmilson/19185aab2347f6ad29f5
*/

func (service *ValidateFacesService) Validate(filePath string) (bool, error) {
	body, contentTypeValue, err := service.prepareBody(filePath)
	if err != nil {
		return false, err
	}

	resp, err := service.sendRequest(body, contentTypeValue)
	if err != nil {
		return false, fmt.Errorf(" Validate: %v", err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	if err := resp.Body.Close(); err != nil {
		return false, err
	}

	// TODO: processResponseBody
	fmt.Println(resp.Status, resp.Body)
	fmt.Println(string(respBody))

	return true, nil
}

func (service *ValidateFacesService) prepareBody(filePath string) (io.Reader, string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("photo", filepath.Base(file.Name()))
	if err != nil {
		return nil, "", err
	}

	if _, err := io.Copy(part, file); err != nil {
		return nil, "", err
	}
	if err := file.Close(); err != nil {
		return nil, "", err
	}
	if err := writer.Close(); err != nil {
		return nil, "", err
	}

	return body, writer.FormDataContentType(), nil
}

func (service *ValidateFacesService) sendRequest(body io.Reader, contentTypeValue string) (*http.Response, error) {
	url := os.Getenv("PARSIV_FACE_VALIDATE_URL")

	r, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("sendRequest: %v", err)
	}

	r.Header.Add("Content-Type", contentTypeValue)
	client := http.Client{}

	return client.Do(r)
}
