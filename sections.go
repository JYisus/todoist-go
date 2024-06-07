package gotodoist

import (
	"fmt"
	"net/http"
)

type Section struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Order     int    `json:"order"`
	ProjectID string `json:"project_id"`
}

func (c *Client) GetSections() ([]Section, error) {
	return doGetRequest[[]Section](
		http.DefaultClient,
		c.apiToken,
		fmt.Sprintf("%s/rest/v2/sections", _apiEndpoint),
		nil,
	)
}

func (c *Client) GetSection(id string) (*Section, error) {
	return doGetRequest[*Section](
		http.DefaultClient,
		c.apiToken,
		fmt.Sprintf("%s/rest/v2/sections/%s", _apiEndpoint, id),
		nil,
	)
}
