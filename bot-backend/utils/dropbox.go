package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type DropboxTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type DropboxClient struct {
	accessToken  *string
	appKey       string
	appSecret    string
	refreshToken string
	baseUrl      string
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

type DropboxFileRequestCreateParams struct {
	Destination string `json:"destination"`
	Open        bool   `json:"open"`
	Title       string `json:"title"`
}

type DropboxFileRequestCreateResponse struct {
	Created  time.Time `json:"created"`
	Deadline struct {
		AllowLateUploads struct {
			Tag string `json:".tag"`
		} `json:"allow_late_uploads"`
		Deadline time.Time `json:"deadline"`
	} `json:"deadline"`
	Description string `json:"description"`
	Destination string `json:"destination"`
	FileCount   int    `json:"file_count"`
	ID          string `json:"id"`
	IsOpen      bool   `json:"is_open"`
	Title       string `json:"title"`
	URL         string `json:"url"`
}

type DropboxFileRequests []DropboxFileRequest

type DropboxFileRequestListResponse struct {
	FileRequests DropboxFileRequests `json:"file_requests"`
}

func (dbx *DropboxClient) GetDropboxAccessToken() (*string, error) {
	if dbx.accessToken == nil {
		// Get access token from Dropbox
		data := url.Values{}
		data.Set("refresh_token", dbx.refreshToken)
		data.Set("grant_type", "refresh_token")

		req, err := http.NewRequest("POST", "https://api.dropbox.com/oauth2/token", strings.NewReader(data.Encode()))
		if err != nil {
			return nil, err
		}

		// set headers
		req.SetBasicAuth(dbx.appKey, dbx.appSecret)

		client := http.Client{
			Timeout: time.Second * 30,
		}

		res, err2 := client.Do(req)
		if err2 != nil {
			log.Printf("Unable to get accesss token from Dropbox: %s\n", err2.Error())
			return nil, err2
		}
		defer res.Body.Close()

		var parsedResponse DropboxTokenResponse
		err = json.NewDecoder(res.Body).Decode(&parsedResponse)
		if err != nil {
			log.Printf("Unable to parse accesss token from Dropbox: %s\n", err.Error())
		}

		dbx.accessToken = &parsedResponse.AccessToken
		return dbx.accessToken, nil
	}
	return dbx.accessToken, nil
}

func (dbx *DropboxClient) DoPostRequest(path string, body any, withAccessToken bool) (*http.Response, error) {
	parsed, err := json.Marshal(body)
	if err != nil {
		log.Println("Unable to parse body for DoPostRequest")
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", dbx.baseUrl, path), bytes.NewBuffer(parsed))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	if withAccessToken {
		accessToken, err := dbx.GetDropboxAccessToken()
		if err != nil {
			return nil, err
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *accessToken))
	}

	client := http.Client{
		Timeout: time.Second * 30,
	}

	return client.Do(req)
}

func (dbx *DropboxClient) GetFileRequests() (*DropboxFileRequests, error) {
	res, err := dbx.DoPostRequest("/file_requests/list", nil, true)
	if err != nil {
		log.Printf("Unable to get file requests: %s\n", err.Error())
		return nil, err
	}
	defer res.Body.Close()

	var parsedResponse DropboxFileRequestListResponse
	err = json.NewDecoder(res.Body).Decode(&parsedResponse)
	if err != nil {
		log.Printf("Unable to decode response to get file requests: %s\n", err.Error())
		return nil, err
	}

	return &parsedResponse.FileRequests, nil
}

func (dbx *DropboxClient) CreateFileRequest(name string) (*DropboxFileRequestCreateResponse, error) {
	params := DropboxFileRequestCreateParams{}
	params.Destination = fmt.Sprintf("/file_requests/%s", name)
	params.Title = name
	params.Open = true
	res, err := dbx.DoPostRequest("/file_requests/create", params, true)
	if err != nil {
		log.Printf("Unable to create file request: %s\n", err.Error())
		return nil, err
	}
	defer res.Body.Close()

	var parsedResponse DropboxFileRequestCreateResponse
	err = json.NewDecoder(res.Body).Decode(&parsedResponse)
	if err != nil {
		log.Printf("Unable to decode response for creating file request: %s\n", err.Error())
		return nil, err
	}

	return &parsedResponse, nil
}

func NewDropboxClient(appKey, appSecret, refreshToken string) *DropboxClient {
	return &DropboxClient{
		accessToken:  nil,
		appKey:       appKey,
		appSecret:    appSecret,
		refreshToken: refreshToken,
		baseUrl:      "https://api.dropboxapi.com/2",
	}
}
