package gotodoist

import (
	"context"
	"fmt"
	"net/http"
)

type Task struct {
	AssigneeID   string       `json:"assignee_id"`
	AssignerID   string       `json:"assigner_id"`
	CommentCount int          `json:"comment_count"`
	Content      string       `json:"content"`
	CreatedAt    string       `json:"created_at"`
	CreatorID    string       `json:"creator_id"`
	Description  string       `json:"description"`
	Due          TaskDue      `json:"due"`
	Duration     TaskDuration `json:"duration"`
	ID           string       `json:"id"`
	IsCompleted  bool         `json:"is_completed"`
	Labels       []string     `json:"labels"`
	Order        int          `json:"order"`
	ParentID     string       `json:"parent_id"`
	Priority     int          `json:"priority"`
	ProjectID    string       `json:"project_id"`
	SectionID    string       `json:"section_id"`
	URL          string       `json:"url"`
}

type TaskDue struct {
	Date        string `json:"date"`
	Datetime    string `json:"datetime"`
	IsRecurring bool   `json:"is_recurring"`
	String      string `json:"string"`
	Timezone    string `json:"timezone"`
}

type TaskDuration struct {
	Amount int    `json:"amount"`
	Unit   string `json:"unit"`
}

type GetTasksOpts struct {
	ProjectID string   `url:"project_id,omitempty"`
	SectionID string   `url:"section_id,omitempty"`
	Label     string   `url:"label,omitempty"`
	Filter    string   `url:"filter,omitempty"`
	Lang      string   `url:"lang,omitempty"`
	IDs       []string `url:"ids,comma,omitempty"`
}

func (c Client) GetTasks(ctx context.Context, opts *GetTasksOpts) ([]Task, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/rest/v2/tasks", _apiEndpoint), nil)
	if err != nil {
		return nil, err
	}

	if err := c.setQueryParams(req, opts); err != nil {
		return nil, err
	}

	var tasks []Task

	if err := c.doRequest(ctx, req, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (c *Client) GetTask(ctx context.Context, id string) (*Task, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/rest/v2/tasks/%s", _apiEndpoint, id), nil)
	if err != nil {
		return nil, err
	}

	var task *Task

	if err := c.doRequest(ctx, req, &task); err != nil {
		return nil, err
	}

	return task, nil
}

type AddTaskOpts struct {
	Content   string `json:"content"`
	ProjectID string `json:"project_id"`
}

func (c Client) AddTask(ctx context.Context, opts *AddTaskOpts) (*Task, error) {
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("%s/rest/v2/tasks", _apiEndpoint), opts)
	if err != nil {
		return nil, err
	}

	var task *Task

	if err := c.doRequest(ctx, req, &task); err != nil {
		return nil, err
	}

	return task, nil
}

func (c Client) CloseTask(ctx context.Context, taskID string) error {
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("%s/rest/v2/tasks/%s/close", _apiEndpoint, taskID), nil)
	if err != nil {
		return err
	}

	if err := c.doRequest(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (c Client) ReopenTask(ctx context.Context, taskID string) error {
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("%s/rest/v2/tasks/%s/reopen", _apiEndpoint, taskID), nil)
	if err != nil {
		return err
	}

	if err := c.doRequest(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (c Client) DeleteTask(ctx context.Context, taskID string) error {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("%s/rest/v2/tasks/%s", _apiEndpoint, taskID), nil)
	if err != nil {
		return err
	}

	if err := c.doRequest(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
