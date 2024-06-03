package gotodoist

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type Comment struct {
	Attachment *CommentAttachment `json:"attachment"`
	Content    string             `json:"content"`
	ID         string             `json:"id"`
	PostedAt   string             `json:"posted_at"`
	ProjectID  *string            `json:"project_id"`
	TaskID     string             `json:"task_id"`
}

type CommentAttachment struct {
	FileName     string `json:"file_name"`
	FileType     string `json:"file_type"`
	FileURL      string `json:"file_url"`
	ResourceType string `json:"resource_type"`
}

type GetCommentsOpts struct {
	TaskID    string `url:"task_id,omitempty"`
	ProjectID string `url:"project_id,omitempty"`
}

func (o GetCommentsOpts) Validate() error {
	if o.TaskID != "" && o.ProjectID != "" {
		return errors.New("either task_id or project_id should be set, not both")
	}

	if o.TaskID == "" && o.ProjectID == "" {
		return errors.New("no id to filter notes provided")
	}

	return nil
}

func (c *Client) GetComments(opts GetCommentsOpts) ([]Comment, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	queryParams, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	return doGetRequest[[]Comment](
		http.DefaultClient,
		c.apiToken,
		fmt.Sprintf("%s/rest/v2/comments", _api_endpoint),
		httpRequestOptions{QueryParams: queryParams},
	)
}
