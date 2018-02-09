package jira

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Instance *instance
	Issues   issues
}

type instance struct {
	instanceURL, username, password string
}

func NewInstance(url, user, pass string) *Client {
	a := &instance{instanceURL: url, username: user, password: pass}
	return injectClient(a)
}

func injectClient(a *instance) *Client {
	client := &Client{Instance: a}
	client.Issues = &Issues{c: client}
	return client
}

func (c *Client) execute(method string, url string) *http.Response {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", c.Instance.instanceURL, url), nil)

	// Set up headers and auth
	req.SetBasicAuth(c.Instance.username, c.Instance.password)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return resp
	}

	fmt.Println("Response Status: ", resp.Status)
	fmt.Println("Response Headers: ", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Body: ", string(body))
	panic("Error with HTTP request")
}
