package gotodoist

import (
	"fmt"
	"net/http"
)

type Label struct {
	Color      string `json:"color"`
	ID         string `json:"id"`
	IsFavorite bool   `json:"is_favorite"`
	Name       string `json:"name"`
	Order      int    `json:"order"`
}

func (c *Client) GetLabels() ([]Label, error) {
	return doGetRequest[[]Label](
		http.DefaultClient,
		c.apiToken,
		fmt.Sprintf("%s/rest/v2/labels", _api_endpoint),
		httpRequestOptions{},
	)
}

func (c *Client) GetLabel(id string) (*Label, error) {
	return doGetRequest[*Label](
		http.DefaultClient,
		c.apiToken,
		fmt.Sprintf("%s/rest/v2/labels/%s", _api_endpoint, id),
		httpRequestOptions{},
	)
}

func (c *Client) GetSharedLabels() ([]string, error) {
	return doGetRequest[[]string](
		http.DefaultClient,
		c.apiToken,
		fmt.Sprintf("%s/rest/v2/labels/shared", _api_endpoint),
		httpRequestOptions{},
	)
}
