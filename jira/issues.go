package jira

import (
	"fmt"
	"net/http"
)

type Issues struct {
	c *Client
}

func (i *Issues) List(projectKey string) *http.Response {
	url := fmt.Sprintf("rest/api/2/search?jql=project=%s", projectKey)
	return i.c.execute("GET", url)
}
