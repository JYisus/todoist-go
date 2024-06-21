package todoist

import (
	"context"
	"errors"
	"fmt"
	"net/http"
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

func (c Client) GetComments(ctx context.Context, opts GetCommentsOpts) ([]Comment, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/rest/v2/comments", _apiEndpoint), nil)
	if err != nil {
		return nil, err
	}

	if err := c.setQueryParams(req, opts); err != nil {
		return nil, err
	}

	var comments []Comment

	if err := c.doRequest(ctx, req, comments); err != nil {
		return nil, err
	}

	return comments, nil
}

func (c Client) GetComment(ctx context.Context, id string) (*Comment, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/rest/v2/comments/%s", _apiEndpoint, id), nil)
	if err != nil {
		return nil, err
	}

	var comment *Comment

	if err := c.doRequest(ctx, req, comment); err != nil {
		return nil, err
	}

	return comment, nil
}

type CreateCommentOpts struct {
	TaskID            string             `json:"task_id,omitempty"`
	ProjectID         string             `json:"project_id,omitempty"`
	Content           string             `json:"content"`
	CommentAttachment *CommentAttachment `json:"attachment,omitempty"`
}

func (o CreateCommentOpts) Validate() error {
	if o.TaskID != "" && o.ProjectID != "" {
		return errors.New("either task_id or project_id should be set, not both")
	}

	if o.TaskID == "" && o.ProjectID == "" {
		return errors.New("no id to filter notes provided")
	}

	return nil
}

func (c Client) CreateComment(ctx context.Context, opts CreateCommentOpts) (*Comment, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("%s/rest/v2/comments", _apiEndpoint), opts)
	if err != nil {
		return nil, err
	}

	var comment *Comment

	if err := c.doRequest(ctx, req, comment); err != nil {
		return nil, err
	}

	return comment, nil
}

type UpdateCommentOpts struct {
	Content string `json:"content"`
}

func (c Client) UpdateComment(ctx context.Context, id string, opts UpdateCommentOpts) (*Comment, error) {
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("%s/rest/v2/comments/%s", _apiEndpoint, id), opts)
	if err != nil {
		return nil, err
	}

	var comment *Comment

	if err := c.doRequest(ctx, req, comment); err != nil {
		return nil, err
	}

	return comment, nil
}

func (c Client) DeleteComment(ctx context.Context, id string) error {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("%s/rest/v2/comments/%s", _apiEndpoint, id), nil)
	if err != nil {
		return err
	}

	var comment *Comment

	if err := c.doRequest(ctx, req, comment); err != nil {
		return err
	}

	return nil
}
