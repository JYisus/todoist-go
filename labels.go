package gotodoist

import (
	"context"
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

func (c *Client) GetLabels(ctx context.Context) ([]Label, error) {
	return doGetRequest[[]Label](
		ctx,
		http.DefaultClient,
		c.apiToken,
		fmt.Sprintf("%s/rest/v2/labels", _apiEndpoint),
		nil,
	)
}

func (c *Client) GetLabel(ctx context.Context, id string) (*Label, error) {
	return doGetRequest[*Label](
		ctx,
		http.DefaultClient,
		c.apiToken,
		fmt.Sprintf("%s/rest/v2/labels/%s", _apiEndpoint, id),
		nil,
	)
}

func (c *Client) GetSharedLabels(ctx context.Context) ([]string, error) {
	return doGetRequest[[]string](
		ctx,
		http.DefaultClient,
		c.apiToken,
		fmt.Sprintf("%s/rest/v2/labels/shared", _apiEndpoint),
		nil,
	)
}
