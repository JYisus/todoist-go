package gotodoist

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
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

func (c *Client) GetTasks(opts *GetTasksOpts) ([]Task, error) {
	queryParams, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	fmt.Println(queryParams.Encode())

	return doGetRequest[[]Task](
		http.DefaultClient,
		c.apiToken,
		fmt.Sprintf("%s/rest/v2/tasks", _api_endpoint),
		httpRequestOptions{
			QueryParams: queryParams,
		},
	)
}

func (c *Client) GetTask(id string) (*Task, error) {
	return doGetRequest[*Task](
		http.DefaultClient,
		c.apiToken,
		fmt.Sprintf("%s/rest/v2/tasks/%s", _api_endpoint, id),
		httpRequestOptions{},
	)
}
