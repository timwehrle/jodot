package api

import (
	"fmt"

	"github.com/timwehrle/alaric/internal/workspace"
)

type Task struct {
	GID          string        `json:"gid"`
	Name         string        `json:"name"`
	DueOn        string        `json:"due_on"`
	CreatedBy    User          `json:"created_by"`
	HtmlNotes    string        `json:"html_notes"`
	Notes        string        `json:"notes"`
	Assignee     User          `json:"assignee"`
	Tags         []Tag         `json:"tags"`
	PermaLink    string        `json:"permalink_url"`
	CustomFields []CustomField `json:"custom_fields"`
	Projects     []Project     `json:"projects"`
}

func (c *Client) GetTasks() ([]Task, error) {
	workspaceGID, _, err := workspace.LoadDefaultWorkspace()
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("/tasks?workspace=%s", workspaceGID)
	endpoint += "&opt_fields=due_on,name,completed"
	endpoint += "&completed_since=now"
	endpoint += "&assignee=me"

	resp, err := c.makeRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Data []Task `json:"data"`
	}

	if err := handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}