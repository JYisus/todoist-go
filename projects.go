package gotodoist

import (
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

func (c *Client) GetProjects() ([]Project, error) {
	return doGetRequest[[]Project](
		http.DefaultClient,
		c.apiToken,
		fmt.Sprintf("%s/rest/v2/projects", _api_endpoint),
		httpRequestOptions{},
	)
}

func (c *Client) GetProject(id string) (*Project, error) {
	return doGetRequest[*Project](http.DefaultClient, c.apiToken, fmt.Sprintf("%s/rest/v2/projects/%s", _api_endpoint, id), httpRequestOptions{})
}
