package todoist

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
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/rest/v2/labels", _apiEndpoint), nil)
	if err != nil {
		return nil, err
	}

	var labels []Label

	if err := c.doRequest(ctx, req, labels); err != nil {
		return nil, err
	}

	return labels, nil
}

func (c *Client) GetLabel(ctx context.Context, id string) (*Label, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/rest/v2/labels/%s", _apiEndpoint, id), nil)
	if err != nil {
		return nil, err
	}

	var label *Label

	if err := c.doRequest(ctx, req, label); err != nil {
		return nil, err
	}

	return label, nil
}

type CreateLabelOpts struct {
	Color      string `json:"color,omitempty"`
	IsFavorite bool   `json:"is_favorite,omitempty"`
	Name       string `json:"name"`
	Order      int    `json:"order,omitempty"`
}

func (c *Client) CreateLabel(ctx context.Context, opts *CreateLabelOpts) (*Label, error) {
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("%s/rest/v2/labels", _apiEndpoint), opts)
	if err != nil {
		return nil, err
	}

	var label *Label

	if err := c.doRequest(ctx, req, label); err != nil {
		return nil, err
	}

	return label, nil
}

type UpdateLabelOpts struct {
	Color      string `json:"color,omitempty"`
	IsFavorite bool   `json:"is_favorite,omitempty"`
	Name       string `json:"name"`
	Order      int    `json:"order,omitempty"`
}

func (c *Client) UpdateLabel(ctx context.Context, id string, opts *UpdateLabelOpts) (*Label, error) {
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("%s/rest/v2/labels/%s", _apiEndpoint, id), opts)
	if err != nil {
		return nil, err
	}

	var label *Label

	if err := c.doRequest(ctx, req, label); err != nil {
		return nil, err
	}

	return label, nil
}

func (c *Client) DeleteLabel(ctx context.Context, id string) error {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("%s/rest/v2/labels/%s", _apiEndpoint, id), nil)
	if err != nil {
		return err
	}

	if err := c.doRequest(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) GetSharedLabels(ctx context.Context) ([]string, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/rest/v2/labels/shared", _apiEndpoint), nil)
	if err != nil {
		return nil, err
	}

	var labels []string

	if err := c.doRequest(ctx, req, labels); err != nil {
		return nil, err
	}

	return labels, nil
}

type RenameSharedLabelOpts struct {
	Name    string `json:"name"`
	NewName string `json:"new_name"`
}

func (c *Client) RenameSharedLabel(ctx context.Context, opts RenameSharedLabelOpts) error {
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("%s/rest/v2/labels/shared/rename", _apiEndpoint), opts)
	if err != nil {
		return err
	}

	if err := c.doRequest(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

type RemoveSharedLabelOpts struct {
	Name string `json:"name"`
}

func (c *Client) RemoveSharedLabel(ctx context.Context, opts RenameSharedLabelOpts) error {
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("%s/rest/v2/labels/shared/remove", _apiEndpoint), opts)
	if err != nil {
		return err
	}

	if err := c.doRequest(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
