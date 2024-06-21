package todoist

import (
	"context"
	"fmt"
	"net/http"
)

type Project struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	CommentCount   int     `json:"comment_count"`
	Order          int     `json:"order"`
	Color          string  `json:"color"`
	IsShared       bool    `json:"is_shared"`
	IsFavorite     bool    `json:"is_favorite"`
	IsInboxProject bool    `json:"is_inbox_project"`
	IsTeamInbox    bool    `json:"is_team_inbox"`
	ViewStyle      string  `json:"view_style"`
	URL            string  `json:"url"`
	ParentID       *string `json:"parent_int"`
}

func (c Client) GetProjects(ctx context.Context) ([]Project, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/rest/v2/projects", _apiEndpoint), nil)
	if err != nil {
		return nil, err
	}

	var projects []Project

	if err := c.doRequest(ctx, req, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

func (c Client) GetProject(ctx context.Context, id string) (*Project, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/rest/v2/projects/%s", _apiEndpoint, id), nil)
	if err != nil {
		return nil, err
	}

	var project *Project

	if err := c.doRequest(ctx, req, &project); err != nil {
		return nil, err
	}

	return project, nil
}

type CreateObjectOpts struct {
	Color      string  `json:"color,omitempty"`
	IsFavorite bool    `json:"is_favorite,omitempty"`
	Name       string  `json:"name"`
	ParentID   *string `json:"parent_id,omitempty"`
	ViewStyle  string  `json:"view_style,omitempty"`
}

func (c Client) CreateProject(ctx context.Context, opts *CreateObjectOpts) (*Project, error) {
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("%s/rest/v2/projects", _apiEndpoint), opts)
	if err != nil {
		return nil, err
	}

	var projects *Project

	if err := c.doRequest(ctx, req, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

type UpdateObjectOpts struct {
	Color      string `json:"color,omitempty"`
	IsFavorite bool   `json:"is_favorite,omitempty"`
	Name       string `json:"name,omitempty"`
	ViewStyle  string `json:"view_style,omitempty"`
}

func (c Client) UpdateProject(ctx context.Context, id string, opts *UpdateObjectOpts) (*Project, error) {
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("%s/rest/v2/projects/id", _apiEndpoint), opts)
	if err != nil {
		return nil, err
	}

	var projects *Project

	if err := c.doRequest(ctx, req, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

func (c Client) DeleteProject(ctx context.Context, id string) error {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("%s/rest/v2/projects/id", _apiEndpoint), nil)
	if err != nil {
		return err
	}

	if err := c.doRequest(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

type Collaborator struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	Name  string `json:"name"`
}

func (c Client) GetAllCollaborators(ctx context.Context, id string) ([]Collaborator, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/rest/v2/projects/%s/collaborators", _apiEndpoint, id), nil)
	if err != nil {
		return nil, err
	}

	var collaborators []Collaborator

	if err := c.doRequest(ctx, req, &collaborators); err != nil {
		return nil, err
	}

	return collaborators, nil
}
