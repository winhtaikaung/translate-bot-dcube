package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type DropboxClient struct {
	accessToken string
	baseUrl     string
}

type DropboxFileRequest struct {
	ID          string    `json:"id"`
	URL         string    `json:"url"`
	Title       string    `json:"title"`
	Destination string    `json:"destination"`
	Created     time.Time `json:"created"`
	IsOpen      bool      `json:"is_open"`
	FileCount   int       `json:"file_count"`
	Description string    `json:"description"`
}

type DropboxFileRequests []DropboxFileRequest

type DropboxFileRequestListResponse struct {
	FileRequests DropboxFileRequests `json:"file_requests"`
}

func (dbx *DropboxClient) DoPostRequest(path string, body any) (*http.Response, error) {
	parsed, err := json.Marshal(body)
	if err != nil {
		log.Println("Unable to parse body for DoPostRequest")
		return nil, err
	}
	return http.Post(fmt.Sprintf("%s%s", dbx.baseUrl, path), "application/json", bytes.NewBuffer(parsed))
}

func (dbx *DropboxClient) GetFileRequests() (*DropboxFileRequests, error) {
	res, err := dbx.DoPostRequest("/file_requests/list", "null")
	if err != nil {
		log.Printf("Unable to get file requests: %s\n", err.Error())
		return nil, err
	}

	var parsedResponse DropboxFileRequestListResponse
	err = json.NewDecoder(res.Body).Decode(&parsedResponse)
	if err != nil {
		log.Printf("Unable to decode response to get file requests: %s\n", err.Error())
		return nil, err
	}

	return &parsedResponse.FileRequests, nil
}

func NewDropboxClient(accessToken string) *DropboxClient {
	return &DropboxClient{
		accessToken: accessToken,
		baseUrl:     "https://api.dropboxapi.com/2",
	}
}
