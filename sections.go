package todoist

import (
	"context"
	"fmt"
	"net/http"
)

type Section struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Order     int    `json:"order,omitempty"`
	ProjectID string `json:"project_id"`
}

type GetSectionsOpts struct {
	ProjectID string `url:"project_id,omitempty"`
}

func (c *Client) GetSections(ctx context.Context, opts *GetSectionsOpts) ([]Section, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/rest/v2/sections", _apiEndpoint), nil)
	if err != nil {
		return nil, err
	}

	if err := c.setQueryParams(req, opts); err != nil {
		return nil, err
	}

	var sections []Section

	if err := c.doRequest(ctx, req, &sections); err != nil {
		return nil, err
	}

	return sections, nil
}

func (c *Client) GetSection(ctx context.Context, id string) (*Section, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/rest/v2/sections/%s", _apiEndpoint, id), nil)
	if err != nil {
		return nil, err
	}

	var section *Section

	if err := c.doRequest(ctx, req, &section); err != nil {
		return nil, err
	}

	return section, nil
}

type CreateSectionOpts struct {
	Name      string `json:"name"`
	ProjectID string `json:"project_id"`
	Order     string `json:"order,omitempty"`
}

func (c Client) CreateSection(ctx context.Context, opts *CreateSectionOpts) (*Section, error) {
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("%s/rest/v2/sections", _apiEndpoint), opts)
	if err != nil {
		return nil, err
	}

	var section *Section

	if err := c.doRequest(ctx, req, &section); err != nil {
		return nil, err
	}

	return section, nil
}

type UpdateSectionOpts struct {
	Name string `json:"name"`
}

func (c Client) UpdateSection(ctx context.Context, id string, opts *UpdateSectionOpts) (*Section, error) {
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("%s/rest/v2/sections/%s", _apiEndpoint, id), opts)
	if err != nil {
		return nil, err
	}

	var section *Section

	if err := c.doRequest(ctx, req, &section); err != nil {
		return nil, err
	}

	return section, nil
}

func (c Client) DeleteSection(ctx context.Context, id string) error {
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("%s/rest/v2/sections/%s", _apiEndpoint, id), nil)
	if err != nil {
		return err
	}

	if err := c.doRequest(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
