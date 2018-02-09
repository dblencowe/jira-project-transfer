package jira

import "net/http"

type issues interface {
	List(projectKey string) *http.Response
}
