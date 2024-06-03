package gotodoist

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	_api_endpoint = "https://api.todoist.com"
)

type Client struct {
	apiToken string
}

func NewClient(apiToken string) *Client {
	return &Client{
		apiToken: apiToken,
	}
}

func getHTTPRequest(apiToken, endpoint string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiToken))

	return req, nil
}

type httpRequestOptions struct {
	QueryParams url.Values
}

func doGetRequest[T any](httpClient *http.Client, apiToken, endpoint string, opts httpRequestOptions) (T, error) {
	var resBody T

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return resBody, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiToken))

	req.URL.RawQuery = opts.QueryParams.Encode()

	res, err := httpClient.Do(req)
	if err != nil {
		return resBody, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resBody, err
	}

	if err := json.Unmarshal(body, &resBody); err != nil {
		return resBody, err
	}

	return resBody, nil
}
