package gotodoist

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

const (
	_apiEndpoint = "https://api.todoist.com"
)

type Client struct {
	apiToken   string
	httpClient *http.Client
}

func NewClient(apiToken string) *Client {
	return &Client{
		apiToken:   apiToken,
		httpClient: http.DefaultClient,
	}
}

func (c Client) newRequest(method, endpoint string, body any) (*http.Request, error) {
	var bodyBuffer io.ReadWriter
	if body != nil {
		bodyBuffer = &bytes.Buffer{}
		if err := json.NewEncoder(bodyBuffer).Encode(body); err != nil {
			return nil, fmt.Errorf("parsing body: %w", err)
		}
	}

	req, err := http.NewRequest(method, endpoint, bodyBuffer)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

func (c Client) setQueryParams(req *http.Request, opts any) error {
	queryParams, err := query.Values(opts)
	if err != nil {
		return err
	}

	req.URL.RawQuery = queryParams.Encode()

	return nil
}

func (c Client) doRequest(ctx context.Context, req *http.Request, v any) error {
	req = req.WithContext(ctx)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		return errors.New(res.Status)
	}

	if res.StatusCode == http.StatusNoContent {
		return nil
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("unknown status code: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &v); err != nil {
		return err
	}

	return nil
}

type httpRequestOptions struct {
	QueryParams url.Values
}

type requestBuilder struct {
	request *http.Request
}

func withAPIToken(apiToken string) requestOptsFunc {
	return func(req *http.Request) {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	}
}

func withQueryParams(queryParams url.Values) requestOptsFunc {
	return func(req *http.Request) {
		req.URL.RawQuery = queryParams.Encode()
	}
}

func withContext(ctx context.Context) requestOptsFunc {
	return func(req *http.Request) {
		req.WithContext(ctx)
	}
}

type requestOptsFunc func(req *http.Request)

func newRequestLegacy(method, endpoint string, optsFunc ...requestOptsFunc) (*http.Request, error) {
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return nil, err
	}

	for _, optFunc := range optsFunc {
		optFunc(req)
	}

	return req, nil
}

func doGetRequest[T any](
	ctx context.Context,
	httpClient *http.Client,
	apiToken,
	endpoint string,
	opts *httpRequestOptions,
) (T, error) {
	var resBody T

	req, err := newRequestLegacy(
		http.MethodGet,
		endpoint,
		withAPIToken(apiToken),
		withContext(ctx),
		withQueryParams(opts.QueryParams),
	)

	// req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return resBody, err
	}

	req.WithContext(ctx)

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

func doPostRequest(
	ctx context.Context,
	httpClient *http.Client,
	apiToken,
	endpoint string,
	opts *httpRequestOptions,
) error {
	req, err := newRequestLegacy(
		http.MethodPost,
		endpoint,
		withAPIToken(apiToken),
		withContext(ctx),
		withQueryParams(opts.QueryParams),
	)

	// req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return err
	}

	req.WithContext(ctx)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiToken))

	req.URL.RawQuery = opts.QueryParams.Encode()

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// body, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	return err
	// }

	// if err := json.Unmarshal(body, &resBody); err != nil {
	// 	return err
	// }

	return nil
}
